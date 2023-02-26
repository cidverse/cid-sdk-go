package cidsdk

import (
	"path"
	"path/filepath"
)

var joinSeparator = ""

func JoinPath(elem ...string) string {
	if joinSeparator == "/" {
		return path.Join(elem...)
	}

	return filepath.Join(elem...)
}
