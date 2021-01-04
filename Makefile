test1Image:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test1 main_test1.go
	docker build -f Dockerfiletest1 -t test1:1.0.0 .

test2Image:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test2 main_test2.go
	docker build -f Dockerfiletest2 -t test2:1.0.0 .

traefikImage:
	docker build -f Dockerfiletraefik -t gateway:1.0.0 .

clean:
	docker rmi -f $$(docker images | grep test1 | awk '{print $$3}') || true
	docker rmi -f $$(docker images | grep test2 | awk '{print $$3}') || true
	docker rmi -f $$(docker images | grep gateway | awk '{print $$3}') || true