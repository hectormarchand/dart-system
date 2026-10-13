FROM golang:1.27.2-alpine3.24 AS build

WORKDIR /build

COPY dart-backend-golang/go.mod dart-backend-golang/go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -C dart-backend-golang/ -o dart-backend

FROM gcr.io/distroless/static-debian13:nonroot

WORKDIR /app

COPY --from=build /build/dart-backend ./

CMD ["./dart-backend"]
