package dtos

type Healthcheck struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewhealthCheck(status int, message string) Healthcheck {
	return Healthcheck{
		Status:  status,
		Message: message,
	}
}
