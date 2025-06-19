# CC - Claude Code Project Generator

A powerful CLI tool that generates repository structures and boilerplate code optimized for [Claude Code](https://docs.anthropic.com/en/docs/claude-code) development. CC creates comprehensive project templates with proper `.claude/` configurations, documentation, CI/CD workflows, and example implementations.

## Features

- **Multiple Project Types**: Support for Python FastAPI, Go, Terraform, Kubernetes, Dagger, and Airflow projects
- **Claude Code Optimized**: Pre-configured with `.claude/` directory, CLAUDE.md, and example workflows
- **GitHub Integration**: Issue templates, PR templates, contributing guidelines, and CI/CD workflows
- **Modern Tooling**: Uses latest best practices for each technology stack
- **ARM64 Compatible**: Built and tested on Apple Silicon Macs
- **Development Ready**: Includes Docker, testing, linting, and formatting configurations

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/onprema/cc.git
cd cc

# Build the binary
go build -o cc

# Move to your PATH (optional)
sudo mv cc /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/onprema/cc@latest
```

## Quick Start

### Create a New Project

```bash
# Create a Python FastAPI project
cc new my-api python-fastapi --github=yourusername

# Create a Go project with GitHub integration
cc new my-service go --github=yourusername --description="My awesome Go service"

# Create a Kubernetes project
cc new my-k8s-app kubernetes --github=yourusername

# See all available project types
cc list
```

### Integrate into Existing Project

```bash
# Add Claude Code features to existing project
cd existing-project
cc integrate --type=go --github=yourusername
```

## Supported Project Types

| Type | Description | Key Features |
|------|-------------|--------------|
| `python-fastapi` | Modern Python FastAPI project | uv, ruff, pytest, Docker, async support |
| `go` | Go project with modern tooling | Gin, structured logging, Docker, proper modules |
| `terraform` | Infrastructure as Code | AWS modules, environments, state management |
| `kubernetes` | Kubernetes with minikube | Kustomize, Helm charts, monitoring, ingress |
| `dagger` | CI/CD pipeline as code | Go SDK, containerized builds, testing |
| `airflow` | Workflow orchestration | DAGs, Docker Compose, monitoring, plugins |

## Command Reference

### Global Flags

```bash
--config string     config file (default is $HOME/.cc.yaml)
--dry-run          show what would be created without creating
-v, --verbose      verbose output
--help             help for cc
--version          version for cc
```

### Commands

#### `cc new <name> <type> [flags]`

Create a new project with the specified type.

**Flags:**
- `-d, --description string`: Project description
- `-g, --github string`: GitHub username for integration
- `-o, --overwrite`: Overwrite existing files

**Examples:**
```bash
cc new my-api python-fastapi --github=myuser --description="REST API for my app"
cc new infra terraform --github=myuser
cc new k8s-demo kubernetes --github=myuser --overwrite
```

#### `cc integrate [flags]`

Integrate Claude Code features into an existing project.

**Flags:**
- `-t, --type string`: Project type (required)
- `-d, --description string`: Project description
- `-g, --github string`: GitHub username for integration
- `-o, --overwrite`: Overwrite existing files

**Examples:**
```bash
cc integrate --type=go --github=myuser
cc integrate --type=python-fastapi --description="Existing API" --overwrite
```

#### `cc list`

List all available project types with descriptions.

## Project Structure

Each generated project includes:

```
project-name/
├── .claude/                    # Claude Code configuration
│   └── examples/              # Project-type specific examples
├── .github/                   # GitHub templates and workflows
│   ├── workflows/             # CI/CD pipelines
│   ├── ISSUE_TEMPLATE/        # Bug reports, feature requests
│   └── pull_request_template.md
├── docs/                      # Documentation
├── tests/                     # Test files
├── CLAUDE.md                  # Claude Code project memory
├── Makefile                   # Common development commands
├── .pre-commit-config.yaml    # Code quality hooks
├── .gitignore                 # Comprehensive gitignore
├── CONTRIBUTING.md            # Contribution guidelines
├── LICENSE                    # MIT license
└── README.md                  # Project documentation
```

## Claude Code Integration

Every project is optimized for Claude Code development:

### CLAUDE.md File
Contains project-specific commands, architecture notes, and development workflows.

### .claude/examples Directory
Includes example prompts and workflows for the specific project type.

### Quick Commands
All projects include a Makefile with standard commands:

```bash
make help      # Show available commands
make install   # Install dependencies
make dev       # Start development environment
make test      # Run tests
make lint      # Run linting and formatting
make build     # Build the project
make clean     # Clean build artifacts
```

## Cheatsheet

### Common Workflows

```bash
# Quick project setup
cc new my-project go --github=myuser
cd my-project
make install && make dev

# Add Claude Code to existing project
cc integrate --type=python-fastapi --github=myuser

# Start development with Claude Code
claude

# Common development commands
make test      # Run tests
make lint      # Format and lint code
make build     # Build for production
```

### Project-Specific Commands

#### Python FastAPI
```bash
uv run fastapi dev             # Start development server
uv run pytest                 # Run tests
uv run ruff check .            # Lint code
docker-compose up              # Start with database
```

#### Go
```bash
go run ./cmd/server            # Start server
go test ./...                  # Run tests
golangci-lint run              # Lint code
docker-compose up              # Start with dependencies
```

#### Kubernetes
```bash
./scripts/setup-minikube.sh    # Setup local cluster
kubectl apply -k k8s/overlays/dev  # Deploy to dev
minikube dashboard             # Open dashboard
./scripts/cleanup.sh           # Clean up resources
```

#### Terraform
```bash
./scripts/terraform.sh dev plan    # Plan dev environment
./scripts/terraform.sh dev apply   # Apply changes
terraform fmt -recursive           # Format files
```

### Claude Code Commands

```bash
claude                         # Start interactive session
claude -p "help with tests"    # Quick query
claude -c                     # Continue last session
claude --dry-run              # Preview actions
```

## Configuration

CC can be configured via:

1. **Config file**: `~/.cc.yaml`
2. **Environment variables**: `CC_*` prefix
3. **Command-line flags**: Override config values

Example config file:

```yaml
# ~/.cc.yaml
github_username: myuser
default_description: "Generated with CC"
verbose: false
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
git clone https://github.com/onprema/cc.git
cd cc
go mod tidy
go build -o cc
./cc --help
```

### Running Tests

```bash
go test ./...
go test -race ./...
go test -cover ./...
```

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [Claude Code Docs](https://docs.anthropic.com/en/docs/claude-code)
- **Issues**: [GitHub Issues](https://github.com/onprema/cc/issues)
- **Examples**: Check the `.claude/examples/` directory in generated projects

---

Built with ❤️ for the Claude Code community.