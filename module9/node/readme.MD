## nodefs 系统分区 工作目录和容器日志

### containerd & docker

`/var/lib/kubelet/pods/ecbbaf5e-c7f1-4f00-8627-454f03cfbb4d`

## imagefs 容器运行时分区 也就是容器运行时的目录 overlayFS作为容器存储系统的驱动 
<!-- 两层结构 容器可写层和镜像层 -->
<!-- 不同的容器运行时目录的区别 -->
### containerd 

`/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs`

### docker

`/var/lib/docker/overlay2/`

## Eviction configuration

```yaml
apiVersion: kubelet.config.k8s.io/v1beta1
kind: KubeletConfiguration
evictionHard:
  memory.available: '500Mi'
  nodefs.available: '1Gi'
  imagefs.available: '100Gi'
evictionMinimumReclaim:
  memory.available: '0Mi'
  nodefs.available: '500Mi'
  imagefs.available: '2Gi'
```
