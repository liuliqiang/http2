## 运行 http/1.1

## 运行 http/2

### 生成 TLS 证书

```
root@liqiang.io# go run $GOROOT/src/crypto/tls/generate_cert.go -host localhost
root@liqiang.io# mv *.pem interface/configs
```

### 运行 Server

```

```
