package models

type Todo struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string `bson:"title,omitempty" json:"title" validate:"required,min=2,max=100"`
	Completed bool   `bson:"completed,omitempty" json:"completed" validate:"required"`
}

func NewTodo() Todo {
	instance := Todo{}
	instance.Title = ""
	instance.Completed = false
	return instance
}
