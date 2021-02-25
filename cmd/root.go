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
	rootCmd.Flags().BoolP("reverse","r", false, "String to base64 value")
	rootCmd.AddCommand(modifyCmd)
}

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Kubernetes Secrets Add/Update value ",
	Run: AddAndUpdate,
	Example: `Add/Update Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks modify key:value 
Add/Update Multiple Secret:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks modify key1:value1 key2:value2`,
}
