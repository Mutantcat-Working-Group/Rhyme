<div align="center">
<img src="https://s2.loli.net/2025/03/30/GQDnK8aPeoBRlmX.jpg" style="width:100px;" width="100"/>
<h2>诗·韵</h2>
</div>

### 一、产品概述

- 像写诗一样实现应用与环境的便捷部署与启动，跨平台一句话完成安装。
- 诗即应用程序，韵即各种环境，每一步操作和问题处理方案都存在诗句里。
- 支持使用本地 poem 文件、文件夹检索，特殊依赖环境可自动拉取与安装。
- 遇到报错自动尝试解决（按错误 ID 前往对应处理方案），每个操作可设最长等待时间。
- 传统运维、自动化运维、快捷学习与编译环境配置均适用，也能大幅简化自动化运维脚本。

核心价值：把「装个环境」从翻搜索引擎、试错半天，收敛成一行命令——依赖自动装，报错自动修。

### 二、软件界面

终端工具，运行即打印 Logo 与执行日志：

```text
██████╗ ██╗  ██╗██╗   ██╗███╗   ███╗███████╗
██╔══██╗██║  ██║╚██╗ ██╔╝████╗ ████║██╔════╝
██████╔╝███████║ ╚████╔╝ ██╔████╔██║█████╗  
██╔══██╗██╔══██║  ╚██╔╝  ██║╚██╔██║██╔══╝  
██║  ██║██║  ██║   ██║   ██║ ╚═╝ ██║███████╗
╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝     ╚═╝╚══════╝
By: mutantcat.org         诗·韵 v1.0.20260920
```

### 三、功能说明

#### 诗歌文件（poem）

- 一个 poem 文件就是一条完整的部署脚本：标题、说明、前置检测、指令行与标点行交替的执行体、成功与失败信息。
- 指令行与标点行必须成对出现，标点行指定该指令的最大等待时间（`-p`，单位秒）与出错时的处理方案（`bad=`）。
- 支持 `need`（前置检测/环境安装）与 `from`（允许被哪些诗歌联调）实现诗歌复用与嵌套，最大联调深度由 `depth` 控制。
- 当前版本以本地 poem 文件与目录检索为主，网络检索诗词与 AI 生成诗词功能后续提供。

#### 自我排错

- 遇到报错时按 bad 指定的错误 ID 前往对应处理诗歌，自动尝试解决。
- 可自行提交 poem 到 `example` 文件夹贡献可用的部署诗，审核后会同步到网络拉取列表。

#### 诗词参数

- 诗词中可定义 `$&{keyword}` 参数，运行时被 `Args` 参数值全局替换；不传参则保留 `$&{keyword}` 形式。
- 预留参数（不建议覆盖）：`OS_TYPE`、`OS_ARCH`、`OS_CORE`、`OS_NAME`、`OS_VER`、`OS_ID`、`OS_LIKE`、`OS_VER_ID`；`NEED_DEPTH` 可覆盖。

#### 诗的结构

```text
# 诗歌中暂不支持注释，这里的 # 仅用于展示
title: CentOS7(amd64)安装Nginx
info: 为CentOS7(amd64)系统安装Nginx
from: check-* install-*
need: check-yum-linux-centos7-amd64

poem:
yum install -y epel-release
-p80 bad=fix-epel-release-linux-centos7-amd64
yum install -y nginx
-p80 bad= check-have-any

good: 安装成功
bad: 安装失败
```

最后指令行的验证指令：

| 指令名称 | 指令说明 | 指令示例 |
| --- | --- | --- |
| `check-have-any` | 是否有任何结果 | `-p80 bad= check-have-any` |
| `check-have-any-r` | 是否有任何结果（取反） | `-p80 bad= check-have-any-r` |
| `check-have-none` | 是否结果为空串 | `-p80 bad= check-have-none` |
| `check-have-none-r` | 是否结果为空串（取反） | `-p80 bad= check-have-none-r` |
| `check-have-all` | 是否与结果完全匹配 | `-p80 bad= check-have-all=mutantcat` |
| `check-have-all-r` | 是否与结果完全匹配（取反） | `-p80 bad= check-have-all-r=mutantcat` |
| `check-have` | 是否包含某内容 | `-p80 bad= check-have=mutantcat` |
| `check-have-r` | 是否包含某内容（取反） | `-p80 bad= check-have-r=mutantcat` |

#### 参数列表

| 参数名称 | 参数类型 | 参数介绍 | 使用示例 |
| --- | --- | --- | --- |
| `file` | 字符串 | 执行要运行的诗歌文件 | `rhyme -file install-nginx-centos7-amd64.poem` |
| `folder` | 字符串 | 搜索诗歌时所在的文件夹 | `rhyme -search -folder ./` |
| `key` | 字符串 | 搜索诗歌时的关键字 | `rhyme -search -folder ./ -key install` |
| `depth` | 数字 | 最大联调深度 | `rhyme -file install-nginx-centos7-amd64.poem -depth 10` |
| `su` | 布尔 | 是否使用管理员权限运行 | `rhyme -file install-nginx-centos7-amd64.poem -su` |
| `search` | 布尔 | 是否启用搜索模式 | `rhyme -search -folder ./` |
| `Args(xx=123)` | 诗词参数 | 替换诗中的 `$&{filename}`、`$&{password}` 等 | `rhyme -file install-nginx-centos7-amd64.poem filename=hello password=123456` |

### 四、安装与下载

最新版本：`1.0.20260920`

从 [Releases](https://github.com/Mutantcat-Working-Group/Rhyme/releases) 下载对应平台安装包（多平台用法统一）：

| 平台 | 架构 | 资产 |
| --- | --- | --- |
| Linux | amd64 / arm64 | `.tar.gz` 或 `.AppImage` |
| macOS | Intel / Apple Silicon | `.tar.gz` 或 `.dmg`（ad-hoc 签名） |
| Windows | amd64 | `.tar.gz` 或 `-setup.exe`（NSIS 安装包） |

另附 `checksums.txt` 供校验。版本号使用纯日期递增（如 `1.0.20260920`），推送同族标签（`v` 前缀可选）后，GitHub Actions 会自动构建三平台安装包并发布 Release。

### 五、快速上手

1. 下载并解压对应平台的包，得到 `org.mutantcat.rhyme` 可执行程序。
2. 运行本地诗歌，一条命令完成安装：

```shell
./org.mutantcat.rhyme -file install-nginx-centos7-amd64.poem -su
```

3. 检索文件夹里可用的诗歌：

```shell
./org.mutantcat.rhyme -search -folder ./ -key install
```

4. 把 `example` 文件夹里的 poem 放进同级目录即可直接运行，或按上面「诗的结构」写自己的部署诗。

### 六、从源码构建

```bash
git clone https://github.com/Mutantcat-Working-Group/Rhyme.git
cd Rhyme
go build -o org.mutantcat.rhyme .
```

### 七、开源协议

本项目基于 Apache-2.0 协议开源。
