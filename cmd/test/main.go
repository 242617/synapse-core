package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)

	go func() {

		// With signed certificate
		cert, err := tls.LoadX509KeyPair("keys/synapse.crawler-1.crt", "keys/synapse.crawler-1.key")
		if err != nil {
			log.Fatal(err)
		}

		caCert, err := ioutil.ReadFile("keys/rootca.crt")
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			RootCAs:      caCertPool,
		}
		tlsConfig.BuildNameToCertificate()
		transport := &http.Transport{TLSClientConfig: tlsConfig}
		client := &http.Client{Transport: transport}

		_, err = client.Get("https://synapse.local:8443/test1")
		if err != nil {
			log.Fatal(err)
		}

		// Without certificate
		_, err := http.Get("https://synapse.local:8443/test2")
		if err != nil {
			log.Fatal(err)
		}

	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %s\n", r.RequestURI)
	})
	log.Fatal(http.ListenAndServeTLS(":8443", "keys/synapse.core.crt", "keys/synapse.core.key", nil))
}
