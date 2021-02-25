package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "ks",
		Short: "Kubernetes Secrets Manager",
		Run: FlagCheck,
		Version: "v0.1.0",
		Example: "kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolP("reverse","r", false, "String to base64 value")
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Kubernetes Secret",
	Run: AddAndUpdate,
	Example: `Update Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks update key:value 
Update Multiple Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks update key1:value1 key2:value2
	`,
}
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Kubernetes Secrets Add value ",
	Run: AddAndUpdate,
	Example: `Add Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks add key:value 
Add Multiple Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks add key1:value1 key2:value2`,
}
