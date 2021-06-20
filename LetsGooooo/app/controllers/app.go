package controllers

import (
	"LetsGooooo/app"
	"LetsGooooo/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type App struct {
	*revel.Controller
}

func init() {
	revel.FilterController(App{}).Insert(authFilter, revel.BEFORE, revel.ActionInvoker)
}

func (c App) Index() revel.Result {

	return c.RenderText("what the hell are you doing here?")
}

func (c App) Register(username string, password string) revel.Result {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		revel.AppLog.Error("Couldn't hash password", err)
		c.Response.Status = http.StatusInternalServerError
		return c.Render()
	}

	user := models.User{}
	app.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		user = models.User{Username: username, Password: hashedPassword}
		app.DB.Create(&user)

		c.Response.Status = http.StatusCreated
		return c.RenderJSON(user)
	}

	c.Response.Status = http.StatusBadRequest
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	return c.RenderText("User already exists by that name" ,  err)
}

var authFilter = func(c *revel.Controller, fc[] revel.Filter) {
	revel.AppLog.Info("Hi there")

	fc[0](c, fc[1:])  // Execute the next filter stage.
}