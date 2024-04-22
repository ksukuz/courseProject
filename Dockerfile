FROM golang:1.17-alpine AS build

WORKDIR /var/www
COPY app.go go.mod go.sum /var/www
RUN CGO_ENABLED=0 go build -o /bin/demo


FROM scratch
COPY --from=build /bin/demo /bin/demo

ENTRYPOINT ["/bin/demo"]
