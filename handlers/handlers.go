package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	cn "webapiingo/connection"
	wrk "webapiingo/worker"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var myDB *gorm.DB

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	var MyWorker wrk.WorkerInfo
	myDB = cn.DataMigration()

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("cache-control", "no-cache")
	w.Header().Add("Content-Type", "application/json")

	if errInDecode := json.NewDecoder(r.Body).Decode(&MyWorker); errInDecode != nil {
		fmt.Fprintln(w, errInDecode)
		log.Println("errInDecode", errInDecode)
	}
	fmt.Println("Myworker", MyWorker)
	if err := myDB.Create(&MyWorker).Error; err != nil {
		//fmt.Println("Not Created :", myDB.Error)
		fmt.Fprintln(w, "Not Created !")
		log.Fatalln("Err:", myDB.Error)
	}

	//table name is worker_infos;

	if err := json.NewEncoder(w).Encode(MyWorker); err != nil {
		fmt.Fprintln(w, "Data not saved ")
		log.Fatalln("err in Encoding:", err)
	}
	fmt.Fprintln(w, "Data has been save successful.")
	defer r.Body.Close()
}

func GetWorker(w http.ResponseWriter, r *http.Request) {
	var workers []wrk.WorkerInfo
	myDB = cn.DataMigration()
	w.Header().Add("Content-Type", "application/json")
	if err := myDB.Find(&workers).Error; err != nil {
		fmt.Println("Not find data in database :", myDB.Error)
		fmt.Fprintln(w, "Data not found ! ")
		log.Fatalln("Err:", myDB.Error)
	}

	if errInEnCode := json.NewEncoder(w).Encode(workers); errInEnCode != nil {
		fmt.Fprintln(w, "Data not fatch ")
		log.Fatalln("err in Encoding:", errInEnCode)
	}
	fmt.Fprintln(w, "Data fetched successful.")
	defer r.Body.Close()

}

func GetWorkerByID(w http.ResponseWriter, r *http.Request) {
	var worker wrk.WorkerInfo
	myDB = cn.DataMigration()
	w.Header().Add("Content-Type", "application/json")
	//myDB.First(&worker,r.URL.Query().Get("id"))			// we can use by url or gorm
	if err := myDB.First(&worker, mux.Vars(r)["id"]).Error; err != nil {
		fmt.Println("Not find data :", myDB.Error)
		fmt.Fprintln(w, "Data not found ! ")
		log.Println("Err:", myDB.Error)
		return
	}

	if errInEnCode := json.NewEncoder(w).Encode(worker); errInEnCode != nil {
		fmt.Fprintln(w, "Data not fatch ")
		log.Fatalln("err in Encoding:", errInEnCode)
	}
	fmt.Fprintln(w, "Data has been fetched.")
	defer r.Body.Close()
}

func UpdateWorker(w http.ResponseWriter, r *http.Request) {
	var worker wrk.WorkerInfo
	myDB = cn.DataMigration()
	w.Header().Add("Content-Type", "application/json")
	//myDB.First(&worker, r.URL.Query().Get("id"))
	if err := myDB.First(&worker, mux.Vars(r)["id"]).Error; err != nil {
		fmt.Fprintln(w, "Data is not there !")
		log.Println("err in find data", err)
		return
	}
	fmt.Println("Find out for Updation", worker)
	if errInDecode := json.NewDecoder(r.Body).Decode(&worker); errInDecode != nil {
		fmt.Fprintln(w, errInDecode)
		log.Println("errInDecode", errInDecode)
	}
	myDB.Save(&worker)
	if errInEnCode := json.NewEncoder(w).Encode(worker); errInEnCode != nil {
		fmt.Fprintln(w, "data not fatch ")
		log.Fatalln("err in Encoding:", errInEnCode)
	}

	fmt.Fprintln(w, "data has been updated.")
	defer r.Body.Close()
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	var worker wrk.WorkerInfo
	myDB = cn.DataMigration()
	w.Header().Add("Content-Type", "application/json")

	//myDB.Delete(&worker, r.URL.Query().Get("id"))
	if errDel := myDB.Delete(&worker, mux.Vars(r)["id"]).Error; errDel != nil {
		fmt.Println("Not find data in database :", myDB.Error)
		fmt.Fprintln(w, "Data not found ! ")
		log.Fatalln("Err:", myDB.Error)
	}

	if errInEnCode := json.NewEncoder(w).Encode(("worker data deleted !!")); errInEnCode != nil {
		fmt.Fprintln(w, "data not delete ")
		log.Fatalln("err in Encoding:", errInEnCode)
	}

	fmt.Fprintln(w, http.StatusOK)
	defer r.Body.Close()
}
