FROM 289208114389.dkr.ecr.us-east-1.amazonaws.com/moonlight-images/golang:1.22.5-alpine3.20-2aaf5373ee9dde132e09e57f52ee46b3d5374d8e AS build

WORKDIR /src

COPY . .

RUN go mod download \
    ; go build -o bin/api main.go

FROM 289208114389.dkr.ecr.us-east-1.amazonaws.com/moonlight-images/alpine:3.20.0-59a126f931cdd1e77750b5c06ddfb2c3ff7e29ef

RUN addgroup -S picpay && adduser -S picpay -G picpay
WORKDIR /home/picpay/app

COPY --from=build /src/docker-entrypoint.sh /src/bin/api ./

RUN chmod +x docker-entrypoint.sh

ENV METRICS_PORT :8181
USER picpay
EXPOSE 8181

ENTRYPOINT ["/home/picpay/app/docker-entrypoint.sh"]
CMD ["/home/picpay/app/api"]
