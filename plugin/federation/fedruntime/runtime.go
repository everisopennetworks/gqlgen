package fedruntime

// Service is the service object that the
// generated.go file will return for the _service
// query
type Federation_Service struct {
	SDL string `json:"sdl"`
}

// Everything with a @key implements this
type Federation_Entity interface {
	IsEntity()
}

// Used for the Link directive
type Link any
