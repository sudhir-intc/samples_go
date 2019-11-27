package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"golang.org/x/net/http2"
)

func main() {

	client := http.Client{}

	postbody, err := ioutil.ReadFile("json/AF_NEF_POST_01.json")
	if err != nil {
		log.Fatalf("Reading json file : %s", err)
		return
	}

	caCert, err := ioutil.ReadFile("certs/root-ca-cert.pem")
	if err != nil {
		log.Fatalf("Reading server certificate : %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}

	res, err := client.Post("https://localhost:8090/3gpp-traffic-influence/v1/AF_01/subscriptions",
		"application/json", bytes.NewBuffer(postbody))
	if err != nil {
		log.Fatalf("Failed go error :%s", err)
	}

	log.Println("Body in the response =>")
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	/* Dump the actual response received */
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", dump)

}
