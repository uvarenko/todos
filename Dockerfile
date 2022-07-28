FROM golang:1.18 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM modules as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /app/api
CMD [ "air", "-c", ".air.toml" ]