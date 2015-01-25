package main

import (
	"log"
	"net/http"

	"github.com/casualjim/go-swagger/examples/2.0/petstore/server/api"
)

func main() {
	petstoreAPI, err := api.NewPetstore()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Serving petstore api on http://0.0.0.0:8344")
	http.ListenAndServe(":8344", petstoreAPI)
}
