package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Kubeconfig struct {
	clientConfig   clientcmd.ClientConfig
	configFilePath string
}

func New() (*Kubeconfig, error) {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})

	return &Kubeconfig{
		clientConfig:   clientConfig,
		configFilePath: rules.GetLoadingPrecedence()[0],
	}, nil
}

func (k *Kubeconfig) ReadContext(name string) (*api.Context, error) {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	obj, ok := rawConfig.Contexts[name]
	if !ok {
		return nil, fmt.Errorf("'%s' context can not read", name)
	}
	if obj == nil {
		return nil, fmt.Errorf("'%s' context was nof found in your kubeconfig", name)
	}

	return obj, nil
}

func (k *Kubeconfig) ReadCluster(name string) (*api.Cluster, error) {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	obj, ok := rawConfig.Clusters[name]
	if !ok {
		return nil, fmt.Errorf("'%s' cluster was nof found in your kubeconfig", name)
	}

	return obj, nil
}

func (k *Kubeconfig) ReadUser(name string) (*api.AuthInfo, error) {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	obj, ok := rawConfig.AuthInfos[name]
	if !ok {
		return nil, fmt.Errorf("'%s' user was nof found in your kubeconfig", name)
	}

	return obj, nil
}

func (k *Kubeconfig) WriteUser(name string, data *api.AuthInfo) error {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return err
	}

	rawConfig.AuthInfos[name] = data

	if err := k.write(rawConfig); err != nil {
		return err
	}

	return nil
}

func (k *Kubeconfig) WriteCluster(name string, data *api.Cluster) error {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return err
	}

	rawConfig.Clusters[name] = data

	if err := k.write(rawConfig); err != nil {
		return err
	}

	return nil
}

func (k *Kubeconfig) WriteContext(name string, data *api.Context) error {
	rawConfig, err := k.clientConfig.RawConfig()
	if err != nil {
		return err
	}

	rawConfig.Contexts[name] = data

	if err := k.write(rawConfig); err != nil {
		return err
	}

	return nil
}

func (k *Kubeconfig) write(rawConfig api.Config) error {
	if err := clientcmd.Validate(rawConfig); err != nil {
		return err
	}

	if err := clientcmd.WriteToFile(rawConfig, k.configFilePath); err != nil {
		return err
	}

	return nil
}
