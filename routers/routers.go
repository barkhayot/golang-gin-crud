package routers

import	(
	"github.com/gin-gonic/gin"
	"editt/handlers"
	"editt/database"	
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	/* routers */
	r.GET("", api.GetPosts)
	r.GET("/:id", api.GetPost)
	r.POST("", api.CreatePost)
	r.PUT("/:id", api.UpdatePost)
	r.DELETE("/:id", api.DeletePost)

	return r

}