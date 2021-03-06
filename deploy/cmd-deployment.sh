kubectl -n braintree run \
           msgparser\
           --image=hemantrumde/msgparser:v2.0\
           --port=4000\
           --replicas=2
