package middleware

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Autentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("TOKEN")
		if secret == "" {
			c.AbortWithStatus(401)
		}

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			c.AbortWithStatus(401)
		}

		auth, err := jwt.Parse(authorization, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !auth.Valid {
			log.Println("Token inválido", err)
			c.AbortWithStatus(401)
			return
		}

		claims, ok := auth.Claims.(jwt.MapClaims)
		if !ok {
			c.Status(400)
			log.Println("Erro ao obter claims", err)
			return
		}

		id, ok := claims["ID"].(float64)
		if !ok {
			c.Status(500)
			log.Println("Erro ao obter id do usuario a partir do token JWT", err)
			return
		}

		/*permissoes, ok := claims["Perm"].(string)
		if !ok {
			c.Status(500)
			log.Println("Erro ao obter id do usuario a partir do token JWT", err)
			return
		}*/

		IDUser := id
		c.Set("id", IDUser)
		c.Next()
	}
}
