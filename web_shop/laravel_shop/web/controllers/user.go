package controllers

import (
	"github.com/kataras/iris"
	"../../models"
	"../../services"
	"github.com/kataras/iris/mvc"
	"golang.org/x/crypto/bcrypt"
	"log"
)


// UserController is our /user controller.
// UserController is responsible to handle the following requests:
// GET  			/user/register
// POST 			/user/register
// GET 				/user/login
// POST 			/user/login
// GET 				/user/me
// All HTTP Methods /user/logout
type UserController struct {
	Ctx iris.Context
	Service services.UserService
	// Session, binded using dependency injection from the main.go.
	//Session *sessions.Session
}


//获取网页上登录的用户
func (c *UserController) getCurrentUserID() []models.User{
	id,err := c.Ctx.URLParamInt64("id")
	if err != nil {
		return nil
	}
	userID := c.Service.Get(id)
	return userID
}

//获取的用户为空，未登录
func (c *UserController) isLoggedIn() bool {
	return c.getCurrentUserID() != nil
}

//退出登录
func (c *UserController) logout() {
	c.Service.Logout()
}

//注册页面
var registerStaticView = mvc.View{
	Name: "views/register.html",
	Data: iris.Map{"Title": "User Registration"},
}


// GetRegister handles GET: http://localhost:8080/user/register.
func (c *UserController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		c.logout()
	}
	return registerStaticView
}

// PostRegister handles POST: http://localhost:8080/user/register.
func (c *UserController) PostRegister() mvc.Result {
	// get firstname, username and password from the form.
	var (
		username  = c.Ctx.FormValue("username")
		password  = c.Ctx.FormValue("password")
		email  = c.Ctx.FormValue("email")
		active = true
		hashpwd,err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	)

	data := models.User{
		Name:username,
		Pwd: password,
		Email:email,
		Active:active,
		HashedPassword:hashpwd,
	}
	// create the new user, the password will be hashed by the service.
	err = c.Service.Create(&data)
	if err != nil {
		log.Panic("注册失败")
	}

	//注册成功跳转个人页面
	return mvc.Response{
		// if not nil then this error will be shown instead.
		Err: err,
		// redirect to /user/me.
		Path: "/user/me",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		// Code: 303,
	}
}

var loginStaticView = mvc.View{
	Name: "user/login.html",
	Data: iris.Map{"Title": "User Login"},
}

// GetLogin handles GET: http://localhost:8080/user/login.
func (c *UserController) GetLogin() mvc.Result {
	if c.isLoggedIn() {
		// if it's already logged in then destroy the previous session.
		c.logout()
	}

	return loginStaticView
}

// PostLogin handles POST: http://localhost:8080/user/register.
func (c *UserController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	found := c.Service.GetByUsernameAndPassword(username, password)

	if !found {
		return mvc.Response{
			Path: "/user/register",
		}
	}

	return mvc.Response{
		Path: "/user/user_center",
	}
}

// GetMe handles GET: http://localhost:8080/user/me.
func (c *UserController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		// if it's not logged in then redirect user to the login page.
		return mvc.Response{Path: "/user/login"}
	}
	id,err := c.Ctx.URLParamInt64("id")
	if err != nil {
		log.Panic("未登录")
	}
	u, found := c.Service.GetByID(id)
	if !found {
		// if the  session exists but for some reason the user doesn't exist in the "database"
		// then logout and re-execute the function, it will redirect the client to the
		// /user/login page.
		c.logout()
		return c.GetMe()
	}

	return mvc.View{
		Name: "user/user_center.html",
		Data: iris.Map{
			"Title": "Profile of ",
			"User":  u,
		},
	}
}

// AnyLogout handles All/Any HTTP Methods for: http://localhost:8080/user/logout.
func (c *UserController) AnyLogout() {
	if c.isLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/user/login")
}


