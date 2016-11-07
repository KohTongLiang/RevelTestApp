package controllers

import (
	"github.com/revel/revel"
    "github.com/myaccount/crudapp/app/models"
    "golang.org/x/crypto/bcrypt"
	)

type App struct {
    GorpController
}

func (c App) Index() revel.Result {
    if c.Session["user"] != "" {
        return c.Render()
    } else {
        return c.Redirect(App.IndexAlt)
    }
}

func (c App) IndexAlt() revel.Result {
    return c.Render()
}

func (c App) Register() revel.Result {
	return c.Render()
}

func (c App) RegisterAcc(username string, password string) revel.Result {
	//user := new(models.User)
	/*rows, err := c.Txn.SelectOne(user, 
        `SELECT username, pasword FROM User WHERE username = ?`, username)
    if rows == 0 {*/
    	//username available
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return c.RenderError(err)
	}

    var user = models.User{
    	Username: username,
    	Password: hashedPassword,
    }

    if err := c.Txn.Insert(&user); err != nil {
        return c.RenderError(err)
    } else {
        return c.Redirect(App.Login)
    }
	return c.Render()
    /*} else {
    	return c.RenderText("Error. Unable to create account.")
    }*/
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) LoginAcc(username string, password string) revel.Result {
	user := new(models.User)
	err := c.Txn.SelectOne(user, 
        `SELECT username, password FROM User WHERE username = ?`, username)
    if err != nil {
        return c.RenderError(err)
    }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
        return c.RenderText("Error. Wrong password or username.")
    }

    c.Session["user"] = username

    return c.Redirect(App.Index)
}

func (c App) Update(id int64) revel.Result {
	return c.Render(id)
}

func (c App) Logout() revel.Result {
    delete(c.Session, "user")
    return c.Redirect(App.Index)
}