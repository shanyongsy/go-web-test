# 这两组命令都可以用于生成自签名证书，具体选择哪个命令取决于您的需求和工作流程。如果您需要一个单一文件包含私钥和证书，第二组命令是更简单的方法。如果您需要分开的私钥和证书文件，可以使用第一组命令。
# 无论哪种方式，最终都会生成自签名的 X.509 证书（server.crt）和相关的私钥文件（server.key），您可以在Go代码中使用它们进行TLS验证。

# 第一组命令

## 生成一个 RSA 私钥文件（server.key）
## openssl genpkey -algorithm RSA -out server.key

## 使用私钥生成证书签名请求文件（server.csr）
## openssl req -new -key server.key -out server.csr

## 使用私钥自签名证书签名请求，生成自签名证书文件（server.crt），默认有效期为 30 天，可以使用 -days 参数指定有效期
## openssl x509 -req -in server.csr -signkey server.key -out server.crt -days 365
## openssl x509 -req -in server.csr -signkey server.key -out server.crt -days 365 -extfile openssl.cnf -extensions v3_req
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout server.key -out server.crt -config openssl.cnf -extensions v3_req

# 第二组命令

## 直接生成包含自签名证书和私钥的文件。这个命令会生成一个包含私钥和自签名证书的文件，无需额外的 CSR 文件
## openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365

