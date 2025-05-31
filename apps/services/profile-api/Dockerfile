FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.work .
COPY go.work.sum .
COPY ./libs ./libs
COPY ./apps/services ./apps/services
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN go build -o /bin/profile-api ./apps/services/profile-api/cmd/server/main.go

FROM alpine:3.21
RUN apk update --no-cache bash curl
WORKDIR /app
COPY --from=builder /bin/profile-api /bin/profile-api
EXPOSE 3000
CMD [ "/bin/profile-api" ]
