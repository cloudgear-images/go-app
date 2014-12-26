package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"os"
)

type Info struct{}

func (i *Info) Hostname() string {
	name, err := os.Hostname()
	if err != nil {
		name = ""
	}
	return name
}

func (i *Info) Env() []string {
	return os.Environ()
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Layout: "layout"}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", &Info{})
	})
	m.RunOnAddr(":5000")
}
