package goroutine_test

import (
	"fmt"
	"math/rand"
	"testing"
)

type Job struct {
	Id int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func TestWorkerPool(t *testing.T) {
	// 1.任务管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(10, jobChan, resultChan)
	// 4. 打印结果
	go printResult(resultChan)
	// 5.生成任务
	genderJob(jobChan)
}

// 打印结果
func printResult(resultChan chan *Result) {
	for result := range resultChan {
		fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
			result.job.RandNum, result.sum)
	}
}

// 生成任务
func genderJob(jobChan chan *Job) {
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建协程，执行任务，获取结果
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i:=0; i<num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				// 将随机数接过来
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 将结果写入result
				r := &Result{
					job: job,
					sum: sum,
				}
				// 放入管道中去
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
