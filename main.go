/**
 * Created by jbisa on 8/20/17.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/patternMiner/adserver/context"
	"github.com/patternMiner/adserver/handlers"
	"flag"
	"crypto/tls"
)

var (
	certPath string
	keyPath string
)

func init() {
	flag.StringVar(&certPath, "cert_path", "data/cert.pem", "SSL Certificate Path")
	flag.StringVar(&keyPath, "key_path", "data/key.pem", "SSL Key Path")
}

func main() {
	flag.Parse()

	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})
	mux.HandleFunc("/info", handlers.InfoHandler)

	mux.HandleFunc("/ad", handlers.AdHandler)

	fs := http.FileServer(http.Dir("client/adtag/dist"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	fmt.Println("Starting up the tester_match https service on port 8080")
	log.Fatal(srv.ListenAndServeTLS(certPath, keyPath))
}
