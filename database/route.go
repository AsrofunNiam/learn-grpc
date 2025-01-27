package database

// import (
// 	"fmt"
// 	"runtime/debug"

// 	"github.com/gin-gonic/gin"
// 	"google.golang.org/grpc"
// 	"gorm.io/gorm"
// )

// func NewRouter(db *gorm.DB,grpc *grpc.ClientConn , ) *gin.Engine {

// 	router := gin.New()

// 	//  exception middleware
// 	router.Use(ErrorHandler())
// 	router.UseRawPath = true

// 	// route path
// 	route.UserRoute(router, db, validate)
// 	route.ProductRoute(router, db, redisClient, validate)

// 	route.ProductRoute(router, db, redisClient, validate)
// 	// route.TransactionRoute(router, db, validate)

// 	return router
// }
