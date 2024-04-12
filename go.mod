module rtcagent

go 1.21

toolchain go1.22.1

require (
	github.com/adubovikov/ebpfmanager v0.4.7
	github.com/cilium/ebpf v0.12.3
	github.com/gogo/protobuf v1.3.2
	github.com/google/gopacket v1.1.19
	github.com/shuLhan/go-bindata v4.0.0+incompatible
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/arch v0.3.0
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa
	golang.org/x/sys v0.14.1-0.20231108175955-e4099bfacb8c
)

require (
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/florianl/go-tc v0.4.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/josharian/native v1.0.0 // indirect
	github.com/mdlayher/netlink v1.7.1 // indirect
	github.com/mdlayher/socket v0.4.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/vishvananda/netlink v1.1.0 // indirect
	github.com/vishvananda/netns v0.0.0-20210104183010-2eb08e3e575f // indirect
	golang.org/x/exp v0.0.0-20230224173230-c95f2b4c22f2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)

replace github.com/google/gopacket v1.1.19 => github.com/cfc4n/gopacket v1.1.20
