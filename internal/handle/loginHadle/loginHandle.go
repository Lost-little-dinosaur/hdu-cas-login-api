package loginHadle

import (
	"github.com/gin-gonic/gin"
	serviceErr "github.com/wujunyi792/crispy-waffle-be/internal/dto/err"
	"github.com/wujunyi792/crispy-waffle-be/internal/dto/login"
	"github.com/wujunyi792/crispy-waffle-be/internal/middleware"
	hducashelper "github.com/wujunyi792/hdu-cas-helper"
	"strings"
)

func HandleGetIHduToken(c *gin.Context) {
	var req login.GetTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Fail(c, serviceErr.RequestErr)
		return
	}
	ticker := hducashelper.CasPasswordLogin(req.Username, req.Password) // 杭电 CAS 账号密码
	iHduLogin := hducashelper.IHduLogin(ticker)
	if iHduLogin.Error() != nil {
		middleware.FailWithCode(c, 40200, "登录iHDU失败")
	}
	tempStr := iHduLogin.GetCookie()
	//log.Println(tempStr)
	middleware.Success(c, login.GetIHduTokenResponse{
		TpUp: tempStr[6 : len(tempStr)-2],
	})
}

func HandleGetNewJWToken(c *gin.Context) {
	var req login.GetTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Fail(c, serviceErr.RequestErr)
		return
	}
	ticker := hducashelper.CasPasswordLogin(req.Username, req.Password) // 杭电 CAS 账号密码
	newJwLogin := hducashelper.NewJWLogin(ticker)
	if newJwLogin.Error() != nil {
		middleware.FailWithCode(c, 40201, "登录NewJW失败")
	}
	tempArr := strings.Split(newJwLogin.GetCookie(), "=")
	//log.Println(tempArr)
	if tempArr[0] == "JSESSIONID" {
		middleware.Success(c, login.GetNewJWTokenResponse{
			JSESSIONID: tempArr[1][0 : len(tempArr[1])-7],
			Route:      tempArr[2][0 : len(tempArr[2])-2],
		})
	} else {
		middleware.Success(c, login.GetNewJWTokenResponse{
			JSESSIONID: tempArr[1][0 : len(tempArr[1])-12],
			Route:      tempArr[2][0 : len(tempArr[2])-2],
		})
	}

}
func HandleGetSKLToken(c *gin.Context) {
	var req login.GetTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.Fail(c, serviceErr.RequestErr)
		return
	}
	ticker := hducashelper.CasPasswordLogin(req.Username, req.Password) // 杭电 CAS 账号密码
	sklLogin := hducashelper.SklLogin(ticker)
	if sklLogin.Error() != nil {
		middleware.FailWithCode(c, 40202, "登录SKL失败")
	}
	middleware.Success(c, login.GetSKLokenResponse{
		Token: sklLogin.GetToken(),
	})
}
