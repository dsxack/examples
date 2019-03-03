all: embeded packed run

embeded:
	cd cmd/embeded; GOOS=linux go build -o ../../resources/data/embeded
	cd resources/data; gzip -f embeded
	cd resources/; rice embed-go

packed:
	cd cmd/packed; GOOS=linux go build -o ../../build/packed

run:
	docker run --rm -v "${PWD}/:/app" -p "8000:8000" ubuntu bash -c "/app/build/packed && sleep 1000"

