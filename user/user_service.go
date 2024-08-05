package user

import (
	"FitnessClass/model"
	"sync"
)

type UserService struct {
	Users map[int]*model.User
	mu    sync.Mutex
}

var userServiceInstance *UserService
var userServiceOnce sync.Once

func GetUserService() *UserService {
	userServiceOnce.Do(func() {
		userServiceInstance = &UserService{
			Users: make(map[int]*model.User),
		}
	})
	return userServiceInstance
}

func (us *UserService) RegisterUser(name, email, password string, tier model.Tier) *model.User {
	us.mu.Lock()
	defer us.mu.Unlock()

	user := model.NewUser(name, email, password, tier)
	us.Users[user.ID] = user
	return user
}
