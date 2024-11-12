package memory

import "admin/microservice/infra/dto"

func PaginateEmployee(employee []dto.Employee, page, limit int) ([]dto.Employee, error) {
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// ตรวจสอบว่า startIndex และ endIndex ไม่เกินความยาวของ slice
	if startIndex >= len(employee) {
		return []dto.Employee{}, nil
	}
	if endIndex > len(employee) {
		endIndex = len(employee)
	}

	// ดึงข้อมูลจาก slice ตามหน้าและ Limit ที่กำหนด

	return employee[startIndex:endIndex], nil
}
func PaginateRole(roles []dto.Role, page, limit int) ([]dto.Role, error) {
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// ตรวจสอบว่า startIndex และ endIndex ไม่เกินความยาวของ slice
	if startIndex >= len(roles) {
		return []dto.Role{}, nil
	}
	if endIndex > len(roles) {
		endIndex = len(roles)
	}

	// ดึงข้อมูลจาก slice ตามหน้าและ Limit ที่กำหนด

	return roles[startIndex:endIndex], nil
}
