### apply gateway with tls redirect setting
```
kubectl apply -f https-gateway.yaml
```
access gateway via http port
```
curl 10.107.158.77/productpage -v -L -k
```
