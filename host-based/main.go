package main

import (
       "os"
       "strings"
       "net/http"
)
var testArg string
var port string
func main() {
       if len(os.Args) > 0 {
              testArg = os.Args[1]
              if len(os.Args) > 1 {
                     port = os.Args[2]
              }
       }else {
              panic("insufficient os.Args")
       }
       if port == ""{
              port = "9058"
       }
       mux := http.NewServeMux()
       mux.HandleFunc("/healthz",healthzHandler)
       http.ListenAndServe(":"+port, mux)
}
func healthzHandler(w http.ResponseWriter, r *http.Request){
       hostname, _ := os.Hostname()

       if strings.Contains(hostname, testArg){
              w.WriteHeader(http.StatusOK)
              w.Write([]byte("OK"))
       }else {
              w.WriteHeader(http.StatusInternalServerError)
              w.Write([]byte("testArg didnt matched"))
       }
}

