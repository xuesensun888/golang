package modules

import (
	"camera/src/log"
	"os/exec"
	"strings"
	"time"

	"fmt"

	"os"
	"path"
)

func Cam_Generate() {

	single_path := "/opt/cargo/camera/tools/3_pictures_Single_orin"
	orinb_1A_PATH := "/opt/cargo/camera/tools/2_pictures_4orin/0_1A_pics/orinb"
	orin_1B_PATH := "/opt/cargo/camera/tools/2_pictures_4orin/1_1B_pics/orinb"
	role := Car_role()
	switch role {
	case "orin_single_A":
		fs, err := os.ReadDir(single_path)
		if err != nil {
			log.Logger.Fatal(err)
		}

		for _, file := range fs {
			if !file.IsDir() && path.Ext(file.Name()) == ".sh" {
				fmt.Println(file.Name())
				err := os.Chdir(single_path)
				if err != nil {
					log.Logger.Println(err)
				}
				cmd := exec.Command("bash", "-c", fmt.Sprintf("./%s", file.Name()))

				cmd.Start()

				pid := cmd.Process.Pid
				fmt.Println(pid)
				time.Sleep(2 * time.Second)

				process, _ := os.FindProcess(pid)

				err = process.Signal(os.Kill)
				if err != nil {
					fmt.Println(err)
				}
				_ = cmd.Wait()
			}
		}
	case "4orin_B":
		hostname, _ := os.Hostname()

		if strings.HasSuffix(hostname, "2B") || strings.HasSuffix(hostname, "2A") {
			log.Logger.Fatalln("请在1A 1B上执行")
		}

		if strings.HasSuffix(hostname, "1A") {
			fs, err := os.ReadDir(orinb_1A_PATH)
			if err != nil {
				log.Logger.Errorln(err)
			}

			for _, file := range fs {
				if !file.IsDir() && path.Ext(file.Name()) == ".sh" {
					fmt.Println(file.Name())
					err := os.Chdir(orinb_1A_PATH)
					if err != nil {
						log.Logger.Fatal("chdir is error", err)
					}

					cmd := exec.Command("bash", "-c", fmt.Sprintf("./%s", file.Name()))

					cmd.Start()

					pid := cmd.Process.Pid

					time.Sleep(2 * time.Second)

					process, _ := os.FindProcess(pid)

					err = process.Signal(os.Kill)
					if err != nil {
						fmt.Println(err)
					}
					_ = cmd.Wait()
				}
			}
		}
		if strings.HasSuffix(hostname, "1B") {
			fs, err := os.ReadDir(orin_1B_PATH)
			if err != nil {
				log.Logger.Errorln(err)
			}

			for _, file := range fs {
				if !file.IsDir() && path.Ext(file.Name()) == ".sh" {
					fmt.Println(file.Name())
					os.Chdir(orinb_1A_PATH)

					cmd := exec.Command("bash", "-c", fmt.Sprintf("./%s", file.Name()))
					cmd.Start()
					//获取命令的pid
					pid := cmd.Process.Pid
					fmt.Println(pid)
					time.Sleep(2 * time.Second)

					process, _ := os.FindProcess(pid)
					process.Signal(os.Kill)
					cmd.Wait()
				}
			}
		}

	}
}
