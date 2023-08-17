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

		perm, ok := claims["Perm"].(string)
		if !ok {
			c.Status(500)
			log.Println("Erro ao obter id do usuario a partir do token JWT", err)
			return
		}

		if perm == "user" && c.Request.URL.Path == "/api/users" || c.Request.URL.Path == "/api/delete-user" || c.Request.URL.Path == "/api/atualizar-user" || c.Request.URL.Path == "/api/mails" || c.Request.URL.Path == "/api/mail-user" {
			c.JSON(401, gin.H{
				"Acesso negado": "Você não tem permissão para acessar esta rota.",
			})
			c.Abort()
		}

		IDUser := id
		c.Set("id", IDUser)
		c.Next()
	}
}

//crud usuario apenas se for admin
