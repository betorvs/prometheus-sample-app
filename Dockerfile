FROM golang:1.17.7-alpine3.15 AS build-env

WORKDIR /tmp/simple-go-app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM scratch

ARG version 

ENV VERSION $version


WORKDIR /app

COPY static/index.html /app/static/index.html
COPY static/style.css /app/static/style.css

COPY --from=build-env /tmp/simple-go-app/prometheus-sample-app /app/prometheus-sample-app

EXPOSE 8080

CMD ["./prometheus-sample-app"]
