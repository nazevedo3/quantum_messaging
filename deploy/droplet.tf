terraform {
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
    }
  }
}

variable "api_token" {
    default = ""
}

provider "digitalocean" {
    token = var.api_token
}

resource "digitalocean_droplet" "web" {
    image = "ubuntu-18-04-x64"
    name = "test-web-vm"
    region = "sfo3"
    size = "s-1vcpu-1gb"
    monitoring = true
    ipv6 = false
    private_networking = false

  
}

