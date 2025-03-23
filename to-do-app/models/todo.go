package models

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusPending TaskStatus = "Pending"
	StatusInProgress TaskStatus = "In Progress"
	StatusCompleted TaskStatus = "Completed"
)

type ToDo struct {
	gorm.Model
	TaskId int `json:"task_id" gorm:"primaryKey"`
	TaskName string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	TaskStatus TaskStatus `json:"task_status"`
	TaskAssignee int `json:"task_assignee"`
}
