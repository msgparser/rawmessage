# How to build source?
The service is implemented in golang. If golang compiler is available on server, then source can be compiled easily. If golang is missing on local server, then source can be compiled in Docker container. In both the cases, we need Docker engine. Docker system can be installed depending on your Linux flavor. 

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
In this exercise TCP port 4000 is hardcoded. Currently os.Args is not used to accept a port by command line arguments. Execute <code>msgparser</code> binary and verify listening port from different window.
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
# Local sandbox testing 
Linux curl command can be utilized to test our service locally. Use POST.sh to test service. In POST.sh change SERVICE_HOST to test on local server, EC2 or EKS
```
$ cd ../smallset
$ ls -l 
total 328
-rw-rw-r--. 1 hemant hemant 47243 Apr  1  2011 20110401_1000mercis_14461469_html.msg
-rw-rw-r--. 1 hemant hemant  8167 Apr  1  2011 20110401_aamarketinginc_14456749_html.msg
-rw-rw-r--. 1 hemant hemant  7133 Apr  1  2011 20110401_aeg_14465739_html.msg
-rw-rw-r--. 1 hemant hemant 51288 Apr  1  2011 20110401_alchemyworx_14461429_multialt.msg
-rw-rw-r--. 1 hemant hemant 34201 Apr  1  2011 20110401_americancollegiatemarketing_14461959_multialt.msg
-rw-rw-r--. 1 hemant hemant  7639 Apr  1  2011 20110401_beliefnet_14461399_html.msg
-rw-rw-r--. 1 hemant hemant 13543 Apr  1  2011 20110401_beliefnet_14464159_html.msg
-rw-rw-r--. 1 hemant hemant 12284 Apr  1  2011 20110401_boydgamingcorporation_14465279_multialt.msg
-rw-rw-r--. 1 hemant hemant  2903 Apr  1  2011 20110401_citibanksingaporelimited_14456499_multialt.msg
-rw-rw-r--. 1 hemant hemant 12638 Apr  1  2011 20110401_cobaltgroup_14464029_html.msg
-rw-rw-r--. 1 hemant hemant 99805 Apr  1  2011 20110401_compostmarketingab_14459379_multialt.msg
-rw-rw-r--. 1 hemant hemant  7242 Apr  1  2011 20110401_corel_14460139_html.msg
-rwxrwxr-x. 1 hemant hemant   278 Mar  6 10:46 POST.sh

$ ./POST.sh
......
======[ 20110401_beliefnet_14461399_html.msg ]====================
{"To":"\u003cbeliefnet@cp.monitor1.returnpath.net\u003e","From":"Announce - Beliefnet Sponsor \u003cspecialoffers@mail.beliefnet.com\u003e","Date":"Fri,  1 Apr 2011 08:12:00 -0600 (MDT)","Subject":"[SP] Grant Funding May Be Available for Top Online Colleges. Get Free Info Today.","MessageID":"\u003c527817310.344.1301667087687.JavaMail.root@mail.beliefnet.com\u003e"}


======[ 20110401_beliefnet_14464159_html.msg ]====================
{"To":"\u003cbeliefnet@cp.monitor1.returnpath.net\u003e","From":"Chicken Soup - Beliefnet Partner \u003cspecialoffers@mail.beliefnet.com\u003e","Date":"Fri,  1 Apr 2011 10:32:42 -0600 (MDT)","Subject":"[SP] The Art of Positive Thinking","MessageID":"\u003c463918295.411.1301674909118.JavaMail.root@mail.beliefnet.com\u003e"}
.....
.....

```

# Compliation in GOLANG Docker container 
If golang compiler is missing on your local server, then golang Docker container can use to build our microservice image.

## ImageBuilder 
```
FROM       golang:latest AS builder
ARG        CFLAGS="-s -w"
WORKDIR    /src
COPY      ./msgparser.go /src
RUN        go build -ldflags "$CFLAGS" /src/msgparser.go

FROM       centos:latest
ARG        VERSION=1.0
WORKDIR    /validity
LABEL      Version=$VERSION
LABEL      WaterMark="Raw Email text message parser to extract Message fields"
LABEL      Author="Hemant Rumde"
COPY     --from=builder /src/msgparser /validity/msgparser
EXPOSE     4000/tcp
ENTRYPOINT /validity/msgparser
```
* **How to build?**
```
$ docker build -t hemantrumde/rawmessage:v1.3 -f ImageBuilder .
$ docker push hemantrumde/rawmessage:v1.3
```
