
# Installation

## Pre-requisites

* Some sort of hypervisor (this one uses VirtualBox)
* Python virtualenv 
* Vagrant

## Steps

* Clone repository to a temporary directory
  * mkdir playground && cd playground
  * git clone https://github.com/kwanlowe/digitalhermit\_website.git
  * cd digitalhermit\_website/linux/minikube

* Install the pre-requisites into local directory
  * make setup
  * source env\_setup

* Create the Virtual Machine
  * make build-box

  (If this step fails, try re-running as it occasionally gets snagged on the network.)


* If the box gets created but provisioning fails, rerun the provision step:
  * make provision-box

  ( Again, this may fail due to network. Just re-run if this happens.)

* Once the box is provisioned, you can ssh into the environment:
  * vagrant ssh


* Once you're logged in, there are a few things to check:
  * groups
    * Should return (wheel, docker)
  * swapon -s
    * Should not return anything
  * free -m
    * Should show a swap size of 0
  
* Setup minikube environment
  *  eval $(minikube -p minikube docker-env)
  *  minikube start --driver=docker
    * This step can take 10 minutes or more to complete, depending on host resources


* Once the minikube environment is up, you can verify with:
  * minikube status

You should see output similar to:

```
[vagrant@localhost ~]$ minikube status
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured

[vagrant@localhost ~]$ kubectl get pods
No resources found in default namespace.
[vagrant@localhost ~]$ kubectl get service
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   24m

```

* Once the environment is ready, you can run a test deployment:
  * `kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4`
  * `kubectl get deployments`
  * `kubectl get pods`
  * `kubectl get events`

You should see something like:
```
[vagrant@localhost ~]$ kubectl get events
LAST SEEN   TYPE      REASON              OBJECT                             MESSAGE
77s         Warning   FailedScheduling    pod/hello-node-7bf657c596-k5k27    no nodes available to schedule pods
2m20s       Normal    SuccessfulCreate    replicaset/hello-node-7bf657c596   Created pod: hello-node-7bf657c596-k5k27
2m20s       Normal    ScalingReplicaSet   deployment/hello-node              Scaled up replica set hello-node-7bf657c596 to 1
```

Again, it may take a while depending on the resources of your environment. If this fails, you may need to increase the memory
in the Vagrantfile in the git repository.



