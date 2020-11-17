package main
import(
  "fmt"
  "net/http"
//  "github.com/kataras/iris/v12"
)

func main(){
  resp,err := http.Get("https://www.baidu.com/")
  if err != nil{
   fmt.Print(err)
  }
  fmt.Print(resp) 
/*  app := iris.Default()
  app.Get("/ping",func(ctx iris.Context){
  ctx.JSON(iris.Map{
     "message": "pong",
   })
})
  app.Run(iris.Addr(":8081"))*/
}
