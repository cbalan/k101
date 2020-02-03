# Services

Think of services as being an layer in front of Pods. A service is a means to expose an application running on a set of pods.

The main benefit of defining a service that sits in front of your Pods is that, rather than having to send requests to a specific Pod,
you can send requests to a service which will forward it on to an available Pod. This means we don't have to know the IP address of the Pod, but rather just the name of the service.

TODO -> resize this fello
![Kubernetes Service Diagram](images/kube-services.jpeg?raw=true "Kubernetes Services")

To view all the services running in your Kubernetes cluster you can run

```kubectl get services --all-namespaces```

