package IjinController

import (
	DBConfig "HRIS-API/Controllers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	I "HRIS-API/Model"

	//"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateIjin(w http.ResponseWriter, r *http.Request) {

	db, err = DBConfig.Config()
	payloads, _ := ioutil.ReadAll(r.Body)

	var ijin I.T_Ijin

	fmt.Println("Endpoint hit: Request Ijin")

	json.Unmarshal(payloads, &ijin)
	db.Create(&ijin)
	db.Save(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Success create ijin"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetIjin(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get data ijin")
	ijin := []I.T_Ijin{}

	db.Raw("SELECT * FROM T_Ijin  ").Scan(&ijin)

	res := I.Result{Code: 200, Data: ijin, Message: "Success get ijin"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetIjinByID(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var ijin I.T_Ijin
	db.Where("ID = ?", ID).First(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Success get ijin"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func UpdateIjin(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var ijinUpdate I.T_Ijin
	json.Unmarshal(payloads, &ijinUpdate)

	var ijin I.T_Ijin
	db.Where("ID = ?", ID).Updates(ijinUpdate)

	res := I.Result{Code: 200, Data: ijin, Message: "Success update ijin"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteIjin(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var ijin I.T_Ijin
	db.Where("ID = ?", ID).Delete(&ijin)

	res := I.Result{Code: 200, Message: "Success delete ijin"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func ApproveIjinHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_Head: "Approved", Done_Head: "true", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Success approve ijin by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveIjinHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_Head2: "Approved", Done_Head2: "true", Approve_Head_Date2: dt.Format("2006-01-02")})

	res := I.Result{Code: 200, Data: ijin, Message: "Success approve ijin by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveIjinHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_HRD: "Approved", Done_HRD: "true", Approve_HRD_Date: dt.Format("2006-01-02")})

	res := I.Result{Code: 200, Data: ijin, Message: "Success approve by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectIjinHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_Head: "Rejected", Done_Head: "false", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Reject ijin by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectIjinHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_Head: "Rejected", Done_Head: "false", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Reject ijin by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectIjinHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	ijinUpdate := []I.T_Ijin{}
	json.Unmarshal(payloads, &ijinUpdate)

	ijin := []I.T_Ijin{}
	db.Where("ID = ?", ID).Updates(I.T_Ijin{Type_Approve_Head: "Rejected", Done_Head: "false", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := I.Result{Code: 200, Data: ijin, Message: "Reject ijin by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
