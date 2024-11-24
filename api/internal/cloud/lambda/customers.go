package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func RegisterCustomerRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi) {
	//LAMBDAS
	createCustomerLambda := awslambda.NewFunction(stack, jsii.String("CreateCustomer"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/customer/create-customer/create-customer.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	loginCustomerLambda := awslambda.NewFunction(stack, jsii.String("LoginCustomer"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/customer/login-customer/login-customer.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createCustomerIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreateCustomerIntegration"), createCustomerLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})
	loginCustomerIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("LoginCustomerIntegration"), loginCustomerLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/customer"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createCustomerIntegration,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/login"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: loginCustomerIntegration,
	})

	table.GrantFullAccess(createCustomerLambda)
	table.GrantFullAccess(loginCustomerLambda)
}
