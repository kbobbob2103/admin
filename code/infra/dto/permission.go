package dto

import "admin/microservice/helpers"

type Permission struct {
	PermissionID     string   `json:"permission_id" bson:"permission_id"`
	PermissionName   string   `json:"permission_name" bson:"permission_name"`
	NavigationBarIds []string `json:"navigation_bar_ids" bson:"navigation_bar_ids"`
	TimeAt           TimeAt   `json:"time_at" bson:"time_at"`
}

func NewPermission() Permission {
	return Permission{
		PermissionID:     helpers.GeneratedUUID(),
		PermissionName:   "",
		NavigationBarIds: make([]string, 0),
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
	}
}

type TimeAt struct {
	CreatedAt              int64  `json:"created_at" bson:"created_at"`
	HumanReadableCreatedAt string `json:"human_readable_created_at" bson:"human_readable_created_at"`
}
