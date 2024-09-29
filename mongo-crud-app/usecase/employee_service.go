package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/akhiltn/learn-golang/golang-mongo-crud-app/model"
	"github.com/akhiltn/learn-golang/golang-mongo-crud-app/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (srvc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmployeeID = uuid.NewString()
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusCreated)
	log.Println("employee inserted with id", insertID, emp)
}

func (srvc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	empId := mux.Vars(r)["id"]
	log.Println("employee id", empId)
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	emp, err := repo.FindEmployeeByID(empId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (srvc *EmployeeService) GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	emp, err := repo.FindAllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (srvc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	empId := mux.Vars(r)["id"]
	log.Println("employee id", empId)
	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
		return
	}
	var emp model.Employee
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = err.Error()
		return
	}
	emp.EmployeeID = empId
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	count, err := repo.UpdateEmployeeID(empId, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (srvc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	empId := mux.Vars(r)["id"]
	log.Println("employee id", empId)
	if empId == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid employee id")
		res.Error = "invalid employee id"
		return
	}
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	count, err := repo.DeleteEmployeeByID(empId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusFound)
}

func (srvc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)
	repo := repository.EmployeeRepo{MongoCollection: srvc.MongoCollection}
	count, err := repo.DeleteAllEmplyee()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error", err)
		res.Error = err.Error()
		return
	}
	res.Data = count
	w.WriteHeader(http.StatusFound)
}
