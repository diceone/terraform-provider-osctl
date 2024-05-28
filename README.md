# Terraform Provider for `osctl`

A Terraform provider for managing and retrieving information from `osctl` API, a command-line tool for administrating Linux operating systems like RHEL, Ubuntu, and SUSE.

## Features

- Manage system services (start, stop, restart)
- Retrieve system statistics (RAM usage, disk usage, etc.)

## Installation

### Building from Source

1. **Clone the Repository**

   ```sh
   git clone https://github.com/diceone/terraform-provider-osctl.git
   cd terraform-provider-osctl
   ```

2. **Build the Provider**

   ```sh
   go build -o terraform-provider-osctl
   ```

3. **Move the Provider Binary**

   Move the compiled binary to the Terraform plugins directory. The exact location depends on your OS and Terraform version. For most Unix-based systems:

   ```sh
   mkdir -p ~/.terraform.d/plugins/yourusername/osctl/1.0.0/linux_amd64
   mv terraform-provider-osctl ~/.terraform.d/plugins/yourusername/osctl/1.0.0/linux_amd64
   ```

## Usage

### Provider Configuration

In your Terraform configuration file, configure the `osctl` provider with the necessary details.

```hcl
provider "osctl" {
  api_url  = "http://your-osctl-api-url"
  username = "admin"
  password = "password"
}
```

### Resources

#### `osctl_service`

Manage system services using the `osctl_service` resource.

```hcl
resource "osctl_service" "example" {
  name   = "apache2"
  action = "start"
}
```

#### Arguments

- `name` (Required) - The name of the service to manage.
- `action` (Required) - The action to perform on the service (start, stop, restart).

### Data Sources

#### `osctl_stats`

Retrieve system statistics using the `osctl_stats` data source.

```hcl
data "osctl_stats" "example" {
}

output "ram_usage" {
  value = data.osctl_stats.example.ram_usage
}

output "disk_usage" {
  value = data.osctl_stats.example.disk_usage
}
```

#### Attributes

- `ram_usage` - The current RAM usage of the system.
- `disk_usage` - The current disk usage of the system.
- Add more attributes as needed.

### Example Configuration

Here is a complete example of using the `osctl` provider in a Terraform configuration:

```hcl
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
```

## Development

### Requirements

- Go 1.18 or later
- Terraform 0.13 or later

### Building the Provider

Clone the repository and build the provider:

```sh
git clone https://github.com/diceone/terraform-provider-osctl.git
cd terraform-provider-osctl
go build -o terraform-provider-osctl
```

### Testing

Before running tests, make sure to set up the environment variables required for the provider configuration.

```sh
export OSCTL_API_URL=http://your-osctl-api-url
export OSCTL_USERNAME=admin
export OSCTL_PASSWORD=password
```

Run the tests:

```sh
go test ./...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This `README.md` provides a comprehensive overview of the Terraform provider for `osctl`, including installation, usage, and development instructions. Adjust the repository URL and other details as needed to match your actual setup.
