package hw8

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("erroeept")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	quit := make(chan interface{})
	taskChan := make(chan Task)
	resultChan := make(chan error)
	errorchan := make(chan error)
	var wg sync.WaitGroup

	go errorchek(resultChan, quit, m, tasks, errorchan)

	for i := 0; i < n; i++ {
		go worker(taskChan, resultChan, quit, &wg)
		wg.Add(1)
	}

	go func() {
		for _, task := range tasks {
			select {
			case <-quit:
				break
			case taskChan <- task:

			}
		}
	}()
	err := <-errorchan
	//close(resultChan)
	//close(taskChan)
	wg.Wait()
	//err := <-errorchan
	//err := <-errorchan
	close(resultChan)
	//close(taskChan)
	return err
}

func errorchek(resultChan chan error, quit chan interface{}, m int, tasks []Task, errorchan chan error) {
	//var err error = nil
	errorsgot := 0
	for i := 0; i < len(tasks); i++ {
		if result := <-resultChan; result != nil {
			errorsgot++
			if errorsgot == m {
				close(quit)
				//	fmt.Println("stop")
				errorchan <- ErrErrorsLimitExceeded //errors.New(fmt.Sprintf("got %d erros\n", errorsgot))
			}
		}
	}
	select {
	case errorchan <- nil:
		close(quit)
	default:
		return
	}

}
func worker(taskChan <-chan Task, resultChan chan<- error, quit <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			//fmt.Println("exit")
			return
		case task := <-taskChan:
			//	fmt.Println(task())
			resultChan <- task()
		}
	}
}
