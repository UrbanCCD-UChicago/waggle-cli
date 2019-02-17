package descriptions

// DateDesc :
const DateDesc = `SSHes into a node and runs "date"

EXAMPLES
========

$ waggle date -n 004
{
  "edge processor": [
    "2019-02-05 17:08:01+00:00"
  ],
  "node controller": [
    "2019-02-05 17:07:59+00:00"
  ],
  "wagman": [
    "2019 2 5 17 7 59"
  ]
}

$ waggle date -n 004 -s ep
{
  "edge processor": [
    "2019-02-05 17:08:13+00:00"
  ]
}
`
