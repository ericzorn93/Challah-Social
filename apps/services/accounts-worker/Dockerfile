FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.work .
COPY go.work.sum .
COPY ./libs ./libs
COPY ./apps/services ./apps/services
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -o /bin/accounts-worker ./apps/services/accounts-worker/cmd/server/main.go

FROM alpine:3.21
RUN apk update --no-cache bash curl
WORKDIR /app
COPY --from=builder /bin/accounts-worker /bin/accounts-worker
EXPOSE 3000
EXPOSE 5000
CMD [ "/bin/accounts-worker" ]
