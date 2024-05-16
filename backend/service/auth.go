// service/auth.go

package service

import (
	"time"
	"github.com/barmin15/mobile-dating-app/backend/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//handles JWT token generation and authentication
type JWTService struct {
    jwtKey []byte
}

//initializes a new JWTService with the provided JWT secret key
func NewJWTService(secretKey string) *JWTService {
    return &JWTService{
        jwtKey: []byte(secretKey),
    }
}

//generates a JWT token for the provided user
func (s *JWTService) GenerateToken(user *model.User) (string, error) {
    // Create the token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = user.Email
    claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expires in 2 hours

    // Sign the token with the secret key
    tokenString, err := token.SignedString(s.jwtKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

//function to validate JWT token
func (s *JWTService) AuthenticateTokenMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.Set("error", "Unauthorized")
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Parse JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return s.jwtKey, nil
        })
        if err != nil || !token.Valid {
            c.Set("error", "Unauthorized")
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Token is valid, continue with the request
        c.Next()
    }
}

