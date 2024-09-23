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

	userRouter.Post("/add", handler.MiddlewareAuth(handler.CreateUserByAdmin))
	userRouter.Get("/{username}", handler.MiddlewareAuth(handler.GetAnotherUser))
	userRouter.Delete("/{username}", handler.MiddlewareAuth(handler.DbDeleteUser))
	userRouter.Get("/list", handler.MiddlewareAuth(handler.GetUserList))
	userRouter.Put("/{user_id}", handler.MiddlewareAuth(handler.UpdateAnotherUser))

	// Bank
	bankRouter := chi.NewRouter()
	bankRouter.Post("/register", handler.MiddlewareAuth(handler.CreateBank))
	bankRouter.Get("/list", handler.GetBankList)
	bankRouter.Put("/{bank_id}", handler.MiddlewareAuth(handler.UpdateBank))

	// Bank Account
	accountRouter := chi.NewRouter()
	accountRouter.Post("/", handler.MiddlewareAuth(handler.CreateBank))
	accountRouter.Get("/", handler.MiddlewareAuth(handler.GetAllAccounts))
	accountRouter.Get("/{account_id}", handler.MiddlewareAuth(handler.GetAccount))
	accountRouter.Put("/{account_id}", handler.MiddlewareAuth(handler.UpdateAccount))

	// Cards
	cardRouter := chi.NewRouter()
	cardRouter.Post("/", handler.MiddlewareAuth(handler.AddCard))
	cardRouter.Get("/", handler.MiddlewareAuth(handler.GetAllCards))
	cardRouter.Get("/{card_id}", handler.MiddlewareAuth(handler.GetCard))
	cardRouter.Put("/{card_id}", handler.MiddlewareAuth(handler.UpdateCard))

	v1Router.Mount("/users", userRouter)
	v1Router.Mount("/banks", bankRouter)
	v1Router.Mount("/accounts", accountRouter)
	v1Router.Mount("/cards", cardRouter)

	app.Mount("/v1", v1Router)
}
