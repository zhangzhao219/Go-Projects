package resource

import (
	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/job"
	global "gitlab.appshahe.com/service-cloud-rec/sort_base.git/base_global"
)

const (
	MMS_UPDATE = "mms_update"
)

func StartNewJobs() {
	schedule := job.NewWithRedis(global.RedisClient)
	schedule.RegisterJob(MMS_UPDATE, job.Build(MMS_UPDATE, AdsUpdateModelConnPool, job.RunSyncAndPanicErr, true, 0, false))
	schedule.ConfigListen(MMS_UPDATE, "mms_update_listen_key", "mms_online_model_conn_pool_path")
	schedule.Start()
}
