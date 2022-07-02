package batch

import (
	"sync"
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

	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	var idUser int64 = 0
	var i int64
	for i = 0; i < pool; i++ {
		wg.Add(1)
		go loadUser(&idUser, &res, &n, mu, wg)
	}
	wg.Wait()
	return
}

func loadUser(id *int64, res *[]user, count *int64, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	for *count > 0 {
		*res = append(*res, getOne(*id))
		*id++
		*count--
	}
	wg.Done()
	mu.Unlock()
}
