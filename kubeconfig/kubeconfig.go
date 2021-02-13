package kubeconfig

type Kubeconfig struct {
}

func NewKubeconfig() (*Kubeconfig, error) {
	return &Kubeconfig{}, nil
}

func (k *Kubeconfig) ReadUser(username string) error {
	return nil
}
