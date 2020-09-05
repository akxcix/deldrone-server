FROM golang:alpine  

WORKDIR /build

ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o main cmd/web/*


ENV DSN "web:pass@/deldrone?parseTime=true"
ENV ADDR 443
ENV TLS false
ENV AUTHENTICATION_KEY "super-secret-key"
ENV ENCRYPTION_KEY "super-secret-key"

CMD /build/main -tls=${TLS} -authkey=${AUTHENTICATION_KEY} -encryptionkey=${ENCRYPTION_KEY} -addr=":${ADDR}" -dsn=${DSN}