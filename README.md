# Server
Code for the WebApp for DelDrone. You can access it over [here](https://deldrone.iamadarshk.com).

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

## Team
- [Adarsh Kumar](https://github.com/iamadarshk)
- [Abhishek Maira](https://github.com/AbhishekMaira10)
- Aditi Garg
- [Aseem Goyal](https://github.com/aseemgoyal200)


## Acknowledgement
- Various code snippets were based on [Alex Edward's](https://www.alexedwards.net) Let's Go book. While I only read the sample, various open source implementations based on
his book helped me in structuring and building the project. The CSS is also provided by him through his website.
