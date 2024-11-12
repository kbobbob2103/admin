package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
	"fmt"
)

type IRoleService interface {
	CreateRoleService(data dto.Role) error
	UpdateRoleService(data dto.Role) error
	FindAllRoleService(query dto.QueryEmployee) ([]dto.Role, error)
	FindOneRoleService(id string) (dto.Role, error)
	DeleteRoleService(id string) error
}
type roleService struct {
	roleRepo          repositoty.IRoleRepo
	permissionService IPermissionService
	navigationBarRepo repositoty.INavigationRepo
}

func NewRoleService(
	roleRepo repositoty.IRoleRepo,
	permissionService IPermissionService,
	navigationBarRepo repositoty.INavigationRepo) IRoleService {
	return roleService{
		roleRepo:          roleRepo,
		permissionService: permissionService,
		navigationBarRepo: navigationBarRepo,
	}
}
func (a roleService) CreateRoleService(data dto.Role) error {
	permission := dto.NewPermission()
	if err := a.permissionService.CreatePermissionService(permission); err != nil {
		return err
	}
	data.PermissionID = permission.PermissionID
	return a.roleRepo.CreateRole(data)
}
func (a roleService) UpdateRoleService(data dto.Role) error {
	role, err := a.roleRepo.FindOneRoleByID(data.RoleID)
	if err != nil {
		return err
	}
	if err := a.permissionService.UpdatePermissionService(dto.Permission{
		PermissionID:   role.PermissionID,
		NavigationBars: data.NavigationBars,
	}); err != nil {
		return err
	}
	role.UpdateRole(data)
	return a.roleRepo.UpdateRole(role)
}
func (a roleService) FindAllRoleService(query dto.QueryEmployee) ([]dto.Role, error) {
	return a.roleRepo.FindAll([]string{}, query.Search, query.Page, query.Limit)
}
func (a roleService) FindOneRoleService(id string) (dto.Role, error) {
	role, err := a.roleRepo.FindOneRoleByID(id)
	if err != nil {

	}
	navigationBar, err := a.navigationBarRepo.FindAllNavigationBar()
	permission, err := a.permissionService.FinaOnePermission(role.PermissionID)
	permissionMapNavigation := make(map[string]dto.PermissionConfigNavigationBar)
	for _, pr := range permission.NavigationBars {
		fmt.Println(pr.NavigationBarID)
		permissionMapNavigation[pr.NavigationBarID] = pr
	}
	navigationBarMap := make([]dto.PermissionConfigNavigationBar, 0)
	for _, na := range navigationBar {
		if _, ok := permissionMapNavigation[na.NavigationBarID]; ok {
			fmt.Println("pppp")
			navigationBarMap = append(navigationBarMap, dto.PermissionConfigNavigationBar{
				NavigationBarID:   na.NavigationBarID,
				NavigationBarName: na.NavigationBarName,
				IsRead:            permissionMapNavigation[na.NavigationBarID].IsRead,
				IsCommand:         permissionMapNavigation[na.NavigationBarID].IsCommand,
			})
		} else {
			navigationBarMap = append(navigationBarMap, dto.PermissionConfigNavigationBar{
				NavigationBarID:   na.NavigationBarID,
				NavigationBarName: na.NavigationBarName,
				IsRead:            false,
				IsCommand:         false,
			})
		}
	}
	role.NavigationBars = navigationBarMap
	return role, nil
}
func (a roleService) DeleteRoleService(id string) error {
	return a.roleRepo.Delete(id)
}
