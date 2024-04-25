### create workload
```
kubectl apply -f serverpod.yaml
```
### create default networkpolicy
```
kubectl apply -f networkpolicy.yaml
```
### check server pod
```
kubectl get po -n calico-demo -owide
```
### start clientpod
```
kubectl apply -f toolbox.yaml
kubectl apply -f toolbox.yaml -n calico-demo
```
### enter toolbox shell
```
kubectl exec it <podname> bash
```
### test connectivity
```
curl <calico-pod-ip>
ping <calico-pod-ip>
```
### create global allow policy
```
kubectl create -f allow-icmp-incluster.yaml
```
### test connectivity from two toolbox pod
```
curl <calico-pod-ip>
ping <calico-pod-ip>
```
### create networkpolicy to allow access
```
kubectl create -f access-calico-demo.yaml
```
### test connectivity from two toolbox pod
```
curl <calico-pod-ip>
ping <calico-pod-ip>
```
<!-- 定义基于集群的策略和 ns 的策略 这些是同事时生效的只有规则重合优先使用 基于ns的策略，其他的啥都是并集的规则 -->