package constant

import "net/http"

type Exception struct {
	Message    string
	StatusCode int
}

var (
	BadRequest          = Exception{"Bad Request", http.StatusBadRequest}
	Unauthorized        = Exception{"Unauthorized", http.StatusUnauthorized}
	LoginFailed         = Exception{"Login Failed", http.StatusBadRequest}
	UnprocessableEntity = Exception{"Unprocessable Entity", http.StatusUnprocessableEntity}
)
