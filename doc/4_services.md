# Services

Think of services as being an layer in front of Pods. A service is a means to expose an application running on a set of pods.

The main benefit of defining a service that sits in front of your Pods is that, rather than having to send requests to a specific Pod,
you can send requests to a service which will forward it on to an available Pod. This means we don't have to know the IP address of the Pod, but rather just the name of the service.

![Kubernetes Service Diagram](images/kube-services.jpeg?raw=true "Kubernetes Services")

To view all the services running in your Kubernetes cluster you can run

```
master $ kubectl get services --all-namespaces
    NAMESPACE     NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
    default       kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP                  8m18s
    kube-system   kube-dns     ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   8m16s
```

For example, let's say you have a pod or set of pods running that are labelled MyApp. Here is a simple Service definition defined in YAML that references these pods.
Essentially, you are saying that this service corresponds to the application running on those pods. Any query to this service, will be forwarded to these pods.

```
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
    - protocol: TCP
      port: 80
      targetPort: 9376
```

We can create this service in the kuberenets cluster by running

```kubetcl apply -f my-service.yaml```

We can get all services to see verify that it has been created
```
master $ kubectl get services
    NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
    my-service   ClusterIP   10.100.246.73   <none>        80/TCP    9s
```

Lets take a deeper look at my-service by running the following:
```
master $ kubectl describe services/my-service
    Name:              my-service
    Namespace:         default
    Labels:            <none>
    Annotations:       kubectl.kubernetes.io/last-applied-configuration:
                         {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"my-service","namespace":"default"},"spec":{"ports":[{"port":80,"p...
    Selector:          app=MyApp
    Type:              ClusterIP
    IP:                10.100.246.73
    Port:              <unset>  80/TCP
    TargetPort:        9376/TCP
    Endpoints:         <none>
    Session Affinity:  None
    Events:            <none>
```

Services can be defined as different Types. If you notice above, the type we created was of Type 'ClusterIp'. There are 4 service Types
- ClusterIp
- NodePort
- LoadBalancer
- ExternalName


The two most common types are ClusterIp & NodePort

ClusterIp
- Exposes a service from **within** the kubernetes cluster only. This means that only requests from within the cluster can reach this service.

NodePort
- Exposes a service from **outside** the kubernetes cluster. For example, from your local machine, if there was connectivity, you could send requests to this
kind of service.




