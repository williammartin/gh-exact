package exact

import (
	"fmt"
	"time"
)

type Versioning string

const (
	PinVersions    Versioning = "pin"
	LatestVersions Versioning = "latest"
)

type App struct {
	ExtensionManager ExtensionManager
	ManifestStorage  ManifestStorage
}

func (a App) Dump() error {
	extensions, err := a.ExtensionManager.List()
	if err != nil {
		return fmt.Errorf("dump: listing extensions: %v", err)
	}

	m := Manifest{
		CreatedAt:  time.Now(),
		Extensions: extensions,
	}

	if err := a.ManifestStorage.Store(m); err != nil {
		return fmt.Errorf("dump: storing manifest: %v", err)
	}

	return nil
}

func (a App) Restore(v Versioning) error {
	m, err := a.ManifestStorage.Load()
	if err != nil {
		return fmt.Errorf("restore: loading manifest: %v", err)
	}

	if err := a.ExtensionManager.Install(m.Extensions, v); err != nil {
		return fmt.Errorf("restore: installing extensions: %v", err)
	}

	return nil
}
