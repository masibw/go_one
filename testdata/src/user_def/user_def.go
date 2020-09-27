package user_def

import (
	"context"
)

type Person struct {
	Name  string
	JobID int
}

type Job struct {
	JobID int
	Name  string
}

func ForStmt() {
	for  {
		ctx, cancel := context.WithTimeout(context.Background(), 9)
		defer cancel()
		ctx.Done()//want "this query is called in a loop"
	}
}
