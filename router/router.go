package router

import (
	"github.com/bytesByHarsh/go-my-info/handler"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *chi.Mux) {
	handler.Init()
	// Middleware
	v1Router := chi.NewRouter()
	v1Router.Get("/", handler.Hello)

	// User
	userRouter := chi.NewRouter()
	userRouter.Post("/register", handler.CreateUser)
	userRouter.Post("/login", handler.LoginUser)

	userRouter.Get("/me", handler.MiddlewareAuth(handler.GetUser))
	userRouter.Put("/me", handler.MiddlewareAuth(handler.UpdateUser))
	userRouter.Delete("/me", handler.MiddlewareAuth(handler.DeleteUser))
	userRouter.Put("/me/password", handler.MiddlewareAuth(handler.UpdateUserPassword))

	userRouter.Get("/{username}", handler.MiddlewareAuth(handler.GetAnotherUser))
	userRouter.Delete("/{username}", handler.MiddlewareAuth(handler.DbDeleteUser))

	v1Router.Mount("/user", userRouter)
	app.Mount("/v1", v1Router)
}
