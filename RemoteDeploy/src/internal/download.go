package internal

import (
	"fmt"
	"os"
	"path"

	"gitee.com/ishmaelwanglin/download"
)

/*
download
下载事传入版本号
下载地址由update地址+版本号组成
M：md5sum
*/
func Download(url string) (fp string, md5 string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r.(string))
		}
	}()
	//fuion-update-furion-1.1.1.deb
	dir := "/var/tmp/" + fmt.Sprintf("fuion-update-%s", path.Base(url))
	if _, err := os.Stat(dir); err != nil {
		_ = os.Mkdir(dir, os.ModePerm)
	}
	fmt.Println("url", url)
	d := download.New(url, dir)
	d.SetThread(2) //d.Threads =2 
	err = d.Down()
	if err != nil {
		return
	}

	fp = d.GetPath()
	md5, err = download.Md5sum(fp)

	return

}
