package LemburController

import (
	DBConfig "HRIS-API/Controllers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	L "HRIS-API/Model"

	//"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateLembur(w http.ResponseWriter, r *http.Request) {

	db, err = DBConfig.Config()
	payloads, _ := ioutil.ReadAll(r.Body)

	var lembur = L.T_Lembur{}

	fmt.Println("Endpoint hit: Request Lembur")

	json.Unmarshal(payloads, &lembur)
	db.Create(&lembur)
	res := L.Result{Code: 200, Data: lembur, Message: "Success create lembur"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetLembur(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get data lembur")
	Lembur := []L.T_Lembur{}

	// db.Raw("SELECT * FROM T_Lembur ").Scan(&Lembur)
	db.Raw("SELECT cast(ID as varchar(max)) as ID, Emp_NIK, Date, Start_Time, End_Time,Lembur_Type,Description,AppByHead," +
		"Type_Approve_Head, Done_Head,Approve_Head_Date,AppByHead2,Type_Approve_Head2,Done_Head2,Approve_Head_Date2,AppByHRD,Type_Approve_HRD,Done_HRD,Approve_HRD_Date,Insert_User,Insert_Date,Update_User,Update_Date,Lembur_File FROM T_Lembur").Scan(&Lembur)

	res := L.Result{Code: 200, Data: Lembur, Message: "Success get lembur"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetLemburByID(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var lembur L.T_Lembur
	db.Where("ID = ?", ID).First(&lembur)
	res := L.Result{Code: 200, Data: lembur, Message: "Success get berita"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func UpdateLembur(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var lemburUpdate L.T_Lembur
	json.Unmarshal(payloads, &lemburUpdate)

	var lembur L.T_Berita
	db.Where("ID = ?", ID).Updates(lemburUpdate)

	res := L.Result{Code: 200, Data: lembur, Message: "Success update lembur"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteLembur(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var lembur L.T_Lembur
	db.Where("ID = ?", ID).Delete(&lembur)

	res := L.Result{Code: 200, Message: "Success delete lembur"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func ApproveLemburHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_Head: "Approved", Done_Head: "true", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveLemburHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_Head2: "Approved", Done_Head2: "true", Approve_Head_Date2: dt.Format("2006-01-02")})

	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveLemburHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_HRD: "Approved", Done_HRD: "true", Approve_HRD_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectLemburHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_Head: "Rejected", Done_Head: "false", Approve_Head_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectLemburHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_Head2: "Rejected", Done_Head2: "false", Approve_Head_Date2: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func RejectLemburHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	lemburUpdate := []L.T_Lembur{}
	json.Unmarshal(payloads, &lemburUpdate)

	lembur := []L.T_Lembur{}
	db.Where("ID = ?", ID).Updates(L.T_Lembur{Type_Approve_HRD: "Rejected", Done_HRD: "false", Approve_HRD_Date: dt.Format("2006-01-02")})
	// db.Save(&ijin)
	res := L.Result{Code: 200, Data: lembur, Message: "Success approve lembur by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
