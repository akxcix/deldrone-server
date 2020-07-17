package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form has an anonymous url field to hold form data and Errors to hold validation errors in form data
type Form struct {
	url.Values
	Errors errors
}

// New initializes a new form using provided url.Values
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks all the fields supplied as string parameters are provided to the form or not
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// MaxLength puts an error if length of field exceeds the d
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, fmt.Sprintf("Field too long. Max allowed is %d", d))
	}
}

// MinLength puts an error if length of field is less than d
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, fmt.Sprintf("Field too short. Min required is %d", d))
	}
}

// Valid returns true if the form does not have any errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
