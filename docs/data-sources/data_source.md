---
page_title: "cognito_user data - terraform-provider-cognito"
subcategory: "AWS"
description: |-
Sample user data in the Terraform provider cognito.
---

# Data `cognito_user`

Sample data in the Terraform provider cognito.

## Example Usage

```terraform
data "cognito_user" "test_user" {
  email    = "joshrose@hey.com"
}
```

## Schema

### Required

- **email** (String) email to find the existing user with


