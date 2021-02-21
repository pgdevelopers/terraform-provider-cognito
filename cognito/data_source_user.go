package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource {
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateEmail,
			},
			"consumer_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsUUID,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(apiClient).Client
	upid := m.(apiClient).UserPoolID
	email := d.Get("email").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	uo, err := c.AdminGetUser(ctx, &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: upid,
		Username:   aws.String(email),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(email)
	err = d.Set("consumer_id",*getConsumerID(uo.UserAttributes))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}