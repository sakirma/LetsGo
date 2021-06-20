package controllers

import (
	"github.com/revel/revel"
)



type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	return c.RenderText("what the hell are you doing here?")
}


type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c App) Login(username string, password string) revel.Result {
	//content := credentials{}
	//err := c.Params.BindJSON(&content)
	//
	//if err != nil {
	//	c.Response.Status = http.StatusBadRequest
	//	return c.RenderText(err.Error())
	//}
	//
	//return c.RenderJSON(content)
	return c.RenderText(username + " " + password)
}