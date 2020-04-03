package ptrace

import (
	"fmt"
	"base"
	"syscall"
)

var AllowSysCall = []uint64{syscall.SYS_READ, syscall.SYS_WRITE, syscall.SYS_OPEN, syscall.SYS_FSTAT, syscall.SYS_LSEEK,
	syscall.SYS_MMAP, syscall.SYS_MUNMAP, syscall.SYS_BRK, syscall.SYS_WRITEV, syscall.SYS_ACCESS, syscall.SYS_EXECVE, syscall.SYS_UNAME,
	syscall.SYS_READLINK, syscall.SYS_ARCH_PRCTL, syscall.SYS_EXIT_GROUP, syscall.SYS_MQ_OPEN, syscall.SYS_IOPRIO_GET,
	syscall.SYS_TIME, syscall.SYS_READ, syscall.SYS_UNAME, syscall.SYS_WRITE, syscall.SYS_OPEN,
    syscall.SYS_CLOSE, syscall.SYS_EXECVE, syscall.SYS_ACCESS, syscall.SYS_BRK, syscall.SYS_MUNMAP,
    syscall.SYS_MPROTECT, syscall.SYS_MMAP, syscall.SYS_FSTAT, syscall.SYS_SET_THREAD_AREA, syscall.SYS_ARCH_PRCTL}

var DangerSysCall = []uint64{syscall.SYS_GETPID}

type SyscallCounter []int

const maxSyscalls = 303

func IsAllowSysCall(id uint64) int64 {
	if id > syscall.SYS_PRLIMIT64 {
		return base.RuntimeError
	}
	for _, v := range AllowSysCall {
		if v == id {
			return base.Normal
		}
	}
	return base.Danger
}

func (s SyscallCounter) Init() SyscallCounter {
	s = make(SyscallCounter, maxSyscalls)
	return s
}

func (s SyscallCounter) Inc(syscallID uint64) error {
	if syscallID > maxSyscalls {
		return fmt.Errorf("invalid syscall ID (%x)", syscallID)
	}

	s[syscallID]++
	return nil
}

/*func (s SyscallCounter) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for k, v := range s {
		if v > 0 {
			name, _ := seccomp.ScmpSyscall(k).GetName()
			fmt.Fprintf(w, "%d\t%s\n", v, name)
		}
	}
	w.Flush()
}

func (s SyscallCounter) GetName(syscallID uint64) string {
	name, _ := seccomp.ScmpSyscall(syscallID).GetName()
	return name
}*/