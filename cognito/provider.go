package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", "us-east-1"),
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_COGNITO_CLIENT_ID", nil),
			},
			"user_pool_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AWS_COGNITO_USER_POOL_ID", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cognito_user": dataSourceUser(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"cognito_user": resourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type apiClient struct {
	Client     *cognitoidentityprovider.Client
	UserPoolID *string
	ClientID   *string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region := d.Get("region").(string)
	clientID := d.Get("client_id").(string)
	userPoolID := d.Get("user_pool_id").(string)

	cfg, err := config.LoadDefaultConfig(ctx, config.WithDefaultRegion(region))
	if err != nil {
		return nil, diag.Errorf("[ERROR] Error initializing the Cognito SDK clients: %v", err)
	}
	return apiClient{
		Client:     cognitoidentityprovider.NewFromConfig(cfg),
		UserPoolID: aws.String(userPoolID),
		ClientID:   aws.String(clientID),
	}, nil
}
