FROM golang:1.13-alpine as build
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /golang
RUN apk add --update --no-cache git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -a -tags netgo -ldflags '-w' -o app .

FROM alpine
LABEL MAINTAINER=1997jirasak@gmail.com
ENV TZ=Asia/Bangkok
WORKDIR /go
RUN apk add --update --no-cache tzdata ca-certificates
COPY --from=build /golang/app app
ENTRYPOINT [ "./app" ]