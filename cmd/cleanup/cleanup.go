package cleanup

import (
	"github.com/devspace-cloud/devspace/cmd/flags"
	"github.com/spf13/cobra"
)

// NewCleanupCmd creates a new cobra command
func NewCleanupCmd(globalFlags *flags.GlobalFlags) *cobra.Command {
	cleanupCmd := &cobra.Command{
		Use:   "cleanup",
		Short: "Cleans up resources",
		Long: `
#######################################################
################## devspace cleanup ###################
#######################################################
	`,
		Args: cobra.NoArgs,
	}

	cleanupCmd.AddCommand(newImagesCmd(globalFlags))

	return cleanupCmd
}
