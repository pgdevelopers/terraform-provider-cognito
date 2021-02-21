---
page_title: "Cognito Provider"
subcategory: "aws"
description: |-
  This provider is designed to help facilitate creation of cognito users that are needed for integration tests
---

# Cognito Provider

## Example Usage

```terraform
provider "cognito" {
  user_pool_id = "us-east-1_ABc123" # user pool id from cognito
  client_id    = "abc1234" # application client id from cognito
}
```

## Schema
### Required

- **user_pool_id** (String) user pool id from cognito
- **client_id** (String) application client id from cognito

### Optional
- **region** (String, Optional) the aws account deploy region, default to us-east-1