# Waggle Admin CLI

A CLI Suite for Managing Waggle Enabled Devices

## Usage

While the tool is under development, the best way to install it is to clone
the repo, pull master and manually install it:

```bash
$ git clone git@github.com:UrbanCCD-UChicago/waggle-cli.git waggle
$ cd waggle
$ go get .
$ go install
```

From there you can use `waggle` like any other system binary (assuming you have
your paths setup in your bash config).

```
$ waggle -h
A CLI suite to manage Waggle enabled devices

Usage:
  waggle [flags]
  waggle [command]

Available Commands:
  connect-and-execute    SSHes into a node and runs a given command
  current-usage          SSHes into a node and runs the current usage script
  date                   SSHes into a node and runs "date"
  disable-edge-processor SSHes into a node and disables its edge processor
  disable-waggle-plugin  SSHes into a node diables a given plugin
  disk-usage             SSHes into a node and runs "df -k"
  eereset-wagman         SSHes into a node and ee-resets the wagman
  enable-edge-processor  SSHes into a node and enables the edge processor
  enable-waggle-plugin   SSHes into a node and enables a given plugin
  fail-counts            SSHes into a node and runs the fail counts script
  help                   Help about any command
  hostname               SSHes into a node and runs "cat /etc/hostname"
  list-nodes             Lists node information stored in the local cache
  list-processes         SSHes into a node and runs "systemctl | grep"
  ls-dev                 SSHes into a node and runs "ls -l /dev/waggle_*"
  modem-info             SSHes into a node and runs the modem-info script
  network-data-usage     SSHes into a node and runs "ifconfig ppp0"
  node-info              Prints detailed information about a node
  put-emmc-mode          SSHes into a node puts it into EMMC mode
  put-sd-mode            SSHes into a node and puts it into SD mode
  refresh-cache          Refreshes the local cache of node information
  reset-wagman           SSHes into a node and resets the wagman
  restart-waggle-plugin  SSHes into a node and restarts a given plugin
  uptime                 SSHes into a node and runs "uptime"

Flags:
  -d, --debug           toggles debug output
  -h, --help            help for waggle
  -n, --node string     selects which node to work on
  -s, --system string   selects which system to target
  -v, --verbose         toggles verbose output

Use "waggle [command] --help" for more information about a command.
```
