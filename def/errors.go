package def

import "errors"

var (
	ErrInvalidDataAsPerSchema = errors.New("Invalid data as per schema")
	ErrInvalidReference       = errors.New("Invalid Reference")
)
