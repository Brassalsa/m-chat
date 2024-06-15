package persister

type Persister interface {
	Get(to string, filter interface{}, decodeTo interface{}) error
	Add(to string, payload interface{}) error
	Delete(from string, filter interface{}) error
	Update(from string, filter interface{}, payload interface{}) error
}
