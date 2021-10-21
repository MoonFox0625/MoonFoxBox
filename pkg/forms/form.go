// Package forms
// Date:  2021/10/21 22:18
// Desc:
package forms

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form Implement a Required method to check that specific fields in the form
// data are present and not blank. If any fields fail this check, add the
// appropriate message to the form errors.
type Form struct {
	url.Values
	Errors errors
}

// New Define a New function to initialize a custom Form struct. Notice that
// this takes the form data as the parameter?
func New(data url.Values) *Form {
	return &Form{
		data,
		make(map[string][]string),
	}
}

// Required Implement a Required method to check that specific fields in the form
// data are present and not blank. If any fields fail this check, add the
// appropriate message to the form errors.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		data := f.Get(field)
		if strings.TrimSpace(data) == "" {
			f.Errors.Add(field, "The field can't be blank")
		}
	}
}

// MaxLength Implement a MaxLength method to check that a specific field in the form
// contains a maximum number of characters. If the check fails then add the
// appropriate message to the form errors.
func (f *Form) MaxLength(field string) {
	data := f.Get(field)
	if len(data) == 0 {
		return
	}

	if utf8.RuneCountInString(data) > 100 {
		f.Errors.Add(field, "The field is too long (maximum is 100 characters)")
	}
}

// PermittedValues Implement a PermittedValues method to check that a specific field in the form
// matches one of a set of specific permitted values. If the check fails
// then add the appropriate message to the form errors.
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	f.Errors.Add(field, "This field is invalid")
}

// Valid Implement a Valid method which returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
