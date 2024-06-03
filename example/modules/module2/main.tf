resource "random_uuid" "test" {
}

module "test3" {
  source = "../module3"
}