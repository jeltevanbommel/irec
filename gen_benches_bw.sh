#!/bin/bash
GAZELLE_MODE=fix
GAZELLE_DIRS=.
dirName=$(date +'%m-%d-%Y-%H%M%S')
mkdir -p "bench_results/latency/${dirName}/"

mkdir -p "bench_bin/"
# BW BENCHMARKS:

# UBPF
bazel run //:gazelle --config=quiet -- update -build_tags=timing -mode=$GAZELLE_MODE -go_naming_convention go_default_library $GAZELLE_DIRS
rm -f bin/*
bazel build //:scion //:scion-ci --define gotags=timing
tar -kxf bazel-bin/scion.tar -C bin
tar -kxf bazel-bin/scion-ci.tar -C bin
cp bin/bench bench_bin/ubpf_bw_ire_bench
cp bin/rac bench_bin/ubpf_bw_ire_rac

# UBPF JIT
bazel run //:gazelle --config=quiet -- update -build_tags=ubpfjit,timing -mode=$GAZELLE_MODE -go_naming_convention go_default_library $GAZELLE_DIRS
rm -f bin/*
bazel build //:scion //:scion-ci --define gotags=ubpfjit,timing
tar -kxf bazel-bin/scion.tar -C bin
tar -kxf bazel-bin/scion-ci.tar -C bin
cp bin/bench bench_bin/ubpfjit_bw_ire_bench
cp bin/rac bench_bin/ubpfjit_bw_ire_rac

# WASM
bazel run //:gazelle --config=quiet -- update -build_tags=wa,timing -mode=$GAZELLE_MODE -go_naming_convention go_default_library $GAZELLE_DIRS
rm -f bin/*
bazel build //:scion //:scion-ci --define gotags=wa,timing
tar -kxf bazel-bin/scion.tar -C bin
tar -kxf bazel-bin/scion-ci.tar -C bin
cp bin/bench bench_bin/wa_bw_ire_bench
cp bin/rac bench_bin/wa_bw_ire_rac

# WASM OPTIMIZATIONS STRIPPED
bazel run //:gazelle --config=quiet -- update -build_tags=waopt,timing -mode=$GAZELLE_MODE -go_naming_convention go_default_library $GAZELLE_DIRS
rm -f bin/*
bazel build //:scion //:scion-ci --define gotags=waopt,timing
tar -kxf bazel-bin/scion.tar -C bin
tar -kxf bazel-bin/scion-ci.tar -C bin
cp bin/bench bench_bin/waopt_bw_ire_bench
cp bin/rac bench_bin/waopt_bw_ire_rac

# NATIVE
bazel run //:gazelle --config=quiet -- update -build_tags=native,timing -mode=$GAZELLE_MODE -go_naming_convention go_default_library $GAZELLE_DIRS
rm -f bin/*
bazel build //:scion //:scion-ci --define gotags=native,timing
tar -kxf bazel-bin/scion.tar -C bin
tar -kxf bazel-bin/scion-ci.tar -C bin
cp bin/bench bench_bin/native_bw_ire_bench
cp bin/rac bench_bin/native_bw_ire_rac
