package main

import (
	"fmt"
	"log"
	"net/http"

	pelanggan_struct "github.com/ahmad/praktikum/ujian/model"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", pelanggan_struct.AmbilPelanggans).Methods("GET")
	r.HandleFunc("/pelanggan/{id}", pelanggan_struct.AmbilPelanggan).Methods("GET")
	r.HandleFunc("/tambahpelanggan", pelanggan_struct.TambahPelanggan).Methods("POST")
	r.HandleFunc("/ubahpelanggan/{id}", pelanggan_struct.UbahPelanggan).Methods("PUT")
	r.HandleFunc("/hapuspelanggan/{id}", pelanggan_struct.HapusPelanggan).Methods("DELETE")
	var portNumber string = ":3000"
	fmt.Println("Server Running di Port", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, r))
}
