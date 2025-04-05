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
