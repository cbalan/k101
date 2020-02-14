# From Git to Kubernetes  
This tutorial assumes that the following browser based **Katacoda** playground is used: 
https://www.katacoda.com/courses/kubernetes/playground

This **Katacoda** playground provides 2 nodes. 
 * `master` - all kubernetes control components are installed. **kubectl** commands can be executed from here.
 * `node01` - regular worker node. Kubernetes can schedule workloads on this node.

To see the list of available nodes, the following **kubectl** command can be used.
```shell script
kubectl get nodes -o wide
```

A private in-cluster docker registry is used in this example. 
To install it, please execute the following on both nodes:

 1. Add `my.registry:31000` to the list of insecure registries. Point `my.registry` to `localhost`. 
```shell script
curl https://raw.githubusercontent.com/cbalan/k101/master/examples/katacoda-local-registry/configure-local-registry.sh | bash -ex

```

 2. Install local docker registry. On the `master` node, apply the docker registry manifest.
 ```shell script
kubectl apply -f https://raw.githubusercontent.com/cbalan/k101/master/examples/katacoda-local-registry/docker-registry.yaml

# Wait for the registry deployment to become ready
kubectl wait --for=condition=available deployment/my -n registry
``` 

Unless specified otherwise, please use the `master` node to follow through the exercise.

## The application
To keep the number of dependencies to a minimum, we'll create a [golang](https://golang.org/) application. 

Please note that Kubernetes apps are not limited to golang applications. 
As long as an application can be packaged as a Docker container, it can be managed via Kubernetes.  

 1. Create the application folder and initialize git repo.
```shell script
mkdir the-app
cd the-app
git init
``` 
 
 2. Create the application code in `main.go` file. Feel free to use your favored text editor. Although `vim main.go` can be used.
The following snippet creates an http server that listens on `31001` port. The `/the-data` path returns the contents of the `data.txt` file.
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/the-data", func(w http.ResponseWriter, r *http.Request) {
        data, err := ioutil.ReadFile("data.txt")
        if err != nil {
            log.Fatal(err)
        }
        count, err := fmt.Fprintf(w, "%s", string(data))
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Request: %#v, Bytes written: %d\n", r, count)
    })

    log.Fatal(http.ListenAndServe(":31001", nil))
}
```

 3. Add sample data.
```shell script
echo 'Sample data' > data.txt
```

 4. Commit application code.
```shell script
git add main.go data.txt    
git commit -a -m "Added the application code"
```

 5. Run the application.
```shell script
docker run -it --name run-the-app -p 31001:31001 -v $(pwd):/the-app -w /the-app  -d golang:1.13 go run main.go

curl http://127.0.0.1:31001/the-data

docker kill run-the-app
``` 

## Build and package
We'll package our application and it's dependencies in a docker image.

1. Create Dockerfile. Note that we are using multi-staged Docker file. https://docs.docker.com/develop/develop-images/multistage-build/
```dockerfile
FROM golang:1.13 as builder
WORKDIR /the-app
COPY main.go main.go
RUN CGO_ENABLED=0 go build -a -o the-app main.go

FROM scratch
COPY --from=builder /the-app/the-app /the-app
COPY data.txt /data.txt
WORKDIR /
ENTRYPOINT ["/the-app"]
```

2. Commit Dockerfile code.

```shell script
git add Dockerfile    
git commit -a -m "Added Dockerfile"
```

3. Build docker package.
```shell script
docker build . -t the-app:v1
```
 
4. Publish the docker image into the private local registry.
```shell script
docker tag the-app:v1 my.registry:31000/the-app:v1
docker push my.registry:31000/the-app:v1
```


## Kubernetes manifests

1. Create the `kubernetes` manifests folder.
```shell script
mkdir kubernetes
```

2. Create namespace manifest in `kubernetes/1_namespace.yaml`.
```yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: the-app
```
      
3. Create deployment manifest in `kubernetes/deployment.yaml`.
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-app
  namespace: the-app
spec:
  selector:
    matchLabels:
      app: the-app
  replicas: 2
  template:
    metadata:
      labels:
        app: the-app
    spec:
      containers:
        - name: the-app
          image: my.registry:31000/the-app:v1
          ports:
            - containerPort: 31001
          livenessProbe:
            httpGet:
              path: /the-data
              port: 31001
          readinessProbe:
            httpGet:
              path: /the-data
              port: 31001
```

4. Create the service manifest in `kubernetes/service.yaml`.
```yaml
    ---
    apiVersion: v1
    kind: Service
    metadata:
      name: the-app
      namespace: the-app
    spec:
      type: ClusterIP
      ports:
        - name: the-app
          protocol: TCP
          port: 31001
          targetPort: 31001
      selector:
        app: the-app
```

5. Commit changes.
```shell script
git add kubernetes 
git commit -a -m "Added kubernetes manifests"   
```


## Install the application in a Kubernetes cluster
Apply application manifests aganist Kubernetes.
```shell script
kubectl apply -f ./kubernetes

# inspect app deployment
kubectl -n the-app describe deployment the-app

# inspect app service
kubectl -n the-app describe service the-app

# inspect app events
kubectl -n the-app get events
```


## Change the app
Apply a change, build and push a new image and update the deployment image value.


 1. Change `main.go` to prefix the response with "the data" string.
```go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/the-data", func(w http.ResponseWriter, r *http.Request) {
        data, err := ioutil.ReadFile("data.txt")
        if err != nil {
            log.Fatal(err)
        }
        count, err := fmt.Fprintf(w, "The data:\n%s", string(data))
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Request: %#v, Bytes written: %d\n", r, count)
    })

    log.Fatal(http.ListenAndServe(":31001", nil))
}
```

 2. Commit changes.
```
git commit -a -m "Added response prefix"
```

 3. Release version 2 of the-app app.
```shell script
docker build . -t the-app:v2
docker tag the-app:v2 my.registry:31000/the-app:v2
docker push my.registry:31000/the-app:v2
```

 4. Update the `kubernetes/deployment.yaml` manifest to point to the new version.
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: the-app
  namespace: the-app
spec:
  selector:
    matchLabels:
      app: the-app
  replicas: 2
  template:
    metadata:
      labels:
        app: the-app
    spec:
      containers:
        - name: the-app
          image: my.registry:31000/the-app:v2
          ports:
            - containerPort: 31001
          livenessProbe:
            httpGet:
              path: /the-data
              port: 31001
          readinessProbe:
            httpGet:
              path: /the-data
              port: 31001
```  

 5. Commit changes
```shell script
git commit -a -m "Bumped the-app to version 2"
```

 6. Apply the changes 
```shell script
kubectl apply -f ./kubernetes

# inspect app deployment
kubectl -n the-app describe deployment the-app

# inspect app service
kubectl -n the-app describe service the-app

# inspect app events
kubectl -n the-app get events
```

## Next topic
We are done. Go to [index](../README.md) 