[go程序会用几个CPU？](https://www.toutiao.com/i6872354717305930254)

```go
// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
func NumCPU() int {
	return int(ncpu)
}
```
ncpu在哪里赋值的呢？go/src/runtime目录运行`grep -Er 'ncpu = ' .`
```
Administrator@DESKTOP-E3H0GN4 MINGW64 /c/Go/src/runtime
$ grep -Er 'ncpu = ' .
./os3_solaris.go:       ncpu = getncpu()
./os_aix.go:    ncpu = int32(sysconf(__SC_NPROCESSORS_ONLN))
./os_darwin.go: ncpu = getncpu()
./os_dragonfly.go:      ncpu = getncpu()
./os_freebsd.go:        ncpu = getncpu()
./os_js.go:     ncpu = 1
./os_linux.go:  ncpu = getproccount()
./os_netbsd.go: ncpu = getncpu()
./os_openbsd.go:        ncpu = getncpu()
./os_plan9.go:          ncpu = 1
./os_plan9.go:  ncpu = getproccount()
./os_windows.go:        ncpu = getproccount()

```
可以看到不同系统的赋值方法

可以设置亲和性 让所有程序都在一些cpu上执行
/sys/fs/cgroup/cpuset/gocpu
# echo '0' > cpuset.mems

# echo '0-1' > cpuset.cpus

# echo $$ > tasks
