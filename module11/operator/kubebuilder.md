### Create a kubebuilder project, which requires an empty folder

```
go mod init github.com/cncamp/demo-operator
```
```sh
kubebuilder init --domain cncamp.io
```

### Check project layout

```sh
cat PROJECT

domain: cncamp.io
layout:
- go.kubebuilder.io/v3
projectName: mysts
repo: github.com/cncamp/demo-operator
version: "3"
```

### Create API, create resource[Y], create controller[Y]

```sh
# 创建新的 api 相当于一个新的 k get apiservice  GKV 三要素 apigroup apiversion kind 
# --kind 创建的 crd 的类型 controller gen /bin 目录下 控制器生成工具
#api 目录下面 有api 文件
# internal/controller 下面有 controller 代码
kubebuilder create api --group apps --version v1beta1 --kind MyDaemonset
```

### Open project with IDE and edit `api/v1alpha1/simplestatefulset_types.go`

```sh
// MyDaemonsetSpec defines the desired state of MyDaemonset
type MyDaemonsetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MyDaemonset. Edit mydaemonset_types.go to remove/update
	Image string `json:"image,omitempty"`
}

// MyDaemonsetStatus defines the observed state of MyDaemonset
type MyDaemonsetStatus struct {
	AvaiableReplicas int `json:"avaiableReplicas,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}
```

### Check Makefile
<!-- make manifests生成对应的 Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects-->
<!-- make install  在~/.kube/config中指定的K8s集群中安装crd。
make build -> make all 整个流程走一遍
make docker-build 构建镜像
make  docker-push 推送镜像
make deploy  将控制器部署到~/.kube/config中指定的K8s集群。
-->

```makefile
Build targets:
    ### create code skeletion
    manifests: generate crd
    generate: generate api functions, like deepCopy

    ### generate crd and install
    run: Run a controller from your host.
    install: Install CRDs into the K8s cluster specified in ~/.kube/config.

    ### docker build and deploy
    docker-build: Build docker image with the manager.
    docker-push: Push docker image with the manager.
    deploy: Deploy controller to the K8s cluster specified in ~/.kube/config.
```

### Edit `controllers/mydaemonset_controller.go`, add permissions to the controller
```go
//+kubebuilder:rbac:groups=apps.cncamp.io,resources=mydaemonsets/finalizers,verbs=update
// Add the following
//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;create;update;patch;delete
```

### Generate crd

```sh
make manifests
```

### Build & install

```sh
make build
# 生成映像并将其推送到以下位置
make docker-build 
# 将控制器部署到具有以下项指定的映像的集群
make docker-push 
# 这里注意必须把 controller 的镜像 push之后才能 deploy controller
# controller 部署在一个单独的 namespace 里面。
make deploy #将 crd 和 controller 都不知道集群当中去
```

## Enable webhooks

### Install cert-manager

```sh
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.1/cert-manager.yaml
```

### Create webhooks

```sh
kubebuilder create webhook --group apps --version v1beta1 --kind MyDaemonset --defaulting --programmatic-validation
```

### Change code

### Enable webhook in `config/default/kustomization.yaml`
### make install 需要安装一些 webhook 的CRD
### Redeploy
