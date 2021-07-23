package UserController

import (
	"encoding/json"
	"fmt"
	"net/http"

	M "HRIS-API/Model"

	//"github.com/jinzhu/gorm"

	DBConfig "HRIS-API/Controllers"

	"gorm.io/gorm"
)

var db *gorm.DB

var err error

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	fmt.Println("Endpoint hit: get Data User")
	datausers := []M.M_Employee{}

	db.Raw("SELECT * FROM M_Employee").Scan(&datausers)
	res := M.Result{Code: 200, Data: datausers, Message: "Get Data Success"}
	results, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db, err = DBConfig.Config()
	//vars := mux.Vars(r)
	emp_nik := r.FormValue("Emp_NIK")

	emp_nickname := r.FormValue("Emp_NickName")
	//password := r.FormValue("password")

	// PasswordHash := md5.New()
	// io.WriteString(PasswordHash, password)
	// fmt.Printf("%x", PasswordHash.Sum(nil))

	var user M.M_Employee
	db.Find(&user, "Emp_NIK = ? AND Emp_NickName = ?", emp_nik, emp_nickname)

	res := M.Result{Code: 200, Data: user, Message: "Login Success"}

	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}
