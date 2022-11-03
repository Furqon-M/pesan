package main

import (
	"net/http"

	"github.com/Furqon-M/pesan/controllers/pelanggancontroller"
)

func main() {

	http.HandleFunc("/", pelanggancontroller.Index)
	http.HandleFunc("/pelanggan", pelanggancontroller.Index)
	http.HandleFunc("/pelanggan/index", pelanggancontroller.Index)
	http.HandleFunc("/pelanggan/add", pelanggancontroller.Add)
	http.HandleFunc("/pelanggan/edit", pelanggancontroller.Edit)
	http.HandleFunc("/pelanggan/delete", pelanggancontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
