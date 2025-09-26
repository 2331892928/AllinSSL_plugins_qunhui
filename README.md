# AllinSSL_plugins_qunhui
AllinSSL插件，将AllinSSL的证书部署到群晖
# 食用方法
## 下载release中的文件
- 根据自己的运行平台选择qunhui插件版本，然后根据你的运行方式选择安装方式
- docker：映射/www/allinssl/plugins到本地，然后把程序直接放进去  
- 源码编译：编译好的程序直接放入allinssl程序同级plugins文件夹，没有自己创建  
## 自行编译
### 自行编译方法
- 请确保已安装Go 1.23+环境：[前往下载地址](https://golang.google.cn/dl/)
- [进入AllinSSL仓库](https://github.com/allinssl/allinssl)
### 克隆仓库并下载依赖
```bash
git clone https://github.com/allinssl/allinssl.git
cd allinssl
```
### 继续执行,设置go镜像  
#### windows
```bash
$env:GO111MODULE = "on"
$env:GOPROXY = "https://goproxy.cn"
```
#### mac or linux
```bash
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```
### 继续执行,下载依赖
```bash
go mod tidy
```
### 编译前期工作
#### 将本项目放入plugins目录
![img.png](img.png)
#### 进入qunhui文件夹
![img_1.png](img_1.png)
图片仅供参考，根据你实际目录
### 开始编译
- 下一步进入编译,根据自己的开发平台选择，根据自己的运行平台执行
#### windows平台
##### windows 编译 windows
```bash
go build
```
##### windows 编译 linux，你是arm平台就把amd64改为arm64
```bash
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build
```
##### windows 编译 mac，你是arm平台就把amd64改为arm64
```bash
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build
```
#### linux平台/mac平台
##### linux 编译 linux
```bash
go build
```
##### linux 编译 windows
```bash
export GOOS = "windows"
export GOARCH = "amd64"
go build
```
##### linux 编译 mac
```bash
export GOOS = "darwin"
export GOARCH = "amd64"
go build
```
### 放入allinssl平台
- 将编译好的文件保存下来，进行下一步选择操作
- docker：映射/www/allinssl/plugins到本地，然后把编译好的程序直接放进去
- 源码编译：编译好的程序直接放入allinssl程序同级plugins文件夹，没有自己创建