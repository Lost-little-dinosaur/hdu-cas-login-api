package loginServiceRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/wujunyi792/crispy-waffle-be/internal/handle/loginHadle"
)

func InitLoginServiceRouter(e *gin.Engine) {
	baseGroup := e.Group("/loginToken")
	{
		baseGroup.POST("/ihdu", loginHadle.HandleGetIHduToken)
		baseGroup.POST("/newjw", loginHadle.HandleGetNewJWToken)
		baseGroup.POST("/skl", loginHadle.HandleGetSKLToken)

	}
}
