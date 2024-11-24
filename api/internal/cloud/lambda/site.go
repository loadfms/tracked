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

func RegisterSiteRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi, authorizer awsapigatewayv2authorizers.HttpLambdaAuthorizer) {
	//LAMBDAS
	createSiteLambda := awslambda.NewFunction(stack, jsii.String("CreateSite"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/site/create-site/create-site.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	querySiteLambda := awslambda.NewFunction(stack, jsii.String("QuerySite"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/site/query-site/query-site.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createSiteIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreateSiteIntegration"), createSiteLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	querySiteIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("QuerySiteIntegration"), querySiteLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/site"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createSiteIntegration,
		Authorizer:  authorizer,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/site/{workspaceUUID}"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: querySiteIntegration,
		Authorizer:  authorizer,
	})

	table.GrantFullAccess(createSiteLambda)
	table.GrantFullAccess(querySiteLambda)
}
