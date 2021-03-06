# Microservice Containerization

Linux Build server may be with or without golang compiler. If golang compiler is avaliable then a binary executable built locally can be containerized by following file
*  **Dockerfile**
```
FROM       centos:latest
WORKDIR    /validity
COPY       ./msgparser /validity/msgparser
EXPOSE     4000/tcp
ENTRYPOINT /validity/msgparser
```
*  **How to build ?**
```
$ ls -l
total 16
-rw-rw-r--. 1 hemant hemant  145 Mar  6 14:00 Dockerfile
-rw-rw-r--. 1 hemant hemant   26 Mar  6 14:00 go.mod
-rw-rw-r--. 1 hemant hemant  486 Mar  6 14:00 ImageBuilder
-rw-rw-r--. 1 hemant hemant 2080 Mar  6 14:00 msgparser.go

$ go build msgparser

$ ls -l msgparser
-rwxrwxr-x. 1 hemant hemant 6624538 Mar  6 14:06 msgparser

$ file msgparser
msgparser: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked (uses shared libs), not stripped

$ docker build -t hemantrumde/msgparser:v1.2 .

$ docker images hemantrumde/msgparser:v1.2
REPOSITORY              TAG                 IMAGE ID            CREATED             SIZE
hemantrumde/msgparser   v1.2                b028e82c0527        14 seconds ago      222MB

```
*  **Quick local test**

Before testing locally, make sure about port 4000. You can use ss, lsof, nc commands to confirm this port. You can use different port in this quick test.
```
$ docker container run -i -p 4000:4000 hemantrumde/msgparser:v1.2 
2021/03/06 19:16:55 Raw Email text Parser Version: 1.0
```
