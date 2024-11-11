package dto

import "admin/microservice/helpers"

type Permission struct {
	PermissionID   string `json:"permission_id" bson:"permission_id"`
	PermissionName string `json:"permission_name" bson:"permission_name"`
}

func NewPermission() Permission {
	return Permission{
		PermissionID:   helpers.GeneratedUUID(),
		PermissionName: "",
	}
}
