terraform {
  required_providers {
    cognito = {
      version = "0.1"
      source   = "github.com/joshuarose/cognito"
    }
  }
}

provider "cognito" {
  user_pool_id = "us-east-1_ABc123" # user pool id from cognito
  client_id    = "abc1234" # application client id from cognito
}

resource "cognito_user" "test_user" {
  email    = "joshrose@hey.com"
  password = "Abcd1234!"
}

data "cognito_user" "test_user" {
  email    = "joshrose+exists@hey.com"
}

output "josh_id" {
  value = data.cognito_user.test_user.consumer_id
}

output "new_josh_id" {
  value = cognito_user.test_user.consumer_id
}