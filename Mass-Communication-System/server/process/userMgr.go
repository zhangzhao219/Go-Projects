package process2

import "fmt"

// 在服务器端实例只有一个，在很多的地方都会使用到

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

// 完成对userMgr初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的增删改查

func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

func (um *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return um.onlineUsers
}

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d不存在", userId)
		return
	}
	return
}
