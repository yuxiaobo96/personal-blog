---
title: 在k8s集群上部署istio,并演示Bookinfo示例
---

### 下载并安装 Istio（使用helm安装Istio）

1. 下载Istio

可去 Istio的发行版本里下载最新的压缩包，[Istio 1.4.0](https://github.com/istio/istio/releases/tag/1.4.0)，解压后，主要包含三部分：

- Kubernetes的安装YAML文件在: install/kubernetes
- 示例应用程序: samples/
- istioctl的二进制文件: /bin/istioctl

将二进制文件 istioctl添加到PATH路径下。如：

```shell
cd bin
cp -i istioctl /usr/local/bin/istioctl
```

2. 安装 helm

可去helm的发行版本里下载二进制可执行文件，[helm 3.0.0](https://github.com/helm/helm/releases/tag/v3.0.0)，解压后，将helm的二进制文件移动到PATH路径下：

```shell
cp -i linux-amd64/helm /usr/local/bin/helm
```

3. 安装 Istio

为istio-system组件创建一个命名空间：

```shell
kubectl create namespace istio-system
```

安装 Istio 的CRD：

```shell
helm template install/kubernetes/helm/istio-init -name istio-init --namespace istio-system | kubectl apply -f -
```

等待所有的 Istio CRD 创建完成：

```shell
kubectl -n istio-system wait --for=condition=complete job --all
```

部署Istio的配置文件：

**注意**：在部署之前，需要对工作节点的配置进行检查，在istio中的pilot组件需要大于2G的内存（建议分配4G），telemetry组件需要大于 1CPU的内核（建议分配2CPU），才可以正常启动。

```shell
helm template install/kubernetes/helm/istio -name istio --namespace istio-system | kubectl apply -f -
```

验证安装是否正确：

参考 Istio 的[组件表](https://istio.io/docs/setup/additional-setup/config-profiles/)，验证与k8s配合使用的组件是否全部安装并启动：

```shell
kubectl get svc -n istio-system
```

确保所调用的 pod均处于 Running 状态：

```shell
kubectl get pods -n istio-system
```

**错误**：在部署istio的默认配置文件后，所有的svc和pod都正常运行，但是istio-galley报错，连接失败，拒绝访问，目测是因为 svc中的istio-ingressgateway 的EXTERNAL-IP处于pending状态，没有外部负载均衡，可以将名为 istio-ingressgateway 的 service 类型修改为 NodePort 类型，随后通过 istio-ingressgateway 所在的服务器IP地址+端口的形式访问服务。

### 部署Bookinfo

istio的默认部署配置中，sidecar自动注入处于开启状态

1. 为bookinfo示例创建命名空间并开启自动注入功能

```shell
kubectl create ns bookinfo
kubectl label namespace bookinfo istio-injection=enabled
```

2. 部署bookinfo

```shell
cd /samples/bookinfo/platform/kube
kubectl apply -f bookinfo.yaml -n bookinfo
```

3. 检查服务和pod均已启动并running

```shell
kubectl get svc -n bookinfo
kubectl get pods -n bookinfo
```

4. 部署网关，使其可在外部访问

```shell
cd /samples/bookinfo/networking
kubectl apply -f  bookinfo-gateway.yaml
```

检查bookinfo-gateway是否运行：

```shell
kubectl get gateway -n bookinfo
```

部署默认的destination规则：

```shell
kubectl apply -f destination-rule-all.yaml
```

**综上**：bookinfo示例的部署已完成，可通过master节点（控制平面节点）的IP地址，加nodePort的端口（31380）进行访问，如：`192.168.43.112:31380/productpage` 。该nodePort的端口号来自`istio-ingressgateway`文件中，可通过命令`kubectl get svc istio-ingressgateway -n istio-system -oyaml`查看。不断刷新界面，就会显示不同的书籍评价级别信息。

### 自定义路由目标规则转发

自定义yaml文件，并配置，这里使其转发到reviews-v3版本，使其永远显示红色的星星

```shell
vim virtual-service-reviews-v3.yaml
# 填入以下内容
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews
spec:
  hosts:
    - reviews
  http:
  - route:
    - destination:
        host: reviews
        subset: v3
```

部署该自定义路由规则的yaml文件：

```shell
kubectl apply -f virtual-service-reviews-v3.yaml
```

刷新页面，即可。

**注意**：也可以使用`/istio-1.4.0/samples/bookinfo/networking`下自带的路由规则转发yaml文件自行测试。
