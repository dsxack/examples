# build stage
FROM registry.dsxack.com/dev-go-build:latest AS build-env

WORKDIR /go/src/app
RUN go get -u github.com/golang/dep/cmd/dep

ADD . .
RUN dep ensure && \
	CGO_ENABLED=0 GOOS=linux go build -a -o main

# final stage
FROM alpine:latest

LABEL maintainer="Smotrov Dmitriy <dsxack@gmail.com>"

RUN apk update && apk add ca-certificates

WORKDIR /app
EXPOSE 80

COPY --from=build-env /go/src/app/main .

CMD ["./main"]
