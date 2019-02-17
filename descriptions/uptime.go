package descriptions

// UptimeDesc :
const UptimeDesc = `SSHes into a node and runs "uptime"

EXAMPLES
========

$ waggle uptime -n 004
{
  "edge processor": [
    " 17:05:43 up 8 days, 5 min,  1 user,  load average: 0.81, 0.51, 0.40"
  ],
  "node controller": [
    " 17:05:41 up 54 days,  1:57,  1 user,  load average: 0.32, 0.34, 0.32"
  ],
  "wagman": [
    "10356146"
  ]
}

$ waggle uptime -n 004 -s nc
{
  "node controller": [
    " 17:06:19 up 54 days,  1:57,  1 user,  load average: 0.18, 0.30, 0.31"
  ]
}
`
