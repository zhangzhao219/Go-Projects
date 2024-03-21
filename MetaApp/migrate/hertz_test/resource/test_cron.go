package resource

import (
	"context"
	"time"

	"gitlab.appshahe.com/service-cloud-rec/rec-utils.git/component/log"
)

func TestUpdate1(ctx context.Context) error {
	log.Base().Infof("test1 execute, time %v\n", time.Now())
	return nil
}

func TestUpdate2(ctx context.Context) error {
	log.Base().Infof("test2 execute, time %v\n", time.Now())
	return nil
}
