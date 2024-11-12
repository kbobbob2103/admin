package dto

import "admin/microservice/helpers"

type Permission struct {
	PermissionID   string                          `json:"permission_id" bson:"permission_id"`
	NavigationBars []PermissionConfigNavigationBar `json:"navigation_bars" bson:"navigation_bars"`
	TimeAt         TimeAt                          `json:"-" bson:"time_at"`
}
type PermissionConfigNavigationBar struct {
	NavigationBarID   string `json:"navigation_bar_id" bson:"navigation_bar_id"`
	NavigationBarName string `json:"navigation_bar_name" bson:"-"`
	IsRead            bool   `json:"is_read" bson:"is_read"`
	IsCommand         bool   `json:"is_command" json:"is_command"`
}

func NewPermission() Permission {
	return Permission{
		PermissionID:   helpers.GeneratedUUID(),
		NavigationBars: []PermissionConfigNavigationBar{},
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
	}
}
func (r *Permission) UpdatePermission(new Permission) {
	r.NavigationBars = new.NavigationBars
	r.TimeAt.UpdatedAt = helpers.GetCurrentUnix()
	r.TimeAt.HumanReadableUpdatedAt = helpers.TimeNowStr()
}

type TimeAt struct {
	CreatedAt              int64  `json:"created_at" bson:"created_at"`
	HumanReadableCreatedAt string `json:"human_readable_created_at" bson:"human_readable_created_at"`
	UpdatedAt              int64  `json:"updated_at" bson:"updated_at"`
	HumanReadableUpdatedAt string `json:"human_readable_updated_at" bson:"human_readable_updated_at"`
}
