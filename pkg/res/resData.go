package res

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() *FormData {
	return &FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type ErrorData struct {
	StatusCode int
	Message    string
}
