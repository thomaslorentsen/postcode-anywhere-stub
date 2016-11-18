# postcode-anywhere-stub
Postcode Anywhere Stub Server

```bash
brew install go
mkdir ~/go/src/thomaslorentsen/
export GOPATH=$HOME/go
```

```bash
cd ~/go/src/thomaslorentsen/
git clone git@github.com:thomaslorentsen/postcode-anywhere-stub.git
```

```bash
go install github.com/thomaslorentsen/postcode-anywhere-stub                                                                                                                                                                             15:21:02 
$GOPATH/bin/hello
```

```bash
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```


```javascript
pca.host = 'localhost:8080'
pca.protocol = 'http:'
```
