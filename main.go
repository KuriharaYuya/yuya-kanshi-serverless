package main

import (
	// "context"
	// "sync"

	"github.com/aws/aws-lambda-go/events"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
// func Handler(ctx context.Context, req Request) (Response, error) {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	// respをポインタ変数として定義
// 	var resp *Response
// 	go func() {
// 		resp = Gateway(req)
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	return *resp, nil
// }

func main() {
  // res, err :=	notionpkg.GetLogProperty()
	// fmt.Println(res)
	// fmt.Println(err)
	// lambda.Start(Handler)
}
