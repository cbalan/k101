# From Git to Kubernetes

To summarize, we've seen how Kubernetes can be used to run workloads via [pods](2_pods.md).

To manage, scale and update multiple pods via [deployments](3_deployments.md). 

To interact with the endpoints exposed by multiple pods via [services](4_services.md).

Next, in order to show case an end to end flow, we'll create, package, install and upgrade a trivial web application into a Kubernetes cluster.  

This tutorial assumes that the following tools are installed on your workstation:
 * Git
 * Your favorite text editor
 * Docker: https://www.docker.com
 * Kind-0.7.0: https://kind.sigs.k8s.io/docs/user/quick-start
 * Kubectl-1.17.0: https://kubernetes.io/docs/tasks/tools/install-kubectl

## The application
To keep the number of dependencies to a minimum, we'll create a golang application. 

Please note that Kubernetes apps are not limited to golang applications. 
As long as an application can be packaged as a Docker container, it can be managed via Kubernetes.  

 1. Create the application folder and initialize git repo
```
mkdir the-app
cd the-app
git init
``` 
 
 2. Create the application code in `main.go` file. Feel free to use your favored text editor. Although `vim main.go` can be used.
The following snippet creates a http server that listens on `31001` port. The `/the-data` path returns the contents of the data.txt file.
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

 3. Add sample data
```
echo 'Sample data' > data.txt
```

 4. Commit application code - optional
```
git add main.go data.txt    
git commit -a -m "Added the application code"
```

 5. Run the application - optional 
```
docker run -it --name run-the-app -p 31001:31001 -v $(pwd):/the-app -w /the-app  -d golang:1.13 go run main.go

curl http://127.0.0.1:31001/the-data

docker kill run-the-app
``` 

## Build and package
We'll wrap our application and it's dependencies in a docker image.

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

2. Build docker package.
```
docker build . -t the-app:v1
```

3. Commit Dockerfile code - optional

```
git add Dockerfile    
git commit -a -m "Added Dockerfile"
```
    
## Kubernetes manifests

1. Create the kubernetes manifests folder
```
mkdir kubernetes
```
  
2. Create namespace manifest

```yaml
---
apiVersion: v1
kind: Namespace
metadata:
  name: the-app
```
      
3. Create deployment manifest
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
              image: catalinbalan/k101-the-app:v1
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

4. Create the service manifest
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
        - name: registry
          protocol: TCP
          port: 31001
          targetPort: 31001
      selector:
        app: the-app
```

5. Commit changes - optional
```
git add kubernetes 
git commit -a -m "Added kubernetes manifests"   
```


## Install application in a kubernetes cluster
Apply application manifests aganist a kubernetes cluster
```
kubectl apply -f ./kubernetes
```

## Next steps
Apply a change, build and push a new image and update the deployment image value

TBA