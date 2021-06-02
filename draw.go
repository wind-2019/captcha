package captcha

// 根据原 图片画图层
import (
	"image"
	"image/color"
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
	c.imBg = image.NewRGBA64(c.imTmpBg.Bounds())
	c.imSlide = image.NewRGBA64(c.imTmpSlide.Bounds())
	for x := 0; x <= c.bgWidth; x++ {
		for y := 0; y <= c.bgHeight; y++ {
			at := c.imTmpBg.At(x, y)
			if (x >= c.x && x <= c.x+c.markWidth) && (y >= c.y && y <= c.y+c.markHeight) {
				r, g, b, a := at.RGBA()
				opacity := uint16(float64(a) * 15.5)
				v := c.imBg.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity})
				rr, gg, bb, aa := v.RGBA()
				c.imBg.SetRGBA64(x, y, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
				c.imSlide.Set(x-c.x, y-c.y, at)
			} else {
				c.imBg.Set(x, y, at)
			}
		}
	}
	return
}
