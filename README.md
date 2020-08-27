# Server
Code for the WebApp for DelDrone. You can access it [here](https://deldrone.iamadarshk.com).

## How to run:

### Command Line Arguments
- `-addr` (string) : HTTP Address (default ":443")
- `-authkey` (string) : Authentication key for sessions. Use 32 or 64 bytes. (default "super-secret-key")
- `-dsn` (string) : Database DSN (default "web:pass@/deldrone?parseTime=true")
- `-encryptionkey` (string) : Encryption key for sessions. Use 16, 24 or 32 bytes (default "super-secret-key")
- `-tls`: Use TLS server
- Example Usage: `-addr=":80"`

### Build
- Using [Docker](https://docker.com):
  ``` bash
  docker build . -t deldrone-server
  ```
- ``` bash
  # PORT is an environment variable holding host's port number
  docker run -p ${PORT}:443 deldrone-server
  # example: docker run -p 8080:80 deldrone-server -addr=":80" -dsn="web:pass@/deldrone?parseTime=true"
  ```
  **Note:** For running with TLS, save certificate as `cert.pem` and private key as `key.pem` in `~/tls`, attach as volume using `-v ~/tls:/build/tls` . 
  Pass `-tls` as argument. <br>
  Example: `docker run -v ~/tls:/build/tls -p ${PORT}:443 deldrone-server -tls -addr=":443"`

## Team
- [Adarsh Kumar](https://github.com/iamadarshk)
- [Abhishek Maira](https://github.com/AbhishekMaira10)
- [Aditi Garg](https://github.com/aditi-08)
- [Aseem Goyal](https://github.com/aseemgoyal200)


## Acknowledgement
- Various code snippets were based on [Alex Edward's](https://www.alexedwards.net) Let's Go book. While I only read the sample, various open source implementations based on
his book helped me in structuring and building the project. The CSS is also provided by him through his website.
