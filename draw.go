package captcha

// 根据原 图片画图层
import (
	"image"
	"image/draw"
)


// destroy 图片销毁
func (c *captcha) destroy() {
	c.imBg = nil
	c.imTmpBg = nil
	c.imSlide = nil
	c.imTmpSlide = nil
}

// 生成图片
func (c *captcha) createImage() {
	c.imSlide = image.NewRGBA64(c.imTmpSlide.Bounds())
	draw.Draw(c.imSlide, c.imSlide.Bounds(), c.imTmpSlide, image.ZP, draw.Over)
	for x := 0; x <= c.markWidth; x++ {
		for y := 0; y <= c.markHeight; y++ {
			r, g, b, a := c.imSlide.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && a != 0 {
				c.imSlide.Set(x, y, c.imTmpBg.At(c.x+x, c.y+y))
			}
		}
	}
	// 获得更换区域
	c.imBg = image.NewRGBA64(c.imTmpBg.Bounds())
	draw.Draw(c.imBg, c.imTmpBg.Bounds(), c.imTmpBg, image.ZP, draw.Over)
	rectangle := image.Rectangle{
		Min: image.Point{
			X: c.x,
			Y: c.y},
		Max: image.Point{
			X: c.x + c.markWidth,
			Y: c.y + c.markHeight}}
	// 获得更换区域
	draw.Draw(c.imBg, rectangle, c.imTmpSlide, image.ZP, draw.Over)

	return

}
