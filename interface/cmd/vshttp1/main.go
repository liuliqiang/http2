package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/liuliqiang/http2/interface/service/http-protocol-compare/h2c"
	"github.com/liuliqiang/http2/interface/service/http-protocol-compare/http1"
	"github.com/liuliqiang/http2/interface/service/http-protocol-compare/http2"
)

func main() {
	var err error
	var args struct {
		Protocol    string `default:"http2"`
		ListenAt    string `default:"localhost:8113"`
		TlsKeyFile  string `arg:"--tls.key"`
		TlsCertFile string `arg:"--tls.cert"`
	}
	arg.MustParse(&args)

	switch args.Protocol {
	case "http1":
		_, _ = fmt.Fprintf(os.Stdout, "Ready to run http1 server at: %s.\n", args.ListenAt)
		if err = http1.RunServer(&http1.Opts{
			ListenAt: args.ListenAt,
			KeyFile:  args.TlsKeyFile,
			CertFile: args.TlsCertFile,
		}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to run http1 server %v.\n", err)
			return
		}
		_, _ = fmt.Fprintf(os.Stdout, "Http1 server exit!\n")
	case "http2":
		_, _ = fmt.Fprintf(os.Stdout, "Ready to run http2 server at: %s.\n", args.ListenAt)
		if err = http2.RunServer(&http2.Opts{
			ListenAt: args.ListenAt,
			KeyFile:  args.TlsKeyFile,
			CertFile: args.TlsCertFile,
		}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to run http2 server %v.\n", err)
			return
		}
		_, _ = fmt.Fprintf(os.Stdout, "Http2 server exit!\n")
	case "h2c":
		_, _ = fmt.Fprintf(os.Stdout, "Ready to run h2c server at: %s.\n", args.ListenAt)
		if err = h2c.RunServer(&h2c.Opts{
			ListenAt: args.ListenAt,
		}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to run h2c server %v.\n", err)
			return
		}
		_, _ = fmt.Fprintf(os.Stdout, "h2c server exit!\n")

	default:
		_, _ = fmt.Fprintf(os.Stderr, "Unsupport protocol %s.\n", args.Protocol)
	}
}
