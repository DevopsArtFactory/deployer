package cmd

import (
	"context"
	"io"

	"github.com/DevopsArtFactory/goployer/pkg/runner"
	"github.com/spf13/cobra"
)

// Create new deploy command
func NewInitCommand() *cobra.Command {
	return NewCmd("init").
		WithDescription("initialize goployer manifest").
		SetFlags().
		RunWithArgs(funcInit)
}

// funcDelete delete stacks
func funcInit(ctx context.Context, _ io.Writer, args []string) error {
	return runWithoutExecutor(ctx, func() error {
		if err := runner.Initialize(args); err != nil {
			return err
		}

		return nil
	})
}
