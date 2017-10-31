package gofsutil

import "context"

// FS provides many filesystem-specific functions, such as mount, format, etc.
type FS struct {

	// ScanEntry is the function used to process mount table entries.
	ScanEntry EntryScanFunc
}

// GetDiskFormat uses 'lsblk' to see if the given disk is unformatted.
func (fs *FS) GetDiskFormat(ctx context.Context, disk string) (string, error) {

	return fs.getDiskFormat(ctx, disk)
}

// FormatAndMount uses unix utils to format and mount the given disk.
func (fs *FS) FormatAndMount(
	ctx context.Context,
	source, target, fsType string,
	options ...string) error {

	return fs.formatAndMount(ctx, source, target, fsType, options)
}

// Mount mounts source to target as fstype with given options.
//
// The parameters 'source' and 'fstype' must be empty strings in case they
// are not required, e.g. for remount, or for an auto filesystem type where
// the kernel handles fstype automatically.
//
// The 'options' parameter is a list of options. Please see mount(8) for
// more information. If no options are required then please invoke Mount
// with an empty or nil argument.
func (fs *FS) Mount(source, target, fsType string, options ...string) error {
	return fs.mount(source, target, fsType, options)
}

// BindMount behaves like Mount was called with a "bind" flag set
// in the options list.
func (fs *FS) BindMount(source, target string, options ...string) error {
	if options == nil {
		options = []string{"bind"}
	} else {
		options = append(options, "bind")
	}
	return fs.mount(source, target, "", options)
}

// Unmount unmounts the target.
func (fs *FS) Unmount(target string) error {
	return fs.unmount(target)
}

// GetMounts returns a slice of all the mounted filesystems.
//
// * Linux hosts use mount_namespaces to obtain mount information.
//
//   Support for mount_namespaces was introduced to the Linux kernel
//   in 2.2.26 (http://man7.org/linux/man-pages/man5/proc.5.html) on
//   2004/02/04.
//
//   The kernel documents the contents of "/proc/<pid>/mountinfo" at
//   https://www.kernel.org/doc/Documentation/filesystems/proc.txt.
//
// * Darwin hosts parse the output of the "mount" command to obtain
//   mount information.
func (fs *FS) GetMounts(ctx context.Context) ([]Info, error) {
	return fs.getMounts(ctx)
}

// GetDevMounts returns a slice of all mounts for the provided device.
func (fs *FS) GetDevMounts(ctx context.Context, dev string) ([]Info, error) {
	return fs.getDevMounts(ctx, dev)
}
