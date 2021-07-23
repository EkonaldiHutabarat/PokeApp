package BeritaController

import (
	DBConfig "HRIS-API/Controllers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	B "HRIS-API/Model"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func CreateBerita(w http.ResponseWriter, r *http.Request) {

	db, err = DBConfig.Config()

	payloads, _ := ioutil.ReadAll(r.Body)

	var berita B.T_Berita

	fmt.Println("Endpoint hit: Insert Berita")

	json.Unmarshal(payloads, &berita)

	db.Create(&berita)
	//db.Raw("INSERT INTO T_Berita(ID,Title,News_Content,Expired_Date,Company_ID,Department_ID,Division_ID,Unit_ID,Insert_User,Insert_Date,Update_User,Update_Date)VALUES (?,?,?,?,?,?,?,?,?,?,?,?)").Scan(&berita)

	db.Save(&berita)
	res := B.Result{Code: 200, Data: berita, Message: "Insert data berita berhasil"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetDataBerita(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: get data berita sukses")
	news := []B.T_Berita{}

	db.Raw("SELECT * FROM T_Berita ").Scan(&news)

	res := B.Result{Code: 200, Data: news, Message: "Get Data Success"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}
func GetBeritaByID(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var berita B.T_Berita
	db.Where("ID = ?", ID).First(&berita)
	res := B.Result{Code: 200, Data: berita, Message: "succes get berita"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func UpdateBerita(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var beritaUpdate B.T_Berita
	json.Unmarshal(payloads, &beritaUpdate)

	var berita B.T_Berita
	db.Where("ID = ?", ID).Updates(beritaUpdate)

	res := B.Result{Code: 200, Data: berita, Message: "Success update berita"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteBerita(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var berita B.T_Berita
	db.Where("ID = ?", ID).Delete(&berita)

	res := B.Result{Code: 200, Message: "Success delete berita"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
