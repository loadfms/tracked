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

func RegisterConsentRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi, authorizer awsapigatewayv2authorizers.HttpLambdaAuthorizer) {
	//LAMBDAS
	createConsentLambda := awslambda.NewFunction(stack, jsii.String("CreateConsent"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/consent/create-consent/create-consent.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	queryConsentLambda := awslambda.NewFunction(stack, jsii.String("QueryConsent"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/consent/query-consent/query-consent.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createConsentIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreateConsentIntegration"), createConsentLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	queryConsentIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("QueryConsentIntegration"), queryConsentLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/consent"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createConsentIntegration,
		Authorizer:  authorizer,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/consent/{siteUUID}"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: queryConsentIntegration,
		Authorizer:  authorizer,
	})

	table.GrantFullAccess(createConsentLambda)
	table.GrantFullAccess(queryConsentLambda)
}
