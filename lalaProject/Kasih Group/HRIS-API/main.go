package main

import (
	"log"
	"net/http"

	DBConfig "HRIS-API/Controllers"
	Berita "HRIS-API/Controllers/BeritaController"
	Cuti "HRIS-API/Controllers/CutiController"
	Ijin "HRIS-API/Controllers/IjinController"
	Lembur "HRIS-API/Controllers/LemburController"
	Master "HRIS-API/Controllers/MasterTypeController"
	User "HRIS-API/Controllers/UserController"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

func main() {
	db, err = DBConfig.Config()
	handleRequest()
}

func handleRequest() {
	log.Println("localhost:8080")
	Router := mux.NewRouter().StrictSlash(true)

	//Router User
	Router.HandleFunc("/api/loginUser", User.LoginUser).Methods("POST")
	Router.HandleFunc("/api/getUser", User.GetUserHandler).Methods("GET")

	//Router Master Type
	Router.HandleFunc("/api/getCutiType", Master.GetCutiType).Methods("GET")
	Router.HandleFunc("/api/getLemburType", Master.GetLemburType).Methods("GET")
	Router.HandleFunc("/api/getIjinType", Master.GetIjinType).Methods("GET")

	//Router Berita
	Router.HandleFunc("/api/createBerita", Berita.CreateBerita).Methods("POST")
	Router.HandleFunc("/api/getBerita", Berita.GetDataBerita).Methods("GET")
	Router.HandleFunc("/api/getBeritaByID/{ID}", Berita.GetBeritaByID).Methods("GET")
	Router.HandleFunc("/api/deleteBerita/{ID}", Berita.DeleteBerita).Methods("DELETE")
	Router.HandleFunc("/api/updateBerita/{ID}", Berita.UpdateBerita).Methods("PUT")

	//Router Cuti
	Router.HandleFunc("/api/createCuti", Cuti.CreateCuti).Methods("POST")
	Router.HandleFunc("/api/getCuti", Cuti.GetCuti).Methods("GET")
	Router.HandleFunc("/api/getCutiByID/{ID}", Cuti.GetCutiByID).Methods("GET")
	Router.HandleFunc("/api/deleteCuti/{ID}", Cuti.DeleteCuti).Methods("DELETE")
	Router.HandleFunc("/api/updateCuti/{ID}", Cuti.UpdateCuti).Methods("PUT")
	Router.HandleFunc("/api/approveCutiHead/{ID}", Cuti.ApproveCutiHead).Methods("PUT")
	Router.HandleFunc("/api/approveCutiHead2/{ID}", Cuti.ApproveCutiHead2).Methods("PUT")
	Router.HandleFunc("/api/approveCutiHRD/{ID}", Cuti.ApproveCutiHRD).Methods("PUT")
	Router.HandleFunc("/api/rejectCutiHead/{ID}", Cuti.RejectCutiHead).Methods("PUT")
	Router.HandleFunc("/api/rejectCutiHead2/{ID}", Cuti.RejectCutiHead2).Methods("PUT")
	Router.HandleFunc("/api/rejectCutiHRD/{ID}", Cuti.RejectCutiHRD).Methods("PUT")

	//Router Lembur
	Router.HandleFunc("/api/createLembur", Lembur.CreateLembur).Methods("POST")
	Router.HandleFunc("/api/getLembur", Lembur.GetLembur).Methods("GET")
	Router.HandleFunc("/api/getLemburByID/{ID}", Lembur.GetLemburByID).Methods("GET")
	Router.HandleFunc("/api/deleteLembur/{ID}", Lembur.DeleteLembur).Methods("DELETE")
	Router.HandleFunc("/api/updateLembur/{ID}", Lembur.UpdateLembur).Methods("PUT")
	Router.HandleFunc("/api/approveLemburHead/{ID}", Lembur.ApproveLemburHead).Methods("PUT")
	Router.HandleFunc("/api/approveLemburHead2/{ID}", Lembur.ApproveLemburHead2).Methods("PUT")
	Router.HandleFunc("/api/approveLemburHRD/{ID}", Lembur.ApproveLemburHRD).Methods("PUT")
	Router.HandleFunc("/api/rejectLemburHead/{ID}", Lembur.RejectLemburHead).Methods("PUT")
	Router.HandleFunc("/api/rejectLemburHead2/{ID}", Lembur.RejectLemburHead2).Methods("PUT")
	Router.HandleFunc("/api/rejectLemburHRD/{ID}", Lembur.RejectLemburHRD).Methods("PUT")

	//Router Ijin
	Router.HandleFunc("/api/createIjin", Ijin.CreateIjin).Methods("POST")
	Router.HandleFunc("/api/getIjin", Ijin.GetIjin).Methods("GET")
	Router.HandleFunc("/api/getIjinByID/{ID}", Ijin.GetIjinByID).Methods("GET")
	Router.HandleFunc("/api/deleteIjin/{ID}", Ijin.DeleteIjin).Methods("DELETE")
	Router.HandleFunc("/api/updateIjin/{ID}", Ijin.UpdateIjin).Methods("PUT")
	Router.HandleFunc("/api/approveIjinHead/{ID}", Ijin.ApproveIjinHead).Methods("PUT")
	Router.HandleFunc("/api/approveIjinHead2/{ID}", Ijin.ApproveIjinHead2).Methods("PUT")
	Router.HandleFunc("/api/approveIjinHRD/{ID}", Ijin.ApproveIjinHRD).Methods("PUT")
	Router.HandleFunc("/api/rejectIjinHead/{ID}", Ijin.RejectIjinHead).Methods("PUT")
	Router.HandleFunc("/api/rejectIjinHead2/{ID}", Ijin.ApproveIjinHead2).Methods("PUT")
	Router.HandleFunc("/api/rejectIjinHRD/{ID}", Ijin.RejectIjinHRD).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", Router))
}
