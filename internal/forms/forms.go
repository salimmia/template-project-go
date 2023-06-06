package forms

import (
	"net/http"
	"net/url"
)

// From creates a custom form sturct and it embeds a url.Values objects
type Form struct {
	url.Values
	Errors errors
}

// Return True if there is no errors, otherwise false
func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}

// New initialize the Form sturct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}


// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool{
	x := r.Form.Get(field)
	if x == ""{
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}