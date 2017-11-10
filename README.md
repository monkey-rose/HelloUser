# HelloUser
## 框架和代码
选择了 martini 框架，特性：

* 使用非常简单
* 无侵入设计
* 可与其他 Go 的包配合工作
* 超棒的路径匹配和路由
* 模块化设计，可轻松添加工具
* 大量很好的处理器和中间件
* 很棒的开箱即用特性
* 完全兼容 http.HandlerFunc 接口

当然，对于我们这个简单的程序来说，不能充分体会到 martini 框架的好处。首先安装 martini 包:

>go get github.com/go-martini/martini

代码：
```
//server.go
package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  //martini.Params服务获取参数，处理器返回字符串显示在page上
  m.Get("/hello/:name", func(params martini.Params) string {
    return "Hello " + " " + params["name"]
  })
  m.Run()
}
```
## curl 测试
#### 首先运行服务，开始监听
![](http://ww1.sinaimg.cn/large/0060lm7Tly1flc6e7mzzqj30m601kt93.jpg)

#### curl
![](http://ww1.sinaimg.cn/large/0060lm7Tly1flc6d7t9ayj30tm0cemzs.jpg)

#### 网页
![](http://ww4.sinaimg.cn/large/0060lm7Tly1flc6eecohnj30gy05smxl.jpg)

### ab测试
参数-n是请求的数量，-c是并发的请求数，-n 1000 -c 100也就是每次同时发100个请求，一共10次。结果中有总时间，失败的请求数，请求的平均响应时间，吞吐率，传输率等，还有请求完成时间中各个部分的情况，有助于分析服务器
![F5dRd.png](https://s1.simimg.com/2017/11/10/F5dRd.png)
