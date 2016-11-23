# Postcode Anywhere Stub Server
## Using with Docker
### Prerequisites
### Installing Docker Container
Build the docker image
```bash
docker build -t postcode-anywhere-stub . 
```
Run the docker image
```bash
docker run -it --rm -p 127.0.0.1:9040:9040 --name postcode-anywhere-stub postcode-anywhere-stub  
```
Open in a browser ```https://localhost:9040/``` and add an exception for the ssl key.
## Local Installation
### Prerequisites
#### Install Go
```bash
brew install go
export GOPATH=$HOME/go
```
### Installation
Clone the service into a new directory
```bash
mkdir ~/go/src/github.com/thomaslorentsen/
cd ~/go/src/github.com/thomaslorentsen/
git clone git@github.com:thomaslorentsen/postcode-anywhere-stub.git
```
### Install SSl Keys
Generate the SSL keys needed to run over https
```bash
mkdir -p ~/go/bin/
cd ~/go/bin/
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```
## Building
Compile the source
```bash
go install github.com/thomaslorentsen/postcode-anywhere-stub
```
### Running Server
Run the server located in the bin folder
```bash
$GOPATH/bin/postcode-anywhere-stub
```
## Using Server
Once Post Code Anywhere has been initiated, the host can be set to use the stub server.
In the browser console run:
```javascript
pca.host = 'localhost:9040'
```
