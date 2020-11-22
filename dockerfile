FROM golang:1.13-alpine as build_base
WORKDIR /build
COPY . .

# Force the go compiler to use modules
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o ./app

FROM scratch  
COPY --from=build_base /build/app /app
ENTRYPOINT ["/app"]
