package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	if n < pool {
		pool = n
	}

	chanUser := make(chan user, n)
	sem := make(chan struct{}, pool)

	var i int64
	for i = 0; i < n; i++ {
		go func(j int64, chanUser chan<- user) {
			sem <- struct{}{}
			chanUser <- getOne(j)
			<-sem
		}(i, chanUser)
	}

	for i = 0; i < n; i++ {
		res = append(res, <-chanUser)
	}
	return
}
