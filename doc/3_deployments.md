Deployments Overview
//TODO 

Deployment Features
//TODO 

Types of Deployment

Recreate - best for development environment
Pro:
application state entirely renewed
Cons:
downtime that depends on both shutdown and boot duration of the application

RollingUpdate
This uses the RollingUpdate strategy provided by kubernetes.   
Pro:
No downtime
Cons:
Takes time


Deployment object encapsulates ReplicaSet and Pod objects
<insert image>

Lets take a look at the deployment yaml file under the resources folder.
The selector field defines how the Deployment finds which Pods to manage. In this case, you simply select a label that is defined in the Pod template

Lets create the deployment

    kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1

Now check if the deployment was applied

    kubectl get deployments
       
Now we can take a look at the deployment yaml on the cluster

    kubectl edit deployments kubernetes-bootcamp
 
Let's check the current deployed ReplicaSets

     kubectl get rs
    
Lets do a rolling upgrade to avoid downtime

    kubectl rollout status deployment.v1.apps/nginx-deployment

Lets scale up the cluster, if not using a local Kube cluster please refrain from scaling to a large number as your cluster may get sluggish. 

    kubectl scale  deployment.v1.apps/kubernetes-bootcamp --replicas=5
    
Lets take a look at our pods 

    kubectl get pods --all-namespaces

    kubectl get deploy
    
Lets scale the cluster back down, but this time edit the yaml
 
     kubectl edit  deployments kubernetes-bootcamp

Look for the number of replicas under the spec parent. Manually edit this to assign 2. Now save the file with :x or :wq 
     

For the curious, Kubernetes Deployment controller code
https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/deployment/deployment_controller.go
 
 
 
 