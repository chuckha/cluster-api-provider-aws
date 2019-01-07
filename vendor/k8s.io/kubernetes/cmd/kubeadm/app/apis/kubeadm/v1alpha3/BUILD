load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "bootstraptokenstring.go",
        "conversion.go",
        "defaults.go",
        "defaults_unix.go",
        "defaults_windows.go",
        "doc.go",
        "register.go",
        "types.go",
        "zz_generated.conversion.go",
        "zz_generated.deepcopy.go",
        "zz_generated.defaults.go",
    ],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha3",
    importpath = "k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/conversion:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//vendor/k8s.io/client-go/tools/bootstrap/token/api:go_default_library",
        "//vendor/k8s.io/client-go/tools/bootstrap/token/util:go_default_library",
        "//vendor/k8s.io/kube-proxy/config/v1alpha1:go_default_library",
        "//vendor/k8s.io/kubelet/config/v1beta1:go_default_library",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm:go_default_library",
        "//vendor/k8s.io/kubernetes/cmd/kubeadm/app/constants:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/config:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/config/scheme:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/proxy/apis/config:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/proxy/apis/config/scheme:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
