package cidsdk

import (
	"github.com/google/uuid"
)

// UUID request
func (sdk SDK) UUID() string {
	return uuid.NewString()
}
