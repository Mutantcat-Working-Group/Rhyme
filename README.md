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
- 本地文件中的诗词中的依赖和bad
- 网络检索诗词和AI生成诗词功能正在开发中...

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

- ```shell
  # 检索文件夹中可用的诗歌
  [root@localhost home]# ./org.mutantcat.rhyme -search -folder ./ -key install
  ██████╗ ██╗  ██╗██╗   ██╗███╗   ███╗███████╗
  ██╔══██╗██║  ██║╚██╗ ██╔╝████╗ ████║██╔════╝
  ██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  
  ██╔══██╗██╔══██║  ╚██╔╝  ██║╚██╔╝██║██╔══╝  
  ██║  ██║██║  ██║   ██║   ██║ ╚═╝ ██║███████╗
  ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝
  By: mutantcat.org         诗·韵 v1.0.20250331
  [./org.mutantcat.rhyme -search -folder ./ -key install]
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
  韵 - 开始搜索诗歌...
  韵 - 找到以下诗歌:
           0 :  .//install-nginx-centos7-amd64.poem
  
  韵 - 您可以使用 rayme -file xxx 来运行指定的诗歌文件
  
  ```

### 四、参数列表

| 参数名称     | 参数类型 | 参数介绍                                                     | 使用示例                                                     |
| ------------ | -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| file         | 字符串   | 执行要运行的诗歌文件                                         | rhyme -file install-nginx-centos7-amd64.poem                 |
| folder       | 字符串   | 搜索诗歌时所在的文件夹                                       | rhyme -search -folder ./                                     |
| key          | 字符串   | 搜索诗歌时的关键字                                           | rhyme -search -folder ./ -key install                        |
| depth        | 数字     | 最大联调深度                                                 | rhyme -file install-nginx-centos7-amd64.poem -depth 10       |
| su           | 布尔     | 是否使用管理员权限运行                                       | rhyme -file install-nginx-centos7-amd64.poem -su             |
| search       | 字符串   | 是否启用搜索模式                                             | rhyme -search -folder ./                                     |
| Args(xx=123) | 诗词参数 | 跟在后面的其他参数(诗中的$&{filename}和$&{password}将被替换) | rhyme -file install-nginx-centos7-amd64.poem filename=hello password=123456 |

### 五、诗词参数说明

- 诗词中可以定义$&{keyword}参数，运行时将会将诗词中全局的$&{keyword}替换为指定的Args参数值

- 若不传参，则参数将保留$&{keyword}形式

- 诗词中有一些预留参数，详见下表

| 参数名称   | 参数说明     | 是否可覆盖 |
| ---------- | ------------ | ---------- |
| OS_TYPE    | 操作系统     | 不建议     |
| OS_ARCH    | 系统架构     | 不建议     |
| OS_CORE    | CPU核心      | 不建议     |
| OS_NAME    | 系统类型     | 不建议     |
| OS_VER     | 版本号       | 不建议     |
| OS_ID      | 系统ID       | 不建议     |
| OS_LIKE    | ID_LIKE      | 不建议     |
| OS_VER_ID  | 主版本       | 不建议     |
| NEED_DEPTH | 最大联调深度 | 是         |

### 六、诗的结构

```
# 诗歌中暂时不支持注释(后续将支持) 这里的#只是用于展示内容
title: CentOS7(amd64)安装Nginx					# 诗歌标题
info: 为CentOS7(amd64)系统安装Nginx				  # 功能与作者说明
from: check-* install-*							 # 允许的来源(用于被need与bad)
need: check-yum-linux-centos7-amd64				 # 前置检测/环境安装

poem:											 # 诗歌正文(要执行的代码)
yum install -y epel-release						 # 指令行 就是一条命令行指令
-p80 bad=fix-epel-release-linux-centos7-amd64	 # 标点行 指定最大等待时间 错误处理方案
yum install -y nginx							 # 指令行
-p80 bad= check-have-any						 # 最后指令行 指定最终是否成功的检查条件

good: 安装成功									  # 成功信息
bad: 安装失败									  # 失败信息
```

- 除了诗歌内容中的指令行和标点行相交替的时候支持换行以外，其他行都是单行的
- 注意诗歌内容中的指令行和标点行必须成对出现
- 最后指令行只有在上面所有操作没报错没超时的情况下会进行检查逻辑决定是否成功，若执行直接出现错误会直接返回失败
- 最终指令行的验证指令如下表

| 指令名称          | 指令说明                | 指令示例                             |
| ----------------- | ----------------------- | ------------------------------------ |
| check-have-any    | 是否有任何结果          | -p80 bad= check-have-any             |
| check-have-any-r  | 是否有任何结果-取反     | -p80 bad= check-have-any-r           |
| check-have-none   | 是否结果为空串          | -p80 bad= check-have-none            |
| check-have-none-r | 是否结果为空串-取反     | -p80 bad= check-have-none-r          |
| check-have-all    | 是否与结果完全匹配      | -p80 bad= check-have-all=mutantcat   |
| check-have-all-r  | 是否与结果完全匹配-取反 | -p80 bad= check-have-all-r=mutantcat |
| check-have        | 是否包含某内容          | -p80 bad= check-have=mutantcat       |
| check-have-r      | 是否包含某内容- 取反    | -p80 bad= check-have-r=mutantcat     |

