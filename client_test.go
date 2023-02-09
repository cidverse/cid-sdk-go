package cidsdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSDKHealth(t *testing.T) {
	if len(os.Getenv("CID_API_ADDR")) == 0 {
		t.Skip("requires local api")
	}

	// sdk
	sdk, sdkErr := NewSDK()
	assert.NoError(t, sdkErr)

	// test
	resp, respErr := sdk.Health()
	assert.NoError(t, respErr)
	assert.Equal(t, "up", resp.Status)
}
