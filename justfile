# Tidy all packages
tidy:
	go mod tidy -C ./pkg/cli
	go mod tidy -C ./pkg/web

# Generate templ file for web
templ:
  templ generate -path ./pkg/web

alias dt := dev-templ
# Listen for changes in templ files
dev-templ:
  watchexec --workdir ./pkg/web -e templ just templ

alias dw := dev-web
# Run web server and listen for changes
dev-web:
  watchexec -r --workdir ./pkg/web -e go go run .

# Launch CLI
cli:
	go run ./pkg/cli
