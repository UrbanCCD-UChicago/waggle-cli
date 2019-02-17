package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"runtime"

	u "github.com/UrbanCCD-UChicago/waggle/utils"
	"github.com/spf13/cobra"
)

var (
	nodeFlag   string                    // --node
	targetFlag string                    // --system
	verbose    bool                      // --verbose
	debug      bool                      // --debug
	env        u.Env                     // computed from environment at runtime
	node       u.Node                    // got from nodeFlag
	targets    = make(map[string]string) // mapped hostnames to targets
	executor   u.Executor                // func we're using to execute commands on targets
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "waggle",
	Short: "A CLI suite to manage Waggle enabled devices",
	Run:   func(cmd *cobra.Command, args []string) {},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// setup env
		// all our nodes run linux, so if the command isn't called from a linux os we can safely
		// assume this is being called from an admin's computer
		if runtime.GOOS != "linux" {
			env = u.AdminEnv
		} else {
			// we have only two types of hardware platforms for our systems: an ODROID XU4 for the
			// edge processor and ODROID C1 for the node controller. we can grep the contents of
			// /proc/cpuinfo to find the value
			contents, err := ioutil.ReadFile("/proc/cpuinfo")
			if err != nil {
				log.Fatal(err)
			}
			strContents := string(contents)

			xu4Regex := regexp.MustCompile(`Hardware\s*:\s*ODROID-XU4`)
			c1Regex := regexp.MustCompile(`Hardware\s*:\s*ODROIDC`)

			if xu4Regex.MatchString(strContents) {
				env = u.EdgeProcessorEnv
			} else if c1Regex.MatchString(strContents) {
				env = u.NodeControllerEnv
			} else {
				// you're an admin and you're running linux -- excellent work!
				env = u.AdminEnv
			}
		}

		if debug {
			u.PrintDebug(fmt.Sprintf("env: %s", env.Name))
		}

		// get the node, if necessary
		if env.Name == u.AdminEnv.Name && nodeFlag != "" {
			n, err := u.GetNode(nodeFlag)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			node = n
		}

		if debug {
			u.PrintDebug(fmt.Sprintf("node: %s", node.Name))
		}

		// set the targets
		if targetFlag == u.NodeControllerTarget.Flag {
			hostname := fmt.Sprintf("%s%s", u.NodeControllerTarget.Host, node.Port)
			targets[u.NodeControllerTarget.Name] = hostname

		} else if targetFlag == u.EdgeProcessorTarget.Flag {
			hostname := fmt.Sprintf("%s%s", u.EdgeProcessorTarget.Host, node.Port)
			targets[u.EdgeProcessorTarget.Name] = hostname

		} else if targetFlag == u.WagmanTarget.Flag {
			hostname := fmt.Sprintf("%s%s", u.WagmanTarget.Host, node.Port)
			targets[u.WagmanTarget.Name] = hostname

		} else { // it's the whole node
			// add nc
			hostname := fmt.Sprintf("%s%s", u.NodeControllerTarget.Host, node.Port)
			targets[u.NodeControllerTarget.Name] = hostname

			// add ep
			hostname = fmt.Sprintf("%s%s", u.EdgeProcessorTarget.Host, node.Port)
			targets[u.EdgeProcessorTarget.Name] = hostname

			// add wgn
			hostname = fmt.Sprintf("%s%s", u.WagmanTarget.Host, node.Port)
			targets[u.WagmanTarget.Name] = hostname
		}

		if debug {
			u.PrintDebug(fmt.Sprintf("targets: %s", targets))
		}

		// set the executor function
		if node.Name != "" {
			executor = u.ExecuteRemote
		} else {
			executor = u.ExecuteLocal
		}

		if debug {
			u.PrintDebug(fmt.Sprintf("executor: %s\n", runtime.FuncForPC(reflect.ValueOf(executor).Pointer()).Name()))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&nodeFlag, "node", "n", "", "selects which node to work on")
	rootCmd.PersistentFlags().StringVarP(&targetFlag, "system", "s", "", "selects which system to target")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "toggles verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "toggles debug output")
}
