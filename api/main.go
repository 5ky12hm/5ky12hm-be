package main

import (
	"api/router"

	"github.com/gin-gonic/gin"
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

	if err := router.SetupRouter(
		gin.Default(),
	).Run(":8080"); err != nil {
		panic(err)
	}
}
