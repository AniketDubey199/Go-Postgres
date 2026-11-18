package middleware

import (
	"net/http"

	"github.com/AniketDubey199/PostgreSQL/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Handlercalling(mux *http.ServeMux, pool *pgxpool.Pool) {
	userhandler := database.NewHandler(pool)
	mux.HandleFunc("POST /createuser", userhandler.CreateUser)
	mux.HandleFunc("PUT /updateuser", userhandler.UpdateUser)
	mux.HandleFunc("GET /getuser", userhandler.GetUser)
	mux.HandleFunc("DELETE /deleteuser", userhandler.DeleteUser)
	mux.HandleFunc("POST /createproduct", userhandler.CreateProduct)
}
