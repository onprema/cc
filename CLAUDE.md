# cc-init

The goal of this project is to create an application that can be used to generate a repository structure and boilerplate code for a new project that utilizes the latest and greatest features of Claude Code.

Source: https://www.anthropic.com/learn/build-with-claude (look recursively)

Requirements:
- Must be able to run on an arm-based macbook
- Should be a standalone application
- Should include a concise README.md that explains how to use it, following the guidelines in the user memory's style guide.
  - Should include a "cheatsheet" that provides a quick reference for common commands and features.
- Should support GitHub integrations (the user's github username could be an environment variable or argument)
- Should support MCP integrations
- Should support the latest features (e.g, plan mode, sound when complete)
- Should follow best practices for tokens and sensitive information
- Generate boilerplat code (e.g., `docs/*`, `CLAUDE.md`, `.claude/*`, plans, should include a text snippet at the top of each file, explaining the purpose of the file and how to use it)
- Should include helpful claude commands and features
- Should include helpful claude integrations


The users of this application will be either:
- creating new projects -> could be a web app, backend api, infrastructure as code, using a variety of languages and technologies
- integrating into an existing project -> could be an existing repo, therefore the application should be smart enough to not overwrite files, and instrument the existing repo with the necessary files and structure that claude code expects.

You can make some reasonable assumptions about the project, like:
- it should have unit tests
- it should have integration tests (where applicable)
- it should have documentation
- you should be able to containerize it and deploy it locally, easily!
- you should include a `.github/workflows` directory with CI/CD workflows for the project.
- custom claude functions that are relevant to the project

**Include an `.claude/examples` directory** -> this should include a few examples of how to populate the generated files from this application. for example, if i ran `cc-init` for a python fastapi project, the `./claude/examples/python-fastapi` directory should be ready-to-go, meaning that the docs, system prompts, configurations, etc should be suitable for a modern python fastapi project (use `uv` and `ruff` for linting, package management, formatting, etc!) -- you can search the internet to find similar projects that have a lot of starts on github and are "industry standards" as inspiration for setting up the project.

beyond python-fastapi, make examples for go, dagger, airflow, terraform, and kubernetes (using minikube).

remember that this all should be able to run on an arm64 architecture.

if you need to create virtual machines, you can use Vagrant and/or Qemu. If you need to use containers, you can use `podman`. Try to use podman instead of docker. Use `Makefile` and `pre-commit` hooks within the project.

so the `examples` directory should be faily robust, where the boilerplate has actually been populated and it should be able to be fully functional. the examples might include comments and explanations about the files and the claude commands to run.
