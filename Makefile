# For http1 vs http2
# Required golang.version >= 1.6
run-http1:
	GODEBUG=http2server=0 go run interface/cmd/vshttp1/main.go \
		--tls.key interface/configs/key.pem \
		--tls.cert interface/configs/cert.pem \
		--listenat localhost:8113 \
		--protocol http1

run-http2:
	GODEBUG=http2server=1 go run interface/cmd/vshttp1/main.go \
		--tls.key interface/configs/key.pem \
		--tls.cert interface/configs/cert.pem \
		--listenat localhost:8114 \
		--protocol http2

run-h2c:
	GODEBUG=http2server=0 go run interface/cmd/vshttp1/main.go \
		--listenat localhost:8115 \
		--protocol h2c
