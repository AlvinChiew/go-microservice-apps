# Getting Started

1. Install [docker](https://docs.docker.com/desktop/windows/wsl/)

1. Install [minikube](https://minikube.sigs.k8s.io/docs/start/) 
    - `brew install minikube`.

1. Setup minikube
    - `minikube start --driver docker --ports=127.0.0.1:9090:9090`

1. tie docker alias to docker in minikube
    - ubuntu `eval $(minikube docker-env)`
    - windows bash `eval $(minikube -p minikube docker-env)`

1. Build docker image & create container
    - `docker build -t go-get-host-details:0.1.0 ./;` 
    - `docker run -d -p 9090:9090 --name get-host-details-0.1.0 go-get-host-details:0.1.0`

# Extra

- Build go exe with `CGO_ENABLED=0` for statically linked binaries to bypass `standard_init_linux.go:228: exec user process caused: no such file or directory`
    - `CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main`

- To debug docker image
    - docker run --rm -it --entrypoint=/bin/bash go-get-host-details:0.1.0