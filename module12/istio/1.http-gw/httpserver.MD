### Deploy simple

```sh
kubectl create ns simple
kubectl create -f simple.yaml -n simple
kubectl create -f istio-specs.yaml -n simple
```

### Check ingress ip

```sh
k get svc -nistio-system

istio-ingressgateway   LoadBalancer   10.107.158.77
```

### Access the simple via ingress

```sh
#这里的 pod 没有注入sidecar 这里是通过 istio-ingress-gateway的 svc来访问的，这里 外部的入站流量会同意走ingress-gateway的istio-proxy 相当于一个  单独的sidecar
export INGRESS_IP=10.107.158.77
curl -H "Host: simple.cncamp.io" $INGRESS_IP/simple/hello -v
```
