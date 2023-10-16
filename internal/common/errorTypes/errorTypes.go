package errorTypes

import "errors"

var (
	ErrDBDataReception = errors.New("error while trying getting data from db")
)
