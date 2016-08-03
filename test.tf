resource "filesystem_file" "test" {
  path = "/opt/terraform/test.txt"
  contents = "derpo"
}
