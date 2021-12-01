package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

var productsURL string

func init() {

	productsURL = os.Getenv("PRODUCT_URL")
	
}

func displayCheckout(w http.ResponseWriter, r *http.Request ) {

	w.Write([]Byte("Olá checkout"))

}

func main() {

	r:= mux.NewRouter()
	r.HandleFunc("/{id}", displayCheckout)

	http.ListAndServe(":8082", r)

}



