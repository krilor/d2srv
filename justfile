run:
  go run main.go

build:
  mkdir -p dist
  GOOS=linux GOARCH=arm64 go build -o dist/d2srv main.go
