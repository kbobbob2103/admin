package exceptions

type throwError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func ThrowNewException(statusCode int, msg string) throwError {
	return throwError{Message: msg, StatusCode: statusCode}
}

//var ErrBillNotFound = error.Error('ไม่พบข้อมูลบิล')
