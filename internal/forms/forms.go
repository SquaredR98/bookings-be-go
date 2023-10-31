package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Error errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string) bool {
	x := f.Get(field)
	return x != ""
}

func (f *Form) Valid() bool {
	return len(f.Error) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Error.Add(field, "This field is required")
		}
	}
}

func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Error.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Error.Add(field, "Invalid Email Address")
	}
}
