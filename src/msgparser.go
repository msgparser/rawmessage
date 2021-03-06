package main

// Author : Hemant Rumde
// Date   : Fri Mar  5 21:42:30 EST 2021

// Description:
// Raw Email message is a feed to restApi /rawmsg

import (
    "encoding/json"
    "log"
    "net/http"
    "io/ioutil"
    "net/mail"
    "strings"
)

type RawMessage struct {
     Name string
}

type EmailFields struct {
     To          string
     From        string
     Date        string
     Subject     string
     MessageID   string
}

func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}

func MessageCreate(w http.ResponseWriter, r *http.Request) {

     body, err := ioutil.ReadAll(r.Body)

     if err != nil {
       log.Printf("Error reading body: %v", err)
       http.Error(w, "can't read body", http.StatusBadRequest)
       return
     }
     log.Printf("Received raw message: IP:%s size:%d\n",ReadUserIP(r),len(body))

     raw := strings.NewReader(string(body))
     MailText, err2 := mail.ReadMessage(raw)
     if err2 != nil {
       log.Print("Invalid Email text")
       http.Error(w, "Invalid raw email text", http.StatusBadRequest)
       return
     }

     header      := MailText.Header
     Date        := header.Get("Date")
     From        := header.Get("From")
     To          := header.Get("To")
     Subject     := header.Get("Subject")
     MessageID   := header.Get("Message-ID")

     log.Printf("Mail Field replied: Date: %s From: %s To: %s Subject: %s Message-ID: %s\n",Date,From,To,Subject,MessageID)

     ReplyFields := EmailFields{To,From,Date,Subject,MessageID}

     w.Header().Set("Content-Type", "application/json")
     if RawErr := json.NewEncoder(w).Encode(ReplyFields); RawErr != nil {
       panic(RawErr)
     }
}

func main() {
    log.Printf("Raw Email text Parser Version: 1.0")
    mux := http.NewServeMux()
    mux.HandleFunc("/rawmsg", MessageCreate)
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
