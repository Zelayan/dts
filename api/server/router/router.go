package router

import (
	"github.com/Zelayan/dts/api/server/router/query"
	"github.com/Zelayan/dts/cmd/colletcor/options"
)

type Register func(opts *options.Options)

func InstallRouter(opts *options.Options) {
	fs := []Register{
		query.NewRouter,
	}
	install(opts, fs...)
}

func install(opts *options.Options, fs ...Register) {
	for _, f := range fs {
		f(opts)
	}
}
