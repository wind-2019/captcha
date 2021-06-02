package captcha

import (
	"bytes"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
	"math/rand"
	"time"
)

//图片验证码
type captcha struct {
	imTmpBg    image.Image   //最初的完整背景
	imTmpSlide image.Image   //最初的动的焦点图片
	imBg       *image.RGBA64 //要滑动的背景
	imSlide    *image.RGBA64 //要滑动的焦点图片
	bgWidth    int           //背景图宽
	bgHeight   int           //背景图高

	markWidth  int //标记宽
	markHeight int //标记高

	x     int //x轴 像素点
	y     int //y轴 像素点
	fault int //容错象素
}

//Version 版本
func (c *captcha) Version() string {
	return "1.0.0"
}

//Check  图片校验参考示例
func (c *captcha) Check(storeX, offset int) (result bool) {
	if storeX == 0 {
		storeX = c.x
	}
	if (storeX - offset) <= c.fault {
		result = true
	}
	return result
}

//GetXY 获取图片x和y坐标
func (c *captcha) GetXY() (int, int) {
	return c.x, c.y
}

//OutImg 输出图片实体
func (c *captcha) OutImg() (image.Image, image.Image) {
	c.createImage()
	defer c.destroy()
	return c.imBg, c.imSlide
}

//OutImgBytes 输出图片字节码
func (c *captcha) OutImgBytes() (error, []byte, []byte) {
	im, imSlide := c.OutImg()
	bufIm := new(bytes.Buffer)
	err := png.Encode(bufIm, im)

	bufImSlide := new(bytes.Buffer)
	err = png.Encode(bufImSlide, imSlide)

	return err, bufIm.Bytes(), bufImSlide.Bytes()
}

//OutImgEncodeString 输出图片编码字符串
func (c *captcha) OutImgEncodeString() (error, string, string) {
	err, bufIm, bufImSlide := c.OutImgBytes()
	return err, base64.StdEncoding.EncodeToString(bufIm), base64.StdEncoding.EncodeToString(bufImSlide)
}

// SetBgImg 设置背景图片
// bgImgPath 背景图片路径  例/bg/1.png
func (c *captcha) SetBgImg(bgImgPath string) (err error) {
	c.imTmpBg, err = imaging.Open(bgImgPath)
	if err != nil {
		return err
	}
	c.bgWidth = c.imTmpBg.Bounds().Max.X
	c.bgHeight = c.imTmpBg.Bounds().Max.Y
	return err
}

// SetBgImgLayer 背景图片设置灰色焦点框图层
// markImgPath   背景图片灰色焦点框浮层图片的路径  例/img/mark.png
func (c *captcha) SetBgImgLayer(markImgPath string) (err error) {
	c.imTmpSlide, err = imaging.Open(markImgPath)
	if err != nil {
		return err
	}
	c.markWidth = c.imTmpSlide.Bounds().Max.X  //横轴为X轴
	c.markHeight = c.imTmpSlide.Bounds().Max.Y //纵轴为Y轴
	c.x = randInt(c.markWidth, c.bgWidth-c.markWidth)
	c.y = randInt(c.markHeight, c.bgHeight-c.markHeight)
	return err
}

//指定随机数
func randInt(min int, max int) int {
	if min >= max || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

//New 创建图片二维码对象
//fault 容错值
func New() *captcha {
	c := &captcha{}
	return c
}
