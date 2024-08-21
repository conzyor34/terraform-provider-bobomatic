terraform {
  required_providers {
    bobomatic = {
      source = "havi.com/enterprisecloud/bobomatic"
    }
  }
}

provider "bobomatic" {}

data "bobomatic_resourcenames" "resourcenames" {
  bu          = "bu"
  purpose     = "purpose"
  region      = "region"
  geography   = "geography"
  client      = "client"
  program     = "program"
  environment = "environment"
  index       = 1
}

output "resourcenames" {
  value = data.bobomatic_resourcenames.resourcenames
}
