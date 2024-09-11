package service

import (
	"fmt"
	"kargo-deploy/conf"
	"kargo-deploy/handler"
	"kargo-deploy/log"

	"net"
	"os"
	"path"
	"sync"
)

func ListenAndServSocket(wg *sync.WaitGroup) {
	defer wg.Done()
	//mkdir socker

	fmt.Println("CREATE DIR")
	dir := path.Dir(conf.Socket)
	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return
		}
	}
	fmt.Println("create compelet")
	//clear residue socket
	_ = os.Remove(conf.Socket)
	fmt.Println("listening service starting...")
	listen, err := net.Listen("unix", conf.Socket)
	if err != nil {
		fmt.Println("net listen error is ", err)

	}
	defer listen.Close()

	_ = os.Chmod(conf.Socket, os.ModePerm)
	fmt.Println("service started")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error", err)
			continue
		}
		wg.Add(1)
		go handler.Handler(&conn, wg)
	}
}
func Update(wg *sync.WaitGroup) {
	log.Logger.Debug("entry updating .....")
	defer wg.Done()
	action := fmt.Sprintf("update/%s/deb", conf.App)
	conn, err := net.Dial("unix", conf.Socket)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if _, err := conn.Write([]byte(action)); err != nil {
		fmt.Println(err)
	}
}
