package dto

import "admin/microservice/helpers"

type Employee struct {
	EmployeeID    string         `json:"employee_id" bson:"employee_id"`
	EmployeeName  string         `json:"employee_name" bson:"employee_name"`
	EmployeePhone string         `json:"employee_phone" bson:"employee_phone"`
	UserName      string         `json:"user_name" bson:"user_name"`
	Password      string         `json:"password" bson:"password"`
	Image         string         `json:"image" bson:"image"`
	PermissionID  string         `json:"permission_id" bson:"permission_id"`
	Status        bool           `json:"status" bson:"status"`
	TimeAt        TimeAt         `json:"time_at" bson:"time_at"`
	StatusAdmin   StatusEmployee `json:"status_admin" bson:"status_admin"`
}

func NewEmployee() Employee {
	return Employee{
		EmployeeID:   helpers.GeneratedUUID(),
		EmployeeName: "",
		PermissionID: "",
		Status:       false,
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
		StatusAdmin: Opening,
	}
}
