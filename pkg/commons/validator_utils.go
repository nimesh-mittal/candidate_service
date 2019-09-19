package commons

import (
	"net/url"
	"sort"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func Validate(entity interface{}) (bool, error) {
	v := validator.New()
	err := v.Struct(entity)

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return false, err
	}

	return err == nil, err
}

func ToString(value url.Values) string {

	if value == nil {
		return Empty
	}
	var buf strings.Builder
	keys := make([]string, 0, len(value))
	for k := range value {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := value[k]
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte(';')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()

}
