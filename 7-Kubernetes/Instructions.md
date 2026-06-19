# Kubernetes (K8s)
Kubernetes is an open-source container orchestration platform. It automates the deployment, scaling, and management of containerized applications, taking the operational burden off developers by automatically handling tasks like load balancing, self-healing, and resource allocation across multiple servers.

Good news, I implore you all to use AI for this session. I do not expect everyone to understand what is going on and you may ask as many questions as you like or go to AI.

## Installation
WARNING: PLENTY OF ISSUES MAY OCCUR WHEN ATTEMPTING TO INSTALL AND RUN DOCKER-DESKTOP. Most of these issues have to do with security settings in the BIOS, lack of hardware support or problems with a host VM.

For Windows, I recommend you download Docker-Desktop with the WSL. For Linux you can use either DockerDesktop or just docker cli. DockerDesktop installation doesn't work well inside of a VM, so if you just want cli, install Kubernetes using "kind" at https://kind.sigs.k8s.io/ or "Minikubes" at https://minikube.sigs.k8s.io/docs/.

https://docs.docker.com/get-started/get-docker/

## Steps
1. Create a Go program that accepts thousands of protobuf objects from another device.
2. Create a Go program that sends thousands of protobuf object to another device.
3. Create 2 Docker images, one that will use program 1 and another one with program 2. (Create Dockerfiles to build your docker images)
4. Then build yaml files for deploying your client and server, you may also have to build a service for pod to pod communications.


## gRPC
Install this to get gRPC.
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Run this on your proto files.
```
protoc --go_out=. --go-grpc_out=. gophertalk.proto
```

Remember to ```go use```.
```
go get google.golang.org/protobuf@latest
go get google.golang.org/grpc@latest
```

## Dockerfile
This is how you build a docker image.
```docker build -f Dockerfile -t projectName:version .```

## DockerDesktop
You have to start the DockerDesktop app, before you can start using the docker cli in WSL.

This is an example of pulling from DockerDesktop images to one that K8s can use.
```
docker save my-server:1.0 | docker exec -i desktop-control-plane ctr -n k8s.io images import -

# Confirm with this
docker exec -it desktop-control-plane ctr -n k8s.io images ls | grep image-name
```


## K8s
Run these commands to implement and destroy your changes.
```
kubectl apply -f <yaml_file>

kubectl delete -f <yaml_file>
```

## K9s
You can install a terminal GUI that works with K8s quite nicely.

```
curl -sS https://webinstall.dev/k9s | bash
```

