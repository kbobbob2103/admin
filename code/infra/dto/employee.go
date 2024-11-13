package dto

import "admin/microservice/helpers"

type Employee struct {
	EmployeeID     string         `json:"employee_id" bson:"employee_id"`
	EmployeeName   string         `json:"employee_name" bson:"employee_name"`
	EmployeePhone  string         `json:"employee_phone" bson:"employee_phone"`
	UserName       string         `json:"user_name" bson:"user_name"`
	Password       string         `json:"password" bson:"password"`
	Image          string         `json:"image" bson:"image"`
	RoleID         string         `json:"role_id" bson:"role_id"`
	Status         bool           `json:"status" bson:"status"`
	RefreshToken   string         `json:"refresh_token" bson:"refresh_token"`
	Token          string         `json:"token" bson:"token"`
	TimeAt         TimeAt         `json:"-" bson:"time_at"`
	StatusEmployee StatusEmployee `json:"status_employee" bson:"status_employee"`
}
type Count struct {
	Count int64 `json:"count"`
}

func NewEmployee() Employee {
	return Employee{
		EmployeeID:    helpers.GeneratedUUID(),
		EmployeeName:  "",
		EmployeePhone: "",
		UserName:      "",
		Password:      "",
		Image:         "",
		RoleID:        "",
		Status:        false,
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
		StatusEmployee: Opening,
	}
}
func (e *Employee) UpdateEmployee(new Employee) {
	e.RoleID = new.RoleID
	e.EmployeeName = new.EmployeeName
	e.EmployeePhone = new.EmployeePhone
	e.UserName = new.UserName
	e.Password = new.Password
	e.Image = new.Image
	e.TimeAt.UpdatedAt = helpers.GetCurrentUnix()
	e.TimeAt.HumanReadableUpdatedAt = helpers.TimeNowStr()
}
