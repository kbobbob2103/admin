package dto

import "admin/microservice/helpers"

type Role struct {
	RoleID   string `json:"role_id" bson:"role_id"`
	RoleName string `json:"role_name" bson:"role_name"`
	TimeAt   TimeAt `json:"time_at" bson:"time_at"`
}

func NewRole() Role {
	return Role{
		RoleID:   helpers.GeneratedUUID(),
		RoleName: "",
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
	}
}
