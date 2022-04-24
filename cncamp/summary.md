历时15个礼拜，终于完成了云原生训练营所有课程，现在我将这段时间所学习的内容总结如下：
本人经过本次培训后，自己搭建了高可用集群，harbor，以及ceph集群。
并且部署了service，ds,deployment到集群中，理论结合实践，对云原生有了更深入的认识，学习到了很多东西，然而训练的内容实在太多，有些内容还来不及消化及实践，比如isto等，在后续的工作和学习中我会继续学习没有理解的内容，并且将所学的内容应用到实际的工作中。
以下是我对本次知识点做的一些笼统的总结，写出来的这些内容基本在理解范围之内，算是对本次培训的一次回顾：

什么是云原生？为什么要用云原生？
云计算正在迅速成为各种规模的企业满足其计算基础设施需求的标准。Kubernetes 是一个可移植的、可扩展的开源平台，用于管理容器化的工作负载和服务，可促进声明式配置和自动化。Kubernetes 拥有一个庞大且快速增长的生态系统。Kubernetes 的服务、支持和工具广泛可用。Kubernetes 这个名字源于希腊语，意为“舵手”或“飞行员”。k8s 这个缩写是因为 k 和 s 之间有八个字符的关系。 Google 在 2014 年开源了 Kubernetes 项目。Kubernetes 建立在 Google 在大规模运行生产工作负载方面拥有十几年的经验 的基础上，结合了社区中最好的想法和实践1，目前Kubernetes已经成为了主导云计算弹性计算的平台。云原生技术有利于各组织在公有云、私有云和混合云等新型动态环境中，构建和运行可弹性扩展的应用。云原生的代表技术包括容器、服务网格、微服务、不可变基础设施和声明式API。

云原生基础技术-GO语言
GO语言的特性，GO是一个可以编译高效，支持高并发，面向垃圾回收的全新语言，适合云原生。
GO的环境搭建，编译。
GO的基本语法，包括控制结构，数组，切片，函数，结构体，接口和反射，错误处理等。
GO的高级特性，包括协程，channel，定时器，上下文，线程加锁，线程调度，Goroutine，MPG的关系，生成消费模型，内存管理，包引用与依赖管理，debig等等。
Go语言GC的触发流程，GOGC 用于控制GC的处发频率， 其值默认为100, 这意味着直到自上次垃圾回收后heap size已经增长了100%时GC才触发运行，live heap size每增长一倍，GC触发运行一次。若设定GOGC=200, 则live heap size 自上次垃圾回收后，增长2倍时，GC触发运行， 总之，其值越大则GC触发运行频率越低， 反之则越高。如果GOGC=off 则关闭GC。
令人印象深刻的是MPG调度模型，以及内存管理，Goroutine调度和内存管理自动化，GO的底层完全包揽了Goroutine的调度和内存的分配，开发者只需要关注应用程序的逻辑，为开发者带来了很大的便利，大大的节省了开发时间。

云原生基础技术-docker
Docker的基本概念，系统架构，Dockerfile最佳实践，和微服务的关系。
Docker的OverlayFS文件系统：
OverlayFS是一个面向Linux的文件系统服务，其实现一个面向其他文件系统的联合挂载。Overlayfs的基本理念只读底层和可写顶层叠加在一起，看起来像一个单一的文件系统。在这种情况下，目录被叠加，文件从顶部看。如果有冲突，对于文件，顶层将优先考虑，而对于目录，顶层将被合并。如果你编辑一个只在底部的文件（只读），在顶部会有一个副本。
Docker可以更高效的利用系统资源，更快速的启动时间，一致的运行环境，持续交付和部署，更轻松的迁移，更轻松的维护和扩展。
容器的主要特性：安全性，隔离性，便携性，可配额。
Namespace是linux内核的资源隔离方案，系统为进程分配不同的namespace，各个不同的namespace互不干扰，每个容器对应一个namespace。
Linux通过Cgroup来给容器配置资源。
最佳实践的目标：易管理，少漏洞，镜像小，层级少，利用缓存。


Kubernetes的相关概念及技术，知识点：
master节点（管理节点）
提供集群的控制
对集群进行全局决策
检测和响应集群事件
主要由：kube-apiserver、kuberproxy、schedular、controllermanage和etcd
node节点（计算节点）
运行容器的实际节点
维护运行pod，并提供具体应用的运行环境
node节点由：kubelet、kube-proxy和容器引擎（如docker）组成
计算节点被设计成水平扩展，该组件在多个节点上运行。

master节点：是整个集群的控制中枢
Kube-APIServer：集群的控制中枢，各个模块之间信息交互都需要经过Kube-APIServer，同时它也是集群管理、资源配置、整个集群安全机制的入口。后端元数据存储于ETCD中（键值对数据库）。
Controller-Manager：集群的状态管理器，保证Pod或其他资源达到期望值，也是需要和APIServer进行通信，在需要的时候创建、更新或删除它所管理的资源。相当于“大总管”。
Scheduler：集群的调度中心，它会根据指定的一系列条件，选择一个或一批最佳的节点，然后部署我们的Pod。相当于“调度室”
Etcd：键值数据库，存储k8s运行中产生的元数据，一般生产环境中建议部署三个以上节点（奇数个）。
node节点：工作节点
Kubelet：负责监听节点上Pod的状态，同时负责上报节点和节点上面Pod的状态，负责与Master节点通信，并管理节点上面的Pod。
Kube-proxy：负责Pod之间的通信和负载均衡，将指定的流量分发到后端正确的机器上。实现service的通信与负载均衡。
容器引擎：负责对容器管理，例如docker。containerd、CRI-O、runC等

Pod
是kubernetes调度的最小部署单位
pod是一个容器的集合,一个pod可以包含一个或多个容器
同个Pod中的容器使用相同的网络命名空间和IP地址, 端口号
同个pod下的容器相互之间通过localhost来发现和通信
每个组内的容器共享一个存储卷(volume)
pod是一个服务的多个进程的聚合单位
pod作为一个独立的部署单位,支持横向扩展和复制
Pod是短暂的
Controllers：控制器是更高级层次对象，用于部署和管理Pod，每个控制器负责不同的任务。
Deployment：无状态应用部署
StatefulSet：有状态应用部署
DaemonSet：确保所有Node运行同一个Pod
Job：一次性任务
Cronjob：定时任务
Service

防止Pod失联
定义一组Pod的访问策略
Label：标签，附加到某个资源上，用于关联对象、查询和筛选

Namespaces：命名空间，将对象逻辑上隔离，系统命名空间有一下四种
default：默认的命名空间，不声明空间的pod都在这里。
kube-node-lease：为高可用提供心跳监视的命名空间
kube-public：公共数据，所有用户都可以读取它。
kube-system：系统服务对象所使用的命名空间
K8S核心Plugin：
Calico：符合CNI标准的网络插件，给每个Pod生成一个唯一的IP地址，并且把每个节点当做一个路由器。
CoreDNS：用于Kubernetes集群内部Service的解析，可以让Pod把Service名称解析成IP地址，然后通过Service的IP地址进行连接到对应的应用上。
Dashboard：基于web的用户接口，用于可视化k8s集群
容器监控系统：例如Prometheus、Heapster等。
集群日志系统：帮助管理员发现和定位问题，常用的是EFK
Ingress Controller：为服务提供外网入口。

Kubernetes 多组件之间的通信原理：
•	apiserver 负责 etcd 存储的所有操作，且只有 apiserver 才直接操作 etcd 集群
•	apiserver 对内（集群中的其他组件）和对外（用户）提供统一的 REST API，其他组件均通过 apiserver 进行通信
o	controller manager、scheduler、kube-proxy 和 kubelet 等均通过 apiserver watch API 监测资源变化情况，并对资源作相应的操作
o	所有需要更新资源状态的操作均通过 apiserver 的 REST API 进行
•	apiserver 也会直接调用 kubelet API（如 logs, exec, attach 等），默认不校验 kubelet 证书，但可以通过 --kubelet-certificate-authority 开启（而 GKE 通过 SSH 隧道保护它们之间的通信）

•	用户通过 REST API 创建一个 Pod
•	apiserver 将其写入 etcd
•	scheduluer 检测到未绑定 Node 的 Pod，开始调度并更新 Pod 的 Node 绑定
•	kubelet 检测到有新的 Pod 调度过来，通过 container runtime 运行该 Pod
•	kubelet 通过 container runtime 取到 Pod 状态，并更新到 apiserver 中

用 kubeadm 搭建集群环境
核心层：Kubernetes 最核心的功能，对外提供 API 构建高层的应用，对内提供插件式应用执行环境
应用层：部署（无状态应用、有状态应用、批处理任务、集群应用等）和路由（服务发现、DNS 解析等）
管理层：系统度量（如基础设施、容器和网络的度量），自动化（如自动扩展、动态 Provision 等）以及策略管理（RBAC、Quota、PSP、NetworkPolicy 等）
接口层：kubectl 命令行工具、客户端 SDK 以及集群联邦
生态系统：在接口层之上的庞大容器集群管理调度的生态系统，可以划分为两个范畴
Kubernetes 外部：日志、监控、配置管理、CI、CD、Workflow等
Kubernetes 内部：CRI、CNI、CVI、镜像仓库、Cloud Provider、集群自身的配置和管理等
在更进一步了解了 k8s 集群的架构后，我使用的是kubeadm工具来进行集群的搭建。
kubeadm是Kubernetes官方提供的用于快速安装Kubernetes集群的工具，通过将集群的各个组件进行容器化安装管理，通过kubeadm的方式安装集群比二进制的方式安装要方便不少，但是目录kubeadm还处于 beta 状态，还不能用于生产环境，Using kubeadm to Create a Cluster文档中已经说明 kubeadm 将会很快能够用于生产环境了。
kubernetes 可以存数据方式有很多，大致有两种，持久化存储与非持久化存储非持久化存储主要是  emptydir 非 emptydir 的基本都是持久存储Kubernetes 默认支持很多种存储，有些是内部原生支持，nfs 是内部原生支持；有些是通过接口支持，通过接口支持的好处是可以对接各家的云。block ebscontainer storageoss 对象存储PV 和 PVCPV：PV 描述的是持久化存储卷，主要定义的是一个持久化存储在宿主机上的目录，比如一个 NFS 的挂载目录。PVC：PVC 描述的是 Pod 所希望使用的持久化存储的属性，比如，Volume 存储的大小、可读写权限等等。
StorageClass PV 和 PVC 方法虽然能实现屏蔽底层存储，但是 PV 创建比较复杂，通常都是由集群管理员管理，这非常不方便。Kubernetes 解决这个问题的方法是提供动态配置 PV 的方法，可以自动创 PV。管理员可以部署 PV 配置器（provisioner），然后定义对应的 StorageClass，这样开发者在创建 PVC 的时候就可以选择需要创建存储的类型，PVC 会把 StorageClass 传递给 PV provisioner，由 provisioner 自动创建 PV。在声明 PVC 时加上 StorageClassName，就可以自动创建 PV，并自动创建底层的存储资源。
基于 nfs 的存储配置https://github.com/kubernetes-csi/csi-driver-nfs/blob/master/docs/install-csi-driver-master.md因为对 nfs 的支持是 kubernetes 内部就实现的，所以大部分文章是静态配置，比如在 pv 里声明 nfs server 地址，或者在 deployment/sts 里直接挂 nfs。

kube-scheduler的工作任务是根据各种调度算法将请求资源绑定到最合适的工作节点，整个调度流程分为两个阶段：预选策略（Predicates）和优选策略（Priorities）。
预选调度
根据预选策略，遍历所有目标Node，筛选出符合要求的候选节点。

•	PodFitsHostPorts：检查Pod容器所需的HostPort是否已被节点上其它容器或服务占用。如果已被占用，则禁止Pod调度到该节点。
•	PodFitsHost：检查Pod指定的NodeName是否匹配当前节点。
•	PodFitsResources：检查节点是否有足够空闲资源（例如CPU和内存）来满足Pod的要求。
•	PodMatchNodeSelector：检查Pod的节点选择器(nodeSelector)是否与节点(Node)的标签匹配
•	NoVolumeZoneConflict：对于给定的某块区域，判断如果在此区域的节点上部署Pod是否存在卷冲突。
•	NoDiskConflict：根据节点请求的卷和已经挂载的卷，评估Pod是否适合该节点。
•	MaxCSIVolumeCount：决定应该附加多少CSI卷，以及该卷是否超过配置的限制。
•	CheckNodeMemoryPressure：如果节点报告内存压力，并且没有配置异常，那么将不会往那里调度Pod。
•	CheckNodePIDPressure：如果节点报告进程id稀缺，并且没有配置异常，那么将不会往那里调度Pod。
•	CheckNodeDiskPressure：如果节点报告存储压力(文件系统已满或接近满)，并且没有配置异常，那么将不会往那里调度Pod。
•	CheckNodeCondition：节点可以报告它们有一个完全完整的文件系统，然而网络不可用，或者kubelet没有准备好运行Pods。如果为节点设置了这样的条件，并且没有配置异常，那么将不会往那里调度Pod。
•	PodToleratesNodeTaints：检查Pod的容忍度是否能容忍节点的污点。
•	CheckVolumeBinding：评估Pod是否适合它所请求的容量。这适用于约束和非约束PVC。

优选调度
确定最优节点，在第1步的基础上，采用优选策略 （ xxx Priority ）计算出每个候选节点的积分，积分最高者胜出。

•	SelectorSpreadPriority：对于属于同一服务、有状态集或副本集（Service，StatefulSet or ReplicaSet）的Pods，会将Pods尽量分散到不同主机上。
•	InterPodAffinityPriority：策略有podAffinity和podAntiAffinity两种配置方式。简单来说，就说根据Node上运行的Pod的Label来进行调度匹配的规则，匹配的表达式有：In, NotIn, Exists, DoesNotExist，通过该策略，可以更灵活地对Pod进行调度。
•	LeastRequestedPriority：偏向使用较少请求资源的节点。换句话说，放置在节点上的Pod越多，这些Pod使用的资源越多，此策略给出的排名就越低。
•	MostRequestedPriority：偏向具有最多请求资源的节点。这个策略将把计划的Pods放到整个工作负载集所需的最小节点上运行。
•	RequestedToCapacityRatioPriority：使用默认的资源评分函数模型创建基于ResourceAllocationPriority的requestedToCapacity。
•	BalancedResourceAllocation：偏向具有平衡资源使用的节点。
•	NodePreferAvoidPodsPriority：根据节点注释scheduler.alpha.kubernet .io/preferAvoidPods为节点划分优先级。可以使用它来示意两个不同的Pod不应在同一Node上运行。
•	NodeAffinityPriority：根据preferredduringschedulingignoredingexecution中所示的节点关联调度偏好来对节点排序。
•	TaintTolerationPriority：根据节点上无法忍受的污点数量，为所有节点准备优先级列表。此策略将考虑该列表调整节点的排名。
•	ImageLocalityPriority：偏向已经拥有本地缓存Pod容器镜像的节点。
•	ServiceSpreadingPriority：对于给定的服务，此策略旨在确保Service的Pods运行在不同的节点上。总的结果是，Service对单个节点故障变得更有弹性。
•	EqualPriority：赋予所有节点相同的权值1。
•	EvenPodsSpreadPriority：实现择优 pod的拓扑扩展约束
优雅终止：
Pod 被删除，状态置为 Terminating。kube-proxy 更新转发规则，将 Pod 从 service 的 endpoint 列表中摘除掉，新的流量不再转发到该 Pod。如果 Pod 配置了 preStop Hook ，将会执行。kubelet 对 Pod 中各个 container 发送 SIGTERM 信号以通知容器进程开始优雅停止。等待容器进程完全停止，如果在  terminationGracePeriodSeconds 内 (默认 30s) 还未完全停止，就发送 SIGKILL 信号强制杀死进程。所有容器进程终止，清理 Pod 资源。
资源需求和QoSQoS 服务质量GuaranteedBurstableBestEffortQoS 类为 Guaranteed 的 Pod：Pod 中的每个容器都必须指定内存限制和内存请求。对于 Pod 中的每个容器，内存限制必须等于内存请求。Pod 中的每个容器都必须指定 CPU 限制和 CPU 请求。对于 Pod 中的每个容器，CPU 限制必须等于 CPU 请求。
服务网格
1.	服务网格为我们提供了一种一致的方式来连接、保护和观察微服务。
1.	网格内的代理捕获了网格内所有通信的请求和指标。
1.	每一次失败、每一次成功的调用、重试或超时都可以被捕获、可视化，并发出警报。
1.	此外，可以根据请求属性做出决定。例如，我们可以检查入站(或出站)请 求并编写规则，将所有具有特定头值的请求路由到不同的服务版本。

所有这些信息和收集到的指标使得一些场景可以合理地直接实现。开发人员和运营商可以配置 和执行以下方案，而不需要对服务进行任何代码修改。  

•	mTLS和自动证书轮换 
•	使用指标识别性能和可靠性问题 
•	在Grafana等工具中实现指标的可视化;这进一步允许改变并与PagerDuty整合，例如使用Jaeger或Zipkin对服务进行调试和追踪
•	基于权重和请求的流量路由，金丝雀部署，A/B测试
•	流量镜像
•	通过超时和重试提高服务的弹性.
•	通过在服务之间注入故障和延迟来进行混沌测试
•	检测和弹出不健康的服务实例的断路器。
常见Mesh技术
amalgam8 - 用于异构微服务的基于版本的路由网格
ambassador - 开源的基于 Envoy proxy 构建的用于微服务的 Kubernetes 原生 API 网关
aspen-mesh - 隶属于 F5 的公司开发的 Service Mesh
conduit - 适用于 Kubernetes 的轻量级 Service Mesh
consul - Consul 一种分布式、高可用的和数据中心感知解决方案，用于跨动态分布式基础架构连接和配置应用程序
dubbo - Apache Dubbo™ (incubating)是一款高性能Java RPC框架。
envoy - C++ 前端/服务代理
istio - 用于连接、保护、控制和观测服务。
kong - 云原生 API 网关 https://konghq.com/install
linkerd - 云原生应用的开源 Service Mesh https://linkerd.io
mesher - 华为开源的基于轻量级基于 go chassis 的 Service Mesh。
mosn - MOSN是由蚂蚁金服开源的一个模块化可观测的智能网络，可用作为 sidecar 部署在 Service Mesh 中。
nginmesh - 基于 Nginx 的 Service Mesh
servicecomb - ServiceComb 是华为开源的微服务框架，提供便捷的在云中开发和部署应用的方式。
tars - Tars 是腾讯开源的基于名称服务的高性能 RPC 框架。使用 tars 协议并提供半自动化运维平台。

Sidecar 注入及流量劫持

将应用程序的功能划分为单独的进程可以被视为 Sidecar 模式。Sidecar 设计模式允许你为应用程序添加许多功能，而无需额外第三方组件的配置和代码。

Kubernetes 通过 Admission Controller 自动注入，或者用户使用 istioctl 命令手动注入 sidecar 容器。
应用 YAML 配置部署应用，此时 Kubernetes API server 接收到的服务创建配置文件中已经包含了 Init 容器及 sidecar proxy。

在 sidecar proxy 容器和应用容器启动之前，首先运行 Init 容器，Init 容器用于设置 iptables（Istio 中默认的流量拦截方式，还可以使用 BPF、IPVS 等方式） 将进入 pod 的流量劫持到 Envoy sidecar proxy。所有 TCP 流量（Envoy 目前只支持 TCP 流量）将被 sidecar 劫持，其他协议的流量将按原来的目的地请求。

启动 Pod 中的 Envoy sidecar proxy 和应用程序容器。这一步的过程请参考通过管理接口获取完整配置。
Sidecar proxy 与应用容器的启动顺序问题
启动 sidecar proxy 和应用容器，究竟哪个容器先启动呢？正常情况是 Envoy Sidecar 和应用程序容器全部启动完成后再开始接收流量请求。但是我们无法预料哪个容器会先启动，那么容器启动顺序是否会对 Envoy 劫持流量有影响呢？答案是肯定的，不过分为以下两种情况。

•	情况1：应用容器先启动，而 sidecar proxy 仍未就绪
这种情况下，流量被 iptables 转移到 15001 端口，而 Pod 中没有监听该端口，TCP 链接就无法建立，请求失败。

•	情况2：Sidecar 先启动，请求到达而应用程序仍未就绪
这种情况下请求也肯定会失败

早期是在程序上判断envoy启动了没有，或者程序延迟启动。实际上,为了解决 sidecar 启动顺序的问题,Kubernetes官方在 1.18 之后特别引入了sidecar container lifecycle的概念,也就是说,通过对 k8s 配置对应的 lifecycle，就能确保 sidecar container 在应用容器启动之前启动。

1.	不论是进入还是从 Pod 发出的 TCP 请求都会被 iptables 劫持，inbound 流量被劫持后经 Inbound Handler 处理后转交给应用程序容器处理，outbound 流量被 iptables 劫持后转交给 Outbound Handler 处理，并确定转发的 upstream 和 Endpoint。

1.	Sidecar proxy 请求 Pilot 使用 xDS 协议同步 Envoy 配置，其中包括 LDS、EDS、CDS 等，不过为了保证更新的顺序，Envoy 会直接使用 ADS 向 Pilot 请求配置更新
Envoy
Envoy 中绝大部分的模块和功能都是围绕着这四个概念展开的：
•	Listener（监听器）：监听器负责监听数据端口，接受下游的连接和请求。作为代理软件，无论是正向代理还是反向代理，肯定要接受来自下游的连接并进行数据处理。Envoy 把相关的功能都抽象在了名为监听器的资源当中。
•	Cluster（集群）：集群是对真实后端服务的抽象和封装，管理后端服务连接池、负责后端服务健康检查、实现服务级熔断等等。
•	Filter（过滤器）：过滤器主要负责将 Listener 接收的客户端二进制数据包解析为结构化的协议数据，比如 HTTP 二进制流解析为具体的 Header、Body、Trailer、Metadata 诸如此类并进行各种流量治理。Envoy 中 Filter 分为多种类型，覆盖不同的层级和场景，是 Envoy 强大功能的源泉。
•	Route（路由）：路由一般是作为某个协议解析 Filter 的一部分存在。筛选器解析出结构化数据后会根据路由中具体规则选择一个 Cluster，最终发数据转发给后端服务。
Downstream 请求自 Listener 进入 Envoy，流经 Filter 被解析、修改、记录然后根据 Route 选择 Cluster 将其发送给 Upstream 服务。

Istio
Istio 是一个功能非常丰富的服务网格，包括以下功能。

•	流量管理
•	策略控制
•	可观察性
•	安全认证

Kubernetes vs Istio vs Envoy
在回顾了 Kubernetes 的 kube-proxy 组件、xDS 和 Istio 对流量管理的抽象后，现在我们仅从流量管理的角度来看看这三个组件 / 协议的比较（注意，三者并不完全等同）

流量管理Istio 中定义了以下 CRD 来帮助用户进行流量管理。VirtualService在 Istio 服务网格中定义路由规则，控制流量路由到服务上的各种行为。规则是按照在 YAML 文件中的顺序执行的。
Gateway为 HTTP/TCP 流量配置了一个负载均衡，多数情况下在网格边缘进行操作，用于启用一个服务的入口（ingress）流量，相当于前端代理。与 Kubernetes 的 Ingress 不同，Istio Gateway 配置四层到六层的功能（例如开放端口或者 TLS 配置），而 Kubernetes 的 Ingress 是七层的。将 VirtualService 绑定到 Gateway 上，用户就可以使用标准的 Istio 规则来控制进入的 HTTP 和 TCP 流量。Gateway 设置了一个集群外部流量访问集群中的某些服务的入口，而这些流量究竟如何路由到那些服务上则需要通过配置 VirtualServcie 来绑定。
ServiceEntry通常用于在 Istio 服务网格之外启用的服务请求，通过ServiceEntry资源，我们可以向Istio的内部服务注册表添加额外的条目，使不属于我们网格的外部服务或内部服务看起来像是我们服务网格的一部分。      当一个服务在服务注册表中时，我们就可以使用流量路由、故障注入和其他网格功能，就像我们对其他服务一样，主要应对的场景是向外部的虚拟机导流
EnvoyFilter描述了针对代理服务的过滤器，用来定制由 Istio Pilot 生成的代理配置。一定要谨慎使用此功能。错误的配置内容一旦完成传播，可能会令整个服务网格陷入瘫痪状态。这一配置是用于对 Istio 网络系统内部实现进行变更的。

多集群：
跨集群同步资源：Federation 提供了在多个集群中保持资源同步的能力。例如，可以保证同一个 deployment 在多个集群中存在。跨集群服务发现：Federation 提供了自动配置 DNS 服务以及在所有集群后端上进行负载均衡的能力。例如，可以提供一个全局 VIP 或者 DNS 记录，通过它可以访问多个集群后端。

Istio 多集群
东西流量不同服务器之间的流量与数据中心或不同数据中心之间的网络流被称为东西流量。简而言之，东西流量是server-server流量。南北流量客户端和服务器之间的流量被称为南北流量。简而言之，南北流量是server-client流量。该命名来自于绘制典型network diagrams的习惯。在图表中，通常核心网络组件绘制在顶部（NORTH），客户端绘制在底部（SOUTH），而数据中心内的不同服务器水平（EAST-WEST）绘制。
网络平面安全networkpolicy用户在使用k8s中，有对网络策略的配置需求，有时候希望不同的namespace之间不能互相访问，但是我们知道k8s中所有的pod之间都是可以互相访问的，这个时候就需要网络策略来帮助我们实现这个诉求，网络策略依赖于cni网络插件，不是所有的网络插件都支持网络策略，calico支持网络策略，而flannel不支持。
入口规则：仅允许 来自（from）命名空间（spec/ingress/from/namespaceSelector）下所有pod策略所在命名空间下pod满足（spec/ingress/from/podSelector）ip段满足（spec/ingress/from/ipblock）的pod以什么协议访问本命名空间下的满足（spec/podSelector）的pod的某个端口或某个段的端口出口规则：仅允许本命名空间下的满足（spec/podSelector）的pod去访问（to）命名空间（spec/ingress/from/namespaceSelector）下的所有pod 本命名空间满足（spec/ingress/from/podSelector）的podip段满足（spec/ingress/from/ipblock）的pod的某个端口或某个段的端口

微服务进一步解耦合SOA体系下某个服务还是偏大，那么进一步拆分提高模块复用性，比如原来每个模块都需要和esb做认证，那么就将认证功能拆分出来作为一个服务容错性高，将错误隔离在单个服务内。双刃剑，复杂度提升了几个等级，运维难度加大技术选型灵活，不同的服务可以根据自己的需求选择不同的技术。自由灵活选择开发语言和框架，使用RPC接口通信易于扩展，可以按需扩展服务，避免资源的浪费。在没有kubernetes之前，需要公司自行研发PaaS平台kubernetes之后，需要公司对kubernetes的特性更为熟悉独立部署，每个服务独立部署，当其中一个服务有需求变更时，可以只编译部署单个应用，减少了对用户的影响。抛弃了原有的开发部署模式小幅度增量


