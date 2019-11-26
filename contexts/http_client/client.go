package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

type person struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	//httpSendWithoutCtx()
	httpSendWithCtx()
	//sudhir := person{"sudhir", 32}
	//httpPutWithCtx(sudhir)
}

func httpSendWithCtx() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancelTimeout(ctx, 2*time.Second)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	requestBody, err := json.Marshal(map[string]string{
		"name":  "sudhir xyz",
		"email": "xyz@intel.com",
	})

	timeout := time.Duration(100 * time.Second)

	client := http.Client{Timeout: timeout}
	http_url := "http://localhost:9080"
	method := "POST"

	/* Check the url type - if its https or http */
	u, error := url.Parse(http_url)
	if error != nil {
		log.Fatal(error)
		return
	}
	log.Printf("Url scheme:%s", u.Scheme)

	req, err := http.NewRequest(method, http_url, bytes.NewBuffer(requestBody))
	req.Header.Set("User-Agent", "NEF-OPENNESS-1912")
	req.Header.Set("Content-Type", "application/json")
	req.WithContext(ctx)
	/* Add user-agent header and content-type header */

	log.Printf("Sending a requst to the server")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	log.Println("Headers in the response =>")
	for k, v := range res.Header {
		log.Printf("%q:%q\n", k, v)
	}
	log.Println("Body in the response =>")
	body, err := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	/* Dump the actual response received */
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", dump)

}

func httpSendWithoutCtx() {
	res, err := http.Get("http2://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	io.Copy(os.Stdout, res.Body)
}
