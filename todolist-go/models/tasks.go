package models

import (
  "gorm.io/gorm"
)

type Task struct {
  gorm.Model
  task string 
  completed bool
} 
