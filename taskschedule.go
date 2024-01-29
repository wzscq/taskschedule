/*
Package taskschedule run task periodically at the specified time.	
The task is executed at the specified time that is specified in the schedule.
The task will be executed at a independent goroutine.
*/
package taskschedule

import (
	"time"
)

type Schedule struct {
	Duration string `json:"duration"`
	RunTime string `json:"runTime"`
}

type Task interface {
	Run()
}

type taskSchedule struct {
	Schedule *Schedule
	Task Task
}

/*
RunTask execute the Run method of the task at the specified time that is specified in the schedule
schedule is used to specify the time when the task is executed
task is the task to be executed
*/
func RunTask(schedule *Schedule,task Task) {
	t := &taskSchedule{
		Schedule: schedule,
		Task: task,
	}
	go t.run()
}

func (t *taskSchedule) waitForRun(){
	if t.Schedule.RunTime!="" {
		duration, _ := time.ParseDuration(t.Schedule.Duration)
		now := time.Now()
		runTime,_ := time.Parse("15:04:05",t.Schedule.RunTime)
		runTime = time.Date(now.Year(),now.Month(),now.Day(),runTime.Hour(),runTime.Minute(),runTime.Second(),0,time.Local)
		for {
			if runTime.Before(now) {
				runTime = runTime.Add(duration)
			} else {
				break
			}
		}
		duration = runTime.Sub(now)
		time.Sleep(duration)
	}
}

func (t *taskSchedule) run() {
	duration, _ := time.ParseDuration(t.Schedule.Duration)
	t.waitForRun()
	for {
			t.Task.Run()
			time.Sleep(duration)
	}
}