FROM golang:1.23-alpine as build

WORKDIR /src

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app

FROM scratch

COPY --from=build /app /app

CMD ["/app"]
