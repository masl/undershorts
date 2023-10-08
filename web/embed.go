package web

import (
	"embed"
	"io/fs"
)

//go:embed static
var static embed.FS
var Static, _ = fs.Sub(static, "static")

//go:embed index.html
var Index embed.FS
