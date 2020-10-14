package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

// Execute uses the args (os.Args[1:] by default)
// and run through the command tree finding appropriate matches
// for commands and then corresponding flags.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	//注册命令
	rootCmd.AddCommand(wordCmd)

	rootCmd.AddCommand(timeCmd)

	rootCmd.AddCommand(sqlCmd)
}
