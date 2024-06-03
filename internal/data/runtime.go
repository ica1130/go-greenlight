package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// strconv.Quote() function will wrap the string in double quotes,
	// making it a valid JSON string.
	quotedValue := strconv.Quote(jsonValue)

	return []byte(quotedValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSON, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSON, " ")
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
