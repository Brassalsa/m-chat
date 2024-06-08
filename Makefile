build:
	go build -o bin/m-chat ./app

run:build
	./bin/m-chat

# install air first the run dev cmd
dev-run: 
	air

css-build: 
	npx tailwindcss -i ./static/css/dev.css -o ./static/css/build.css
	
css-dev: 
	npx tailwindcss -i ./static/css/dev.css -o ./static/css/build.css --watch

dev:
	make -j dev-run css-dev
