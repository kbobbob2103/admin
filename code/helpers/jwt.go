package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type SignedDetail struct {
	EmployeeId   string
	EmployeeName string
	UserName     string
	RoleID       string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(employeeID, employeeName, userName, roleID string) (string, string, error) {
	claims := &SignedDetail{
		EmployeeId:   employeeID,
		EmployeeName: employeeName,
		UserName:     userName,
		RoleID:       roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	refreshClaims := SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(15 * 24 * time.Hour).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {

	}
	return token, refreshToken, err
}
