package main

import (
	"context"
	user_feature "kitextest/kitex_gen/user_feature"
	"time"
)

// UserFeatureRPCServiceImpl implements the last service interface defined in the IDL.
type UserFeatureRPCServiceImpl struct{}

// GetUserFeature implements the UserFeatureRPCServiceImpl interface.
func (s *UserFeatureRPCServiceImpl) GetUserFeature(ctx context.Context, req *user_feature.Request) (resp *user_feature.UserFeatureItem, err error) {
	// TODO: Your code here...
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	time.Sleep(time.Duration(time.Second * 15))
	// klog.Info("fgfgfgf")
	resp = &user_feature.UserFeatureItem{}
	return
}
