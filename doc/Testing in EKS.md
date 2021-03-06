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
# Deployment 
I have kept yaml files in deploy folder. I used these files to deploy this service in EKS. I used Route 53 subdomain for an instance running in Docker container on EC2.
Successfully tested with AWS load balancer. 

* **Deployment**
```
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    technical-exercise: "Raw Message text parser to extract fields"
  labels:
    run: msgparser
  name: msgparser
  namespace: braintree
spec:
  replicas: 2
  selector:
    matchLabels:
      run: msgparser
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      labels:
        run: msgparser
    spec:
      containers:
      - image: hemantrumde/msgparser:v2.0
        imagePullPolicy: Always
        name: msgparser
        ports:
        - containerPort: 4000
          protocol: TCP
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      securityContext: {}
```
* **Service for Node port**
```
apiVersion: v1
kind: Service
metadata:
  labels:
    run: msgparser
  name: svc-msgparser-nodeport
  namespace: braintree
spec:
  ports:
  - nodePort: 32577
    port: 4000
    protocol: TCP
    targetPort: 4000
  selector:
    run: msgparser
  type: NodePort
```
