package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"os"
	"runtime"
)

type Info struct {
	Hostname   string   `json:"hostname"`
	Env        []string `json:"env"`
	NumCPU     int      `json:"num_cpu"`
	GoMaxProcs int      `json:"go_max_procs"`
}

func NewInfo() *Info {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = ""
	}

	return &Info{
		Hostname:   hostname,
		Env:        os.Environ(),
		NumCPU:     runtime.NumCPU(),
		GoMaxProcs: runtime.GOMAXPROCS(0),
	}
}

func getEnv(key, defaultValue string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return defaultValue
}

func main() {
	port := getEnv("PORT", "5000")
	info := NewInfo()

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{Layout: "layout"}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", info)
	})
	m.Get("/json", func(r render.Render) {
		r.JSON(200, info)
	})

	m.RunOnAddr(":"+port)
}
