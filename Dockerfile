FROM golang:alpine AS build

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN go build -v -o /app

# FROM scratch

# COPY --from=build /app /app

# CMD ["/app"]

ENV RESEND_API_KEY=
ENV NATS_URL=
ENV NATS_USER=
ENV NATS_PASSWORD=

CMD ["go", "run", "main.go"]