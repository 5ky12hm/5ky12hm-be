package main

import (
	"api/router"
	"api/utility"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var ginLambda *ginadapter.GinLambda
//
// func init() {
// 	ginLambda = ginadapter.New(
// 		router.SetupRouter(gin.Default()),
// 	)
// }
//
// func Handler(c context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	return ginLambda.ProxyWithContext(c, req)
// }

func main() {
	// lambda.Start(Handler)

	if err := godotenv.Load(
		fmt.Sprintf(".env.%s", os.Getenv(utility.EnvKeyApiEnv)),
	); err != nil {
		panic(err)
	}
	if err := router.SetupRouter(
		gin.Default(),
	).Run(":8080"); err != nil {
		panic(err)
	}
}
