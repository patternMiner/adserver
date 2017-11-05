load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "handlers",
    srcs = [
        "ad_handler.go",
        "common.go",
        "info_handler.go",
    ],
    deps = ["//context"],
    importpath = "github.com/patternMiner/adserver/handlers",
    visibility = ["//visibility:public"],
)
