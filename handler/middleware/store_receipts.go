package middleware

import "github.com/gin-gonic/gin"

func StoreReceipts() gin.HandlerFunc {
	return func(c *gin.Context) {
		// time.Sleep(6 * time.Second)
		c.Next()
	}
}
