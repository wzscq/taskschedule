package taskschedule

import (
	"fmt"
	"testing"
	"time"
)

type testTask struct {
	Count int
}

func (t *testTask) Run() {
	t.Count++
	fmt.Println(t.Count," ",time.Now().Format("2006-01-02 15:04:05"))
}

func TestRunTask(t *testing.T) {
	//获取当前时间10s后的时间
	now := time.Now()
	runTime := now.Add(10*time.Second)

	schedule := &Schedule{
		Duration: "1s",
		RunTime: runTime.Format("15:04:05"),
	}
	task := &testTask{}
	RunTask(schedule,task)
	time.Sleep(10*time.Second)
	if task.Count>1 {
		t.Error("task run count is not 0 ",task.Count)
	}
	time.Sleep(10*time.Second)
	if task.Count<10 {
		t.Error("task count is not 10 ",task.Count)
	}
}