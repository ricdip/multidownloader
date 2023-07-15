package cmd

import (
	"multidownloader/src/flags"
	"multidownloader/src/supervisor"
	"multidownloader/src/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     ProgramName,
	Long:    ProgramDesc,
	Version: ProgramVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := utils.CheckDownloadLinks(*flags.Links); err != nil {
			return err
		}

		supervisor.Exec()
		return nil
	},
}

func init() {
	flags.Links = rootCmd.Flags().StringArrayP(
		"link",
		"l",
		[]string{},
		"each link to download (-l <link_1> -l <link_2> ...)",
	)
	rootCmd.MarkFlagRequired("link")
	rootCmd.Flags().BoolVarP(
		&flags.AutoYes,
		"yes",
		"y",
		false,
		"automatic yes to prompt",
	)
	rootCmd.Flags().BoolVarP(
		&flags.DebugLog,
		"debug",
		"d",
		false,
		"enable debug logging",
	)
	rootCmd.Flags().BoolVarP(
		&flags.ColoredLog,
		"color",
		"c",
		false,
		"enable colored logging",
	)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
