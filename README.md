# Kubernetes introduction workshop
For the next ~1h 30m we'll be looking at managing applications with **Kubernetes**.

This is a hands-on type of workshop. No slides are available.

Feel free to follow along as the presenter goes through the prepared content. 

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
 * **k8s**, **kube** - shortcuts, references to the **Kubernetes** project 
 * **kube control**, **kube C T L**, **"kube cuddle"** - reference to **kubectl** command line
 * **SIG** - Special Interest Group. Kubernetes SIG list available here: https://github.com/kubernetes/community/blob/master/sig-list.md
 

## Resources
- https://12factor.net/
- https://kubernetes.io/docs/tutorials/
- https://www.katacoda.com/courses/kubernetes
- https://www.oreilly.com/library/view/kubernetes-up-and/9781491935668/
- https://www.oreilly.com/library/view/kubernetes-patterns/9781492050278/
- https://www.manning.com/books/kubernetes-in-action
- https://github.com/kelseyhightower/kubernetes-the-hard-way
