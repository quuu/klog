package core

import (
	"github.com/stretchr/testify/assert"
	"klog/cli"
	"klog/testutil"
	"testing"
)

func TestErrorIfPathDoesNotExist(t *testing.T) {
	testutil.WithDisk(func(path string) {
		code := Execute(path+"asdf1234", []string{"create"})
		assert.Equal(t, cli.PROJECT_PATH_INVALID, code)
	})
}

func TestCreateProject(t *testing.T) {
	testutil.WithDisk(func(path string) {
		code := Execute(path, []string{"create"})
		assert.Equal(t, cli.OK, code)
	})
}