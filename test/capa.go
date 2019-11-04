package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1alpha2"
)

type Generator struct {
	Ref     string
	Version string
}

func (g *Generator) GetName() string {
	return fmt.Sprintf("Cluster API Provider AWS version %s", g.Version)
}

func (g *Generator) releaseYAMLPath() string {
	return fmt.Sprintf("https://github.com/kubernetes-sigs/cluster-api-provider-aws/releases/download/%s/infrastructure-components.yaml", g.Version)
}

func (g *Generator) ProviderComponents() ([]byte, error) {
	resp, err := http.Get(g.releaseYAMLPath())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()
	if os.Getenv("AWS_B64ENCODED_CREDENTIALS") == "" {
		fmt.Fprintln(os.Stderr, "AWS_B64ENCODED_CREDENTIALS must be set")
		fmt.Fprintln(os.Stderr, "export AWS_B64ENCODED_CREDENTIALS=$(clusterawsadm alpha bootstrap encode-aws-credentials)")
		return nil, errors.New("AWS_B64ENCODED_CREDENTIALS is not set")
	}
	return []byte(os.ExpandEnv(string(out))), nil
}

func (g *Generator) CreateCluster() runtime.Object {
	return &v1alpha2.Cluster{}
}
