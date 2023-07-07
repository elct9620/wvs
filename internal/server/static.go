//go:build !embed
// +build !embed

package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/evanw/esbuild/pkg/api"
	"go.uber.org/zap"
)

func mustNewAssetsContext(logger *zap.Logger) api.BuildContext {
	ctx, err := api.Context(api.BuildOptions{
		EntryPoints: []string{
			ScriptDir + "/app.ts",
		},
		Define: map[string]string{
			"DEV": "true",
		},
		Bundle: true,
		Outdir: StaticDir + "/js",
	})
	if err != nil {
		for _, msg := range err.Errors {
			logger.Error(msg.Text)
		}
		logger.Fatal("unable to setup esbuild")
	}

	return ctx
}

func mustWatchAssets(ctx api.BuildContext, logger *zap.Logger) {
	err := ctx.Watch(api.WatchOptions{})
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func mustServeAssets(ctx api.BuildContext, logger *zap.Logger) *api.ServeResult {
	res, err := ctx.Serve(api.ServeOptions{
		Servedir: StaticDir,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &res
}

func mustCreateProxy(serverAddr string, logger *zap.Logger) *httputil.ReverseProxy {
	proxyURL, err := url.Parse(serverAddr)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetXForwarded()
			r.SetURL(proxyURL)
		},
	}

}

func WithRoot(sessions SessionStore, logger *zap.Logger) HTTPOptionFn {
	ctx := mustNewAssetsContext(logger)
	server := mustServeAssets(ctx, logger)
	mustWatchAssets(ctx, logger)

	proxy := mustCreateProxy(fmt.Sprintf("http://%s:%d", server.Host, server.Port), logger)

	return func(mux *http.ServeMux) {
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			http.SetCookie(w, sessions.Renew(req))
			proxy.ServeHTTP(w, req)
		})
	}
}
