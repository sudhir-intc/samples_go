package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//httpSendWithoutCtx()
	httpSendWithCtx()
}

func httpSendWithCtx() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sending a requst to the server")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}

func httpSendWithoutCtx() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout, res.Body)
}
