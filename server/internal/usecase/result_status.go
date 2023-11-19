package usecase

type ResultStatus struct {
	StatusCode int
	Err        error
}

func NewResultStatus(code int, err error) *ResultStatus {
	return &ResultStatus{
		StatusCode: code,
		Err:        err,
	}
}
