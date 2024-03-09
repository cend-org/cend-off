package route

import (
	"duval/internal/configuration"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	engine.Use(gin.Recovery())
}

func Serve() (err error) {
	err = attach(engine)
	if err != nil {
		panic(err)
	}

	err = engine.Run(configuration.App.Host + ":" + configuration.App.Port)
	if err != nil {
		panic(err)
	}

	return err
}
