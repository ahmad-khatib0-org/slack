package model

import "github.com/ahmad-khatib0-org/slack/server/public/shared/mlog"

type BundleInfo struct {
	Path string

	Manifest      *Manifest
	ManifestPath  string
	ManifestError error
}

func (b *BundleInfo) WrapLogger(logger *mlog.Logger) *mlog.Logger {
	if b.Manifest != nil {
		return logger.With(mlog.String("plugin_id", b.Manifest.Id))
	}
	return logger.With(mlog.String("plugin_path", b.Path))
}

// Returns bundle info for the given path. The return value is never nil.
func BundleInfoForPath(path string) *BundleInfo {
	m, mpath, err := FindManifest(path)
	return &BundleInfo{
		Path:          path,
		Manifest:      m,
		ManifestPath:  mpath,
		ManifestError: err,
	}
}
