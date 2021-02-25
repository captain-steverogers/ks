package cmd

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	b64 "encoding/base64"
	"strings"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"context"
	"k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1Types "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)
var secretsClient coreV1Types.SecretInterface
var ctx context.Context
type Secret map[string]interface{}

func FlagCheck(cmd *cobra.Command, args []string) {
		bytes := ReadInputs()
		bools,_ := cmd.Flags().GetBool("reverse")
		if bools {
			value := SecretsConvert(bytes)
			fmt.Print(value)
		} else {
			SecretsGet(bytes)
		}
}


func SecretsConvert(bytes []byte) string{
	b := strings.Replace(string(bytes), "\n", "", -1)
	return b64.StdEncoding.EncodeToString([]byte(b))
}

func SecretsGet(bytes []byte) {
	var spec Secret
	var m interface{}
	err := yaml.Unmarshal(bytes, &spec)
	if err != nil {
	  panic(err.Error())
	}	
	m = spec["data"]
	Mapkeys := m.(map[string]interface{})
	for k,v := range Mapkeys {
		value := fmt.Sprintf("%s",v)
		data,_:= b64.StdEncoding.DecodeString(value)
		Mapkeys[k] = string(data)
	}
	s,err := yaml.Marshal(spec)
	fmt.Printf("%s\n",s)
}

func UpdateSecrets(spec v1.Secret) *v1.Secret {
	secret, err := secretsClient.Update(context.TODO(),&spec, metaV1.UpdateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("succesfully Secret Updated")
	return secret
}
func initClient(namespace string) {
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		kubeconfig = os.Getenv("HOME") + "/.kube/config"
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secretsClient = clientset.CoreV1().Secrets(namespace)
}

func AddAndUpdate(cmd *cobra.Command, args []string) {
		bytes := ReadInputs()
	    var spec v1.Secret
		err := yaml.Unmarshal(bytes, &spec)
		if err != nil {
		  panic(err.Error())
		}
		if spec.Data == nil {
			spec.Data = map[string][]byte{}
		}
		for _,s := range args {
			args := strings.Split(s,":")
			spec.Data[args[0]] = []byte(args[1])
		}
		initClient(spec.Namespace)
		UpdateSecrets(spec)
}

func ReadInputs() []byte {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return bytes
}