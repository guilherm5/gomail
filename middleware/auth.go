package middleware

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func autentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Erro ao carregar variaveis de ambiente", err)
			c.Status(500)
			return
		}
		secret := os.Getenv("TOKEN")
		authorization := c.GetHeader("Authorization")

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

		IDUser := id
		c.Set("id", IDUser)
		c.Next()
	}
}
