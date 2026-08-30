package main

import(
	"fmt"
	"sort"
)

func fcfs(){
	type process struct{
		pid 	string
		arrival	int
		burst	int
	}
	timeline:=make(map[int]string)
	waiting:=make(map[string]int)
	turnaround:=make(map[string]int)
	processes:=[]process{}
	for{
		var string pid
		var int arrival
		var int burst
		fmt.Print("Enter pid:")
		fmt.Scanln(&pid)
		fmt.Print("Enter arrival:")
		fmt.Scanln(&arrival)
		fmt.Print("Enter burst:")
		fmt.Scanln(&burst)
		processes=append(processes,process{pid,arrival,burst})
	}
	sort.Slice(processes, func(i, j int) bool {
        return processes[i].arrival < processes[j].arrival
    })
	time:=0
	for i:=0;i<len(processes);i++{
		x:=processes[i]
		if x.arrival>=time{
			time=x.arrival
			timeline[time]=x.pid+" Started"
			timeline[time+x.burst]=x.pid+" Finished executing"
		}else{
			timeline[time]=x.pid+" Arrived"
			timeline[time+x.burst]=x.pid+" Finished executing"
		}
		time+=x.burst
		turnaround[x.pid]=time+x.burst-x.arrival
		waiting[x.pid]=turnaround[x.pid]-x.burst
	}
	avgturn:=0
	avgwait=0
	fmt.Print("Turnarounds")
	for pid,time:=range turnaround{
		fmt.Print(pid+"\t\t"+time)
		avgturn+=time
	}
	fmt.Print("Average turnaround =",avgturn/len(turnaround))
	fmt.Print("Waitings")
	for pid,time:=range waiting{
		fmt.Print(pid+"\t\t"+time)
		avgwait+=time
	}
	fmt.Print("Average Waiting =",avgwait/len(waiting))
	fmt.Print("Timeline")
	for time,sent:=range timeline{
		fmt.print(time+"\t\t"+sent)
	}
}

func sjf(){
	type process struct{
		pid 	string
		arrival	int
		burst	int
	}
	timeline:=make(map[int]string)
	processes:=[]process{}
	for{
		var string pid
		var int arrival
		var int burst
		fmt.Print("Enter pid:")
		fmt.Scanln(&pid)
		fmt.Print("Enter arrival:")
		fmt.Scanln(&arrival)
		fmt.Print("Enter burst:")
		fmt.Scanln(&burst)
		processes=append(processes,process{pid,arrival,burst})
	}
	sort.Slice(processes, func(i, j int) bool {
        return processes[i].burst < processes[j].burst
    })
}

func rr(){
	type process struct{
		pid 	string
		arrival	int
		burst	int
		quantum	int
	}
	timeline:=make(map[int]string)
	processes:=[]process{}
	for{
		var string pid
		var int arrival
		var int burst
		var int quantum
		fmt.Print("Enter pid:")
		fmt.Scanln(&pid)
		fmt.Print("Enter arrival:")
		fmt.Scanln(&arrival)
		fmt.Print("Enter burst:")
		fmt.Scanln(&burst)
		fmt.Print("Enter quantum:")
		fmt.Scanln(&quantum)
		processes=append(processes,process{pid,arrival,burst,quantum})
	}
	sort.Slice(processes, func(i, j int) bool {
        return processes[i].arrival < processes[j].arrival
    })
}

func main(){
	for{
		var choice int
		fmt.Print(`Enter choice
		[1]First Come First Serve (FCFS)
		[2]Shortest Job First (SJF - Non-Preemptive)
		[3]Round Robin (RR)
		Click anything else to exit
		`)
		fmt.Scanln(&choice)
		if choice==1{
			fcfs()
			break
		}else if choice==2{
			sjf()
			break
		}else if choice==3{
			rr()
			break
		}else{
			fmt.Print("Choice not found")
		}
	}
}