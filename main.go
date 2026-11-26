package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("wrong")
	}
}

func run() {
	fmt.Printf("parent running %v as pid %d\n", os.Args[2], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...) 
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
 
	must(cmd.Run())
}

func child() {
	fmt.Printf("child running %v as pid %d\n", os.Args[2], os.Getpid())

	cg()

	must(syscall.Sethostname([]byte("made container")))
	must(syscall.Chroot("/container-fs"))
	must(syscall.Chdir("/"))

	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd := exec.Command(os.Args[2], os.Args[3:]...) 
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
 
	must(cmd.Run())

	must(syscall.Unmount("/proc", 0))
}

func cg() {
	cgroup := "/sys/fs/cgroup/demo-cont"
	os.Mkdir(cgroup, 0755)
	must(os.WriteFile(filepath.Join(cgroup, "pids.max"), []byte("20"), 0700))
	must(os.WriteFile(filepath.Join(cgroup, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}