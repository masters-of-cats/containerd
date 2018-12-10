#!/usr/bin/env bash
set -euox pipefail

testname=$1

killall -9 containerd-shim || true
killall -9 containerd-shim-runc-v1 || true

umount  /run/containerd-test/io.containerd.runtime.v1.linux/testing/$testname/rootfs || true
umount /run/containerd-test/io.containerd.runtime.v2.task/testing/$testname/rootfs || true

rm -rf /run/containerd-test/io.containerd.runtime.v1.linux/testing/$testname || true
rm -rf /run/containerd-test/io.containerd.runtime.v2.task/testing/$testname || true

rm -rf /run/containerd-test || true
rm -rf /run/containerd/runc/testing || true

rm -rf /run/containerd-test/io.containerd.runtime.v1.linux/testing/$testname


