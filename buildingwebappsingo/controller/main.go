package main

import (
	"net/http"
	"gopkg.in/unrolled/render.v1"
	"github.com/wagnergaldino/learning/buildingwebappsingo/controller/basecontroller"
)

type MyController struct {
	basecontroller.AppController
	*render.Render
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8080", c.Action(c.Index))
}
