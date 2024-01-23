package runner

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginAdapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

type LambdaRunner struct {
	adapter *ginAdapter.GinLambdaV2
}

func (r LambdaRunner) Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return r.adapter.ProxyWithContext(ctx, req)
}

func (r LambdaRunner) Run() {
	lambda.Start(r.Handler)
}

func NewLambdaRunner(g *gin.Engine) *LambdaRunner {
	return &LambdaRunner{
		adapter: ginAdapter.NewV2(g),
	}
}
