---
page_title: "cognito_user Resource - terraform-provider-cognito"
subcategory: "AWS"
description: |-
  Sample user resource in the Terraform provider cognito.
---

# Resource `cognito_user`

Sample resource in the Terraform provider scaffolding.

## Example Usage

```terraform
resource "cognito_user" "test_user" {
  email    = "joshrose@hey.com"
  password = "Abcd1234!"
}
```

## Schema

### Required

- **email** (String) email to register user into cognito user pool with
- **password** (String) password for the new cognito user


