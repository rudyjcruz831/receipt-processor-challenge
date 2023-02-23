package app

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rudyjcruz831/receipt-processor-challenge/handler"
)

// will initialize a handler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func inject() (*gin.Engine, error) {
	log.Println("Injecting data sources")

	/*
	 * repository layer
	 */

	/*
	 * service layer
	 */
	// userService := services.NewUserService(&services.USConfig{
	// 	UserRepository:  userRepository,
	// 	ImageRepository: imageRepository,
	// 	DocsRepository:  DocsRepository,
	// })

	// initialize gin.Engine
	router := gin.Default()

	// read in project baseURL from environment variable
	baseURL := os.Getenv("BASE_URL")

	//
	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	fmt.Println(handlerTimeout)
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	fmt.Println(ht)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}

	handler.NewHandler(&handler.Config{
		R:               router,
		BaseURL:         baseURL,
		TimeoutDuration: time.Duration(time.Duration(ht) * time.Second),
		MaxBodyBytes:    1024 * 1024 * 1024,
	})

	return router, nil
}
