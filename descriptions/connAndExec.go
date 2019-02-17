package descriptions

// ConnAndExecDesc :
const ConnAndExecDesc = `SSHes into a node and runs a given command

EXAMPLES
========

$ waggle connect-and-execute -n 004 -s nc 'ls -l'
{
  "node controller": [
    "total 76",
    "-rw-r--r-- 1 root root     0 Dec 17 14:16 fs_locked",
    "-r-------- 1 root root  1675 Nov  9 21:41 id_rsa_waggle_registration",
    "-rw-r--r-- 1 root root  3187 Oct 25 14:30 rabbitmq-release-signing-key.asc",
    "-rw-r--r-- 1 root root 67813 Nov  9 21:41 report.txt"
  ]
}

CAVEATS
=======

* Use of the --system flag is required for this command
`
