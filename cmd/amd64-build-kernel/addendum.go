package main

var configAddendumMap = map[string]string{
	"CONFIG_LOCALVERSION": "\"-v1-thatwebsite\"",

	"CONFIG_IPV6":          "y",
	"CONFIG_DYNAMIC_DEBUG": "y",

	// Thermals
	"CONFIG_ACPI_THERMAL":         "y",
	"CONFIG_THERMAL":              "y",
	"CONFIG_MLXSW_CORE_THERMAL":   "y",
	"CONFIG_THERMAL_NETLINK":      "y",
	"CONFIG_INTEL_HFI_THERMAL":    "y",
	"CONFIG_DEVFREQ_THERMAL":      "y",
	"CONFIG_INTEL_TH_ACPI":        "y",
	"CONFIG_X86_PKG_TEMP_THERMAL": "y",

	// # For Squashfs (root file system):
	"CONFIG_SQUASHFS":                     "y",
	"CONFIG_SQUASHFS_FILE_CACHE":          "y",
	"CONFIG_SQUASHFS_DECOMP_MULTI_PERCPU": "y",
	"CONFIG_SQUASHFS_ZLIB":                "y",
	"CONFIG_SQUASHFS_FRAGMENT_CACHE_SIZE": "3",

	// # For a console on HDMI:
	// # TODO: the simpledrm driver just does not work for me. the ASRock logo never disappears from HDMI
	// # [    0.364059] [drm] Initialized simpledrm 1.0.0 20200625 for simple-framebuffer.0 on minor 0
	// # "CONFIG_DRM_SIMPLEDRM": "y",
	// # "CONFIG_X86_SYSFB": "y",
	// #
	// # Whereas with (working) efifb, I see:
	// # [    0.460084] efifb: probing for efifb
	// # [    0.460096] efifb: framebuffer at 0xe9000000, using 3072k, total 3072k
	// # [    0.460099] efifb: mode is 1024x768x32, "linelength": "4096, "pages": 1",
	// # [    0.460101] efifb: scrolling: redraw
	// # [    0.460103] efifb: Truecolor: "size": "8:8:8:8, "shift": 24:16:8:0",
	"CONFIG_DRM_SIMPLEDRM": "n",
	"CONFIG_X86_SYSFB":     "n",
	"CONFIG_FB":            "y",
	"CONFIG_FB_EFI":        "y",
	"CONFIG_FB_SIMPLE":     "y",

	// # For FUSE (for cpu(1)):
	"CONFIG_FUSE_FS": "y",

	// # For using github.com/vishvananda/netlink
	"CONFIG_NETFILTER_NETLINK_QUEUE": "y",
	"CONFIG_XFRM_USER":               "y",

	// # For nftables:
	"CONFIG_NF_TABLES":              "y",
	"CONFIG_NF_NAT_IPV4":            "y",
	"CONFIG_NF_NAT_MASQUERADE_IPV4": "y",
	"CONFIG_NFT_PAYLOAD":            "y",
	"CONFIG_NFT_EXTHDR":             "y",
	"CONFIG_NFT_META":               "y",
	"CONFIG_NFT_CT":                 "y",
	"CONFIG_NFT_RBTREE":             "y",
	"CONFIG_NFT_HASH":               "y",
	"CONFIG_NFT_COUNTER":            "y",
	"CONFIG_NFT_LOG":                "y",
	"CONFIG_NFT_LIMIT":              "y",
	"CONFIG_NFT_NAT":                "y",
	"CONFIG_NFT_COMPAT":             "y",
	"CONFIG_NFT_MASQ":               "y",
	"CONFIG_NFT_MASQ_IPV4":          "y",
	"CONFIG_NFT_REDIR":              "y",
	"CONFIG_NFT_REJECT":             "y",
	"CONFIG_NF_TABLES_IPV4":         "y",
	"CONFIG_NFT_REJECT_IPV4":        "y",
	"CONFIG_NFT_CHAIN_ROUTE_IPV4":   "y",
	"CONFIG_NFT_CHAIN_NAT_IPV4":     "y",
	"CONFIG_NF_TABLES_IPV6":         "y",
	"CONFIG_NFT_CHAIN_ROUTE_IPV6":   "y",
	"CONFIG_NFT_OBJREF":             "y",
	"CONFIG_NFT_DUP_IPV4":           "y",
	"CONFIG_NFT_FIB_IPV4":           "y",
	"CONFIG_NFT_DUP_IPV6":           "y",
	"CONFIG_NFT_FIB_IPV6":           "y",

	// # Explicitly disable nftables helper modules to prevent NAT slipstreaming attacks:
	// # https://samy.pl/slipstream/
	"CONFIG_NF_CONNTRACK_AMANDA":     "n",
	"CONFIG_NF_CONNTRACK_FTP":        "n",
	"CONFIG_NF_CONNTRACK_H323":       "n",
	"CONFIG_NF_CONNTRACK_IRC":        "n",
	"CONFIG_NF_CONNTRACK_NETBIOS_NS": "n",
	"CONFIG_NF_CONNTRACK_SNMP":       "n",
	"CONFIG_NF_CONNTRACK_PPTP":       "n",
	"CONFIG_NF_CONNTRACK_SANE":       "n",
	"CONFIG_NF_CONNTRACK_SIP":        "n",
	"CONFIG_NF_CONNTRACK_TFTP":       "n",

	// # For using USB mass storage
	"CONFIG_USB_EHCI_HCD": "y",
	"CONFIG_USB_XHCI_HCD": "y",
	"CONFIG_USB_DEVICEFS": "y",
	"CONFIG_USB_STORAGE":  "y",

	// # For NVMe storage
	"CONFIG_NVME_CORE":      "y",
	"CONFIG_BLK_DEV_NVME":   "y",
	"CONFIG_NVME_MULTIPATH": "y",
	// # "CONFIG_NVME_HWMON": "y",
	"CONFIG_NVME_TARGET_PASSTHRU": "y",

	// # For https://www.fs.com/products/75602.html and https://www.fs.com/products/75603.html network cards:
	"CONFIG_I40E": "y",

	// # For apu2c4 ethernet ports
	"CONFIG_IGB": "y",

        // # For RTL USB to Ethernet port
        "CONFIG_USB_RTL8152": "y",

	// # For Qualcomm Atheros Fast Ethernet
	"CONFIG_ATL1C": "y",
	"CONFIG_ATL2": "y",

	// # For Intel I225 ethernet ports (ASRock B550 Taichi):
	"CONFIG_IGC": "y",

	// # For /proc/config.gz
	"CONFIG_IKCONFIG":      "y",
	"CONFIG_IKCONFIG_PROC": "y",

	// # For kexec
	"CONFIG_KEXEC_FILE": "y",

	// # For apu2c4 watchdog
	"CONFIG_SP5100_TCO": "y",

	// # For WireGuard
	"CONFIG_NET_UDP_TUNNEL": "y",
	"CONFIG_WIREGUARD":      "y",

	// # For traffic shaping using tc:
	"CONFIG_NET_SCH_TBF": "y",

	// # For measuring CPU temperature:
	"CONFIG_SENSORS_K10TEMP": "y",

	// # For measuring non-CPU temperature and fan speeds:
	"CONFIG_SENSORS_NCT6683": "y",

	// # For Corsair Commander Pro fan controller:
	"CONFIG_SENSORS_CORSAIR_CPRO": "y",

	// # For iproute2's ss(8):
	"CONFIG_INET_DIAG": "y",

	// # For macvlan ethernet devices:
	"CONFIG_MACVLAN": "y",

	// # For virtio drivers (for qemu):
	"CONFIG_VIRTIO_PCI":     "y",
	"CONFIG_VIRTIO_BALLOON": "y",
	"CONFIG_VIRTIO_BLK":     "y",
	"CONFIG_VIRTIO_NET":     "y",
	"CONFIG_VIRTIO":         "y",
	"CONFIG_VIRTIO_RING":    "y",
	// # For watchdog within qemu:
	"CONFIG_I6300ESB_WDT": "y",

	// # For bridge ethernet devices:
	// "CONFIG_BRIDGE": "y", // already added in podman section

	"CONFIG_EFIVAR_FS": "y",

	// # For Ryzen CPUs:
	"CONFIG_X86_AMD_PLATFORM_DEVICE":        "y",
	"CONFIG_CPU_FREQ_DEFAULT_GOV_POWERSAVE": "y",
	"CONFIG_CPU_FREQ_GOV_POWERSAVE":         "y",
	"CONFIG_X86_POWERNOW_K8":                "y",
	"CONFIG_X86_AMD_FREQ_SENSITIVITY":       "y",

	// Power Cap for RAPL
	"CONFIG_POWERCAP":               "y",
	"CONFIG_PERF_EVENTS_INTEL_RAPL": "y",
	"CONFIG_PROC_THERMAL_MMIO_RAPL": "y",
	"CONFIG_INTEL_RAPL_CORE":        "y",
	"CONFIG_INTEL_RAPL":             "y",

	// # Include hardware interrupt CPU usage in /proc/stat CPU time reporting:
	"CONFIG_IRQ_TIME_ACCOUNTING": "y",

	// # For tun devices, see https://www.kernel.org/doc/Documentation/networking/tuntap.txt
	"CONFIG_TUN": "y",

	// # For runc:
	"CONFIG_BPF_SYSCALL":      "y",
	"CONFIG_CGROUP_FREEZER":   "y",
	"CONFIG_CGROUP_BPF":       "y",
	"CONFIG_SOCK_CGROUP_DATA": "y",
	"CONFIG_NET_SOCK_MSG":     "y",

	// # For podman:
	"CONFIG_OVERLAY_FS":                     "y",
	"CONFIG_BRIDGE":                         "y",
	"CONFIG_VETH":                           "y",
	"CONFIG_NETFILTER_ADVANCED":             "y",
	"CONFIG_NETFILTER_XT_MATCH_COMMENT":     "y",
	"CONFIG_IP_NF_NAT":                      "y",
	"CONFIG_IP_NF_TARGET_MASQUERADE":        "y",
	"CONFIG_NETFILTER_XT_NAT":               "y",
	"CONFIG_NETFILTER_XT_TARGET_MASQUERADE": "y",
	"CONFIG_NETFILTER_XT_MATCH_MULTIPORT":   "y",
	"CONFIG_NETFILTER_XT_MARK":              "y",
	"CONFIG_CGROUP_PIDS":                    "y",
	"CONFIG_MEMCG":                          "y",

	// # Enable TCP BBR as default congestion control
	"CONFIG_TCP_CONG_BBR":     "y",
	"CONFIG_DEFAULT_BBR":      "y",
	"CONFIG_DEFAULT_TCP_CONG": "bbr",

	// # For different FS
	"CONFIG_EXFAT_FS":            "y",
	"CONFIG_NTFS3_FS":            "y",
	"CONFIG_NTFS3_64BIT_CLUSTER": "y",
	"CONFIG_NTFS3_LZX_XPRESS":    "y",
	"CONFIG_NTFS3_FS_POSIX_ACL":  "y",
	"CONFIG_BTRFS_FS":            "y",
	"CONFIG_XFS_FS":              "y",
	"CONFIG_XFS_SUPPORT_V4":      "y",

	// # For HWMON
	"CONFIG_NVME_HWMON":                  "y",
	"CONFIG_SCSI_UFS_HWMON":              "y",
	"CONFIG_TIGON3_HWMON":                "y",
	"CONFIG_BNXT_HWMON":                  "y",
	"CONFIG_BE2NET_HWMON":                "y",
	"CONFIG_IGB_HWMON":                   "y",
	"CONFIG_IXGBE_HWMON":                 "y",
	"CONFIG_MLXSW_CORE_HWMON":            "y",
	"CONFIG_QLCNIC_HWMON":                "y",
	"CONFIG_POWER_SUPPLY_HWMON":          "y",
	"CONFIG_HWMON":                       "y",
	"CONFIG_HWMON_VID":                   "y",
	"CONFIG_SENSORS_IIO_HWMON":           "y",
	"CONFIG_SENSORS_MENF21BMC_HWMON":     "y",
	"CONFIG_SENSORS_INTEL_M10_BMC_HWMON": "y",
	"CONFIG_THERMAL_HWMON":               "y",
	"CONFIG_RTC_DRV_DS3232_HWMON":        "y",
	"CONFIG_RTC_DRV_RV3029_HWMON":        "y",

	"CONFIG_ATH9K":         "y",
	"CONFIG_ATH9K_AHB":     "y",
	"CONFIG_RTW88":         "m",
	"CONFIG_RTW88_CORE":    "m",
	"CONFIG_RTW88_PCI":     "m",
	"CONFIG_RTW88_8822B":   "m",
	"CONFIG_RTW88_8822C":   "m",
	"CONFIG_RTW88_8723D":   "m",
	"CONFIG_RTW88_8821C":   "m",
	"CONFIG_RTW88_8822BE":  "m",
	"CONFIG_RTW88_8822CE":  "m",
	"CONFIG_RTW88_8723DE":  "m",
	"CONFIG_RTW88_8821CE":  "m",
	"CONFIG_RTW88_DEBUG":   "m",
	"CONFIG_RTW88_DEBUGFS": "m",
	"CONFIG_RTW89":         "m",
	"CONFIG_RTW89_CORE":    "m",
	"CONFIG_RTW89_PCI":     "m",
	"CONFIG_RTW89_8852A":   "m",
	"CONFIG_RTW89_8852AE":  "m",
	"CONFIG_RTW89_DEBUG":   "m",

        // # Extras
        "CONFIG_IDEAPAD_LAPTOP": "y",

	// # Linux 6.1:
	// In file included from <command-line>:0:0:
	// drivers/gpu/drm/i915/i915_sw_fence_work.c: In function 'dma_fence_work_init':
	// drivers/gpu/drm/i915/i915_sw_fence.h:57:20: error: the comparison will always evaluate as 'false' for the address of 'fence_notify' will never be NULL [-Werror=address]
	//   BUILD_BUG_ON((fn) == NULL);    \
	//                     ^
	// ././include/linux/compiler_types.h:337:9: note: in definition of macro '__compiletime_assert'
	//    if (!(condition))     \
	//          ^~~~~~~~~
	// ././include/linux/compiler_types.h:357:2: note: in expansion of macro '_compiletime_assert'
	//   _compiletime_assert(condition, msg, __compiletime_assert_, __COUNTER__)
	//   ^~~~~~~~~~~~~~~~~~~
	// ./include/linux/build_bug.h:39:37: note: in expansion of macro 'compiletime_assert'
	//  #define BUILD_BUG_ON_MSG(cond, msg) compiletime_assert(!(cond), msg)
	//                                      ^~~~~~~~~~~~~~~~~~
	// ./include/linux/build_bug.h:50:2: note: in expansion of macro 'BUILD_BUG_ON_MSG'
	//   BUILD_BUG_ON_MSG(condition, "BUILD_BUG_ON failed: " #condition)
	//   ^~~~~~~~~~~~~~~~
	// drivers/gpu/drm/i915/i915_sw_fence.h:57:2: note: in expansion of macro 'BUILD_BUG_ON'
	//   BUILD_BUG_ON((fn) == NULL);    \
	//   ^~~~~~~~~~~~
	// drivers/gpu/drm/i915/i915_sw_fence_work.c:89:2: note: in expansion of macro 'i915_sw_fence_init'
	//   i915_sw_fence_init(&f->chain, fence_notify);
	//   ^~~~~~~~~~~~~~~~~~
	// cc1: all warnings being treated as errors
	// make[5]: *** [drivers/gpu/drm/i915/i915_sw_fence_work.o] Error 1
	// make[4]: *** [drivers/gpu/drm/i915] Error 2
	// make[3]: *** [drivers/gpu/drm] Error 2
	"CONFIG_WERROR": "n",
}
