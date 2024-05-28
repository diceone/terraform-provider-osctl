provider "osctl" {
  api_url  = "http://your-osctl-api-url"
  username = "admin"
  password = "password"
}

resource "osctl_service" "example" {
  name   = "apache2"
  action = "start"
}

data "osctl_stats" "example" {
}

output "ram_usage" {
  value = data.osctl_stats.example.ram_usage
}

output "disk_usage" {
  value = data.osctl_stats.example.disk_usage
}
