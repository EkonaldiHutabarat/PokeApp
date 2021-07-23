package MasterTypeController

import (
	DBConfig "HRIS-API/Controllers"
	T "HRIS-API/Model"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetCutiType(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get Data Cuti Type")
	cutiType := []T.M_Lembur_Type{}

	db.Raw("SELECT * FROM M_Cuti_Type ").Scan(&cutiType)

	res := T.Result{Code: 200, Data: cutiType, Message: "Success get cuti type"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetLemburType(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get Data Lembur Type")
	lemburType := []T.M_Lembur_Type{}

	db.Raw("SELECT * FROM M_Lembur_Type ").Scan(&lemburType)

	res := T.Result{Code: 200, Data: lemburType, Message: "Success get lembur type"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetIjinType(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get Data Cuti Type")
	ijinType := []T.M_Ijin_Type{}

	db.Raw("SELECT * FROM M_Ijin_Type ").Scan(&ijinType)

	res := T.Result{Code: 200, Data: ijinType, Message: "Success get Ijin type"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
