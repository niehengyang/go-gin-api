# Lets Projects Template
> 基于Lets框架的业务开发项目模板


## 安装

* git clone https://github.com/niehengyang/go-gin-api.git {你的项目名称}
* 进入{项目目录}，执行以下命令，移除git
```bash
find . -name ".git" | xargs rm -Rf
```
* 修改项目目录根目录中的 go.mod 文件

```conf
//后半部分更改为letsframework所在的目录
//拉取letsframework-> git clone https://github.com/niehengyang/letsframework.git
replace go.ebupt.com/lets => ../../../go/src/lets/ 
```
* 运行 go run main.go

* Web服务端口为 8901 , 访问 http://127.0.0.1:8901/ 可看见欢迎页面
