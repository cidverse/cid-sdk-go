package cidsdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinPathForTestCases(t *testing.T) {
	joinSeparator = "/"
	assert.Equal(t, `my-dir/my-file`, JoinPath("my-dir", "my-file"))
}

func TestJoinPath(t *testing.T) {
	joinSeparator = ""
	assert.Equal(t, fmt.Sprintf(`my-dir%cmy-file`, os.PathSeparator), JoinPath("my-dir", "my-file"))
}
