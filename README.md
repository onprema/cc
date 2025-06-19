# cc - Claude Code Project Generator

A standalone CLI application for generating repository structures and boilerplate code optimized for [Claude Code](https://docs.anthropic.com/en/docs/claude-code) development. This tool helps bootstrap new projects or integrate Claude Code features into existing repositories.

## Features

- **Project Generation**: Creates new projects with Claude Code optimizations
- **Existing Project Integration**: Adds Claude Code features to existing repositories without overwriting files  
- **ARM64 Compatible**: Designed to run on Apple Silicon Macs
- **Multiple Project Types**: Planned support for Python FastAPI, Go, Terraform, Kubernetes, Dagger, and Airflow
- **GitHub Integration**: Configurable GitHub username integration
- **MCP Support**: Built-in Model Context Protocol integrations
- **Modern Tooling**: Follows latest best practices for security and token management

## Installation

### From Source

```bash
# Clone the repository  
git clone <repository-url>
cd cc

# Build the binary
go build -o cc

# Move to your PATH (optional)
sudo mv cc /usr/local/bin/
```

## Usage

This application is currently in development. The planned usage will be:

```bash
# Create a new project (planned)
cc new <project-name> --type=<project-type> --github=<username>

# Integrate into existing project (planned)  
cd existing-project
cc integrate --type=<project-type> --github=<username>
```

## Planned Project Types

The following project types are planned for implementation:

| Type | Description | Key Features |
|------|-------------|--------------|
| `python-fastapi` | Modern Python FastAPI project | uv, ruff, pytest, Podman, async support |
| `go` | Go project with modern tooling | Standard library focus, structured logging, Podman |
| `terraform` | Infrastructure as Code | Modules, environments, state management |
| `kubernetes` | Kubernetes with minikube | Kustomize, Helm charts, monitoring, ingress |
| `dagger` | CI/CD pipeline as code | Go SDK, containerized builds, testing |
| `airflow` | Workflow orchestration | DAGs, Podman Compose, monitoring, plugins |

## Planned Features

The following features are planned for implementation:

- **Project Generation**: Create new projects with comprehensive boilerplate
- **Repository Integration**: Add Claude Code features to existing repositories
- **Template Examples**: Pre-configured examples in `.claude/examples/` directory
- **GitHub Integration**: Automated setup with workflows and templates
- **MCP Integrations**: Built-in Model Context Protocol support
- **Security Best Practices**: Proper handling of tokens and sensitive information
- **Development Tooling**: Makefiles, pre-commit hooks, and CI/CD workflows

## Planned Project Structure

Each generated project will include:

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

## Planned Claude Code Integration

Generated projects will be optimized for Claude Code development:

### CLAUDE.md File
Will contain project-specific commands, architecture notes, and development workflows.

### .claude/examples Directory  
Will include example prompts and workflows for the specific project type.

### Standard Development Commands
All projects will include a Makefile with standard commands:

```bash
make help      # Show available commands
make install   # Install dependencies  
make dev       # Start development environment
make test      # Run tests
make lint      # Run linting and formatting
make build     # Build the project
make clean     # Clean build artifacts
```

## Development Status

This project is currently in early development. The goal is to create a robust CLI tool that generates comprehensive project structures optimized for Claude Code development.

### Current Structure

```
cc/
├── main.go                    # Entry point
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── CLAUDE.md                  # Project instructions
├── cmd/                       # CLI commands
├── internal/                  # Internal packages
└── bin/                       # Built binaries
```

### Development Setup

```bash
git clone <repository-url>
cd cc
go mod tidy
go build -o cc
./cc --help
```

### Contributing

This project follows the guidelines specified in CLAUDE.md. Contributions should focus on:

- ARM64 compatibility
- Security best practices for tokens and sensitive information
- Modern tooling and development practices
- Comprehensive project templates with working examples

## License

This project will be licensed under the MIT License.

## Resources

- **Claude Code Documentation**: [https://docs.anthropic.com/en/docs/claude-code](https://docs.anthropic.com/en/docs/claude-code)
- **Build with Claude**: [https://www.anthropic.com/learn/build-with-claude](https://www.anthropic.com/learn/build-with-claude)