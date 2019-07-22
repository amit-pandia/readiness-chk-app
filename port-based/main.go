package main

import (
       "os"
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
              port = "9057"
       }
       mux := http.NewServeMux()
       mux.HandleFunc("/healthz",healthzHandler)
       http.ListenAndServe(":"+port, mux)
}
func healthzHandler(w http.ResponseWriter, r *http.Request){

       if port == "6666"{
              w.WriteHeader(http.StatusOK)
              w.Write([]byte("OK"))
       }else {
              w.WriteHeader(http.StatusInternalServerError)
              w.Write([]byte("port didnt matched"))
       }
}

