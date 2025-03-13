package cidsdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EnvOverwriteStruct struct {
	Key   string `env:"KEY"`
	Value string `env:"VALUE"`
}

func TestEnvOverwrite(t *testing.T) {
	os.Clearenv()
	val := EnvOverwriteStruct{Key: "hello", Value: "world"}
	PopulateFromEnv(&val, map[string]string{
		"KEY":   "hi",
		"VALUE": "mom",
	})

	assert.Equal(t, "hi", val.Key)
	assert.Equal(t, "mom", val.Value)
}
