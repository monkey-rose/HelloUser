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
