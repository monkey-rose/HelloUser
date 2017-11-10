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
![](https://simimg.com/i/F5dRd)
