package mock

import (
	exact "github.com/williammartin/gh-exact"
)

type ExtensionManager struct {
	ListExtensionsStub   []exact.Extension
	InstallExtensionsSpy []exact.Extension
}

func (e ExtensionManager) List() ([]exact.Extension, error) {
	return e.ListExtensionsStub, nil
}

func (e ExtensionManager) Install(extensions []exact.Extension, v exact.Versioning) error {
	e.InstallExtensionsSpy = extensions
	return nil
}
