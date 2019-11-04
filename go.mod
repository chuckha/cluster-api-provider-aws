module sigs.k8s.io/cluster-api-provider-aws

go 1.12

require (
	github.com/aws/aws-sdk-go v1.25.16
	github.com/awslabs/goformation/v3 v3.0.0
	github.com/go-logr/logr v0.1.0
	github.com/golang/mock v1.2.0
	github.com/onsi/ginkgo v1.10.1
	github.com/onsi/gomega v1.7.0
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	golang.org/x/net v0.0.0-20191021144547-ec77196f6094
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/apiserver v0.0.0-20190918200908-1e17798da8c1
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/cluster-bootstrap v0.0.0-20190516232516-d7d78ab2cfe7
	k8s.io/component-base v0.0.0-20190918200425-ed2f0867c778
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20190809000727-6c36bc71fc4a
	sigs.k8s.io/cluster-api v0.2.6-0.20191018152358-b01a13ef030f
	sigs.k8s.io/cluster-api/test/infrastructure/docker v0.0.0-20191015161153-1dca5eea774a
	sigs.k8s.io/controller-runtime v0.3.0
)

replace sigs.k8s.io/cluster-api => ../cluster-api

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
