package model

type Employee struct {
	EmployeeID string `json:"employee_id" bson:"employee_id"`
	Name       string `json:"name" bson:"name"`
	Department string `json:"department" bson:"department"`
}
