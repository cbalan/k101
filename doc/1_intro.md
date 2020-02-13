# Kubernetes overview

## What is Kubernetes?
Production-Grade Container Scheduling and Management
 - https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/
 - https://kubernetes.io/docs/concepts/overview/components/
 
 Kubernetes is an open source 'Production-Grade' container Orchestration platform. It can be used to automate
 deployment, scaling and management of containerized applications. With modern web services, users expect
 applications to be available 24/7, and developers expect to deploy new versions of those applications
 several times a day. Containerization helps package software to serve these goals, enabling
 applications to be released and updated in an easy and fast way without downtime.
 Kubernetes helps you make sure those containerized applications run where and when
 you want, and helps them find the resources and tools they need to work
 
 One of the advantages of it being open source is that it gives you the freedom to run your workloads anywhere
 (on-premises, hybrid, or public cloud infrastructure).
 
 In this workshop we are going to focus on the basics but the platform is fully extendable to suit a various number
 of use cases and work loads.

## Kubernetes Basic Architecture
![Kubernetes Architecture Diagram](images/kube-arch.png?raw=true "Kubernetes Basic Architecture ")

More details about each components can be found at https://kubernetes.io/docs/concepts/overview/components/

Above is a basic diagram of some of the components that make up the platform. The master node is the brain of
the platform. Users use the cli tool kubectl to interact with master node and it will provision your containers across
the worker nodes. Each worker node is running an agent called kubelete which is interacting with the master node
provision containers as necessary until the platform matches the state that you define. You will rarely ever
interact with the worker nodes directly.

If you are looking to understand how to create your own kubernetes cluster read
Kelsey Hightower's [Kubernetes the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way)

## Kubernetes Basics

All interactions in this workshop will be performed using the kubectl tool. 
https://kubernetes.io/docs/reference/kubectl/overview/

Use the following syntax to run commands

`kubectl [command] [TYPE] [NAME] [flags]`

You can also look at this as operations on a resource


For example:

```
kubectl     [Operations]    [Resource]
kubectl     get         nodes
```

There are a lot of action to choose from but some of the most commonly uses are 

```
apply       - Apply a configuration change to a resource from a file or stdin.
create      - Create one or more resources from a file or stdin.
describe    - Display the detailed state of one or more resources.
get         - List one or more resources.
logs        - Print the logs for a container in a pod.
```


We will be going through various types of resources throughout this talk:

----
Let get a feel for some of these commands and what they do. Don't worry about understanding what is happening
these are just to give you a feel for some of the actions that are performed using the commands.

`kubectl apply -f resources/intro.yaml`

To see everything 
```shell
master $ kubectl get all
NAME                                    READY   STATUS    RESTARTS   AGE
pod/intro-deployment-785df54c5c-5lbz6   1/1     Running   0          94s
pod/intro-deployment-785df54c5c-c6vvc   1/1     Running   0          94s
pod/intro-deployment-785df54c5c-m7wfp   1/1     Running   0          94s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   21m

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/intro-deployment   3/3     3            3           94s

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/intro-deployment-785df54c5c   3         3         3       94s
```

You can can get all of a specific resource. (adding -owide give you more information)
```
master $ kubectl get deployments -owide
NAME               READY   UP-TO-DATE   AVAILABLE   AGE    CONTAINERS   IMAGES    SELECTOR
intro-deployment   3/3     3            3           107s   intro        busybox   app=hello-world
```

You can use describe to find out more details about something running
```
master $ kubectl describe deployment intro-deployment
Name:                   intro-deployment
Namespace:              default
....
```

You can view the logs of a running pod
```
master $ kubectl get pods
NAME                                READY   STATUS    RESTARTS   AGE
intro-deployment-785df54c5c-5lbz6   1/1     Running   0          2m48s
intro-deployment-785df54c5c-c6vvc   1/1     Running   0          2m48s
intro-deployment-785df54c5c-m7wfp   1/1     Running   0          2m48s

master $ kubectl logs intro-deployment-785df54c5c-5lbz6
Hello Kubernetes!
```


Clean up 
```
master $ k delete -f resources/into.yaml
deployment.apps "nginx-deployment" deleted
```





