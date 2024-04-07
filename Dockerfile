FROM golang:1.17-alpine AS build

WORKDIR /home/ksenia/courseProject
COPY go.mod go.sum app.go /home/ksenia/courseProject
RUN CGO_ENABLED=0 go build -o /bin/demo


FROM scratch
COPY --from=build /bin/demo /bin/demo

EXPOSE 8888
ENTRYPOINT ["/bin/demo"]
