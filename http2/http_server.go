package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	serverHTTP2 := &http.Server{Addr: ":10080"}
	err := http2.ConfigureServer(serverHTTP2, &http2.Server{})
	if err != nil {
		log.Print("failed at configuring HTTP2 server")
	}

	/*
		_, cancel := context.WithCancel(context.Background())

		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		go func() {
			sig := <-osSignals
			log.Printf("Received signal: %#v", sig)
		}()
	*/
	http.HandleFunc("/", handlerWithCtx)
	log.Print("Starting the http server")
	go func() {
		log.Print("Starting the http server 2")
		log.Fatal(http.ListenAndServe(":9080", nil))
	}()
	log.Fatal(serverHTTP2.ListenAndServeTLS("certs/server-cert.pem",
		"certs/server-key.pem"))

}

func handlerWithoutCtx(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler Started")
	defer log.Printf("Handler Completed")
	time.Sleep(5 * time.Second)
	fmt.Fprintln(w, "Hello")
}

func handlerWithCtx(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	/* Dump the request received */
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	log.Println(string(dump))

	defer log.Printf("Handler Completed")
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "Hello Thanks !!!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
