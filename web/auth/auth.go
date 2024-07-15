package auth

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/Outtech105k/GPS-Reminder-Server/web/response"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewJWTMiddleware(db *sql.DB) (*jwt.GinJWTMiddleware, error) {
	jwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "reminder",
		Key:        []byte(os.Getenv("JWT_KEY")),
		Timeout:    time.Hour * 3,
		MaxRefresh: time.Hour * 24 * 7,
		SendCookie: false,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			return jwt.MapClaims{
				jwt.IdentityKey: data,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var user AccountRequest

			if err := c.ShouldBindJSON(&user); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if !user.isValid(db) {
				return "", jwt.ErrFailedAuthentication
			}

			return user.Username, nil
		},
	})

	if err != nil {
		return nil, err
	}

	err = jwtMiddleware.MiddlewareInit()

	if err != nil {
		return nil, err
	}

	return jwtMiddleware, nil
}

type AccountRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user AccountRequest) isValid(db *sql.DB) bool {
	// パスワード検証
	var storedPassHash string
	if err := db.QueryRow("SELECT hashed_pass FROM users WHERE name=?", user.Username).Scan(&storedPassHash); err != nil {
		return false
	}

	// ハッシュパスワードの照合
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassHash), []byte(user.Password)); err != nil {
		return false
	}

	return true
}

func GetUsernameInJWT(ctx *gin.Context) string {
	claims := jwt.ExtractClaims(ctx)
	return claims[jwt.IdentityKey].(string)
}
