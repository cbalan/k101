# Kubernetes introduction workshop
For the next ~1h 30m we'll be looking at managing applications with [Kubernetes](https://kubernetes.io).

This is a hands-on type of workshop. No slides are available.

Feel free to follow along as the presenter goes through the prepared content. 
```shell script
git clone git@github.com:cbalan/k101.git
```

At any moment during the workshop, if you have any questions, please interrupt the presenter. 

The content provided by this workshop assumes that the audience has limited experience with Kubernetes. 


## Agenda
 * [Kubernetes overview](doc/1_intro.md)
 * Deep dive into 
   - [Pods](doc/2_pods.md)
   - [Deployments](doc/3_deployments.md)
   - [Services](doc/4_services.md)
 * [From Git to Kubernetes exercise](doc/5_git2kube.md) 


## Workstation setup
All samples assume that the audience already has access to a kubernetes cluster. 

To keep things simple, the following browser based **Katacoda** playground can be used: 
https://www.katacoda.com/courses/kubernetes/playground

Alternatively, the following tools can be used:
 * Git
 * Your favorite text editor
 * Docker: https://www.docker.com
 * Kind-0.7.0: https://kind.sigs.k8s.io/docs/user/quick-start
 * Kubectl-1.17.0: https://kubernetes.io/docs/tasks/tools/install-kubectl


## Terminology
Some of these may or may not be obvious:
 * **manifests** - All Kubernetes api objects are defined via YAML format
 * **k8s**, **kube** - shortcuts, references to the **Kubernetes** project 
 * **kube control**, **kube C T L**, **"kube cuddle"** - reference to **kubectl** command line
 * **SIG** - Special Interest Group. Kubernetes SIG list available here: https://github.com/kubernetes/community/blob/master/sig-list.md


## I'm done! What's next?
This workshop is barely scratching the surface. 
Please refer to https://kubernetes.io/docs/concepts/ for advanced topics like security, storage, networking and cluster administration. 

Check out how other projects are being deployed on Kubernetes via helm charts. 
https://github.com/helm/charts/tree/master/stable/

Probably check out first what is helm at https://helm.sh/ 

Check out https://landscape.cncf.io/

Check out https://operatorhub.io/  

Learn how to use managed Kubernetes powered by a cloud provider:
 - GKE - https://cloud.google.com/kubernetes-engine/docs/quickstart
 - EKS - https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html
 - AKS - https://docs.microsoft.com/en-us/azure/aks/kubernetes-walkthrough
 - https://www.digitalocean.com/resources/kubernetes/ 
 - google for others :) 
  
Or build a Kubernetes cluster from scratch: https://github.com/kelseyhightower/kubernetes-the-hard-way


## Resources
- https://12factor.net/
- https://kubernetes.io/docs/tutorials/
- https://www.katacoda.com/courses/kubernetes
- https://www.oreilly.com/library/view/kubernetes-up-and/9781491935668/
- https://www.oreilly.com/library/view/kubernetes-patterns/9781492050278/
- https://www.manning.com/books/kubernetes-in-action
