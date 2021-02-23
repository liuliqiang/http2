package h2c

import (
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Opts struct {
	ListenAt string
}

func RunServer(opts *Opts) (err error) {
	h2s := &http2.Server{}

	server := &http.Server{
		Addr: opts.ListenAt,
		Handler: h2c.NewHandler(
			http.FileServer(
				http.Dir("interface/assets"),
			),
			h2s,
		),
	}

	if err = server.ListenAndServe(); err != nil {
		return errors.Wrap(err, "listen and serve")
	}
	return nil
}
