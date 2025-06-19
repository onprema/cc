package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type ProjectConfig struct {
	Name           string
	Type           string
	Description    string
	GitHubUsername string
	Overwrite      bool
	DryRun         bool
	Verbose        bool
	Integration    bool // true if integrating into existing project
}

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

// GenerateProject is deprecated - use InitializeProject instead
func (g *Generator) GenerateProject(config *ProjectConfig) error {
	return g.InitializeProject(config)
}

func (g *Generator) InitializeProject(config *ProjectConfig) error {
	// Create Claude Code optimization structure
	if err := g.createClaudeStructure(".", config); err != nil {
		return fmt.Errorf("failed to create Claude structure: %w", err)
	}

	// Generate Claude Code files
	if err := g.generateClaudeFiles(".", config); err != nil {
		return fmt.Errorf("failed to generate Claude files: %w", err)
	}

	// Generate development workflow files
	if err := g.generateDevelopmentFiles(".", config); err != nil {
		return fmt.Errorf("failed to generate development files: %w", err)
	}

	// Generate GitHub integration if requested
	if config.GitHubUsername != "" {
		if err := g.generateGitHubIntegration(".", config); err != nil {
			return fmt.Errorf("failed to generate GitHub integration: %w", err)
		}
	}

	return nil
}

func (g *Generator) IntegrateProject(config *ProjectConfig) error {
	// Delegate to InitializeProject for now - same functionality
	return g.InitializeProject(config)
}

func (g *Generator) CheckExistingClaudeFiles(dir string) []string {
	claudeFiles := []string{
		"CLAUDE.md",
		".claude/",
	}

	var existing []string
	for _, file := range claudeFiles {
		if _, err := os.Stat(filepath.Join(dir, file)); err == nil {
			existing = append(existing, file)
		}
	}

	return existing
}

// Removed type validation - cc now works with any project type

// Removed createBaseStructure - using createClaudeStructure instead

func (g *Generator) createClaudeStructure(projectPath string, config *ProjectConfig) error {
	dirs := []string{
		".claude",
		".github",
		".github/workflows",
	}

	for _, dir := range dirs {
		dirPath := filepath.Join(projectPath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
	}

	return nil
}

// Removed generateFiles - functionality moved to InitializeProject and generateDevelopmentFiles

func (g *Generator) generateDevelopmentFiles(projectPath string, config *ProjectConfig) error {
	// Generate basic development files
	if err := g.generateGitIgnore(projectPath, config); err != nil {
		return err
	}

	if err := g.generateGenericMakefile(projectPath, config); err != nil {
		return err
	}

	if err := g.generatePreCommitConfig(projectPath, config); err != nil {
		return err
	}

	if err := g.generateGenericGitHubWorkflow(projectPath, config); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateClaudeFiles(projectPath string, config *ProjectConfig) error {
	// Generate CLAUDE.md
	if err := g.generateClaudeMD(projectPath, config); err != nil {
		return err
	}

	// Generate .claude directory with README
	if err := g.generateClaudeExamples(projectPath, config); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateClaudeMD(projectPath string, config *ProjectConfig) error {
	tmpl := `# {{.Name}}

{{.Description}}

This project has been optimized for Claude Code development.

## Quick Commands

` + "```bash" + `
# Development
make dev          # Start development environment
make test         # Run all tests
make lint         # Run linting and formatting
make build        # Build the project

# Claude Code Integration
claude            # Start Claude Code interactive session
claude -p "help"  # Quick help
claude -c         # Continue last session
` + "```" + `

## Project Structure

- ` + "`.claude/`" + ` - Claude Code configuration
- ` + "`.github/workflows/`" + ` - CI/CD pipelines
- ` + "`Makefile`" + ` - Development commands
- ` + "`.pre-commit-config.yaml`" + ` - Code quality hooks

## Development Workflow

1. Use ` + "`make install`" + ` to install dependencies
2. Use ` + "`make dev`" + ` to start development
3. Run ` + "`make test`" + ` before committing
4. Use ` + "`claude`" + ` for AI assistance
5. Commit with conventional commit messages

## Claude Code Features

This project includes:
- Pre-configured project memory (this file)
- Integration with development tools via Makefile
- GitHub workflows and templates
- Pre-commit hooks for code quality
- Claude Code configuration in ` + "`.claude/`" + ` directory

## Getting Started

1. Install dependencies: ` + "`make install`" + `
2. Start development: ` + "`make dev`" + `
3. Run tests: ` + "`make test`" + `
4. Open Claude Code: ` + "`claude`" + `

## Useful Claude Code Commands

- ` + "`claude -p \"explain the project structure\"`" + ` - Get project overview
- ` + "`claude -p \"help with testing\"`" + ` - Get testing assistance
- ` + "`claude -p \"review my changes\"`" + ` - Code review help
- ` + "`claude --dry-run`" + ` - Preview actions without making changes

For more information, see the ` + "`.claude/README.md`" + ` file for Claude Code configuration options.

---
*Generated by cc on {{.Date}}*
`

	t, err := template.New("claude-md").Parse(tmpl)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := struct {
		Name        string
		Description string
		Date        string
	}{
		Name:        config.Name,
		Description: config.Description,
		Date:        time.Now().Format("2006-01-02"),
	}

	file, err := os.Create(filepath.Join(projectPath, "CLAUDE.md"))
	if err != nil {
		return fmt.Errorf("failed to create CLAUDE.md: %w", err)
	}
	defer file.Close()

	return t.Execute(file, data)
}

// Removed getTypeSpecificNotes - using generic approach now

func (g *Generator) generateGitIgnore(projectPath string, config *ProjectConfig) error {
	content := `# Claude Code
.claude/local/
*.claude-session

# Common
.env
.env.local
*.log
.DS_Store
.vscode/
.idea/

# Dependencies
node_modules/
venv/
__pycache__/
*.pyc

# Build artifacts
dist/
build/
*.egg-info/
target/

# Test coverage
.coverage
htmlcov/
.pytest_cache/

# OS specific
Thumbs.db

# Add project-specific ignores below this line
`

	return g.writeFile(filepath.Join(projectPath, ".gitignore"), content)
}

// Removed getTypeSpecificGitIgnore - using generic approach

// Removed old generateMakefile - using generateGenericMakefile instead

// Removed MakeCommands and getTypeSpecificMakeCommands - using generic approach

func (g *Generator) generatePreCommitConfig(projectPath string, config *ProjectConfig) error {
	content := `# Pre-commit configuration for code quality
# Install with: pre-commit install

repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.4.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: check-yaml
      - id: check-added-large-files
      - id: check-merge-conflict
      
  - repo: local
    hooks:
      - id: test
        name: run tests
        entry: make test
        language: system
        pass_filenames: false
        
      - id: lint
        name: run linting
        entry: make lint
        language: system
        pass_filenames: false

# Add project-specific pre-commit hooks below this line
`

	return g.writeFile(filepath.Join(projectPath, ".pre-commit-config.yaml"), content)
}

// Removed getTypeSpecificPreCommitHooks - using generic approach

// Removed old generateGitHubWorkflow - using generateGenericGitHubWorkflow instead

// Removed getTypeSpecificGitHubWorkflow - using generic approach

// Removed generateTypeSpecificFiles - cc now focuses on Claude Code optimization only

func (g *Generator) generateClaudeExamples(projectPath string, config *ProjectConfig) error {
	// Just create the .claude directory - no examples needed
	claudeDir := filepath.Join(projectPath, ".claude")
	if err := os.MkdirAll(claudeDir, 0755); err != nil {
		return fmt.Errorf("failed to create .claude directory: %w", err)
	}

	// Create a simple .claude/README.md explaining the directory
	content := `# .claude Directory

This directory contains Claude Code configuration and project-specific settings.

## What goes here?

- Custom Claude Code configurations
- Project-specific prompts and workflows
- Local Claude Code settings (not committed to git)
- Integration configurations for MCP servers

## Getting Started

This directory is automatically created by the cc tool. You can customize it
based on your project's specific needs.
`

	return g.writeFile(filepath.Join(claudeDir, "README.md"), content)
}

func (g *Generator) writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func (g *Generator) generateGitHubIntegration(projectPath string, config *ProjectConfig) error {
	// Generate .github/ISSUE_TEMPLATE directory
	issueTemplateDir := filepath.Join(projectPath, ".github", "ISSUE_TEMPLATE")
	if err := os.MkdirAll(issueTemplateDir, 0755); err != nil {
		return fmt.Errorf("failed to create issue template directory: %w", err)
	}

	// Generate bug report template
	bugReportContent := `---
name: Bug report
about: Create a report to help us improve
title: ''
labels: bug
assignees: ` + config.GitHubUsername + `

---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Screenshots**
If applicable, add screenshots to help explain your problem.

**Environment (please complete the following information):**
 - OS: [e.g. iOS]
 - Version [e.g. 22]

**Additional context**
Add any other context about the problem here.
`

	if err := g.writeFile(filepath.Join(issueTemplateDir, "bug_report.md"), bugReportContent); err != nil {
		return err
	}

	// Generate feature request template
	featureRequestContent := `---
name: Feature request
about: Suggest an idea for this project
title: ''
labels: enhancement
assignees: ` + config.GitHubUsername + `

---

**Is your feature request related to a problem? Please describe.**
A clear and concise description of what the problem is. Ex. I'm always frustrated when [...]

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of any alternative solutions or features you've considered.

**Additional context**
Add any other context or screenshots about the feature request here.
`

	if err := g.writeFile(filepath.Join(issueTemplateDir, "feature_request.md"), featureRequestContent); err != nil {
		return err
	}

	// Generate pull request template
	prTemplateContent := `## Description

Please include a summary of the changes and the related issue. Please also include relevant motivation and context.

Fixes # (issue)

## Type of change

Please delete options that are not relevant.

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] This change requires a documentation update

## How Has This Been Tested?

Please describe the tests that you ran to verify your changes. Provide instructions so we can reproduce.

- [ ] Test A
- [ ] Test B

## Checklist:

- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] Any dependent changes have been merged and published
`

	if err := g.writeFile(filepath.Join(projectPath, ".github", "pull_request_template.md"), prTemplateContent); err != nil {
		return err
	}

	// Generate CONTRIBUTING.md
	contributingContent := `# Contributing to ` + config.Name + `

First off, thank you for considering contributing to ` + config.Name + `! It's people like you that make ` + config.Name + ` such a great tool.

## Where do I go from here?

If you've noticed a bug or have a feature request, make sure to check our [Issues](https://github.com/` + config.GitHubUsername + `/` + config.Name + `/issues) if there's something similar to what you have in mind. If there isn't, feel free to open a new issue!

## Fork & create a branch

If this is something you think you can fix, then fork ` + config.Name + ` and create a branch with a descriptive name.

A good branch name would be:

` + "```" + `
git checkout -b 325-add-japanese-translations
` + "```" + `

## Get the test suite running

Make sure you're using a recent version of the development tools:

` + "```bash" + `
make install
make test
` + "```" + `

## Implement your fix or feature

At this point, you're ready to make your changes! Feel free to ask for help; everyone is a beginner at first.

## View your changes

Make sure to take a look at your changes in a real environment.

## Get the style right

Your patch should follow the same conventions & pass the same code quality checks as the rest of the project.

` + "```bash" + `
make lint
` + "```" + `

## Make a Pull Request

At this point, you should switch back to your main branch and make sure it's up to date with the latest ` + config.Name + ` main branch:

` + "```bash" + `
git remote add upstream git@github.com:` + config.GitHubUsername + `/` + config.Name + `.git
git checkout main
git pull upstream main
` + "```" + `

Then update your feature branch from your local copy of main, and push it!

` + "```bash" + `
git checkout 325-add-japanese-translations
git rebase main
git push --set-upstream origin 325-add-japanese-translations
` + "```" + `

Finally, go to GitHub and make a Pull Request!

## Keeping your Pull Request updated

If a maintainer asks you to "rebase" your PR, they're saying that a lot of code has changed, and that you need to update your branch so it's easier to merge.

## Merging a PR (maintainers only)

A PR can only be merged into main by a maintainer if:

* It is passing CI.
* It has been approved by at least two maintainers. If it was a maintainer who opened the PR, only one extra approval is needed.
* It has no requested changes.
* It is up to date with current main.

Any maintainer is allowed to merge a PR if all of these conditions are met.
`

	if err := g.writeFile(filepath.Join(projectPath, "CONTRIBUTING.md"), contributingContent); err != nil {
		return err
	}

	// Generate LICENSE file (MIT)
	licenseContent := `MIT License

Copyright (c) 2024 ` + config.GitHubUsername + `

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`

	return g.writeFile(filepath.Join(projectPath, "LICENSE"), licenseContent)
}

func (g *Generator) generateGenericMakefile(projectPath string, config *ProjectConfig) error {
	content := `# Makefile for ` + config.Name + `
# Generated by cc - Claude Code optimization tool

.PHONY: help install dev test lint build clean

help:
	@echo "Available commands:"
	@echo "  make install   - Install dependencies"
	@echo "  make dev       - Start development environment"
	@echo "  make test      - Run tests"
	@echo "  make lint      - Run linting and formatting"
	@echo "  make build     - Build the project"
	@echo "  make clean     - Clean build artifacts"

install:
	@echo "Installing dependencies..."
	@echo "Add your dependency installation commands here"

dev:
	@echo "Starting development environment..."
	@echo "Add your development startup commands here"

test:
	@echo "Running tests..."
	@echo "Add your test commands here"

lint:
	@echo "Running linting and formatting..."
	@echo "Add your linting commands here"

build:
	@echo "Building project..."
	@echo "Add your build commands here"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf dist/ build/ *.egg-info/ target/
	find . -type d -name __pycache__ -exec rm -rf {} + 2>/dev/null || true
	find . -type f -name "*.pyc" -delete 2>/dev/null || true
`

	return g.writeFile(filepath.Join(projectPath, "Makefile"), content)
}

func (g *Generator) generateGenericGitHubWorkflow(projectPath string, config *ProjectConfig) error {
	workflowDir := filepath.Join(projectPath, ".github", "workflows")
	if err := os.MkdirAll(workflowDir, 0755); err != nil {
		return fmt.Errorf("failed to create workflow directory: %w", err)
	}

	content := `name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Run tests
        run: make test
        
      - name: Run linting  
        run: make lint
        
      - name: Build project
        run: make build
`

	return g.writeFile(filepath.Join(workflowDir, "ci.yml"), content)
}

// Type-specific file generators are implemented in separate files

