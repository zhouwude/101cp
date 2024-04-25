### define allow nothing rule in foo
```
<!-- 不接受任何请求 也不能访问别人-->
kubectl apply -f - <<EOF
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: allow-nothing
  namespace: foo
spec:
  {}
EOF
```
### access from bar.sleep to foo.httpbin
### access from foo.sleep to foo.httpbin
```
kubectl exec "$(kubectl get pod -l app=sleep -n bar -o jsonpath={.items..metadata.name})" -c sleep -n bar -- curl http://httpbin.foo:8000/headers
kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl http://httpbin.foo:8000/headers
```
### allow accessing to foo.httpbin
```
kubectl apply -f - <<EOF
apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: "httpbin-viewer"
  namespace: foo
spec:
  selector:
    matchLabels:
      app: httpbin
  action: ALLOW
  rules:
  - to:
    - operation:
        methods: ["GET"]
EOF
```

### access from bar.sleep to foo.httpbin
### access from foo.sleep to foo.httpbin
```
kubectl exec "$(kubectl get pod -l app=sleep -n bar -o jsonpath={.items..metadata.name})" -c sleep -n bar -- curl http://httpbin.foo:8000/headers
kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl http://httpbin.foo:8000/headers
```

### allow accessing from bar.sleep to foo.httpbin
```
kubectl apply -f - <<EOF
apiVersion: security.istio.io/v1
kind: AuthorizationPolicy
metadata:
  name: "httpbin-viewer"
  namespace: foo
spec:
  selector:
    matchLabels:
      app: httpbin
  action: ALLOW
  rules:
  - from:
    - source:
        namespaces: ["bar"]
    to:
    - operation:
        methods: ["GET"]
        paths: ["/info*"]
EOF
```
```
kubectl apply -f - <<EOF
apiVersion: security.istio.io/v1
kind: AuthorizationPolicy
metadata:
  name: "httpbin-viewer"
  namespace: foo
spec:
  action: ALLOW
  rules:
  - from:
    - source:
        principals: ["cluster.local/ns/bar/sa/sleep"]
    - source:
        namespaces: ["bar"]
    to:
    - operation:
        methods: ["GET"]
        paths: ["/info*"]
    - operation:
        methods: ["POST"]
        paths: ["/data"]
EOF
```
### access from bar.sleep to foo.httpbin
### access from foo.sleep to foo.httpbin
```
kubectl exec "$(kubectl get pod -l app=sleep -n bar -o jsonpath={.items..metadata.name})" -c sleep -n bar -- curl http://httpbin.foo:8000/headers
kubectl exec "$(kubectl get pod -l app=sleep -n foo -o jsonpath={.items..metadata.name})" -c sleep -n foo -- curl http://httpbin.foo:8000/headers
```
### clean up
```
k get authorizationpolicies -A
k get pa -A
k get ra -A
k get gw -A
k get vs -A
```
