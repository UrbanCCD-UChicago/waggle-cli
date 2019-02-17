package descriptions

// DiskUsage :
const DiskUsage = `SSHes into a node and runs "df -k"

EXAMPLES
========

$ waggle disk-usage -n 004
{
  "edge processor": [
    "/dev/mmcblk1p2  7.2G  2.0G  4.9G  29% /",
    "/dev/mmcblk1p1  129M   14M  115M  11% /media/boot",
    "/dev/mmcblk1p3  7.3G  1.1G  5.9G  15% /wagglerw"
  ],
  "node controller": [
    "/dev/mmcblk0p2  7.2G  1.7G  5.1G  25% /",
    "/dev/mmcblk0p1  129M   12M  118M   9% /media/boot",
    "/dev/mmcblk0p3  7.3G  4.5G  2.4G  66% /wagglerw"
  ]
}
`
