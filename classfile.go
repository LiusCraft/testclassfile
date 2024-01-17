/*
This is an example of a go+classfile
go+ version: 1.2.0.pre.1
*/
package lss

import "fmt"

var (
	GopPackage = true
)

type Apper interface {
	initApp()
}

type App struct {
	*Lss
}

func (p *App) initApp() {
	p.Lss = &Lss{}
}

func Gopt_App_Main(app Apper) {
	app.initApp()
	app.(interface{ MainEntry() }).MainEntry()
}

func (p *App) Run(name string) {
	fmt.Println(name)
}
