package uid

import "github.com/teris-io/shortid"

func New() string {
	id, _ := shortid.Generate()
	return id
}
