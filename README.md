# Docker-Ci

###  Mac 安装
  * 安装Go环境
  * 设定GOPATH和GOBIN环境变量
  * docker环境，版本1.10以上
  * 目前推荐使用Dlite下安装的docker环境
  * docker in docker环境，负责编译代码，和Ci环境的使用。（说明：主要是因为Mac在编译的时候如果设置GOBIN路径则不能进行交叉编译，Docker in Docker环境推荐使用Ubuntu安装Go)

### 开发日志

#### Version V1:
  * 测试代码Build
  * 测试Docker镜像的Build
  * Push Docker镜像到指定的Repo


#### 等待修复内容
  * 基础Docker镜像不能在本地没有的时候主动去拉取
  * 不能在测试出错的情况下更换基础镜像尽心调试
  * 垃圾回收(镜像和容器的回收)不完善


test