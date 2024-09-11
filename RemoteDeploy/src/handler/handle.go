package handler

import (
	"errors"
	"fmt"
	"io"
	"kargo-deploy/conf"
	"kargo-deploy/internal"
	"kargo-deploy/log"
	"kargo-deploy/request"
	"net"
	"strings"
	"sync"
	"time"
)

/*
update/funci/deb
*/
func Handler(conn *net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	p, err := io.ReadAll(*conn)
	if err != nil {
		fmt.Printf("get command %s from conn", string(p))
		return
	}
	s := string(p)

	cmd := strings.Split(s, "/")
	if len(cmd) != 3 {
		fmt.Println("invalid cmd format")
		return
	}
	if cmd[0] != "update" {
		fmt.Println("invalid action")
		return
	}
	switch cmd[0] {
	case "update":
		log.Logger.Infof("new request for update %s %s", cmd[2], cmd[1])
		Upgrade(cmd[1], cmd[2])

	default:
		fmt.Println("invalid aciton")
	}
}
func Upgrade(app, typ string) {
	log.Logger.Infoln("upgrade")
	var err error
	switch typ {
	case "deb":
		log.Logger.Infoln("upgrade(app)")
		err = UpgradeDeb(app)
		if err != nil {
			log.Logger.Error(err)
		}
	case "test":
	case "binary":
	default:
	}

}

func UpgradeDeb(app string) error {
	log.Logger.Infoln("upgradedeb")
	var resp *request.RespVersion
	var err error
	for i := 0; i < conf.Reconnect; i++ {

		log.Logger.Infoln("request headversionlatest")
		resp, err = request.HeadVersionLatest(conf.App)

		if err != nil {
			if i == conf.Reconnect-1 {
				return err
			}
			time.Sleep(time.Minute * 1)
			continue
		}
		if resp.ErrCode != 0 {
			err = fmt.Errorf("errcode: %d", resp.ErrCode)
			return err
		}
		if resp.Type != "deb" {
			err = fmt.Errorf("type of app not deb,is %s", resp.Type)
			return err
		}
		break
	}
	log.Logger.Infoln("resp result")
	di := internal.DebQuery(app)
	if di.Version == "" {
		err := errors.New("query deb info faild")
		return err
	}
	//不需要更新
	if !internal.CompareVersion(internal.IsNew, resp.Version, di.Version) {
		log.Logger.Infof("app %s need not update", app)
		return nil
	}
	fmt.Println("start download....")
	fmt.Printf("app local version: %s, app latest version: %s\n", di.Version, resp.Version)
	var url string
	switch {
	case strings.HasPrefix(resp.Path, "http"):
		url = resp.Path
		//http://127.0.0.1:9092/api/v1/io/app/download/furion-1.1.1.deb
	default:
		url = conf.UriUpdate + resp.Path

	}
	for i := 0; i < conf.Redownload; i++ {
		fp, md5, err := internal.Download(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if resp.Md5Sum != md5 {
			fmt.Println("md5sum error")
			continue
		}
		fmt.Printf("download %s successful", app)
		//拿apt 文件锁
		flock := internal.NewLock("/var/lib/dpkg/lock-frontend")
		if err = flock.WLock(); err != nil {
			return err
		}
		defer flock.Unlock()

		cmd := fmt.Sprintf("dpkg -i %s", fp)
		// 安装deb
		_, stderr, exitcode := internal.Command(cmd)
		if exitcode != 0 {
			err = errors.New(stderr)
			return err
		}
		break
	}
	fmt.Printf("install %s successful", app)
	return nil

}
