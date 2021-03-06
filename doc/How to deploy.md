# How to delploy
I used AWK EKS to test this service in cloud. I used my home linux server to access EKS. 

## Deploy Kubernetes deployment
```
$ cd rawmessage/deploy
$ kubectl create -f deployment.yaml
$ kubectl create -f LB-service.yaml
```
* Use Load balancer IP addresses in Route 53 record to access this service from internet. 

## Local single node cluster 
* use deploy/Node-Service.yaml to test this service locally.
