package modules

import (
	"os"
	"strings"

	"camera/src/log"
)

func Car_role() (role string) {
	role_dir := "/opt/cargo/acu/info/type"
	file, err := os.Open(role_dir)
	if err != nil {
		log.Logger.Println(err)
		return
	}
	defer file.Close()
	b := make([]byte, 15)
	n, err := file.Read(b)
	if err != nil {
		log.Logger.Println(err)
		return
	}

	role = strings.TrimSpace(string(b[:n]))
	return

}
