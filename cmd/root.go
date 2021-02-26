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
	rootCmd.AddCommand(deleteCmd)
}

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Kubernetes Secrets Add/Update value ",
	Run: AddAndUpdate,
	Example: `Add/Update Secret Data:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks modify key:value 
Add/Update Multiple Secret Data:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks modify key1:value1 key2:value2`,
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Kubernetes Secrets Delete value ",
	Run: DeleteValueFromSecret,
	Example: `Delete Secret Data:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks delete key
Delete Multiple Secret Data:
kubetcl get secrets <Secret Name> -n <Namespace> -o yaml | ks delete key1 key2`,
}