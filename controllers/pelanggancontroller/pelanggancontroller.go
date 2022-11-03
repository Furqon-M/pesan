package pelanggancontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Furqon-M/pesan/entities"
	"github.com/Furqon-M/pesan/libraries"
	"github.com/Furqon-M/pesan/models"
)

var validation = libraries.NewValidation()
var pelangganModel = models.NewPelangganModel()

func Index(response http.ResponseWriter, request *http.Request) {

	pelanggan, _ := pelangganModel.FindAll()

	data := map[string]interface{}{
		"pelanggan": pelanggan,
	}

	temp, err := template.ParseFiles("views/pelanggan/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/pelanggan/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var pelanggan entities.Pelanggan
		pelanggan.NoPesanan = request.Form.Get("no_pesanan")
		pelanggan.Pesanan = request.Form.Get("pesanan")
		pelanggan.Harga = request.Form.Get("harga")
		pelanggan.Jumlah = request.Form.Get("jumlah")
		pelanggan.Total = request.Form.Get("total")
		pelanggan.Meja = request.Form.Get("meja")
		pelanggan.Status = request.Form.Get("status")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pelanggan)

		if vErrors != nil {
			data["pelanggan"] = pelanggan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Pesanan Berhasil Disimpan "
			pelangganModel.Create(pelanggan)
		}
		temp, _ := template.ParseFiles("views/pelanggan/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var pelanggan entities.Pelanggan
		pelangganModel.Find(id, &pelanggan)

		data := map[string]interface{}{
			"pelanggan": pelanggan,
		}

		temp, err := template.ParseFiles("views/pelanggan/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var pelanggan entities.Pelanggan
		//pelanggan.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		pelanggan.NoPesanan = request.Form.Get("no_pesanan")
		pelanggan.Pesanan = request.Form.Get("pesanan")
		pelanggan.Harga = request.Form.Get("harga")
		pelanggan.Jumlah = request.Form.Get("jumlah")
		pelanggan.Total = request.Form.Get("total")
		pelanggan.Meja = request.Form.Get("meja")
		pelanggan.Status = request.Form.Get("status")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(pelanggan)

		if vErrors != nil {
			data["pelanggan"] = pelanggan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Pesanan Berhasil Diperbarui "
			pelangganModel.Update(pelanggan)
		}
		temp, _ := template.ParseFiles("views/pelanggan/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	pelangganModel.Delete(id)

	http.Redirect(response, request, "/pelanggan", http.StatusSeeOther)

}
