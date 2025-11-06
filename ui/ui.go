package ui

import (
	"embed"
	"io/fs"
)

//go:embed v1/dist/* v1/dist/assets/*
var managementUI embed.FS
var ManagementUI fs.FS

func init() {
	var err error
	ManagementUI, err = fs.Sub(managementUI, "v1/dist")
	if err != nil {
		panic(err)
	}
}
