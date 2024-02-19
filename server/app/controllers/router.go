package controllers

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/heatmap-creator/server/app/graph/generated"
	"github.com/m-butterfield/heatmap-creator/server/app/graph/resolvers"
	"github.com/m-butterfield/heatmap-creator/server/app/lib"
	"github.com/m-butterfield/heatmap-creator/server/app/static"
	"net/http"
	"os"
	"path"
)

func router() (*gin.Engine, error) {
	r, err := lib.BaseRouter()
	if err != nil {
		return nil, err
	}

	staticFS := http.FS(static.FS{})
	fileServer := http.FileServer(staticFS)
	addStaticHandler(r, "/js", fileServer)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.Use(auth)
	r.NoRoute(index)
	r.GET("/chat", chat)
	r.GET("/chat/info", info)

	graphql := r.Group("/graphql")
	graphql.Use(ginContextToContextMiddleware)

	graphql.POST("/query", makeGraphQLHandler())
	if os.Getenv("GQL_PLAYGROUND") != "" {
		graphql.GET("/", makePlayGroundHandler())
	}

	return r, nil
}

func addStaticHandler(r *gin.Engine, prefix string, fileServer http.Handler) {
	h := func(c *gin.Context) { fileServer.ServeHTTP(c.Writer, c.Request) }
	pattern := path.Join(prefix, "/*filepath")
	r.GET(pattern, h)
	r.HEAD(pattern, h)
}

func makePlayGroundHandler() func(*gin.Context) {
	playgroundHandler := playground.Handler("GraphQL", "/graphql/query")
	return func(c *gin.Context) {
		playgroundHandler.ServeHTTP(c.Writer, c.Request)
	}
}

func makeGraphQLHandler() func(*gin.Context) {
	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{
					DS: ds,
				},
			},
		),
	)
	return func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	}
}

func ginContextToContextMiddleware(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), lib.GinContextKey, c)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
