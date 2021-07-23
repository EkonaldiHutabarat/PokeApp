package CutiController

import (
	DBConfig "HRIS-API/Controllers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	C "HRIS-API/Model"

	//"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateCuti(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	payloads, _ := ioutil.ReadAll(r.Body)

	var cuti = C.T_Cuti{}

	fmt.Println("Endpoint hit: Request Cuti")

	json.Unmarshal(payloads, &cuti)

	db.Create(&cuti)
	//db.Raw("INSERT INTO T_Cuti(ID,Emp_NIK,Start_Date,End_Date,Length_Days,Description,Note,AppByHead,Type_Approve_Head,Done_Head,Approve_Head_Date,AppByHead2,Type_Approve_Head2,Done_Head2,Approve_Head_Date2,AppByHRD,Type_Approve_HRD,Done_HRD,Approve_HRD_Date,Insert_User,Update_User,Insert_Date,Update_Date,Cuti_File,SisaCuti)VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)").Scan(&cuti)
	res := C.Result{Code: 200, Data: cuti, Message: "Success create cuti"}
	db.Save(&cuti)

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetCuti(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: Get data cuti")
	Cuti := []C.T_Cuti{}

	db.Raw("SELECT * FROM T_Cuti ").Scan(&Cuti)

	res := C.Result{Code: 200, Data: Cuti, Message: "Success get cuti"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetCutiByID(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var cuti C.T_Cuti
	// db.First(&cuti, ID)
	db.Where("ID = ?", ID).First(&cuti)
	res := C.Result{Code: 200, Data: cuti, Message: "succes get cuti"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func UpdateCuti(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var cutiUpdate C.T_Cuti
	json.Unmarshal(payloads, &cutiUpdate)

	var cuti C.T_Cuti
	db.Where("ID = ?", ID).Updates(cutiUpdate)
	// db.First(&cuti, ID)
	// db.Model(&cuti).Updates(cutiUpdate)

	res := C.Result{Code: 200, Data: cuti, Message: "Success update cuti"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteCuti(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	var cuti C.T_Cuti
	//db.First(&cuti, ID)
	//db.Delete(&cuti)
	db.Where("ID = ?", ID).Delete(&cuti)

	res := C.Result{Code: 200, Message: "succes delete cuti"}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func ApproveCutiHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_Head: "Approved", Done_Head: "true", Approve_Head_Date: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success approve cuti by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveCutiHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_Head2: "Approved", Done_Head2: "true", Approve_Head_Date2: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success approve cuti by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func ApproveCutiHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_HRD: "Approved", Done_HRD: "true", Approve_HRD_Date: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success approve cuti by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectCutiHead(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_Head: "Rejected", Done_Head: "false", Approve_Head_Date: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success reject cuti by Head"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RejectCutiHead2(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_Head2: "Rejected", Done_Head2: "false", Approve_Head_Date2: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success reject cuti by Head2"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
func RejectCutiHRD(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	vars := mux.Vars(r)
	ID := vars["ID"]

	payloads, _ := ioutil.ReadAll(r.Body)

	dt := time.Now()

	cutiUpdate := []C.T_Cuti{}
	json.Unmarshal(payloads, &cutiUpdate)

	cuti := []C.T_Cuti{}
	db.Where("ID = ?", ID).Updates(C.T_Cuti{Type_Approve_HRD: "Rejected", Done_HRD: "false", Approve_HRD_Date: dt.Format("2006-01-02")})

	res := C.Result{Code: 200, Data: cuti, Message: "Success reject cuti by HRD"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
