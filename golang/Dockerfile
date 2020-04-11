FROM golang:1.14.2

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

# copy code into container
COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 5000

# command to run when starting the continner
CMD ["/dist/main"]
