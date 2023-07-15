package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/notaryproject/notation-plugin-base/errors"
	"github.com/notaryproject/notation-plugin-base/internal/io"
	"github.com/notaryproject/notation-plugin-base/plugin"
	"github.com/spf13/cobra"
)

type Helper struct {
	p plugin.Plugin
}

func NewHelper(pl plugin.Plugin) *Helper {
	return &Helper{p: pl}
}

func (c Helper) Execute() {
	cmd := &cobra.Command{
		Use:           "plugin for Notation E2E test",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		c.getPluginMetadataCommand(),
		c.describeKeyCommand(),
		c.generateSignatureCommand(),
		c.generateEnvelopeCommand(),
		c.verifySignatureCommand(),
	)

	if err := cmd.Execute(); err != nil {
		plgErr, ok := err.(*errors.PluginError)
		var errString string
		if ok {
			errString = plgErr.Error()
		} else {
			errString = errors.NewGenericError(err.Error()).Error()
		}

		_, _ = fmt.Fprintf(os.Stderr, errString)
		os.Exit(1)
	}
}

func (c Helper) getPluginMetadataCommand() *cobra.Command {
	return &cobra.Command{
		Use: "get-plugin-metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &plugin.GetMetadataRequest{}
			if err := io.UnmarshalRequest(req); err != nil {
				return err
			}

			resp, err := c.p.GetMetadata(context.Background(), req)
			if err != nil {
				return err
			}
			return io.PrintResponse(resp)
		},
	}
}

func (c Helper) describeKeyCommand() *cobra.Command {
	return &cobra.Command{
		Use: "describe-key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &plugin.DescribeKeyRequest{}
			if err := io.UnmarshalRequest(req); err != nil {
				return err
			}

			resp, err := c.p.DescribeKey(context.Background(), req)
			if err != nil {
				return err
			}

			return io.PrintResponse(resp)
		},
	}
}

func (c Helper) generateEnvelopeCommand() *cobra.Command {
	return &cobra.Command{
		Use: "generate-envelope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &plugin.GenerateEnvelopeRequest{}
			if err := io.UnmarshalRequest(req); err != nil {
				return err
			}

			resp, err := c.p.GenerateEnvelope(context.Background(), req)
			if err != nil {
				return err
			}

			return io.PrintResponse(resp)
		},
	}
}

func (c Helper) generateSignatureCommand() *cobra.Command {
	return &cobra.Command{
		Use: "generate-signature",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &plugin.GenerateSignatureRequest{}
			if err := io.UnmarshalRequest(req); err != nil {
				return err
			}
			resp, err := c.p.GenerateSignature(context.Background(), req)
			if err != nil {
				return err
			}

			return io.PrintResponse(resp)
		},
	}
}

func (c Helper) verifySignatureCommand() *cobra.Command {
	return &cobra.Command{
		Use: "verify-signature",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := &plugin.VerifySignatureRequest{}
			if err := io.UnmarshalRequest(req); err != nil {
				return err
			}

			resp, err := c.p.VerifySignature(context.Background(), req)
			if err != nil {
				return err
			}

			return io.PrintResponse(resp)
		},
	}
}
