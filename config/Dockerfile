FROM golang:1.16.7 AS build
WORKDIR /go/src/github.com/zbwang163/
RUN git clone https://github.com/zbwang163/ad_account_server.git\
    && cd ad_account_server \
    && go mod tidy \
    && echo '11111' \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o account_server . \
    && echo '22222'

FROM alpine:latest as final
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/zbwang163/ad_account_server/account_server .
EXPOSE 50001
CMD ["./account_server"]

