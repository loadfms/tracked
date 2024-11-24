package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func RegisterCookieRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi, authorizer awsapigatewayv2authorizers.HttpLambdaAuthorizer) {
	//LAMBDAS
	createCookieLambda := awslambda.NewFunction(stack, jsii.String("CreateCookie"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/cookie/create-cookie/create-cookie.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	queryCookieLambda := awslambda.NewFunction(stack, jsii.String("QueryCookie"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/cookie/query-cookie/query-cookie.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createCookieIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreateCookieIntegration"), createCookieLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	queryCookieIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("QueryCookieIntegration"), queryCookieLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/cookie"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createCookieIntegration,
		Authorizer:  authorizer,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/cookie/{siteUUID}"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: queryCookieIntegration,
		Authorizer:  authorizer,
	})

	table.GrantFullAccess(createCookieLambda)
	table.GrantFullAccess(queryCookieLambda)
}
