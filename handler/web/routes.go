package web

func (me *Web) Routes() {
	me.Router.LoadHTMLGlob("views/*")
	me.Router.Static("/assets", "./assets")
	me.Router.GET("/", me.Home)
	me.Router.GET("/login", me.Login)
	me.Router.GET("/logout", me.Logout)
	me.Router.GET("/help", me.Help)
	me.Router.GET("/done", me.Done)
}
