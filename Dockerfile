FROM golang:1.15-alpine as build

WORKDIR /app
ENV GO111MODULE=on

RUN apk --no-cache add ca-certificates

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/user cmd/user.go

FROM scratch
COPY --from=build /bin/user /
ENTRYPOINT ["/user"]
