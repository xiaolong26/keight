package handle

import (
	"github.com/gin-gonic/gin"
	"k8s/jwt"
)

type Resource struct {
	Name	string
	Namespace	string
	User	string
 	Kind 	string
}

func GenerateCommonParams(c *gin.Context)*Resource {
	cliam,_ := jwt.ParseToken(c.Query("token"))
	user := cliam.Username
	commonParams := &Resource{
		Name:     c.Query("name"),
		Namespace: c.Query("namespace"),
		User:      user,
		Kind:      c.Query("kind"),
	}
	return commonParams
}