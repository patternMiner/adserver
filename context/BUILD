load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "context",
    srcs = [
        "adset.go",
        "common.go",
        "context.go",
        "data_fetch_event.go"
    ],
    deps = ["@patternMiner_async//:async"],
    importpath = "github.com/patternMiner/adserver/context",
    visibility = ["//visibility:public"],
)

go_test(
    name = "context_test",
    srcs = ["context_test.go"],
    embed = [":context"],
    data = ["//data"],
    importpath = "github.com/patternMiner/adserver/context",
)