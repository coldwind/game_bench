我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# game bench
## 已编译文件
```
bin/client.xxx
bin/server.xxx
```
## 配置文件说明
### 目录
> bin/etc目录
### 文件说明
#### token_list
> 用于存放游戏token数据（公司内游戏使用）

#### request_url.yaml
> 用于配置压测目标
```
request:
  -
   address: "http://www.domain.com/game?cmd=enter" // 要压测的地址
   param: "{}" // 要输入的参数 JSON格式 如：{"param":"value"}
  -
   address: "http://www.domain.com/game?cmd=enter"
   param: "{}"
  -
   address: "http://www.domain.com/game?cmd=enter"
   param: "{}"
```

## 压测流程
### 压测的两种方式
* 通过配置文件
* 通过-a参数
### 压测的-s参数说明
> 当客户端使作-s参数时，服务器端需先启动压测的server程序进行监听  
> 客户端会收到服务器端的session id  
> 服务器端在压测期间会把CPU 内存 磁盘 网络使用情况以文本形式记录在/tmp目录下，命名格式为：bench:xxxxxxx

## Client命令行参数
> -a http://xxx.xxx.com/xxx.html 压测的URL（使用此项 则不会读取bin/etc/request_url.yaml）  
> -c 100 并发数  
> -n 1000 总数  
> -s http://xxx.xxx.com/xxx.html 监控端服务器地址（监控服务器软件需要部署在被压测的机器上，目前不支持多台服务器同时压测监控）  
> -t yes|no 是否使用token文件

## Server使用方法
> 监控服务器软件需要部署在被压测的机器上，目前不支持多台服务器同时压测监控
```
# ./server
```
