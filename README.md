# Server
Server for DelDrone

## How to run:
For running without TLS(easier to get going):
- With Docker
    - `docker build . -t deldrone-server`
    - `docker run -p 4000:443 deldrone -tls=false`
    - browse to http://localhost:4000
- Without Docker
    - install go
    - `go mod download`
    - `go run cmd/web/* -addr=":4000" -tls=false`
    - browse to http://localhost:4000

For running with TLS:
- save certificate as `cert.pem` and private key as `key.pem` in `./tls`. To generate 
using go, execute: `go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`. Note: this command may vary depending on your system.
- With Docker
    - `docker build . -t deldrone-server`
    - `docker run -p 4000:443 deldrone`
    - browse to https://localhost:4000
- Without Docker:
    - install go
    - `go mod download`
    - `go run cmd/web/* -addr=":4000"`
    - browse to https://localhost:4000
