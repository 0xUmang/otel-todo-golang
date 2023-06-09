package Routes

import (
	"Backend/Controllers"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(otelgin.Middleware("todo"))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)
	
	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetAllTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodobyID)
		v1.PUT("todo/:id", Controllers.UpdateATodobyID)
		v1.DELETE("todo/:id", Controllers.DeleteATodobyID)
	}

	return r
}
