package main

import (
	"github.com/anon/go-api/gin"
	"github.com/anon/go-api/internal"
	"github.com/anon/go-api/mariadb"
)

func main() {
	env := internal.GetEnvironment()
	dbClient := internal.GetDbClient(&env)

	us := &mariadb.UserService{Db: dbClient}

	var h gin.Handler
	h.UserService = us

	h.ServeHttp(&env)
}
