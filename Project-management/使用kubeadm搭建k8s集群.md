---
title: 在 ubuntu 虚拟机上搭建 k8s 集群（使用  virtualbox安装虚拟机）
content: l流程和坑
---

## 在 virtualbox 上安装 ubuntu 虚拟机

1. 教程很多，自行百度
这里贴一个相对完善的教程链接 [安装 virtualbox 及 ubuntu 虚拟机](https://blog.csdn.net/u012732259/article/details/70172704)

2. ubuntu 桌面版最新版下载地址
[Ubuntu 18.04.3 LTS](https://ubuntu.com/download/desktop)

3. 在教程中有与主机共享文件夹这一步骤，个人觉得没多大用

4. 在安装虚拟机时，可能会遇到`Kernel driver not installed`的问题（即内核驱动没有安装或者启动）。
使用命令`lsmod`查看virtualbox的内核驱动是否加载，包含四个驱动，分别是：`vboxpci`、`vboxnetadp`、`vboxnetflt`、`vboxdrv`。
按照以下步骤进行安装和启动：

首先安装 build tools和 kernel header

```shell
sudo apt-get install build-essential module-assistant
sudo m-a prepare
```

安装virtualbox-dkms

```shell
sudo apt-get install virtualbox-dkms
```

重新配置 dkms并加载vboxdrv 模块

```shell
sudo dpkg-reconfigure virtualbox-dkms
sudo modprobe vboxdrv
```

然后 virtualbox 就能正常启动了

## 配置 ubuntu 虚拟机

### 实现主机与虚拟机双向复制黏贴

打开虚拟机，在虚拟机的最上面有一排管理栏，
点击*设备*，然后*共享黏贴板*，最后点击*双向*
即可完成（如不行。重启虚拟机）

### 虚拟机终端文件夹为中文解决办法

1. 进入设置，点击*区域和语言*，将上端的语言修改为英语，重启

2. 重启后，系统会提示是否将文件名改为英文，选择右面状态栏修改，
修改完成

3. 此时系统中的控制页面也变成了英文的，若想修改回中文，
再次执行*步骤1*的操作，这次选择 Chinese，重启后不替换文件使用语言，
保留英文的形式

### 终端命令行使用 apt-get install 安装软件时报错

当使用 `apt-get install` 命令安装软件时报错：

```shell
无法获得锁 /var/lib/apt/lists/lock-frontend - open (11: 资源暂时不可用)
无法获取 dpkg 前端锁 (/var/lib/dpkg/lock-frontend)，是否有其他进程占用它?
```

解决方法：

1. 杀死占用 apt-get 的进程，释放系统锁
查找进程 PID :

```shell
ps -aux|grep apt-get
```

然后把查找的 apt-get 相关的进程全部 kill 掉：

```shell
sudo kill PID #PID 即占用 apt-get 的进程
```

1. 强制解锁，删除文件夹

当使用第一种方法时，没有找到 apt-get 的相关进程；
或者找到相应的 PID，但 kill 时系统提示无此进程；
此时可强制解锁：

```shell
sudo rm /var/cache/apt/archives/lock
sudo rm /var/lib/dpkg/lock
```

1. 再次执行 apt-get 安装或更新命令即可

### 虚拟机的网络设置

1. 在虚拟机控制台设置的网络配置中，默认虚拟机的网络连接模式是网络地址转换（NAT）

2. 将其改为桥接网卡，并将下一列的高级选项中的混杂模式修改为全部允许。

3. 重启虚拟机

## 安装搭建 k8s 集群需要的工具及组件

### 清除可能与 k8s 冲突的设置

1. 关闭防火墙：

```shell
ufw disable
```

2. 禁用 SELinux：

```shell
apt install -y selinux-utils
setenforce 0
```

3. 清除 iptables（可不进行此操作）
避免自定义 iptables 影响 k8s 服务工作，建议自定义需求通过 k8s 实现，如果满足不了，安装好 k8s 后，再配置自定义 iptables 。

```shell
iptables -t filter -F
iptables -t nat -F
iptables -t mangle -F
apt remove --purge ufw
apt remove --purge iptables-persist
```

4. 关闭虚拟交换

```shell
# 把根目录文件系统设为可读写

sudo mount -n -o remount,rw /

# 用vi修改/etc/fstab文件，在swap分区这行前加 # 禁用掉，保存退出

vi /etc/fstab

i      #进入insert 插入模式

:wq   #保存退出


# 重新启动电脑，使用free -m查看分区状态

reboot

sudo free -m

```

5. 控制平面节点（即主节点）需分配两个cpu；node节点一个就好

6. 安装ssh客户端及服务端：进行主从节点间的访问

SSH分客户端openssh-client和服务端openssh-server，ubuntu默认情况下是安装了客户端openssh-client，可以通过命令`dpkg -l | grep ssh`查询。

安装ssh-server服务端和ssh-client服务端（如果没有默认安装）

```shell
sudo apt-get install openssh-server
sudo apt-get install openssh-client
```

通过命令`ps -e | grep ssh`查看客户端与服务端是否启动，若看到下列输出即为启动：

```shell
1035 ?        00:00:00 sshd # 服务端
2153 ?        00:00:00 ssh-agent # 客户端
```

如若没启动，可以使用以下命令启动：

```shell
sudo /etc/init.d/ssh start
或
sudo service ssh start
```

登录SSH

ssh username@机器IP地址

其中，username为对应IP地址机器上的用户，需要输入密码。

断开连接：exit

7. 记得修改主机名，更好记忆

如：
主节点：k8s-master

从节点：k8s-node1

从节点：k8s-node2

### 安装及配置 docker

1. 目前使用的最新版本为 5:19.03.4~3-0~ubuntu-bionic

- 使用命令 `apt-cache madison docker-ce` 查看所有可用版本
- 使用命令 `sudo apt-get install docker-ce=<VERSION>` 安装指定版本
- 安装教程：
[docker 安装及配置](https://yeasy.gitbooks.io/docker_practice/install/ubuntu.html)

2. 对建立 docker 用户组进行补充说明：

- 在新建用户组之前，先查看用户组中是否有 docker 组，使用命令：
`sudo cat /etc/group | grep docker`
- 建立 docker 组，使用命令：
`sudo groupadd docker`
- 将当前用户加入 docker 组，使用命令：
`sudo usermod -aG docker $USER` 
其中的 `$USER` 是指当前用户的用户名
- 检查创建是否有效：
`cat /etc/group`
- 重启 docker 使权限生效：
`sudo systemctl restart docker`
- 此时在用户界面执行docker命令，测试是否配置完成，如：
`docker version`

**注**：
如果提示 get ......dial unix /var/run/docker.sock权限不够，则修改/var/run/docker.sock权限：
`sudo chmod a+rw /var/run/docker.sock`

3. 设置docker开机自启动

```shell
systemctl enable docker && systemctl start docker
```

4. 配置镜像，阿里镜像加速（可有可无）

```shell
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["阿里云加速器地址"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker
```

### 配置k8s源信息来下载 kubeadm,kubectl.kubelet

kubectl 可按照官网上的 [步骤](https://kubernetes.io/docs/tasks/tools/install-kubectl/)安装

安装 kubeadm、kubelet 需要配置国内源，使用以下命令即可（可以使用的国内源）：

```shell
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm
sudo apt-mark hold kubelet kubeadm
```

设置 kubelet 开机自启动

```shell
systemctl enable kubelet && systemctl start kubelet
```

### 配置镜像源来拉取镜像

1. 使用以下命令查看需要的镜像：

```shell
kubeadm config images list
```

2. 先创建一个以`.sh`结尾的文件，如`shell.sh`，给予该文件可执行权限：

```shell
chmod u+x shell.sh
```

3. 以下是 作者的shell 脚本，需要根据步骤1命令所列出的配置版本自行更改，
shell脚本：

```shell
images=( #根据步骤1列出的镜像来更改相应的镜像版本
    kube-apiserver:v1.16.2
     kube-controller-manager:v1.16.2
     kube-scheduler:v1.16.2
     kube-proxy:v1.16.2
     pause:3.1
     etcd:3.3.15-0
     coredns:1.6.2
)

for imageName in ${images[@]} ; do
    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName
    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/$imageName k8s.gcr.io/$imageName
done
```

4. 然后在该文件的目录下输入命令`./shell.sh`，运行此文件，等待一段时间，即可拉取需要的镜像。

### 更改cgroup driver

1. 在docker配置文件`/etc/docker/daemon.json.bk`中更改cgroup driver（警告信息），如下

```shell
{
  "registry-mirrors": ["阿里云镜像地址"], # 配置的阿里云镜像
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "100m"
  },
  "storage-driver": "overlay2"
}

```

然后重启docker：

```shell
systemctl daemon-reload
systemctl restart docker
```

**注意**：到这为止，单个节点的k8s集群的配置就完成了。
可以使用命令 `kubeadm init`进行初始化，再根据输出的内容进行操作。

也可搭建多节点集群，推荐一master几点，两node节点，每个节点的基本配置都相同，即以上步骤。

### 给master节点安装网络插件（推荐Calico）

1. 这一步是为了与其他node节点通信,
首先需使用命令`kubeadm reset`使主机恢复原状。
然后使用以下命令初始化主机：

```shell
sudo kubeadm init --pod-network-cidr=192.168.0.0/16 #为了避免网络插件的IP地址网段与k8s集群所使用的IP地址网段冲突，这里建议使用 10.0.0.0/16
```

三大内网：

C类：192.168.0.0-192.168.255.255
B类：172.16.0.0-172.31.255.255
A类：10.0.0.0-10.255.255.255

**注意**：记录该命令的输出内容，方便下面的步骤使用

2. 执行以下命令来配置kubectl （此命令是`kubeadm init`的输出内容）:

```shell
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

3. 使用以下命令安装Calico：

```shell
kubectl apply -f https://docs.projectcalico.org/v3.10/manifests/calico.yaml
```

4. 使用命令`kubectl get pods --all-namespaces`查看每个pod的状态，
等待一段时间，直到pod的状态均为`Running`，即代表配置成功。如下所示：

```shell
NAMESPACE     NAME                                       READY   STATUS    RESTARTS   AGE
kube-system   calico-kube-controllers-6d85fdfbd8-2nrzs   1/1     Running   0          69m
kube-system   calico-node-cx9jn                          1/1     Running   0          69m
kube-system   coredns-5644d7b6d9-cbftb                   1/1     Running   0          75m
kube-system   coredns-5644d7b6d9-jcwxj                   1/1     Running   0          75m
kube-system   etcd-yuxiaobo-ubuntu                       1/1     Running   0          74m
kube-system   kube-apiserver-yuxiaobo-ubuntu             1/1     Running   0          74m
kube-system   kube-controller-manager-yuxiaobo-ubuntu    1/1     Running   0          74m
kube-system   kube-proxy-6mxnd                           1/1     Running   0          75m
kube-system   kube-scheduler-yuxiaobo-ubuntu             1/1     Running   0          74m

```

### 将node节点加入集群

前提：
在将node节点加入到集群之前，需要在master节点上获得token和sha256码（即`--discovery-token-ca-cert-hash`）。
token在k8s集群创建后会自动生成，但仅24小时内有效，可以用以下命令重新创建：

```shell
kubeadm token list # 列出已存在的token
kubeadm token create # 如果token已过期，则重新创建
```

获得集群的sha256码，命令如下：

```shell
openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
```

1. 查看之前`kubeadm init`输出的最后一行内容（以下是本集群初始化后生成的内容），即为加入节点的命令：

```shell
kubeadm join 192.168.1.177:6443 --token sm6uji.x9ze9c54qopm4r3c \
    --discovery-token-ca-cert-hash sha256:ff2a7cb4df8dc6e4767051e8a69f14837a0a3ef8164f8c2d52a8970f6463d069
```

2. 在node节点上执行该命令：

```shell
kubeadm join 192.168.1.177:6443 --token sm6uji.x9ze9c54qopm4r3c \
    --discovery-token-ca-cert-hash sha256:ff2a7cb4df8dc6e4767051e8a69f14837a0a3ef8164f8c2d52a8970f6463d069
```

3. 等待5分钟，查看集群状态（仅master）

```shell
kubectl get pods
```

**注意**：在加入节点前，需要检查node机器是否可以与master主机的网络通信。

综上，使用kubeadm搭建多节点k8s集群完成。

4. 若想在node节点上控制集群，需将控制平面节点（即master节点）上的 `/etc/kubernetes/admin.conf`文件复制到 node 节点上（此`admin.conf`文件为用户提供对集群的超级用户特权），命令如下：

```shell
cd /etc/kubernetes
scp admin.conf root@node-host-ip:/etc/kubernetes/admin.conf
```

**注意**：`admin.conf`有文件权限的限制，所以需要先使用 chomd 改变文件的权限，再进行复制

### 在控制平面节点上安装 dash-board

这里不做叙述，请参考[0x06节安装 Dashboard 插件](https://tomoyadeng.github.io/blog/2018/10/12/k8s-in-ubuntu18.04/index.html)。

### 其他问题

1. 以下是搭建过程中需要的工具或者一些疑问解答

- 阿里云镜像仓库地址：

registry.cn-hangzhou.aliyuncs.com
registry.aliyuncs.com

- 当使用apt-get 去安装kubeadm、kubectl和kubelet时，显示无法定位软件：

```shell
E: 无法定位软件包 kubelet
E: 无法定位软件包 kubeadm

```

因为上面配置的是apt的源
所以要用apt去install，不能用apt-get去安装
应该使用命令

```shell
apt install -y kubelet kubeadm
```
