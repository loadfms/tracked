package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func GetRouteAuthorizers(stack awscdk.Stack) *awsapigatewayv2authorizers.HttpLambdaAuthorizer {
	//LAMBDAS
	authorizerLambda := awslambda.NewFunction(stack, jsii.String("Authorizer"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/authorizer/authorizer.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	authorizer := awsapigatewayv2authorizers.NewHttpLambdaAuthorizer(jsii.String("Authorizer"), authorizerLambda, &awsapigatewayv2authorizers.HttpLambdaAuthorizerProps{})

	return &authorizer
}
