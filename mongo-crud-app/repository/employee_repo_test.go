package repository

import (
	"context"
	"github.com/akhiltn/learn-golang/golang-mongo-crud-app/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"testing"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://tnakhil:p8FztRhRwEHdOiwy@cluster0.tyi6ris.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))
	if err != nil {
		log.Fatal("error while connecting mongodb", err)
	}
	log.Println("mongodb successfully connected.")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("ping success")
	return mongoTestClient
}

func TestMongoOperation(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())
	emp1 := uuid.New().String()
	// emp2 := uuid.New().String()
	coll := mongoTestClient.Database("companydb").Collection("employee_test")
	empRepo := EmployeeRepo{MongoCollection: coll}
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Tony Stark",
			Department: "Physics",
			EmployeeID: emp1,
		}
		result, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatal("insert 1 operation failed", err)
		}
		t.Log("insert 1 operation failed", result)
	})
	t.Run("Get Employee 1", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("emp 1", result.Name)
	})
	t.Run("Get all employee", func(t *testing.T) {
		result, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("emplyees", result)
	})
	t.Run("Update Employee name", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Tony Stark aka Iron Man",
			Department: "Physics",
			EmployeeID: emp1,
		}
		result, err := empRepo.UpdateEmployeeID(emp1, &emp)
		if err != nil {
			log.Fatal("upate operation failed", err)
		}
		t.Log("update count", result)
	})
	t.Run("Get Employee 1 After update", func(t *testing.T) {
		result, err := empRepo.FindEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("emp 1", result.Name)
	})
	t.Run("Delete Employee 1", func(t *testing.T) {
		result, err := empRepo.DeleteEmployeeByID(emp1)
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("emplyees", result)
	})
	t.Run("Get All Employee After delete", func(t *testing.T) {
		results, err := empRepo.FindAllEmployee()
		if err != nil {
			t.Fatal("get operation failed", err)
		}
		t.Log("employees", results)
	})
	t.Run("Delete all employees for clean up", func(t *testing.T) {
		result, err := empRepo.DeleteAllEmplyee()
		if err != nil {
			log.Fatal("delete operation failed", err)
		}
		t.Log("deleted count", result)
	})

}
