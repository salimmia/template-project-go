package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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

// Required checks for required fields
func (f *Form) Required(feilds ...string){
	for _, feild := range feilds{
		value := f.Get(feild)

		if strings.TrimSpace(value) == ""{
			f.Errors.Add(feild, "This field cannot be blank")
		}
	}
}


// Has checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool{
	x := r.Form.Get(field)
	
	return x != ""
}

// MinLength checks for strings minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool{
	x := f.Errors.Get(field)

	if len(x) < length{
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}