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

func RegisterPrivacyPolicyRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi, authorizer awsapigatewayv2authorizers.HttpLambdaAuthorizer) {
	//LAMBDAS
	createPrivacyPolicyLambda := awslambda.NewFunction(stack, jsii.String("CreatePrivacyPolicy"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/privacypolicy/create-privacypolicy/create-privacypolicy.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	queryPrivacyPolicyLambda := awslambda.NewFunction(stack, jsii.String("QueryPrivacyPolicy"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/privacypolicy/query-privacypolicy/query-privacypolicy.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createPrivacyPolicyIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreatePrivacyPolicyIntegration"), createPrivacyPolicyLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	queryPrivacyPolicyIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("QueryPrivacyPolicyIntegration"), queryPrivacyPolicyLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/privacypolicy"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createPrivacyPolicyIntegration,
		Authorizer:  authorizer,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/privacypolicy/{siteUUID}"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: queryPrivacyPolicyIntegration,
		Authorizer:  authorizer,
	})

	table.GrantFullAccess(createPrivacyPolicyLambda)
	table.GrantFullAccess(queryPrivacyPolicyLambda)
}
