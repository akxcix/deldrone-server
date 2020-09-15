# Server
Code for the WebApp for DelDrone. You can access it [here](https://deldrone.iamadarshk.com).

## How to run:

### Command Line Arguments
- `-addr` (string) : HTTP Address (default ":443")
- `-authkey` (string) : Authentication key for sessions. Use 32 or 64 bytes. (default "super-secret-key")
- `-dsn` (string) : Database DSN (default "user:pass@/database?parseTime=true")
- `-encryptionkey` (string) : Encryption key for sessions. Use 16, 24 or 32 bytes (default "super-secret-key")
- `-tls`: Use TLS server
- Example Usage: `-addr=":80"`

### Build
Using [Docker](https://docker.com) and docker-compose:
``` bash
# docker-compose is assuming that the command line arguments are passed as environment variables
docker-compose up --build
```
The mapping for Environment Variables is displayed in the following table:

| Environment Variable | Command Line Argument | Example |
| :---: | :---: | :---: |
| TLS | -tls | `true` |
| ENCRYPTION_KEY | -encryptionkey | `"oomu4Eiv1biefoij"` |
| AUTHENTICATION_KEY | -authkey | `"ceePais6aef7eep1iefe4pah5Pahshee"` |
| DSN | -dsn | `"user:pass@/database?parseTime=true"` |
| ADDR | -addr | `4000` |

1. The `ADDR` variable refers to your host port to which docker would bind. It just also happens to match containers port to same `ADDR`.
2. In case you're using TLS, put `cert.pem` and `key.pem` in `~/tls` on host as docker-compose attaches that directory to container as a volume.

## Team
- [Adarsh Kumar](https://github.com/iamadarshk)
- [Abhishek Maira](https://github.com/AbhishekMaira10)
- [Aditi Garg](https://github.com/aditi-08)
- [Aseem Goyal](https://github.com/aseemgoyal200)

## Acknowledgement
- Various code snippets were based on [Alex Edward's](https://www.alexedwards.net) Let's Go book. While I only read the sample, various open source implementations based on
his book helped me in structuring and building the project. The CSS is also provided by him through his website.
