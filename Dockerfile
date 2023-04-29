FROM golang:alpine AS build

WORKDIR /port-svc

ADD . .

RUN go build -o port-svc ./cmd/port-svc/*.go


FROM alpine

COPY --from=build /port-svc/port-svc /bin/port-svc

ENTRYPOINT /bin/port-svc
