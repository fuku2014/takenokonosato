terraform {
  required_providers {
    takenokonosato = {
      source = "fuku2014/takenokonosato"
    }
  }
}

provider "takenokonosato" {
  kinoko = true
}

resource "takenokonosato_ichigoazi" "example" {
  name = "oyatsu"

  price = 100
}
