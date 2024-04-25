### Install ingress controller

```sh
kubectl create -f nginx-ingress-deployment.yaml
```

### Generate key-cert

```sh
# use openssl (version >= 1.1.1f) on Linux, e.g. Ubuntu 20.04
# don't run on macOS, which is using LibreSSL
# instead, you can `brew install openssl` on macOS
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=cncamp.com/O=cncamp" -addext "subjectAltName = DNS:cncamp.com"
```

### Create secret

```sh
kubectl create secret tls cncamp-tls --cert=./tls.crt --key=./tls.key
```

### Create a ingress

```sh
kubectl create -f ingress.yaml
```

### Test the result
<!-- 10.96.129.187 是ingress创建 service -->
```sh
curl -H "Host: cncamp.com" https://10.96.129.187 -v -k
curl -H "Host: argocd.myapp.com" http://10.96.129.187 -v
```

### Understand why ingress is not enough?
<!-- 理解为什么ingress是不够的?  极简的需求可以用  不能真正达到生产需求
ingress 不支持 header rewrite 也不支持 url rewrite 很难满足生产环境。
-->
- tls: cypher, dhkey, TLSVersion
- header based L7 rule
- rewriting?
  - header rewriting
  - URI rewriting
<!-- 在nginx ingress中如何支持重写 ingress 不支持  counter lstio-->
### How rewrite is supported in nginx ingress

```
这都是社区提供的 provider 通过添加 annotation 来解决 ingress 的不足，但由于分散不能同意
Name	Description	Values
nginx.ingress.kubernetes.io/rewrite-target	Target URI where the traffic must be redirected	string
nginx.ingress.kubernetes.io/ssl-redirect	Indicates if the location section is accessible SSL only (defaults to True when Ingress contains a Certificate)	bool
nginx.ingress.kubernetes.io/force-ssl-redirect	Forces the redirection to HTTPS even if the Ingress is not TLS Enabled	bool
nginx.ingress.kubernetes.io/app-root	Defines the Application Root that the Controller must redirect if it's in '/' context	string
nginx.ingress.kubernetes.io/use-regex	Indicates if the paths defined on an Ingress use regular expressions	bool
```
