package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

// if a receiver is a value, method calls can be invoked on pointers and values
// however pointer methods can only be invoked on pointers.
// in order for this function to be implemented when converting
// function singature must satisfy the json.Marshaler interface.
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	// strconv.Quote() function will wrap the string in double quotes,
	// making it a valid JSON string.
	quotedValue := strconv.Quote(jsonValue)

	return []byte(quotedValue), nil
}
