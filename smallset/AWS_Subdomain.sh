for M in $(ls *.msg)
do
   echo "======[ $M ]===================="
   DATA="'$(cat $M)'"
   curl --request POST\
        --header 'Content-Type: text/plain'\
          http://msgparser.hemantrumde.com:4000/rawmsg \
         -d "$DATA"
   echo
   echo
done

