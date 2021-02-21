package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

func confirmUser(ctx context.Context, cognitoSvc *cognitoidentityprovider.Client, userPoolID string, email string) {
	_, _ = cognitoSvc.AdminConfirmSignUp(ctx, &cognitoidentityprovider.AdminConfirmSignUpInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(email),
	})
}

func getConsumerID(attrs []types.AttributeType) *string {
	for _, v := range attrs {
		if *v.Name == "sub" {
			return v.Value
		}
	}
	return nil
}

func checkUserExists(ctx context.Context, cognitoSvc *cognitoidentityprovider.Client, userPoolID string, email string) *string {
	uo, err := cognitoSvc.AdminGetUser(ctx, &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(email),
	})
	if err != nil {
		return nil
	}
	confirmUser(ctx, cognitoSvc, userPoolID, email)
	return getConsumerID(uo.UserAttributes)
}