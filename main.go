/**
 * Created by jbisa on 8/20/17.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/patternMiner/adserver/context"
	"github.com/patternMiner/adserver/handlers"
)

func main() {
	err := context.InitContext()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/info", handlers.InfoHandler)

	http.HandleFunc("/ad", handlers.AdHandler)

	fs := http.FileServer(http.Dir("client/adtag/dist"))
	http.Handle("/static/",
		http.StripPrefix("/static/", fs))

	fmt.Println("Starting up the adserver service on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
