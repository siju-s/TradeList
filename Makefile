run-server:
	cd ./backend/cmd && go mod download && go run main.go

build-client:
	cd ./frontend && npm install

run-client:
	cd ./frontend && npm start