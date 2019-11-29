package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	client := http.Client{}

	postbody, err := ioutil.ReadFile("json/AF_NEF_POST_01.json")
	if err != nil {
		log.Fatalf("Reading json file : %s", err)
		return
	}
	postSmfbody, err := ioutil.ReadFile("json/SMF_NEF_NOTIF_01.json")
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

	/* Post towards AF */
	log.Print("Triggering AF POST : https://localhost:8090/3gpp-traffic-influence/v1/AF_01/subscriptions")
	res, err := client.Post("https://localhost:8090/3gpp-traffic-influence/v1/AF_01/subscriptions",
		"application/json", bytes.NewBuffer(postbody))
	if err != nil {
		log.Fatalf("Failed go error :%s", err)
	}
	log.Println("Body in the response =>")
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	/* Wait for 10 seconds */
	time.Sleep(5 * time.Second)

	/* Post from SMF */
	log.Print("Triggering SMF POST : https://localhost:8090/3gpp-traffic-influence/v1/notification/upf")
	res, err = client.Post("https://localhost:8090/3gpp-traffic-influence/v1/notification/upf",
		"application/json", bytes.NewBuffer(postSmfbody))
	if err != nil {
		log.Fatalf("Failed go error :%s", err)
	}

	log.Println("Body in the response =>")
	body, _ = ioutil.ReadAll(res.Body)
	log.Println(string(body))

	/* Dump the actual response received */
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", dump)

}
