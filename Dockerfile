FROM golang:1.27.2-alpine AS build
WORKDIR /src
COPY go.mod .
COPY cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/catalog ./cmd/catalog

FROM scratch
COPY --from=build /out/catalog /catalog
USER 10001:10001
EXPOSE 8080
ENTRYPOINT ["/catalog"]
