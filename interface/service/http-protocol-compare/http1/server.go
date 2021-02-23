package http1

import (
	"net/http"

	"github.com/pkg/errors"
)

type Opts struct {
	ListenAt string
	KeyFile  string
	CertFile string
}

func RunServer(opts *Opts) (err error) {
	http.Handle("/", http.FileServer(http.Dir("interface/assets")))

	if err = http.ListenAndServeTLS(opts.ListenAt, opts.CertFile, opts.KeyFile, nil); err != nil {
		return errors.Wrap(err, "listen and server")
	}
	return nil
}
