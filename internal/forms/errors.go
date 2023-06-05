package forms

type errors map[string][]string

// Add adds an error for the given form field
func (e errors) Add(feild, message string) {
	e[feild] = append(e[feild], message)
}

// Get returns the error for the given form field
func (e errors) Get(feild string) string{
	es := e[feild]
	
	if len(es) == 0 {
        return ""
    }

    return es[0]
}