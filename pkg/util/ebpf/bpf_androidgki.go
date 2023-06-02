//go:build androidgki
// +build androidgki

package ebpf

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

const (
	BootConfigPath      = "/proc/config.gz"
	ConfigDebugInfoBtf  = "CONFIG_DEBUG_INFO_BTF"
	SysKernelBtfVmlinux = "/sys/kernel/btf/vmlinux"
)

var (
	// use same list of locations as libbpf
	// https://android.googlesource.com/platform/external/libbpf/

	locations = []string{
		"/sys/kernel/btf/vmlinux",
	}
)

func GetSystemConfig() (map[string]string, error) {
	return getAndroidConfig(BootConfigPath)
}

func getAndroidConfig(filename string) (map[string]string, error) {
	var KernelConfig = make(map[string]string)
	// Open file bootConf.
	f, err := os.Open(filename)
	if err != nil {
		return KernelConfig, err
	}
	defer f.Close()

	// check if the file is gzipped
	var magic []byte
	var i int
	magic = make([]byte, 2)
	i, err = f.Read(magic)
	if err != nil {
		return KernelConfig, err
	}
	if i != 2 {
		return KernelConfig, fmt.Errorf("read %d bytes, expected 2", i)
	}

	var s *bufio.Scanner
	_, err = f.Seek(0, 0)
	if err != nil {
		return KernelConfig, err
	}

	var reader *gzip.Reader
	//magic number for gzip is 0x1f8b
	if magic[0] == 0x1f && magic[1] == 0x8b {
		// gzip file
		reader, err = gzip.NewReader(f)
		if err != nil {
			return KernelConfig, err
		}
		s = bufio.NewScanner(reader)
	} else {
		// not gzip file
		s = bufio.NewScanner(f)
	}

	if err = parse(s, KernelConfig); err != nil {
		return KernelConfig, err
	}
	return KernelConfig, nil
}

// IsContainedInCgroup returns true if the process is running in a container.
func IsContainer() (bool, error) {
	return false, nil
}
