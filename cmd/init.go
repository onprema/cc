package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/onprema/cc/internal/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Claude Code optimization for current project",
	Long: `Initialize adds Claude Code optimization to your current project.
It creates .claude/ directory, CLAUDE.md file, development workflows,
and other files to make your project work seamlessly with Claude Code.

This command is safe to run multiple times and will not overwrite
existing files unless --overwrite is specified.`,
	Example: `  cc init                                    # Basic Claude Code setup
  cc init --github=username                 # Add GitHub integration  
  cc init --description="My project"        # Add project description
  cc init --overwrite                       # Overwrite existing files`,
	RunE: runInit,
}

var (
	description string
	github      string
	overwrite   bool
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&description, "description", "d", "", "Project description")
	initCmd.Flags().StringVarP(&github, "github", "g", "", "GitHub username for integration")
	initCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "Overwrite existing files")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Get current directory name as project name
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	
	projectName := filepath.Base(cwd)
	
	// Set default description if not provided
	if description == "" {
		description = fmt.Sprintf("A project optimized for Claude Code development")
	}

	// Check for existing Claude Code files
	gen := generator.New()
	existing := gen.CheckExistingClaudeFiles(".")
	
	if len(existing) > 0 && !overwrite {
		fmt.Printf("Found existing Claude Code files: %v\n", existing)
		fmt.Println("Use --overwrite to replace existing files")
		fmt.Println("Or run with different flags to add missing files")
	}

	config := &generator.ProjectConfig{
		Name:           projectName,
		Description:    description,
		GitHubUsername: github,
		Overwrite:      overwrite,
		DryRun:         viper.GetBool("dry-run"),
		Verbose:        viper.GetBool("verbose"),
		Integration:    true, // Always integration mode for init
	}

	if config.Verbose {
		fmt.Printf("Initializing Claude Code optimization for: %s\n", projectName)
		fmt.Printf("Description: %s\n", config.Description)
		if config.GitHubUsername != "" {
			fmt.Printf("GitHub integration: %s\n", config.GitHubUsername)
		}
	}

	if config.DryRun {
		fmt.Println("DRY RUN - No files will be created")
		return nil
	}

	// Initialize Claude Code optimization
	if err := gen.InitializeProject(config); err != nil {
		return fmt.Errorf("failed to initialize Claude Code optimization: %w", err)
	}

	fmt.Printf("✅ Successfully initialized Claude Code optimization for %s\n", projectName)
	fmt.Println("\nNext steps:")
	fmt.Println("1. Review the generated CLAUDE.md file")
	fmt.Println("2. Check the .claude/ directory for examples")
	fmt.Println("3. Run 'claude' to start using Claude Code")
	
	if config.GitHubUsername != "" {
		fmt.Println("4. Commit and push your changes to GitHub")
	}

	return nil
}