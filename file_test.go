package cidsdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	file := NewFile("/project/project.sbom.json")

	assert.Equal(t, "/project/project.sbom.json", file.Path)
	assert.Equal(t, "project.sbom.json", file.Name)
	assert.Equal(t, "project", file.NameShort)
	assert.Equal(t, ".sbom.json", file.Extension)
}
