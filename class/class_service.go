package class

import (
	"FitnessClass/model"
	"sync"
	"time"
)

type ClassService struct {
	Classes map[int]*model.Class
	mu      sync.Mutex
}

var classServiceInstance *ClassService
var classServiceOnce sync.Once

func GetClassService() *ClassService {
	classServiceOnce.Do(func() {
		classServiceInstance = &ClassService{
			Classes: make(map[int]*model.Class),
		}
	})
	return classServiceInstance
}

func (cs *ClassService) CreateClass(classType model.ClassType, capacity int, schedule time.Time) *model.Class {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	class := model.NewClass(classType, capacity, schedule)
	cs.Classes[class.ID] = class
	return class
}
