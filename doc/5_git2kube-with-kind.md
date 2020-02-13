# From Git to Kubernetes  
This tutorial assumes that the following tools are installed on your workstation:
 * Git
 * Your favorite text editor
 * Docker: https://www.docker.com
 * Kind-0.7.0: https://kind.sigs.k8s.io/docs/user/quick-start
 * Kubectl-1.17.0: https://kubernetes.io/docs/tasks/tools/install-kubectl

Once KIND is installed, please create a new KIND cluster. 
```shell script
kind create cluster
```

## The application
To keep the number of dependencies to a minimum, we'll create a golang application. 

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

1. Create Dockerfile. Note that we are using multi-staged Docker file.
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
 
4. Load docker image into KIND in order to allow Kubernetes to use it.
```shell script
kind load docker-image the-app:v1
```


## Kubernetes manifests

1. Create the kubernetes manifests folder.
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
          image: the-app:v1
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

# port forward the-app service to localhost 
kubectl -n the-app port-forward svc/the-app 31001:31001

# open the application endpoint in a browser window
open http://127.0.0.1:31001/the-data
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
kind load docker-image the-app:v2
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
          image: the-app:v2
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

# port forward the-app service to localhost 
kubectl -n the-app port-forward svc/the-app 31001:31001

# open the application endpoint in a browser window
open http://127.0.0.1:31001/the-data
```

## Clean up
```shell script
# delete app resources
kubectl delete -f ./kubernetes

# delete kind cluster
kind delete cluster
```

## Next topic
We are done. Go to [index](../README.md) 