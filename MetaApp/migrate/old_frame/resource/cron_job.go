package resource

import (
	"context"
	"fmt"
	"math/rand"

	toolbox "github.com/beego/beego/v2/task"

	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/log"
)

type TimingTasks struct {
	tasks []SpecTask
}

type SpecTask struct {
	Name string
	Spec string
	Task func(ctx context.Context) error
}

func (task *TimingTasks) MustInitTask() {
	err := task.InitTask()
	if err != nil {
		panic(err)
	}
}

func (task *TimingTasks) InitTask() error {
	task.tasks = []SpecTask{
		{"test_update_1", fmt.Sprintf("0/%d * * * * *", rand.Intn(10)), TestUpdate1},
		{"test_update_2", fmt.Sprintf("0/%d * * * * *", rand.Intn(5)), TestUpdate2},
	}
	return nil
}

func (task *TimingTasks) AddTasks(t []SpecTask) {
	task.tasks = append(task.tasks, t...)
}

func (task *TimingTasks) MustExecuteSuccess() {
	err := task.Execute()
	if err != nil {
		panic(err)
	}
}

func (task *TimingTasks) Execute() error {
	for _, task := range task.tasks {
		tk := toolbox.NewTask(task.Name, task.Spec, task.Task)
		if err := tk.Run(context.Background()); err != nil {
			log.Base().Error(task.Name, err)
			return fmt.Errorf("execute task %s failed:%s", task.Name, err)
		} else {
			log.Base().Infof("execute task %s success: %s", task.Name, task.Spec)
		}
		toolbox.AddTask(task.Name, tk)
	}
	log.Base().Info("init task successful!")
	return nil
}
