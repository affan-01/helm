package manager

import (
	"boot.dev/helm/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	Event         map[string]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
}

func (m *Manager) UpdateTasks() {
}

func (m *Manager) SendWork() {

}
