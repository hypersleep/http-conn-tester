# http-conn-tester
This tool helps you to debug network troubles with http connections and collect results distributed by time.

# Usage

Clone and complie:

```
git clone git@github.com:hypersleep/http-conn-tester.git
go get
go build
```

Print help:

```
$ ./http-conn-tester --help
Usage of ./http-conn-tester:
  -d="8.8.8.8": DNS server
  -i=1s: Interval between requests
  -o="output": Output file
  -t=1s: Request timeout
  -u="https://google.com": URL
```

Let's test our resource!

```
$ ./http-conn-tester -d 8.8.8.8 -i 1s -t 1s -u https://qwe.asd/stats
2015/12/16 18:03:31 https://146.155.170.51/stats
2015/12/16 18:03:31
2015/12/16 18:03:31 &{7.359089ms 262.704685ms 104.566µs 270.16834ms OK}
2015/12/16 18:03:32 https://146.155.170.51/stats
2015/12/16 18:03:32
2015/12/16 18:03:32 &{30.388352ms 228.856082ms 53.335µs 259.297769ms OK}
2015/12/16 18:03:33 https://146.155.170.51/stats
2015/12/16 18:03:33
2015/12/16 18:03:33 &{6.905023ms 214.277717ms 41.801µs 221.224541ms OK}
2015/12/16 18:03:34 https://146.155.170.51/stats
2015/12/16 18:03:34
2015/12/16 18:03:34 &{6.912308ms 216.67633ms 91.68µs 223.680318ms OK}
2015/12/16 18:03:35 https://188.226.150.79/stats
2015/12/16 18:03:35
2015/12/16 18:03:35 &{394.134464ms 284.094748ms 131.496µs 678.360708ms OK}
2015/12/16 18:03:36 https://188.226.150.79/stats
2015/12/16 18:03:36
2015/12/16 18:03:36 &{28.961265ms 218.001909ms 84.232µs 247.047406ms OK}
```
