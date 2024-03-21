package resource

import (
	"context"
	"fmt"
	"math/rand"

	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/cron_job"
	global "gitlab.appshahe.com/service-cloud-rec/sort_base.git/base_global"
)

func StartSchedule() {
	taskList := []cron_job.Task{
		cron_job.DefaultCronTask("test_update_1", TestUpdate1, fmt.Sprintf("0/%d * * * * *", rand.Intn(10))),
		cron_job.DefaultCronTask("test_update_2", TestUpdate2, fmt.Sprintf("0/%d * * * * *", rand.Intn(5))),
		cron_job.DefaultCronTask("action_map_updater", AdsUpdateModelConnPool, fmt.Sprintf("0 %d 8 * * *", rand.Intn(20))).WithListen("mms_update_listen_key", "mms_online_model_conn_pool_path"),
	}
	var listenMap = map[string]cron_job.ListenTask{}
	for _, task := range taskList {
		task = task.TaskInit(Gr.Env)
		if task.ListenChannel == "" {
			continue
		}
		topic, ok := listenMap[task.ListenChannel]
		if !ok {
			topic.Topic = task.ListenChannel
			topic.PayloadMap = make(map[string]cron_job.PayloadAndMessage)
		}
		payload, ok := topic.PayloadMap[task.ListenPayload]
		if !ok {
			payload.Name = task.ListenPayload
		}
		payload.JobList = append(payload.JobList, task)
		topic.PayloadMap[task.ListenPayload] = payload
		listenMap[task.ListenChannel] = topic
	}
	for channel, topic := range listenMap {
		pubSub := global.RedisClient.Subscribe(context.Background(), channel)
		go cron_job.Listen(pubSub.Channel(), channel, topic.PayloadMap)
	}
}
