//go:build !embed
// +build !embed

package server

import (
	"net/http"

	esbuildfs "github.com/elct9620/esbuild-fs"
	"github.com/evanw/esbuild/pkg/api"
	"go.uber.org/zap"
)

func WithAssets(logger *zap.Logger) HTTPOptionFn {
	assets, sse, err := esbuildfs.Serve(api.BuildOptions{
		EntryPoints: []string{
			ScriptDir + "/app.ts",
		},
		Define: map[string]string{
			"DEV": "true",
		},
		Bundle: true,
		Outdir: StaticDir + "/js",
	}, esbuildfs.WithHandlerPrefix("assets"))

	if err != nil {
		logger.Fatal(err.Error())
	}

	return func(mux *http.ServeMux) {
		mux.Handle("/esbuild", sse)
		mux.Handle("/assets/", assets)
	}
}
