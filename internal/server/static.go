//go:build !embed
// +build !embed

package server

import (
	"net/http"

	esbuildfs "github.com/elct9620/esbuild-fs"
	"github.com/evanw/esbuild/pkg/api"
	"go.uber.org/zap"
)

var BuildOptions = api.BuildOptions{
	EntryPoints: []string{
		"assets/app.ts",
	},
	Define: map[string]string{
		"DEV": "true",
	},
	Bundle: true,
	Outdir: "dist",
}

func WithAssets(logger *zap.Logger) HTTPOptionFn {
	assets, sse, err := esbuildfs.Serve(BuildOptions, esbuildfs.WithHandlerPrefix("assets"))

	if err != nil {
		logger.Fatal(err.Error())
	}

	return func(mux *http.ServeMux) {
		mux.Handle("/esbuild", sse)
		mux.Handle("/assets/", assets)
	}
}
