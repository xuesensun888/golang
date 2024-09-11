package internal

import (
	"os"
	"syscall"
)

// Level of process LOCK
type FLock struct {
	Fp   string
	fd   uintptr
	Flag int
	Perm os.FileMode
	Lock *syscall.Flock_t
}

func NewLock(fp string) *FLock {
	// 定义文件锁的信息
	lock := syscall.Flock_t{
		Whence: 0,
		Start:  0,
		Len:    0,
		Pid:    int32(os.Getpid()),
	}
	return &FLock{
		Fp:   fp,
		Lock: &lock,
	}
}

func (fl *FLock) WLock() error {
	file, err := os.OpenFile(fl.Fp, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件句柄
	fl.fd = file.Fd()
	// 写锁
	fl.Lock.Type = syscall.F_WRLCK
	// 加锁, 阻塞锁
	err = syscall.FcntlFlock(fl.fd, syscall.F_SETLKW, fl.Lock)

	return err
}

func (fl *FLock) Unlock() error {
	fl.Lock.Type = syscall.F_UNLCK
	err := syscall.FcntlFlock(fl.fd, syscall.F_SETLK, fl.Lock)
	return err
}
