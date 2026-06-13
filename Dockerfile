FROM golang:1.26.4 AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o clima-cep .

FROM scratch
WORKDIR /app
COPY --from=build /app/clima-cep .
ENTRYPOINT ["./clima-cep"]