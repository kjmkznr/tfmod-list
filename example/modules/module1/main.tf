resource "random_uuid" "test" {
}

module "test2" {
  source = "../module2"
}

module "test3" {
  source = "../module3"
}