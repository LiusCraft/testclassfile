package main

import "github.com/liuscraft/lss"

type home struct {
	lss.App
}
//line demo/gox/home_lss.gox:1
func (this *home) MainEntry() {
//line demo/gox/home_lss.gox:1:1
	this.Run("hello world")
}
func main() {
	lss.Gopt_App_Main(new(home))
}
