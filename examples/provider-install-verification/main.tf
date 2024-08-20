terraform {
  required_providers {
    bobomatic = {
      source = "havi.com/enterprisecloud/bobomatic"
    }
  }
}

provider "bobomatic" {}

data "bobomatic_resourcenames" "resourcenames" {
  bu          = "ct"
  purpose     = "automationhub"
  region      = null
  geography   = "global"
  client      = null
  program     = null
  environment = "dev"
  index       = 1
}

output "resourcenames" {
  value = data.bobomatic_resourcenames.resourcenames
}
