# Deployments Overview
Deployments represent a set of multiple, identical Pods with no unique identities. 
A Deployment runs multiple replicas of your application and automatically replaces any instances that fail or become unresponsive. 
Deployments are managed by the Kubernetes Deployment controller. A Deployment provides declarative updates for Pods and ReplicaSets.

Lets create an nginx deployment. 

    kubectl apply -f https://k8s.io/examples/application/deployment.yaml

Now check if the deployment was applied

    kubectl get deployments
    or
    kubectl get deploy
    

Deployment object encapsulates ReplicaSet and Pod objects
![Deployment Object Structure](images/deployment-object.png?raw=true "Deployment Object Structure ")
       
Now we can take a look at the deployment yaml on the cluster  The selector field defines how the Deployment finds which Pods to manage. 
In this case, you simply select a label that is defined in the Pod template

    kubectl describe deployment nginx-deployment
    
    kubectl edit deployments nginx-deployment

# Replica Sets
A ReplicaSet’s purpose is to maintain a stable set of replica Pods running at any given time. 
As such, it is often used to guarantee the availability of a specified number of identical Pods.

Check the current running pods
    
    kubectl get pods -o wide
 
Now try deleting a pod. Copy and paste one of your pods listed into the following delete command

    kubectl delete pod **********
 
Now Check the running pods again. What has happened?
    
    kubectl get pods -o wide

Let's check the current deployed ReplicaSets

     kubectl get rs

# Deployment Features
You describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. 
You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.
- Upgrade
- Rollback
- Scale up or Down
- Pause and resume


# Types of Deployment Update Strategies

## Recreate Strategy
Shut down version A, then start version B. Best for development environment

Pros:
 - application state entirely renewed
 
Cons:
 - downtime that depends on both shutdown and boot duration of the application

## RollingUpdate Strategy
This uses the RollingUpdate strategy provided by kubernetes. Default strategy.    

Pros:
- No downtime
- Rollback functionality

Cons:
- Takes time

   
Lets scale up the cluster, if not using a local Kube cluster please refrain from scaling to a large number as your cluster may get sluggish. 

    kubectl scale  deployment.v1.apps/nginx-deployment --replicas=5
    
Lets take a look at our pods 

    kubectl get pods -o wide

Look for the number of replicas under the spec parent. Manually edit this to assign 2. Now save the file with :x or :wq 

Now lets do a rolling upgrade to avoid downtime. We will update the docker image version on nginx in our deployment. First lets have a look at the current image defined in the deployment. 

    kubectl edit deployments nginx-deployment

Now lets set the image version using kubectl command. This could be done by editing the yaml directly either. 
This action will kick off a rolling update and the Kube API will detect the desired version is not set.

    kubectl set image deployment nginx-deployment nginx=nginx:1.9.1

Be quick!  lets check the status. The more replicas defined the longer the rollout will take. 

    kubectl rollout status deployment.v1.apps/nginx-deployment

Okay lets break some stuff :0 
Set a new nginx image to something random to try pul    l a docker image verion that does not exist 

      kubectl set image deployment nginx-deployment nginx=nginx:1.9.10000000

Check the pods. Look for errors. What do you see?

      kubectl get pods -o wide
    
We can check the rollout history 

    kubectl rollout history deployment.v1.apps/nginx-deployment

If something has gone wrong in the update to the new docker image, we can undo the update
     
    kubectl rollout undo  deployment.v1.apps/nginx-deployment 
  

For the curious, Kubernetes Deployment controller code. This is used to manger the desired state
https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/deployment/deployment_controller.go
 
 
 

## Next topic 
[Services](4_services.md)