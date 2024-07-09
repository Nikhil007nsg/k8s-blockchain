# This image is a microservice in golang for the Degree chaincode
FROM golang:1.19-alpine AS build

WORKDIR /go/src/github.com/basic

COPY ./ /go/src/github.com/basic

COPY packaging /go/src/github.com/basic

COPY assetTransfer.go /go/src/github.com/basic

#COPY fabcar.go /go/src/github.com/basic


COPY go.mod /go/src/github.com/basic

COPY go.sum /go/src/github.com/basic

COPY .dockerignore /go/src/github.com/basic

#COPY ./ /go/src/github.com/basic

# Build application
RUN go mod tidy
RUN go mod vendor
RUN go build -o chaincode -v .

# Production ready image
# Pass the binary to the prod image
FROM alpine:3.11 as prod
WORKDIR /app
COPY --from=build /go/src/github.com/basic/chaincode /app/chaincode

# Set CORE_CHAINCODE_ID_NAME
#ENV CORE_CHAINCODE_ID=peer0-org1

#ENV CORE_PEER_ADDRESS=peer0.org1:7051

#ENV CORE_CHAINCODE_LOGLEVEL=debug

USER 1000

CMD ./chaincode