build:css-build
	go build -o bin/m-chat ./cmd/main.go

run:build
	./bin/m-chat

test: 
	go test ./... -v

# install air first the run dev cmd
dev-run: 
	air

css-build: 
	npx tailwindcss -i ./web/static/css/dev.css -o ./web/static/css/build.css
	
css-dev: 
	npx tailwindcss -i ./web/static/css/dev.css -o ./web/static/css/build.css --watch

dev:
	make -j dev-run css-dev
