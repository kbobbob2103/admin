package dto

import "admin/microservice/helpers"

type Role struct {
	RoleID         string                          `json:"role_id" bson:"role_id"`
	RoleName       string                          `json:"role_name" bson:"role_name"`
	PermissionID   string                          `json:"permission_id" bson:"permission_id"`
	IsDeleted      bool                            `json:"is_deleted" bson:"is_deleted"`
	NavigationBars []PermissionConfigNavigationBar `json:"navigation_bars" bson:"-"`
	TimeAt         TimeAt                          `json:"-" bson:"time_at"`
}

func NewRole() Role {
	return Role{
		RoleID:       helpers.GeneratedUUID(),
		RoleName:     "",
		PermissionID: "",
		TimeAt: TimeAt{
			CreatedAt:              helpers.GetCurrentUnix(),
			HumanReadableCreatedAt: helpers.TimeNowStr(),
		},
		IsDeleted: false,
	}
}
func (r *Role) UpdateRole(new Role) {
	r.RoleName = new.RoleName
	r.TimeAt.UpdatedAt = helpers.GetCurrentUnix()
	r.TimeAt.HumanReadableUpdatedAt = helpers.TimeNowStr()
}
