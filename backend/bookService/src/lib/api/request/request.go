package request

type Response struct {
	Status string `json:"status"` //ok or error
	Error  string `json:"error,omitempty"`
}

type StatusLikes struct {
	Likes    uint     `json:"likes"`
	DisLikes uint     `json:"dislikes"`
	Response Response `json:"response"`
}

const (
	StatusOk    = "OK"
	StatusError = "ERROR"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func Error(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}
