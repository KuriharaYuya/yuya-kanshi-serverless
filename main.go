package main

import (
	// "context"
	// "sync"

	"context"
	"fmt"
	"os"
	"sync"

	"github.com/KuriharaYuya/yuya-kanshi-serverless/usecase"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req Request) (Response, error) {
	var wg sync.WaitGroup
	wg.Add(1)
	// respをポインタ変数として定義
	var resp *Response
	go func() {
		resp = Gateway(req)
		wg.Done()
	}()

	wg.Wait()
	return *resp, nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)

	}
}

func main() {
	debugMode := os.Getenv("DEBUG_MODE")
	if debugMode == "true" {
		fmt.Println("デバッグモードです")
		// run some tgt funcs below
		usecase.CheckDailyLog()

	} else {
		lambda.Start(Handler)
	}

}
