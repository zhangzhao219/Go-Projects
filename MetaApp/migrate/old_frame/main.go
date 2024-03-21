package main

import (
	"flag"
	"old_frame/resource"
	"old_frame/router"

	"github.com/beego/beego/v2/task"

	"github.com/beego/beego/v2/server/web"
	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/log"
)

var (
	env       string
	configDir string
)

func main() {
	flag.StringVar(&env, "env", "dev", "config env name")
	flag.StringVar(&configDir, "config_dir", "./config/config_file", "config file dir")
	flag.Parse()
	resource.Gr.Env.Store(env)

	if err := resource.Init(env, configDir); err != nil {
		log.Base().Error("global init: ", err)
		return
	}

	task.StartTask()
	defer task.StopTask()
	resource.StartNewJobs()

	router.Load(resource.Gr.ServerConfig)

	web.Run()
}
