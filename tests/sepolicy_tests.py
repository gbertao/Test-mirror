from optparse import OptionParser
from optparse import Option, OptionValueError
import os
import policy
import re
import sys

#############################################################
# Tests
#############################################################
def TestRootfsViolations(pol):
    known_roots = [
        "/acct",
        "/adb_keys",
        "/apex",
        "/bin",
        "/bugreports",
        "/build.prop",
        "/cache/",
        "/charger",
        "/config",
        "/d",
        "/data",
        "/data_mirror",
        "/debug_ramdisk",
        "/default.prop",
        "/dev/",
        "/etc",
        "/fstab",
        "/init",
        "/lib",
        "/lost+found",
        "/mapping_sepolicy.cil",
        "/metadata/",
        "/mnt",
        "/nonplat_file_contexts",
        "/nonplat_hwservice_contexts",
        "/nonplat_property_contexts",
        "/nonplat_seapp_contexts",
        "/nonplat_sepolicy.cil",
        "/nonplat_service_contexts",
        "/odm/",
        "/odm_dlkm",
        "/oem/",
        "/plat_file_contexts",
        "/plat_hwservice_contexts",
        "/plat_keystore2_key_contexts",
        "/plat_property_contexts",
        "/plat_seapp_contexts",
        "/plat_sepolicy.cil",
        "/plat_service_contexts",
        "/postinstall",
        "/postinstall/",
        "/proc",
        "/product/",
        "/product_file_contexts",
        "/product_property_contexts",
        "/res",
        "/sbin",
        "/sdcard",
        "/sdcard/",
        "/seapp_contexts",
        "/second_stage_resources",
        "/selinux_version",
        "/sepolicy",
        "/sepolicy",
        "/sys",
        "/system/",
        "/ts_snap.txt", # FIXME: from cuttlefish
        "/ueventd",
        "/vendor",
        "/vendor_dlkm",
        "/vendor_file_contexts",
        "/vendor_hwservice_contexts",
        "/vendor_property_contexts",
        "/vendor_seapp_contexts",
        "/vendor_service_contexts",
        "/verity_key",
        "/vndservice_contexts",
    ]
    return pol.AssertPathTypesHaveAttr(["/"], known_roots, "no_fs_type")

def TestDataTypeViolations(pol):
    return pol.AssertPathTypesHaveAttr(["/data/"], [], "data_file_type")

def TestSystemTypeViolations(pol):
    partitions = ["/system/", "/system_ext/", "/product/"]
    exceptions = [
        # devices before treble don't have a vendor partition
        "/system/vendor/",

        # overlay files are mounted over vendor
        "/product/overlay/",
        "/product/vendor_overlay/",
        "/system/overlay/",
        "/system/product/overlay/",
        "/system/product/vendor_overlay/",
        "/system/system_ext/overlay/",
        "/system_ext/overlay/",
    ]

    return pol.AssertPathTypesHaveAttr(partitions, exceptions, "system_file_type")

def TestProcTypeViolations(pol):
    return pol.AssertGenfsFilesystemTypesHaveAttr("proc", "proc_type")

def TestSysfsTypeViolations(pol):
    ret = pol.AssertGenfsFilesystemTypesHaveAttr("sysfs", "sysfs_type")
    ret += pol.AssertPathTypesHaveAttr(["/sys/"], ["/sys/kernel/debug/",
                                    "/sys/kernel/tracing"], "sysfs_type")
    return ret

def TestDebugfsTypeViolations(pol):
    ret = pol.AssertGenfsFilesystemTypesHaveAttr("debugfs", "debugfs_type")
    ret += pol.AssertGenfsFilesystemTypesHaveAttr("tracefs", "debugfs_type")
    ret += pol.AssertPathTypesHaveAttr(["/sys/kernel/debug/",
                                    "/sys/kernel/tracing"], [], "debugfs_type")
    return ret

def TestVendorTypeViolations(pol):
    partitions = ["/vendor/", "/odm/"]
    exceptions = [
        "/vendor/etc/selinux/",
        "/vendor/odm/etc/selinux/",
        "/odm/etc/selinux/",
    ]
    return pol.AssertPathTypesHaveAttr(partitions, exceptions, "vendor_file_type")

def TestCoreDataTypeViolations(pol):
    return pol.AssertPathTypesHaveAttr(["/data/"], ["/data/vendor",
            "/data/vendor_ce", "/data/vendor_de"], "core_data_file_type")

def TestPropertyTypeViolations(pol):
    return pol.AssertPropertyOwnersAreExclusive()


###
# extend OptionParser to allow the same option flag to be used multiple times.
# This is used to allow multiple file_contexts files and tests to be
# specified.
#
class MultipleOption(Option):
    ACTIONS = Option.ACTIONS + ("extend",)
    STORE_ACTIONS = Option.STORE_ACTIONS + ("extend",)
    TYPED_ACTIONS = Option.TYPED_ACTIONS + ("extend",)
    ALWAYS_TYPED_ACTIONS = Option.ALWAYS_TYPED_ACTIONS + ("extend",)

    def take_action(self, action, dest, opt, value, values, parser):
        if action == "extend":
            values.ensure_value(dest, []).append(value)
        else:
            Option.take_action(self, action, dest, opt, value, values, parser)

Tests = [
    "TestDataTypeViolators",
    "TestProcTypeViolations",
    "TestSysfsTypeViolations",
    "TestSystemTypeViolators",
    "TestDebugfsTypeViolations",
    "TestVendorTypeViolations",
    "TestCoreDataTypeViolations",
    "TestPropertyTypeViolations"
]

if __name__ == '__main__':
    usage = "sepolicy_tests -l $(ANDROID_HOST_OUT)/lib64/libsepolwrap.so "
    usage += "-f vendor_file_contexts -f "
    usage +="plat_file_contexts -p policy [--test test] [--help]"
    parser = OptionParser(option_class=MultipleOption, usage=usage)
    parser.add_option("-f", "--file_contexts", dest="file_contexts",
            metavar="FILE", action="extend", type="string")
    parser.add_option("-p", "--policy", dest="policy", metavar="FILE")
    parser.add_option("-l", "--library-path", dest="libpath", metavar="FILE")
    parser.add_option("-t", "--test", dest="test", action="extend",
            help="Test options include "+str(Tests))

    (options, args) = parser.parse_args()

    if not options.libpath:
        sys.exit("Must specify path to libsepolwrap library\n" + parser.usage)
    if not os.path.exists(options.libpath):
        sys.exit("Error: library-path " + options.libpath + " does not exist\n"
                + parser.usage)

    if not options.policy:
        sys.exit("Must specify monolithic policy file\n" + parser.usage)
    if not os.path.exists(options.policy):
        sys.exit("Error: policy file " + options.policy + " does not exist\n"
                + parser.usage)

    if not options.file_contexts:
        sys.exit("Error: Must specify file_contexts file(s)\n" + parser.usage)
    for f in options.file_contexts:
        if not os.path.exists(f):
            sys.exit("Error: File_contexts file " + f + " does not exist\n" +
                    parser.usage)

    pol = policy.Policy(options.policy, options.file_contexts, options.libpath)

    results = ""
    # If an individual test is not specified, run all tests.
    if options.test is None or "TestRootfsViolations" in options.test:
        results += TestRootfsViolations(pol)
    if options.test is None or "TestDataTypeViolations" in options.test:
        results += TestDataTypeViolations(pol)
    if options.test is None or "TestProcTypeViolations" in options.test:
        results += TestProcTypeViolations(pol)
    if options.test is None or "TestSysfsTypeViolations" in options.test:
        results += TestSysfsTypeViolations(pol)
    if options.test is None or "TestSystemTypeViolations" in options.test:
        results += TestSystemTypeViolations(pol)
    if options.test is None or "TestDebugfsTypeViolations" in options.test:
        results += TestDebugfsTypeViolations(pol)
    if options.test is None or "TestVendorTypeViolations" in options.test:
        results += TestVendorTypeViolations(pol)
    if options.test is None or "TestCoreDataTypeViolations" in options.test:
        results += TestCoreDataTypeViolations(pol)
    if options.test is None or "TestPropertyTypeViolations" in options.test:
        results += TestPropertyTypeViolations(pol)

    if len(results) > 0:
        sys.exit(results)
