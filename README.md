# Server
Code for the WebApp for DelDrone available [here](https://deldrone.iamadarshk.com).

## How to run:
Without TLS:
- ```
  docker build . -t deldrone-server
  docker run -p 4000:443 deldrone-server
  ```
- Browse to http://localhost:4000

For running with TLS:
- Save certificate as `cert.pem` and private key as `key.pem` in `./tls`. 
- ```
  docker build . -t deldrone-server
  docker run -p 4000:443 deldrone-server -tls
  ```
- Browse to https://localhost:4000
