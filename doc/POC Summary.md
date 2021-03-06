# POC Summary 
```
kubectl get ns
NAME              STATUS   AGE
default           Active   12m
kube-node-lease   Active   12m
kube-public       Active   12m
kube-system       Active   12m

kubectl create namespace braintree
namespace/braintree created

kubectl get ns
NAME              STATUS   AGE
braintree         Active   31s
default           Active   13m
kube-node-lease   Active   13m
kube-public       Active   13m
kube-system       Active   13m

kubectl -n braintree run \
           msgparser\
           --image=hemantrumde/msgparser:v2.0\
           --port=4000\
           --replicas=2

kubectl -n braintree get deployments
NAME        READY   UP-TO-DATE   AVAILABLE   AGE
msgparser   2/2     2            2           3m20s

kubectl -n braintree get pods
NAME                         READY   STATUS    RESTARTS   AGE
msgparser-68cc789459-st9fb   1/1     Running   0          22s
msgparser-68cc789459-tlmtd   1/1     Running   0          22s

kubectl -n braintree logs -f msgparser-68cc789459-st9fb
2021/03/06 16:14:07 Raw Email text Parser Version: 1.0

kubectl -n braintree logs -f msgparser-68cc789459-tlmtd
2021/03/06 16:14:08 Raw Email text Parser Version: 1.0

kubectl -n braintree expose \
           deployment msgparser \
           --type=NodePort \
           --name=svc-msgparser-nodeport

kubectl -n braintree get svc
NAME                     TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
svc-msgparser-nodeport   NodePort   10.100.56.220   <none>        4000:32252/TCP   5m12s

kubectl -n braintree describe svc/svc-msgparser-nodeport
Name:                     svc-msgparser-nodeport
Namespace:                braintree
Labels:                   run=msgparser
Annotations:              <none>
Selector:                 run=msgparser
Type:                     NodePort
IP:                       10.100.56.220
Port:                     <unset>  4000/TCP
TargetPort:               4000/TCP
NodePort:                 <unset>  32252/TCP
Endpoints:                192.168.28.80:4000,192.168.55.189:4000
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```
* **Testing on local server, Docker container and EKS**
```
#===========================================================
# Login to worker nodes and access Microservice 
#export SERVICE_HOST="localhost"

#Node 1
#export SERVICE_HOST="192.168.28.80"

#Node 2
#export SERVICE_HOST="192.168.55.189"

#Cluser Service
export SERVICE_HOST="10.100.56.220"

for M in $(ls *.msg)
do
   echo "======[ $M ]===================="
   DATA="'$(cat $M)'"
   curl --request POST\
        --header 'Content-Type: text/plain'\
          http://${SERVICE_HOST}:4000/rawmsg \
         -d "$DATA"
   echo
   echo
done
#===========================================================
```
