package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/ghodss/yaml"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/tuuturu/pager-email-client/pkg/core"
	"github.com/tuuturu/pager-email-client/pkg/email/imap"
	"github.com/tuuturu/pager-email-client/pkg/filtering"
)

var (
	cfg core.Config

	filterConfigPath string

	filter core.Filter
)

var rootCmd = &cobra.Command{
	Use:   "pager-email-client",
	Short: "pager-email-client sends notifications to a Pager service upon receiving emails matching certain criteria",
	Long:  `A simple CLI tool to send email notifications to a Pager service`,
	PreRunE: func(cmd *cobra.Command, args []string) (err error) {
		cfg, err = core.LoadConfig()
		if err != nil {
			return err
		}

		var input io.Reader
		if filterConfigPath == "-" {
			input = os.Stdin
		} else {
			fs := &afero.Afero{Fs: afero.NewOsFs()}

			input, err = fs.Open(filterConfigPath)
			if err != nil {
				return fmt.Errorf("opening file at path %s: %w", filterConfigPath, err)
			}
		}

		filter, err = filtering.ParseFilterConfig(input)
		if err != nil {
			return fmt.Errorf("parsing filter config: %w", err)
		}

		return cfg.Validate()
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		messages, err := imap.RetrieveEmails(cfg.IMAPServerURL, cfg.Username, cfg.Password)
		if err != nil {
			return fmt.Errorf("retrieving emails: %w", err)
		}

		for _, message := range messages {
			if filter.Test(message) {
				raw, _ := yaml.Marshal(message)

				_, _ = os.Stdout.Write(raw)
			}
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	flags := rootCmd.Flags()

	flags.StringVarP(
		&filterConfigPath,
		"filter-config-path",
		"f",
		"-",
		"set filter config path",
	)
}
