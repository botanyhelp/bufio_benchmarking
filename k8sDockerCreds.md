### bufio_benchmarking cheat sheet steps

* login to GCP cloudshell
* clone the repo
* see that we have **go**
* **main.go** is the only thing that will be run by the docker container

```
git clone https://github.com/botanyhelp/bufio_benchmarking.git
git clone git@github.com:botanyhelp/bufio_benchmarking.git
cd bufio_benchmarking
go version
go run main.go
```

* get and run existing public docker image from docker hub
* remove image and container

```
docker image ls
docker container ls
docker pull botanyhelp/bufio_benchmarking:1.0.1
docker image ls
docker container ls
docker run b4ebee7dce04
docker container ls
docker container ls -a
docker container rm 4c02162f9bd1
docker container ls -a
docker image ls
docker image rm b4ebee
docker image ls
```

* view existing Dockerfile
* modify main.go and Dockerfile 
* build and push new image to docker hub
* from a remote command line like we have in GCP cloudshell, we will need to login via terminal because no browser is available:
* **docker login -u botanyhelp**

```
cat Dockerfile
docker image ls
docker build -t bufio_benchmarking .
docker image ls
docker tag bufio_benchmarking:latest botanyhelp/bufio_benchmarking:1.0.2
docker push botanyhelp/bufio_benchmarking:1.0.2
docker login -u botanyhelp
docker push botanyhelp/bufio_benchmarking:1.0.2
```

* start minikube
* setup kubectl
* run a testpod
* run our recently created docker image botanyhelp/bufio_benchmarking:1.0.2
* kubectl is NOT logged into docker, but docker command on local gcp cloudshell linux is logged-in
* logout of docker hub from linux

```
minikube start
source <(kubectl completion bash)
alias k=kubectl
complete -o default -F __start_kubectl k
alias kdry="kubectl --dry-run=client -o yaml"
kdry run testpod --image=nginx:1.23.4-alpine --restart=Never -- /bin/sh -c "date && sleep 30 && date && sleep 20 && date && sleep 10"
kdry run testpod --image=nginx:1.23.4-alpine --restart=Never -- /bin/sh -c "date && sleep 30 && date && sleep 20 && date && sleep 10" > pod.yaml
cat pod.yaml
k apply -f pod.yaml
kdry run bufio-benchmarking-pod --image=botanyhelp/bufio_benchmarking:1.0.2 --restart=Never > pod_bufio_benchmarking.yaml
k apply -f pod_bufio_benchmarking.yaml
cat ~/.docker/config.json|grep A4 auths
docker logout
cat ~/.docker/config.json|grep A4 auths
```

* return to docker hub and make our **botanyhelp/bufio_benchmarking** private
* https://hub.docker.com/repositories/botanyhelp
* delete pod from cluster
* notice that we can still recreate **botanyhelp/bufio_benchmarking** even though we are NOT logged-in and it is now private

```
k delete pod/bufio-benchmarking-pod
k get pod -o wide
k apply -f pod_bufio_benchmarking.yaml
```

* we visit kubernetes docs to see how to allow kubectl to pull images from a private registry, which is what we are now doing:
* https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
* we ask google why we are still able to get that private image:
* **does kubectl delete pod actually delete the downloaded image?**
* we try to login to the node and delete 

```
minikube ssh
crictl rmi --prune
k get pod -o wide
k apply -f pod_bufio_benchmarking.yaml
k get pod -o wide
k logs bufio-benchmarking-pod
```

* we stop and delete minikube
* ..and we restart fresh and see that our image is no longer pullable, and we get **ErrImagePull**

```
minikube stop
minikube delete
minikube start
k apply -f pod_bufio_benchmarking.yaml
k get pod -o wide
```

* now we are able to apply the docker-login to kubectl so that kubectl can access docker hub with login credentials:
* https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/

```
docker login -u botanyhelp
kubectl create secret docker-registry dockercreds --from-file=/home/botanyhelp/.docker/config.json
kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "dockercreds"}]}'
k apply -f pod_bufio_benchmarking.yaml
k get pod -o wide
k logs bufio-benchmarking-pod
minikube delete
```
