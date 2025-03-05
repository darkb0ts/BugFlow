package utils

var blacklistedCommands = []string{
	// File and Directory Deletion/Destruction
	"rm",     // Blocks 'rm' alone
	"rm -rf", // Blocks 'rm -rf' (recursive force delete)
	"del",    // Blocks Windows file deletion
	"rmdir",  // Blocks Windows directory removal
	"unlink", // Blocks Unix file removal

	// Remote Access and Networking
	"ssh",    // Blocks SSH commands
	"telnet", // Blocks Telnet connections
	"nc",     // Blocks netcat (potential backdoor tool)
	"curl",   // Blocks curl (could fetch malicious content)
	"wget",   // Blocks wget (could download harmful files)
	"scp",    // Blocks secure copy over SSH

	// System Modification and Privilege Escalation
	"sudo",   // Blocks sudo commands
	"su",     // Blocks switching to superuser
	"chmod",  // Blocks changing file permissions
	"chown",  // Blocks changing file ownership
	"passwd", // Blocks password changes

	// Disk and Filesystem Operations
	"dd",     // Blocks dangerous disk-wiping commands
	"mkfs",   // Blocks filesystem creation (e.g., mkfs.ext4)
	"fdisk",  // Blocks disk partitioning
	"parted", // Blocks partition editing
	"format", // Blocks Windows disk formatting
	"fsck",   // Blocks filesystem checking (could corrupt if misused)

	// System Shutdown/Reboot
	"reboot",   // Blocks system reboot
	"shutdown", // Blocks system shutdown
	"halt",     // Blocks system halt
	"poweroff", // Blocks power off

	// Process and Resource Manipulation
	"kill",    // Blocks killing processes
	"killall", // Blocks killing processes by name
	"pkill",   // Blocks killing processes by pattern
	"nice",    // Blocks changing process priority
	"renice",  // Blocks adjusting running process priority

	// Potentially Malicious Tools
	"bash",   // Blocks direct bash invocation (if scripts are a concern)
	"sh",     // Blocks direct shell invocation
	"python", // Blocks Python (could run malicious scripts)
	"perl",   // Blocks Perl (could run malicious scripts)
	"ruby",   // Blocks Ruby (could run malicious scripts)

	// System Configuration and Info Exposure
	"ifconfig",        // Blocks network config changes
	"ip",              // Blocks IP command (e.g., ip link)
	"route",           // Blocks routing table changes
	"iptables",        // Blocks firewall rule changes
	"ufw",             // Blocks uncomplicated firewall changes
	"cat /etc/shadow", // Blocks reading sensitive files
	"whoami",          // Blocks user identity checks (optional)
}
