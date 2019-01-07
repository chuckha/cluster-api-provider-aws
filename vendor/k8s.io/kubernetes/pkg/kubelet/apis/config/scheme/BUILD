load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["scheme.go"],
    importmap = "sigs.k8s.io/cluster-api-provider-aws/vendor/k8s.io/kubernetes/pkg/kubelet/apis/config/scheme",
    importpath = "k8s.io/kubernetes/pkg/kubelet/apis/config/scheme",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/config:go_default_library",
        "//vendor/k8s.io/kubernetes/pkg/kubelet/apis/config/v1beta1:go_default_library",
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
