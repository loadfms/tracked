package main

import (
	"tracked/internal/cloud/dynamodb"
	"tracked/internal/cloud/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TrackedStackProps struct {
	awscdk.StackProps
}

func TrackedStack(scope constructs.Construct, id string, props *TrackedStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//HTTP API
	httpApi := awsapigatewayv2.NewHttpApi(stack, jsii.String("TrackedApi"), &awsapigatewayv2.HttpApiProps{
		ApiName:     jsii.String("TrackedApi"),
		Description: jsii.String("Tracked API"),
	})

	//DYNAMO
	table := dynamodb.CreateDynamo(stack)

	//AUTHORIZER
	authorizer := lambda.GetRouteAuthorizers(stack)

	//ROUTES
	lambda.RegisterCustomerRoutes(stack, table, httpApi)
	lambda.RegisterWorkspaceRoutes(stack, table, httpApi, *authorizer)
	lambda.RegisterSiteRoutes(stack, table, httpApi, *authorizer)
	lambda.RegisterPrivacyPolicyRoutes(stack, table, httpApi, *authorizer)
	lambda.RegisterCookieRoutes(stack, table, httpApi, *authorizer)
	lambda.RegisterConsentRoutes(stack, table, httpApi, *authorizer)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	TrackedStack(app, "TrackedStack", &TrackedStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	//return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		//Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region: jsii.String("sa-east-1"),
	}
}
