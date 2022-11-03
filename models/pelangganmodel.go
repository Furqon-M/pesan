package models

import (
	"database/sql"
	"fmt"

	"github.com/Furqon-M/pesan/config"
	"github.com/Furqon-M/pesan/entities"
)

type PelangganModel struct {
	conn *sql.DB
}

func NewPelangganModel() *PelangganModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PelangganModel{
		conn: conn,
	}
}

func (p *PelangganModel) FindAll() ([]entities.Pelanggan, error) {

	rows, err := p.conn.Query("select * from pelanggan")
	if err != nil {
		return []entities.Pelanggan{}, err
	}
	defer rows.Close()

	var dataPelanggan []entities.Pelanggan
	for rows.Next() {
		var pelanggan entities.Pelanggan
		rows.Scan(&pelanggan.Id,
			&pelanggan.NoPesanan,
			&pelanggan.Pesanan,
			&pelanggan.Harga,
			&pelanggan.Jumlah,
			&pelanggan.Total,
			&pelanggan.Meja,
			&pelanggan.Status)

		if pelanggan.Status == "1" {
			pelanggan.Status = "berhasil"
		} else {
			pelanggan.Status = "menunggu"
		}
		dataPelanggan = append(dataPelanggan, pelanggan)
	}

	return dataPelanggan, nil

}

func (p *PelangganModel) Create(pelanggan entities.Pelanggan) bool {
	result, err := p.conn.Exec("insert into pelanggan(no_pesanan, pesanan, harga, jumlah, total, meja, status) values(?,?,?,?,?,?,?)",
		pelanggan.NoPesanan, pelanggan.Pesanan, pelanggan.Harga, pelanggan.Jumlah, pelanggan.Total, pelanggan.Meja, pelanggan.Status, pelanggan.Id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0

}

func (p *PelangganModel) Find(id int64, pelanggan *entities.Pelanggan) error {

	return p.conn.QueryRow("select * from pelanggan where id = ?", id).Scan(
		&pelanggan.Id,
		&pelanggan.NoPesanan,
		&pelanggan.Pesanan,
		&pelanggan.Harga,
		&pelanggan.Jumlah,
		&pelanggan.Total,
		&pelanggan.Meja,
		&pelanggan.Status)

}

func (p *PelangganModel) Update(pelanggan entities.Pelanggan) error {

	_, err := p.conn.Exec(
		"update pelanggan set no_pesanan = ?, pesanan = ?, harga = ?, jumlah = ?, total = ?, meja = ?, status = ? where",
		pelanggan.NoPesanan, pelanggan.Pesanan, pelanggan.Harga, pelanggan.Jumlah, pelanggan.Total, pelanggan.Meja, pelanggan.Status, pelanggan.Id)

	if err != nil {
		return err
	}

	return nil

}

func (p *PelangganModel) Delete(id int64) {
	p.conn.Exec("delete from pasien where id = ?", id)
}
