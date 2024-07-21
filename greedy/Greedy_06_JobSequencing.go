package greedy

import (
	"sort"
)

// jobs:= []greedy.Job{
// 		{Id: "a",Deadline: 2, Profit: 100},
// 		{Id: "b",Deadline: 1, Profit: 19},
// 		{Id: "c",Deadline: 2, Profit: 27},
// 		{Id: "d",Deadline: 1, Profit: 25},
// 		{Id: "e",Deadline: 3, Profit: 15},
// 	}

type Job struct {
	Id       string
	Deadline int
	Profit   int
}

//Return max profit and sequence of jobs
//TC - nlogn + n + (n*maxdeadLine) 
//SC - O(maxDeadline)
func JobSequencing(jobs []Job) (int,[]string) {
    n:=len(jobs)
	maxProfit:=0
    
	sort.Slice(jobs,func(i, j int) bool {
        if jobs[i].Profit==jobs[j].Profit{
			return jobs[i].Deadline<jobs[j].Deadline
		}
		return jobs[i].Profit>jobs[j].Profit
	})

	maxDeadline:=0

	for _,job:=range jobs{
		if job.Deadline>maxDeadline{
			maxDeadline = job.Deadline
		}
	}

	slots:= make([]string,maxDeadline+1)

	for i:=0;i<n;i++{
        for j:=jobs[i].Deadline;j>0;j--{
		    if slots[j]==""{
               slots[j] = jobs[i].Id
			   maxProfit+=jobs[i].Profit
			   break
			}
		}
	}

	jobsList := []string{}

	for _,val:=range slots{
		if val!=""{
			jobsList = append(jobsList, val)
		}
	}
   
	return maxProfit,jobsList

}