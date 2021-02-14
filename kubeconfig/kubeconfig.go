package kubeconfig

import (
	"fmt"
	"log"

	"k8s.io/client-go/tools/clientcmd"
)

type Kubeconfig struct {
	clientConfig clientcmd.ClientConfig
}

func NewKubeconfig() (*Kubeconfig, error) {
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	return &Kubeconfig{clientConfig: clientConfig}, nil
}

func (k *Kubeconfig) ReadByUser(username string) error {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return err
	}
	users := rawConfig.AuthInfos

	u, ok := users[username]
	if ok != true {
		log.Fatalf("User with username %s not found in kube config", username)
	}

	// なんかの構造体に格納する
	fmt.Printf("%+v", u)

	return nil
}
