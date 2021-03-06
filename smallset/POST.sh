# Change following HOST to test the service from localhost, EC2 and EKS
# Service Unit Test

export SERVICE_HOST="localhost"
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
