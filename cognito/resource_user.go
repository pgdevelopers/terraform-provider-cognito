package cognito

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceUserRead,
		CreateContext: resourceUserCreate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"consumer_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
		},
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	uid := *getConsumerID(uo.UserAttributes)
	d.SetId(uid)
	err = d.Set("consumer_id", uid)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(apiClient).Client
	upid := m.(apiClient).UserPoolID
	cid := m.(apiClient).ClientID
	email := d.Get("email").(string)
	pwd := d.Get("password").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	uo, err := c.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: cid,             // The App Client ID of the UserPool
		Password: aws.String(pwd), // The Password of the user
		Username: aws.String(email),
	})
	if err != nil {
		return diag.FromErr(err)
	}
	confirmUser(ctx, c, *upid, email)
	err = d.Set("consumer_id", *uo.UserSub)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(*uo.UserSub)
	resourceUserRead(ctx, d, m)
	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(apiClient).Client
	upid := m.(apiClient).UserPoolID
	email := d.Get("email").(string)

	var diags diag.Diagnostics
	eid := checkUserExists(ctx, c, *upid, email)
	if eid != nil {
		_, err := c.AdminDeleteUser(ctx, &cognitoidentityprovider.AdminDeleteUserInput{
			UserPoolId: upid,
			Username:   aws.String(email),
		})
		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId("")

	return diags
}
