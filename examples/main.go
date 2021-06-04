package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/wind-2019/captcha"
	"math/rand"
	"strconv"
	"time"
)

//定义临时存放验证码坐标轴的字典
var captchaData map[string]int

func main() {
	captchaData = make(map[string]int, 0)
	app := iris.Default()
	crs := cors.New(cors.Options{
		//开启跨域，配合联合调试//Access-Control-Allow-Origin
		AllowedOrigins:   []string{"http://localhost", "*"},                           //允许通过的主机域名
		AllowedMethods:   []string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"}, //通过的方式
		AllowedHeaders:   []string{"Accept", "content-type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Screen"},
		AllowCredentials: true,
	})
	r := app.Party("/", crs)
	{
		r.Get("/getImgTest", hero.Handler(GetImgTest))
		r.Get("/check", hero.Handler(Check))
	}
	// 运行
	_ = app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
func Check(ctx iris.Context) iris.Map {
	id := ctx.URLParamDefault("id", "")
	left := ctx.URLParamIntDefault("left", 0)
	if _, ok := captchaData[id]; !ok {
		return iris.Map{"code": 500, "info": "id不存在"}
	}

	if (captchaData[id] - left) >= 10 {
		return iris.Map{"code": 504, "info": "验证失败"}
	}
	return iris.Map{"code": 200, "info": ""}
}

func GetImgTest(ctx iris.Context) iris.Map {
	cap := captcha.New()
	n := strconv.Itoa(rand.Intn(7) + 10)
	if err := cap.SetBgImg("./examples/target/" + n + ".jpg"); err != nil {
		fmt.Println(err)
	}
	if err := cap.SetBgImgLayer("./examples/target/hycdn.png"); err != nil {
		fmt.Println(err)
	}
	_, im, imSlide := cap.OutImgEncodeString()
	x, y := cap.GetXY()
	id := time.Now().Format("150405") + strconv.Itoa(y)
	captchaData[id] = x
	return iris.Map{"id": id, "y": y, "im": im, "imSlide": imSlide}
}
