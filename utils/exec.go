package utils

import (
	"fmt"
	"os/exec"

	a "github.com/logrusorgru/aurora"
)

func PrintError(message string) {
	fmt.Println(a.Red(message))
}

func PrintDebug(message string) {
	fmt.Println(a.Cyan(message))
}

func PrintVerbose(message string) {
	fmt.Println(a.Green(message))
}

type Executor func(string, string, bool)

func ExecuteRemote(command string, hostname string, verbose bool) {
	cmd := fmt.Sprintf("'%s'", command)

	if verbose {
		PrintVerbose(fmt.Sprintf("Connecting to %s; Executing %s", hostname, cmd))
	}

	out, err := exec.Command("ssh", "-q", hostname, "bash", "-c", cmd).CombinedOutput()
	if err != nil {
		PrintError(fmt.Sprintf("Error [%s] : %s", cmd, err))
	}

	fmt.Print(string(out))
}

func ExecuteLocal(command string, _hostname string, _verbose bool) {
	cmd := fmt.Sprintf("'%s'", command)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		PrintError(fmt.Sprintf("Error [%s] : %s", cmd, err))
	}

	fmt.Print(string(out))
}
