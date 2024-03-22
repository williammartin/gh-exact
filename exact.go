package exact

import "time"

type ExtensionManager interface {
	List() ([]Extension, error)
	Install(extensions []Extension, versioning Versioning) error
}

type ManifestStorage interface {
	Store(manifest Manifest) error
	Load() (Manifest, error)
}

type Manifest struct {
	CreatedAt  time.Time
	Extensions []Extension
}

type Extension struct {
	Name    string
	Repo    string
	Version string
}

func (e Extension) IsLocal() bool {
	// TODO: Do this at parse time into a struct
	return e.Name != "" && e.Repo == "" && e.Version == ""
}
