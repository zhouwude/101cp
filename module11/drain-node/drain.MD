### Evict all pods from a node

```sh
# 驱逐节点---
kubectl drain cadmin 
```

### Make a node scheduable again

```sh

kubectl cordon
kubectl uncordon
```
<!-- nginx-deployment-drain-7854ff8877-778mb 是 replicaset 创建 evict之后会重建 -->
```sh
curl -v -H 'Content-type: application/json' --key admin.key --cert admin.crt https://172.16.0.3:6443/api/v1/namespaces/default/pods/nginx-deployment-drain-7854ff8877-778mb/eviction -d @eviction.json -k

{
  "apiVersion": "policy/v1",
  "kind": "Eviction",
  "metadata": {
    "name": "nginx-app",
    "namespace": "default"
  }
}
```
