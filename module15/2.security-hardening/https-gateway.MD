### enable sidecar auto inject
```
kubectl label namespace default istio-injection=enabled
```
### cd istio root
```
cd istio-1.12.0/
```
### create bookinfo app
```
kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
kubectl get pods
```
### create expose to gateway
```
kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
```
### check ingress service http nodePort
```
kubectl get svc -n istio-system
```
### access productpage
```
<!-- 换成自己的nodeid 和 nodeport -->
curl http://172.16.0.3:31783/productpage
```

### secure the gateway by https protocol
```
<!-- /O=cncamp Inc./CN=172.16.0.3 o是 organazation 相当于用户组 common name 相当于用户 -->
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=172.16.0.3' -keyout bookinfo.key -out bookinfo.crt
kubectl create -n istio-system secret tls bookinfo-credential --key=bookinfo.key --cert=bookinfo.crt
kubectl apply -f https-gateway.yaml
```
### access product page via safari(chrome blocks self sign certs)
```
https://172.16.0.3:31106/productpage
  curl https://172.16.0.3:31106/productpage -k (-k 模拟 https)
```
