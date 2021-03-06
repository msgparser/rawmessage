kubectl -n braintree expose \
           deployment msgparser \
           --type=NodePort \
           --name=svc-msgparser-nodeport
