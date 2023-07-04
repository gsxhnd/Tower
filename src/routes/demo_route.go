package routes

func (r *Routes) newDemoRoute() {
	demoRoutes := r.Engine.Group("/demo")
	demoRoutes.GET("", r.DemoHandle.Ping)
}
