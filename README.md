<div align="center">
<img src="https://s2.loli.net/2025/03/30/GQDnK8aPeoBRlmX.jpg" style="width:100px;"/>
<h2>诗·韵</h2>
</div>

### 一、功能简述

- 承上启下，惩前毖后

- 像写诗一样实现应用/环境的便捷部署与启动
- 实现跨平台一句话部署应用/环境，操作和信息都在诗与韵中
- 诗-->应用程序、韵-->各种环境
- 支持使用本地的poem文件或者网络检索或指定地址文件
- 支持遇到特殊依赖环境自动拉取与安装
- 支持遇到报错自动尝试解决
- 每个操作必须设置最长等待时间
- 欢迎提交pr将可用的poem上传到example文件夹
- 也欢迎将错误信息新建一个错误ID上传到example文件夹
- 提交的poem在我们审核过后很快也会同步到网路拉取列表中

### 二、应用场景

- 一般安装一种环境或复杂安装的软件的时候我们可能会选择Docker安装，但是有很多基础应用软件并没有合适的Docker镜像或者虚拟机镜像，当然包括安装Docker本身或者虚拟机本身都可能会遇到多种多样的问题，这时候我们可能会去个各种搜索引擎搜索一步一步的操作流程进行操作，但是操作也不一定成功，而且文章质量也参差不齐，甚至有直接复制别人文章的一半就发出来了，为了节约时间，减少麻烦，"诗·韵"应运而生，我们将每一步操作和问题处理方案存到诗句中，遇到问题或依赖的环境会根据问题ID自动前往访问能解决的操作中
- 传统运维场景适用，并且因为具有一定的自我排错能力，能很大程度上减少时间成本
- 自动化运维场景适用，通过调用Rhyme能极大简化自动化运维脚本的大小
- 快捷学习场景或者编译场景配置适用，一条指令安装好编译环境和运行环境

### 三、快速开始

- 下载最新版本的诗·韵可执行程序（多平台用法统一）

- ```shell
  # 下载诗歌以及相关依赖诗歌 使用一条指令运行
  [root@localhost home]# ./org.mutantcat.rhyme -file install-nginx-centos7-amd64.poem -su
  ██████╗ ██╗  ██╗██╗   ██╗███╗   ███╗███████╗
  ██╔══██╗██║  ██║╚██╗ ██╔╝████╗ ████║██╔════╝
  ██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  
  ██╔══██╗██╔══██║  ╚██╔╝  ██║╚██╔╝██║██╔══╝  
  ██║  ██║██║  ██║   ██║   ██║ ╚═╝ ██║███████╗
  ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝
  By: mutantcat.org         诗·韵 v1.0.20250331
  [./org.mutantcat.rhyme -file install-nginx-centos7-amd64.poem -su]
  ↓↓↓↓↓↓↓↓↓↓↓ 韵 - 系统环境信息自检 ↓↓↓↓↓↓↓↓↓↓↓
  韵 - 操作系统:   linux
  韵 - 系统架构:   amd64
  韵 - CPU核心:    1
  韵 - 系统类型:   CentOS Linux
  韵 - 版本号:     7 (Core)
  韵 - ID:         centos
  韵 - ID_LIKE:    rhel fedora
  韵 - 主版本:     7
  ↑↑↑↑↑↑↑↑↑↑↑ 韵 - 系统环境信息自检 ↑↑↑↑↑↑↑↑↑↑↑
  ★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★★
  ↓↓↓↓↓↓↓↓↓↓↓ 诗 - 开始运行解析内容 ↓↓↓↓↓↓↓↓↓↓↓
  诗 - 开始执行诗句: CentOS7(amd64)安装Nginx
  
  诗 - 需要的前置验证: check-yum-linux-centos7-amd64
  诗 - 开始执行诗句: 检查yum是否可用
  
  诗 - 需要的前置验证: check-uname-a
  诗 - 开始执行诗句: 检查一下uname能否执行
  
  诗 - 正在执行命令 check-uname-a 1 : uname -a (最长 10 秒)
  
  诗 - 命令执行成功 check-uname-a 1 : Linux localhost.localdomain 3.10.0-693.el7.x86_64 #1 SMP Tue Aug 22 21:09:27 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
  
  诗 - 执行结果: 验证yum成功
  诗 - 前置验证成功: check-uname-a
  
  诗 - 正在执行命令 check-yum-linux-centos7-amd64 1 : yum --version (最长 10 秒)
  
  诗 - 命令执行成功 check-yum-linux-centos7-amd64 1 : 3.4.3
    已安装： rpm-4.11.3-25.el7.x86_64 在 2025-02-16 01:39
    构建    ：CentOS BuildSystem <http://bugs.centos.org> 在 2017-08-03 03:48
    已提交：Panu Matilainen <pmatilai@redhat.com> ，共 2017-03-17 
  
    已安装： yum-3.4.3-154.el7.centos.noarch 在 2025-02-16 01:39
    构建    ：CentOS BuildSystem <http://bugs.centos.org> 在 2017-08-05 19:13
    已提交：CentOS Sources <bugs@centos.org> ，共 2017-08-01 
  
    已安装： yum-plugin-fastestmirror-1.1.31-42.el7.noarch 在 2025-02-16 01:39
    构建    ：CentOS BuildSystem <http://bugs.centos.org> 在 2017-08-11 10:23
    已提交：Valentina Mukhamedzhanova <vmukhame@redhat.com> ，共 2017-03-21 
  
  诗 - 执行结果: 验证yum成功
  诗 - 前置验证成功: check-yum-linux-centos7-amd64
  
  诗 - 正在执行命令 install-nginx-centos7-amd64 1 : yum install -y epel-release (最长 80 秒)
  
  诗 - 命令执行成功 install-nginx-centos7-amd64 1 : 已加载插件：fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * epel: d2lzkl7pfhq30w.cloudfront.net
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  软件包 epel-release-7-14.noarch 已安装并且是最新版本
  无须任何处理
  
  诗 - 正在执行命令 install-nginx-centos7-amd64 2 : yum install -y nginx (最长 80 秒)
  
  诗 - 命令执行成功 install-nginx-centos7-amd64 2 : 已加载插件：fastestmirror, langpacks
  Loading mirror speeds from cached hostfile
   * base: mirrors.aliyun.com
   * epel: d2lzkl7pfhq30w.cloudfront.net
   * extras: mirrors.aliyun.com
   * updates: mirrors.aliyun.com
  软件包 1:nginx-1.20.1-10.el7.x86_64 已安装并且是最新版本
  无须任何处理
  
  诗 - 执行结果: 安装成功
  ↑↑↑↑↑↑↑↑↑↑↑ 诗 - 完成运行解析内容 ↑↑↑↑↑↑↑↑↑↑↑
  
  ```

- 

