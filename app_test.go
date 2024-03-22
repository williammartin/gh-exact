package exact_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	exact "github.com/williammartin/gh-exact"
	"github.com/williammartin/gh-exact/mock"
)

func TestAppDumpAndRestore(t *testing.T) {
	// Given we have a list of installed extensions
	mockExtensionManager := mock.ExtensionManager{
		ListExtensionsStub: []exact.Extension{
			{Name: "foo", Repo: "bar", Version: "v1.0.0"},
			{Name: "baz", Repo: "qux", Version: "v2.0.0"},
		},
	}

	app := exact.App{
		ExtensionManager: &mockExtensionManager,
		ManifestStorage:  &mock.ManifestStorage{},
	}

	// When we dump the extensions
	require.NoError(t, app.Dump())

	// And restore them
	require.NoError(t, app.Restore(exact.PinVersions))

	// Then the listed extensions should be installed
	require.Equal(
		t,
		mockExtensionManager.ListExtensionsStub,
		mockExtensionManager.InstallExtensionsSpy,
	)
}
