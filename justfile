# Tidy all packages
tidy:
	go mod tidy -C ./cmd/cli
	go mod tidy -C ./cmd/web

# Generate templ file for web
templ:
  templ generate -path ./cmd/web

alias dt := dev-templ
# Listen for changes in templ files
dev-templ:
  watchexec --workdir ./cmd/web -e templ just templ

alias dw := dev-web
# Run web server and listen for changes
dev-web:
  watchexec -r -e go go run ./cmd/web

# Launch CLI
cli:
	go run ./cmd/cli
