package request

import (
	"encoding/json"
	"fmt"
	"io"
	"kargo-deploy/conf"
	"net/http"
)

type AppVersion struct {
	Id        string `json:"id"`
	AppId     string `json:"app_id"`
	Version   string `json:"version"`
	Type      string `json:"type"`
	Md5Sum    string `json:"md5"`
	Path      string `json:"url"`
	TimeStamp string `json:"date"`
}
type RespVersion struct {
	ErrCode     int `json:"errcode"`
	*AppVersion `json:"data"`
}

// 获取版本信息
func HeadVersionLatest(app string) (resp *RespVersion, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r.(string))
		}
	}()
	response, err := http.Get(conf.UriUpdate + "/" + app)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("resp code %d", response.StatusCode)
		return
	}
	defer response.Body.Close()

	p, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	resp = &RespVersion{
		AppVersion: &AppVersion{},
	}
	err = json.Unmarshal(p, resp)
	return
}
