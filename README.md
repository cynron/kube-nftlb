# kube-nftlb

##### Author: Víctor Manuel Oliver Acosta



## Description

`kube-nftlb` is a local Docker container able to communicate the Kubernetes API Server, using a Debian image with nftlb/nftables installed.

So far, this project only can request information from the API Server such as new, updated or deleted Services, using an official Kubernetes client (known as `client-go`).


## Software required before proceeding

* Docker
* Docker-machine
* Minikube [**v0.29.0**](https://github.com/kubernetes/minikube/releases/tag/v0.29.0) _(already started)_ 
* Golang
* `client-go`

`It is assumed that you are able to install everything on your own, following the official installation guides.`


## Getting the cluster ready

**You must only do these steps if you have NOT done it before, and if you meet the specified conditions mentioned in each point.** Otherwise, you can skip this section.

* This is a mandatory step if you started Minikube with `--vm-driver=none`, and you mustn't do it if that's not your case. `coredns` won't be able to resolve external hostnames unless you run this command:
```
root@pc: kubectl apply -f yaml/give_internet_access_to_pods.yaml
```
* The cluster needs a `nftlb` privileged rol, because in order to use `nftlb` for communicating the API Server, it needs to be recognised and authenticated by the API Server. Run this command:
```
root@pc: kubectl apply -f yaml/authentication_system_level_from_pod.yaml
```


## Project test: steps to follow

1. Download the project locally in your computer and get inside the directory. Additionally, log into your terminal as root.
```
user@pc: git clone https://github.com/zevenet/kube-nftlb
user@pc: cd kube-nftlb
user@pc: su
```

2. The script `build.sh` will compile `main.go` and will build a Docker container to put it inside the cluster. **Before running it, you MUST read the script**. Once you have read it and adapted it to your use case, run:
```
root@pc: sh build.sh
```

3. Once the script has finished, the `nftlb` Pod will be made as [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/). Inside `yaml` there's a file ready for this, apply it to the cluster running this:
```
root@pc: kubectl apply -f yaml/create_nftlb_as_daemonset.yaml
```

4. You must need to know the name of the `nftlb` Pod to do this step. Access to the cluster dashboard, with `Namespace: kube-system`, in `Workloads > Daemon Sets > nftlb`. You will see a name with a pattern similar to `nftlb-xxxxx`. Copy that name and replace it in the following command:
```
root@pc: kubectl exec -n kube-system -it nftlb-xxxxx /app
```

5. The test will be made with a [Ghost](https://ghost.org/) instance, exposing, editing and deleting a Service. Open another terminal as root (like you did in step 1) and run:
```
root@pc: kubectl run ghost --image=ghost --port=2368
```

6. The `ghost` Pod will be exposed through a Service with this command (pay attention to the terminal where you are connected to `nftlb`):
```
root@pc: kubectl expose deployment ghost --type=NodePort
```
If you see in the `nftlb` terminal a message like `Added Service: ...` followed by a JSON object, congrats! You succeeded.

7. Update the Service with this command, changing the port from 2368 to 2369, and save the file:
```
root@pc: kubectl edit svc ghost
```
If you see in the `nftlb` terminal a message like `Updated Service: ...` followed by two JSON objects, congrats! You succeeded.

8. Delete the Service with this command:
```
root@pc: kubectl delete svc ghost
```
If you see in the `nftlb` terminal a message like `Deleted Service: ...` followed by a JSON object, congrats! You succeeded.


## FAQ

* **I've done everything already, how can I exit `nftlb`?**

Press `Control` + `C`.

* **I have followed the guide and I've got no errors. But, how can I delete the `nftlb` Pod to test the project again from the start?**

Run this command as root:
```
root@pc: kubectl delete -f yaml/create_nftlb_as_daemonset.yaml
```

* **How can I also delete the `ghost` Pod? The guide explains how to delete its Service, but not its Pod.**

Run this command as root:
```
root@pc: kubectl delete deployment ghost
```