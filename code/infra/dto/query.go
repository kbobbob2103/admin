package dto

type QueryEmployee struct {
	RoleID string `form:"role_id"`
	Search string `form:"search"`
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
}
