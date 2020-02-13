#Pods
https://kubernetes.io/docs/concepts/workloads/pods/pod-overview/
-----

A pod is the most basic unit in the kubernetes. You provide the API with a yaml template definition 

You can think of a pod as a wrapper around a single container, although it can 
incorporate multiple containers which is commonly referred to as a side car. You define
in yaml what you would like to be deployed inside of your pod and some meta data that you would like
to associate with it. Kubernetes will orchestrate a location for this pod to be deployed 

Follow the next steps to deploy your first pod!

Example yaml:
 ```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
spec:
  containers:
  - name: myapp-container
    image: busybox
    command: ['sh', '-c', 'echo Hello Kubernetes! && sleep 3600']
```

Let create a pod:

`kubectl apply -f templates/simple_pod.yaml`

Something was created... let find out what 

```kubectl get pods
master $ k get pods
NAME        READY   STATUS    RESTARTS   AGE
myapp-pod   1/1     Running   0          4m38s
```

Its alive

You can view the logs of a running pod with:

```
master $ kubectl logs myapp-pod
Hello Kubernetes!
```

```
master $ kubectl describe pod/myapp-pod
   Name:               myapp-pod
   Namespace:          default
   Priority:           0
   PriorityClassName:  <none>
   Node:               node01/172.17.0.36
   Start Time:         Fri, 31 Jan 2020 16:25:36 +0000
   Labels:             app=myapp
   Annotations:        kubectl.kubernetes.io/last-applied-configuration:
                         {"apiVersion":"v1","kind":"Pod","metadata":{"annotations":{},"labels":{"app":"myapp"},"name":"myapp-pod","namespace":"default"},"spec":{"c...
   Status:             Running
   IP:                 10.32.0.2
   Containers:
     myapp-container:
       Container ID:  docker://c795cb9094a00e46e6b3c3f0699d3c6785758b93228d6413e2d52aa81360f27b
       Image:         busybox
       Image ID:      docker-pullable://busybox@sha256:6915be4043561d64e0ab0f8f098dc2ac48e077fe23f488ac24b665166898115a
       Port:          <none>
       Host Port:     <none>
       Command:
         sh
         -c
         echo Hello Kubernetes! && sleep 3600
       State:          Running
         Started:      Fri, 31 Jan 2020 16:25:39 +0000
       Ready:          True
       Restart Count:  0
       Environment:    <none>
       Mounts:
         /var/run/secrets/kubernetes.io/serviceaccount from default-token-9t9k5 (ro)
   Conditions:
     Type              Status
     Initialized       True
     Ready             True
     ContainersReady   True
     PodScheduled      True
   Volumes:
     default-token-9t9k5:
       Type:        Secret (a volume populated by a Secret)
       SecretName:  default-token-9t9k5
       Optional:    false
   QoS Class:       BestEffort
   Node-Selectors:  <none>
   Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                    node.kubernetes.io/unreachable:NoExecute for 300s
   Events:
     Type    Reason     Age   From               Message
     ----    ------     ----  ----               -------
     Normal  Scheduled  10m   default-scheduler  Successfully assigned default/myapp-pod to node01
     Normal  Pulling    10m   kubelet, node01    Pulling image "busybox"
     Normal  Pulled     10m   kubelet, node01    Successfully pulled image "busybox"
     Normal  Created    10m   kubelet, node01    Created container myapp-container
     Normal  Started    10m   kubelet, node01    Started container myapp-container```
```

Getting into a running container:
```
master $ k exec -it myapp-pod /bin/sh
/ #
```


Destroying a POD
```
master $ k delete -f pod.yaml
   pod "myapp-pod" deleted
```

but you can also give the pod name

```
master $ k delete pod myapp-pod
pod "myapp-pod" deleted

```

#Probes
https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes

You can configure probes to better protect your pods from failure there are three main types of probes that 
you can define that each serve a different purpose

##Liveness probe
As the name suggest this probe will check if you your pod is running and restart it if necessary

```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: liveness
  name: liveness-exec
spec:
  containers:
  - name: liveness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    - touch /tmp/healthy; sleep 30; rm -rf /tmp/healthy; sleep 600
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5

```

The liveness probe will poll based on the defined periodSeconds and execute the check that you
define. If the the check fails it will automatically restart the pod. Run the above command and in ~30
seconds you will see that the pod has restarted.

```
kubectl apply -f resources/liveness_probe.yaml

master $ watch kubectl get pods
NAME            READY   STATUS    RESTARTS   AGE
liveness-exec   1/1     Running   1          79s

master $ kubectl describe pods/liveness-exec
...
  Warning  Unhealthy  12s (x6 over 97s)    kubelet, node01    Liveness probe failed: cat: can't open '/tmp/healthy': No such file or directory
  Normal   Killing    12s (x2 over 87s)    kubelet, node01    Container liveness failed liveness probe, will be restarted
```

You can create liveness probes that will run against http endpoints for example

```yaml
livenessProbe:
    httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3
```


##Readiness probe

A readiness probe can be used to set conditions that your probe needs to meet before being set to ready.
```yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: readiness
  name: readiness-exec
spec:
  containers:
  - name: readiness
    image: k8s.gcr.io/busybox
    args:
    - /bin/sh
    - -c
    -  sleep 30; touch /tmp/healthy; sleep 600
    readinessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5
```

```
kubectl apply -f resources/readiness_probe.yaml

master $ watch kubectl get pods
NAME            READY   STATUS    RESTARTS   AGE
readiness-exec   1/1     Running   0          79s
```

Once the file has been created it will be set to a ready state. This is useful when you are starting a pod that
take a long time to initialise.
