FROM golang:1.25 as build-stage


WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . ./


RUN GOOS=linux go build -o /sfs-bridge

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /sfs-bridge /sfs-bridge

USER nonroot:nonroot
CMD ["/sfs-bridge"]