package plugin

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ahmad-khatib0-org/slack/server/public/model"
	"github.com/ahmad-khatib0-org/slack/server/public/plugin/utils"
	"github.com/ahmad-khatib0-org/slack/server/public/shared/mlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSupervisor(t *testing.T) {
	for name, f := range map[string]func(*testing.T){
		"Supervisor_InvalidExecutablePath":     testSupervisorInvalidExecutablePath,
		"Supervisor_NonExistentExecutablePath": testSupervisorNonExistentExecutablePath,
		"Supervisor_StartTimeout":              testSupervisorStartTimeout,
	} {
		t.Run(name, f)
	}
}

func testSupervisorInvalidExecutablePath(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	os.WriteFile(filepath.Join(dir, "plugin.json"), []byte(`{"id": "foo", "server": {"executable": "/foo/../../backend.exe"}}`), 0o600)

	bundle := model.BundleInfoForPath(dir)
	logger := mlog.CreateConsoleTestLogger(t)
	supervisor, err := newSupervisor(bundle, nil, nil, logger, nil)
	assert.Nil(t, supervisor)
	assert.Error(t, err)
}

func testSupervisorNonExistentExecutablePath(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	os.WriteFile(filepath.Join(dir, "plugin.json"), []byte(`{"id": "foo", "server": {"executable": "thisfileshouldnotexist"}}`), 0o600)

	bundle := model.BundleInfoForPath(dir)
	logger := mlog.CreateConsoleTestLogger(t)
	supervisor, err := newSupervisor(bundle, nil, nil, logger, nil)
	require.Error(t, err)
	require.Nil(t, supervisor)
}

// If plugin development goes really wrong, let's make sure plugin activation won't block forever.
func testSupervisorStartTimeout(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	backend := filepath.Join(dir, "backend.exe")
	utils.CompileGo(t, `
		package main

		func main() {
			for {
			}
		}
	`, backend)

	os.WriteFile(filepath.Join(dir, "plugin.json"), []byte(`{"id": "foo", "server": {"executable": "backend.exe"}}`), 0o600)

	bundle := model.BundleInfoForPath(dir)
	logger := mlog.CreateConsoleTestLogger(t)
	supervisor, err := newSupervisor(bundle, nil, nil, logger, nil)
	require.Error(t, err)
	require.Nil(t, supervisor)
}
