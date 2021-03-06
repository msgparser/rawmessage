# How to build source?
The service is implemented in golang. If golang compiler is available on server, then it can be compiled easily. If golang is missing on local server, then it can be
compiled in Docker container. In both the cases, we need Docker engine. Docker system can be installed depending on your Linux flavor. 

## Compiling by local golang compiler 

```
$ git clone https://github.com/msgparser/rawmessage.git
$ cd rawmessage/src
$ ls -l 
total 16
-rw-rw-r--. 1 hemant hemant  145 Mar  6 09:48 Dockerfile
-rw-rw-r--. 1 hemant hemant   26 Mar  6 09:48 go.mod
-rw-rw-r--. 1 hemant hemant  486 Mar  6 09:48 ImageBuilder
-rw-rw-r--. 1 hemant hemant 2080 Mar  6 09:48 msgparser.go
```
## Build 
```
$ go build msgparser
$ ls -l msgparser
-rwxrwxr-x. 1 hemant hemant 6624538 Mar  6 09:52 msgparser
```
## Local Execution 
In this exercise TCP port 4000 is hardcoded. Currently os.Args is not used to accept port by command line arguments. Rum msgparser binary and verify listening port from
different window.
```
$ ./msgparser
2021/03/06 10:23:27 Raw Email text Parser Version: 1.0

```
## Verify listening port by Linux commands
```
$ ss -ltn
State      Recv-Q Send-Q   Local Address:Port                  Peer Address:Port
LISTEN     0      128                  *:22                               *:*
LISTEN     0      100          127.0.0.1:25                               *:*
LISTEN     0      128               [::]:4000                          [::]:*
LISTEN     0      128               [::]:2375                          [::]:*
LISTEN     0      128               [::]:22                            [::]:*
LISTEN     0      100              [::1]:25                            [::]:*
LISTEN     0      64                [::]:3001                          [::]:*

$  nc -v localhost 4000
Ncat: Version 7.50 ( https://nmap.org/ncat )
Ncat: Connected to ::1:4000.

$ lsof -i:4000
COMMAND     PID   USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
msgparser 22663 hemant    3u  IPv6 311600      0t0  TCP *:terabase (LISTEN)

```
