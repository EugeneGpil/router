dev:
	code --folder-uri="vscode-remote://dev-container+$$(pwd | tr -d '\n' | xxd -p)/workspaces/$$(basename "$$(pwd)")"

# inside container
test:
	go test ./...
