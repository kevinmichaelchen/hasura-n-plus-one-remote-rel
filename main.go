package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/gin-gonic/gin"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/graphql"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/graphql/generated"
)

func main() {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graphql.Resolver{},
			},
		),
	)

	r := gin.Default()

	r.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}
