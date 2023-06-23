package routes

func (r *Routes) NewDemoRoute() {
	demoRoutes := r.Engine.Group("/demo")
	demoRoutes.GET("", r.DemoHandle.Ping)
}
