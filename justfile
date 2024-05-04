# Tidy all packages
tidy:
	go mod tidy -C ./pkg/cli
	go mod tidy -C ./pkg/web

# Run web server
web:
	go run ./pkg/web

# Launch CLI
cli:
	go run ./pkg/cli
