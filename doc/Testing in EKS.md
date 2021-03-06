# Testing in EKS
## EKS cluster
AWS EKS small cluster with two workers can be built to test **rawmessage** microservice. Add IAM user, get access-key, secret-key and copy on your local server to access
EKS remotely. Generate your SSH key to access this cluster.
```
$ ssh-keygen -t rsa -b 4096 -f mango

$ eksctl create cluster \
    --name mango \
    --version 1.16 \
    --region us-east-1 \
    --nodegroup-name mango-workers \
    --node-type t3.small \
    --ssh-access \
    --ssh-public-key ./mango.pub \
    --managed

```
## Docker Registry
I have added a public docker repository for this microservice. It does not need password. Hence Kubernete can pull images without secret in deployment. This will help
 end user to test this service without exchanging password.

**Public repository**
```
docker push hemantrumde/msgparser:tagname
docker pull hemantrumde/msgparser:tagname
 ```
