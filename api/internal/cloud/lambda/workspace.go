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

func RegisterWorkspaceRoutes(stack awscdk.Stack, table awsdynamodb.Table, httpApi awsapigatewayv2.HttpApi, authorizer awsapigatewayv2authorizers.HttpLambdaAuthorizer) {
	//LAMBDAS
	createWorkspaceLambda := awslambda.NewFunction(stack, jsii.String("CreateWorkspace"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/workspace/create-workspace/create-workspace.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	queryWorkspaceLambda := awslambda.NewFunction(stack, jsii.String("QueryWorkspace"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/workspace/query-workspace/query-workspace.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	createWorkspaceIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("CreateWorkspaceIntegration"), createWorkspaceLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	queryWorkspaceIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("QueryWorkspaceIntegration"), queryWorkspaceLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/workspace"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: createWorkspaceIntegration,
		Authorizer:  authorizer,
	})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/workspace"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: queryWorkspaceIntegration,
		Authorizer:  authorizer,
	})

	table.GrantFullAccess(createWorkspaceLambda)
	table.GrantFullAccess(queryWorkspaceLambda)
}
