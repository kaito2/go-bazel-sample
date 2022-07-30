# Generate `BUILD.bazel` files with gazelle

## Generate files

```
bazel run //:gazelle
```

## Build

```
bazel build //...
```

## Run

```
$ bazel run //app1/cmd:cmd
INFO: Analyzed target //app1/cmd:cmd (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //app1/cmd:cmd up-to-date:
  bazel-bin/app1/cmd/main
INFO: Elapsed time: 0.177s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
Hello Gazelle
```
