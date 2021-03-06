package kubeconfig

import (
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
	clientcmd "k8s.io/client-go/tools/clientcmd/api/v1"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/tls"
	"github.com/openshift/installer/pkg/types"
)

type kubeconfig struct {
	Config *clientcmd.Config
	File   *asset.File
}

// generate generates the kubeconfig.
func (k *kubeconfig) generate(
	rootCA tls.CertKeyInterface,
	clientCertKey tls.CertKeyInterface,
	installConfig *types.InstallConfig,
	userName string,
	kubeconfigPath string,
) error {
	k.Config = &clientcmd.Config{
		Clusters: []clientcmd.NamedCluster{
			{
				Name: installConfig.ObjectMeta.Name,
				Cluster: clientcmd.Cluster{
					Server: fmt.Sprintf("https://%s-api.%s:6443", installConfig.ObjectMeta.Name, installConfig.BaseDomain),
					CertificateAuthorityData: []byte(rootCA.Cert()),
				},
			},
		},
		AuthInfos: []clientcmd.NamedAuthInfo{
			{
				Name: userName,
				AuthInfo: clientcmd.AuthInfo{
					ClientCertificateData: []byte(clientCertKey.Cert()),
					ClientKeyData:         []byte(clientCertKey.Key()),
				},
			},
		},
		Contexts: []clientcmd.NamedContext{
			{
				Name: userName,
				Context: clientcmd.Context{
					Cluster:  installConfig.ObjectMeta.Name,
					AuthInfo: userName,
				},
			},
		},
		CurrentContext: userName,
	}

	data, err := yaml.Marshal(k.Config)
	if err != nil {
		return errors.Wrap(err, "failed to Marshal kubeconfig")
	}

	k.File = &asset.File{
		Filename: kubeconfigPath,
		Data:     data,
	}

	return nil
}

// Files returns the files generated by the asset.
func (k *kubeconfig) Files() []*asset.File {
	if k.File != nil {
		return []*asset.File{k.File}
	}
	return []*asset.File{}
}

// load returns the kubeconfig from disk.
func (k *kubeconfig) load(f asset.FileFetcher, name string) (found bool, err error) {
	file, err := f.FetchByName(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	config := &clientcmd.Config{}
	if err := yaml.Unmarshal(file.Data, config); err != nil {
		return false, errors.Wrapf(err, "failed to unmarshal")
	}

	k.File, k.Config = file, config
	return true, nil
}
