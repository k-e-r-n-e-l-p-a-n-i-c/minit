FROM golang:latest
WORKDIR /usr/src/minit
COPY go.mod ./
RUN go mod download && go mod verify
COPY main.go .
RUN go build -v -o /usr/local/bin/minit ./...
CMD ["minit","sleep","10000"]

