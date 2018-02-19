build: clean build-app build-image

clean : 
	rm -rf ./dist

build-app :
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/main ./src

build-image :
	docker build -t example-scratch -f Dockerfile .

local-host : 
	docker run -it -p 8000:8000 example-scratch

run : build local-host
