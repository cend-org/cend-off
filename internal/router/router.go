package router

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cend-org/duval/graph/generated"
	"github.com/cend-org/duval/graph/resolver"
	"github.com/cend-org/duval/internal/router/cors"
	"github.com/cend-org/duval/internal/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
)

func Serve() {
	r := gin.Default()
	r.Use(middleware.Middleware())
	r.Use(cors.Set())
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	ExtendRoute(r)
	r.POST("/query", graphqlHandler())
	r.GET("/playground", playgroundHandler())
	err := r.Run(":8087")
	if err != nil {
		panic(err)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	var mb int64 = 1 << 20
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})
	h.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})
	return func(c *gin.Context) {
		c.Writer.Header().Set("test", "test")
		h.ServeHTTP(c.Writer, c.Request)
	}

}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
