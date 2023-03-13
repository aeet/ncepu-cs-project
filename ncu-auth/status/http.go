package status

type HttpStatus struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Http200(message string, body interface{}) (int, HttpStatus) {
	return 200, HttpStatus{
		Status:  200,
		Message: message,
		Body:    body,
	}
}

func Http401(message string, body interface{}) (int, HttpStatus) {
	return 401, HttpStatus{
		Status:  401,
		Message: message,
		Body:    body,
	}
}

func Http400(message string, body interface{}) (int, HttpStatus) {
	return 400, HttpStatus{
		Status:  400,
		Message: message,
		Body:    body,
	}
}

func Http500(message string, body interface{}) (int, HttpStatus) {
	return 500, HttpStatus{
		Status:  500,
		Message: message,
		Body:    body,
	}
}
