package descriptions

// LsDev :
const LsDev = `SSHes into a node and runs "ls -l /dev/waggle_*"

EXAMPLES
========

$ waggle ls-dev -n 004
{
  "edge processor": [
    "lrwxrwxrwx 1 root root  6 Jan 28 17:00 /dev/waggle_cam_bottom -\u003e video1",
    "lrwxrwxrwx 1 root root  6 Jan 28 17:00 /dev/waggle_cam_top -\u003e video0",
    "lrwxrwxrwx 1 root root 13 Jan 28 17:00 /dev/waggle_microphone -\u003e snd/controlC1"
  ],
  "node controller": [
    "lrwxrwxrwx 1 root root 7 Feb  1 16:44 /dev/waggle_coresense -\u003e ttyACM8",
    "lrwxrwxrwx 1 root root 7 Dec 13 15:09 /dev/waggle_sysmon -\u003e ttyACM1"
  ]
}

$ waggle ls-dev -n 004 | jq '."node controller"'
[
  "lrwxrwxrwx 1 root root 7 Feb  1 16:44 /dev/waggle_coresense -> ttyACM8",
  "lrwxrwxrwx 1 root root 7 Dec 13 15:09 /dev/waggle_sysmon -> ttyACM1"
]

CAVEATS
=======

* This does not work on Wagman
`
