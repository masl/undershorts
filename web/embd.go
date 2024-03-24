package embd

import (
	"embed"
	"io/fs"
)

//go:embed dist/*
var webFS embed.FS

func WebFS() (fs.FS, error) {
	f, err := fs.Sub(webFS, "dist")
	if err != nil {
		return nil, err
	}

	return f, nil
}
