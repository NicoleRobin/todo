package serializer

import (
	"github.com/nicolerobin/todo/model"
)

type Task struct {
	Id        uint   `json:"id" exmaple:"1"`
	Title     string `json:"title" exmaple:"吃饭"`
	Content   string `json:"content" exmaple:"中午吃螺蛳粉"`
	View      uint64 `json:"view" example:"32"`
	Status    int    `json:"status" example:"0"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	return tasks
}
