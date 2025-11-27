package repository

import (
	"fmt"
	"myapp/events"
)

type Repo struct{}

func (r *Repo) Save(data string) error {
	fmt.Println("Saving to DB:", data)
	return nil
}

func (r *Repo) DoSomething() {
	// We pass our method as the function argument (Dependency Injection)
	events.Trigger("some data", r.Save)
}
