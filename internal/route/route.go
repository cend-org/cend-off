package route

import (
	"duval/internal/configuration"
	"duval/internal/route/docs"
	"github.com/gin-gonic/gin"
	"net/http"
)

var engine *gin.Engine

func init() {
	if configuration.App.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	//engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())

}

func Serve() (err error) {
	err = attach(engine)
	if err != nil {
		panic(err)
	}

	engine.Routes()

	err = engine.Run(configuration.App.Host + ":" + configuration.App.Port)
	if err != nil {
		panic(err)
	}

	return err
}

func attach(g *gin.Engine) (err error) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, docs.ParseDocumentation(RootRoutesGroup))
	})

	for i := 0; i < len(RootRoutesGroup); i++ {
		group := g.Group(RootRoutesGroup[i].Group)
		err = docs.GenerateDocumentation(group, RootRoutesGroup[i].Paths)
		if err != nil {
			panic(err)
		}
	}

	return err
}
