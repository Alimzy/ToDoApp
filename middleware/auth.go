package middleware

 
import (
    "net/http"
    "os"
    "strings"
 
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "gotask/db"
    "gotask/models"
)
 
func RequireAuth(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
        return
    }
 
    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil || !token.Valid {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        return
    }
 
    // Extract the user ID from the token claims
    claims, _ := token.Claims.(jwt.MapClaims)
    userID := uint(claims["sub"].(float64))
 
    var user models.User
    if err := db.DB.First(&user, userID).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
        return
    }
 
    // Store user on context so handlers can access it
    c.Set("currentUser", user)
    c.Next()
}
