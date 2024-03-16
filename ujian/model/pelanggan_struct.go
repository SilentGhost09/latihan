package pelanggan_struct

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	config_db "github.com/ahmad/praktikum/ujian/config"
	_ "github.com/lib/pq"
)

type Pelanggan struct {
	Pel_Id      string
	Pel_Nama    string
	Pel_Jk      int8
	Pel_Alamat  string
	Pel_Hp      string
	Pel_Ktp     string
	Pel_Tgllhr  time.Time
	Pel_Tmptlhr string
	Pel_Create  time.Time
	Pel_Update  time.Time
}

func getpsqlinfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config_db.DBHost, config_db.DBPort, config_db.DBUser, config_db.DBPass, config_db.DBName)
}
func AmbilPelanggans(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlinfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	var pelanggans []Pelanggan
	result, err := db.Query("SELECT * FROM pelanggan")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var pelanggan Pelanggan
		err := result.Scan(&pelanggan.Pel_Id, &pelanggan.Pel_Nama, &pelanggan.Pel_Jk, &pelanggan.Pel_Alamat, &pelanggan.Pel_Hp,
			&pelanggan.Pel_Ktp, &pelanggan.Pel_Tgllhr, &pelanggan.Pel_Tmptlhr, &pelanggan.Pel_Create, &pelanggan.Pel_Update)
		if err != nil {
			panic(err.Error())
		}
		pelanggans = append(pelanggans, pelanggan)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelanggans)
}

func AmbilPelanggan(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlinfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	params := mux.Vars(r)
	pel_id := params["id"]
	var pelanggan Pelanggan
	sqlStatement := "SELECT * FROM pelanggan WHERE pel_id = " + pel_id
	result, err := db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&pelanggan.Pel_Id, &pelanggan.Pel_Nama, &pelanggan.Pel_Jk, &pelanggan.Pel_Alamat, &pelanggan.Pel_Hp,
			&pelanggan.Pel_Ktp, &pelanggan.Pel_Tgllhr, &pelanggan.Pel_Tmptlhr, &pelanggan.Pel_Create, &pelanggan.Pel_Update)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelanggan)
}
func TambahPelanggan(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlinfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	var pelanggan Pelanggan
	_ = json.NewDecoder(r.Body).Decode(&pelanggan)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelanggan)
}
func UbahPelanggan(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlinfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	var pelanggan Pelanggan
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&pelanggan)
	pelanggan.Pel_Id = params["id"]
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelanggan)
}
func HapusPelanggan(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlinfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	params := mux.Vars(r)
	pel_id := params["id"]
	var pelanggan Pelanggan
	sqlStatement := "DELETE * FROM pelanggan WHERE pel_id = " + pel_id
	result, err := db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&pelanggan.Pel_Id, &pelanggan.Pel_Nama, &pelanggan.Pel_Jk, &pelanggan.Pel_Alamat, &pelanggan.Pel_Hp,
			&pelanggan.Pel_Ktp, &pelanggan.Pel_Tgllhr, &pelanggan.Pel_Tmptlhr, &pelanggan.Pel_Create, &pelanggan.Pel_Update)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pelanggan)
}
