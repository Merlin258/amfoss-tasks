package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Process struct {
	id  string
	at  int // arrival time
	bt  int // burst time
	rt  int // remaining time
	ct  int // completion time
	wt  int // waiting time
	tat int // turn around time
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("=================================")
	fmt.Println(" GRAND LINE CPU SCHEDULER")
	fmt.Println(" (pirate crew edition)")
	fmt.Println("=================================")

	fmt.Print("How many pirate crews? ")
	n := readInt()

	var processes []Process

	for i := 0; i < n; i++ {
		fmt.Println("\nCrew", i+1)
		fmt.Print("Crew name: ")
		name := readString()
		if name == "" {
			name = "P" + strconv.Itoa(i+1)
		}

		fmt.Print("Arrival time: ")
		at := readInt()

		fmt.Print("Burst time: ")
		bt := readInt()

		p := Process{id: name, at: at, bt: bt, rt: bt}
		processes = append(processes, p)
	}

	fmt.Println("\nWhich algorithm do you want to run?")
	fmt.Println("1. FCFS")
	fmt.Println("2. SJF (non preemptive)")
	fmt.Println("3. Round Robin")
	fmt.Print("Choice: ")
	choice := readInt()

	if choice == 1 {
		fcfs(processes)
	} else if choice == 2 {
		sjf(processes)
	} else if choice == 3 {
		fmt.Print("Time quantum: ")
		tq := readInt()
		roundRobin(processes, tq)
	} else {
		fmt.Println("that's not 1, 2 or 3, bye")
	}
}

// keeps asking until it gets a real number
func readInt() int {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	val, err := strconv.Atoi(line)
	if err != nil {
		fmt.Print("not a number, try again: ")
		return readInt()
	}
	return val
}

func readString() string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

//FCFS 

func fcfs(processes []Process) {
	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].at < processes[j].at
	})

	time := 0
	var ganttID []string
	var ganttStart []int
	var ganttEnd []int

	for i := 0; i < len(processes); i++ {
		if time < processes[i].at {
			// cpu has nothing to do, waiting for next ship to dock
			ganttID = append(ganttID, "IDLE")
			ganttStart = append(ganttStart, time)
			ganttEnd = append(ganttEnd, processes[i].at)
			time = processes[i].at
		}

		start := time
		time = time + processes[i].bt

		processes[i].ct = time
		processes[i].tat = processes[i].ct - processes[i].at
		processes[i].wt = processes[i].tat - processes[i].bt

		ganttID = append(ganttID, processes[i].id)
		ganttStart = append(ganttStart, start)
		ganttEnd = append(ganttEnd, time)
	}

	fmt.Println("\n--- FCFS RESULT ---")
	printGantt(ganttID, ganttStart, ganttEnd)
	printTable(processes)
}

//SJF (non preemptive)

func sjf(processes []Process) {
	n := len(processes)
	done := make([]bool, n)
	completed := 0
	time := 0

	var ganttID []string
	var ganttStart []int
	var ganttEnd []int

	for completed < n {
		idx := -1
		for i := 0; i < n; i++ {
			if done[i] {
				continue
			}
			if processes[i].at > time {
				continue
			}
			if idx == -1 || processes[i].bt < processes[idx].bt {
				idx = i
			}
		}

		if idx == -1 {
			// nobody has arrived yet, skip ahead to whoever comes next
			next := -1
			for i := 0; i < n; i++ {
				if done[i] {
					continue
				}
				if next == -1 || processes[i].at < processes[next].at {
					next = i
				}
			}
			ganttID = append(ganttID, "IDLE")
			ganttStart = append(ganttStart, time)
			ganttEnd = append(ganttEnd, processes[next].at)
			time = processes[next].at
			continue
		}

		start := time
		time = time + processes[idx].bt

		processes[idx].ct = time
		processes[idx].tat = processes[idx].ct - processes[idx].at
		processes[idx].wt = processes[idx].tat - processes[idx].bt
		done[idx] = true
		completed = completed + 1

		ganttID = append(ganttID, processes[idx].id)
		ganttStart = append(ganttStart, start)
		ganttEnd = append(ganttEnd, time)
	}

	fmt.Println("\n--- SJF RESULT ---")
	printGantt(ganttID, ganttStart, ganttEnd)
	printTable(processes)
}

//Round Robin

func roundRobin(processes []Process, tq int) {
	n := len(processes)

	sort.SliceStable(processes, func(i, j int) bool {
		return processes[i].at < processes[j].at
	})

	arrived := make([]bool, n)
	var queue []int
	time := 0
	completed := 0

	var ganttID []string
	var ganttStart []int
	var ganttEnd []int

	if n > 0 {
		time = processes[0].at
	}
	for i := 0; i < n; i++ {
		if processes[i].at <= time && !arrived[i] {
			queue = append(queue, i)
			arrived[i] = true
		}
	}

	for completed < n {
		if len(queue) == 0 {
			next := -1
			for i := 0; i < n; i++ {
				if !arrived[i] {
					if next == -1 || processes[i].at < processes[next].at {
						next = i
					}
				}
			}
			ganttID = append(ganttID, "IDLE")
			ganttStart = append(ganttStart, time)
			ganttEnd = append(ganttEnd, processes[next].at)
			time = processes[next].at

			for i := 0; i < n; i++ {
				if processes[i].at <= time && !arrived[i] {
					queue = append(queue, i)
					arrived[i] = true
				}
			}
			continue
		}

		idx := queue[0]
		queue = queue[1:]

		run := tq
		if processes[idx].rt < tq {
			run = processes[idx].rt
		}

		start := time
		time = time + run
		processes[idx].rt = processes[idx].rt - run

		ganttID = append(ganttID, processes[idx].id)
		ganttStart = append(ganttStart, start)
		ganttEnd = append(ganttEnd, time)

		for i := 0; i < n; i++ {
			if processes[i].at <= time && !arrived[i] {
				queue = append(queue, i)
				arrived[i] = true
			}
		}

		if processes[idx].rt > 0 {
			queue = append(queue, idx)
		} else {
			processes[idx].ct = time
			processes[idx].tat = processes[idx].ct - processes[idx].at
			processes[idx].wt = processes[idx].tat - processes[idx].bt
			completed = completed + 1
		}
	}

	fmt.Println("\n--- ROUND ROBIN RESULT (quantum =", tq, ") ---")
	printGantt(ganttID, ganttStart, ganttEnd)
	printTable(processes)
}

//generating the results

func printGantt(ids []string, starts []int, ends []int) {
	fmt.Println("\nGantt Chart:")

	line := "|"
	for i := 0; i < len(ids); i++ {
		line = line + " " + ids[i] + " |"
	}
	fmt.Println(line)

	timeline := strconv.Itoa(starts[0])
	for i := 0; i < len(ends); i++ {
		timeline = timeline + "\t" + strconv.Itoa(ends[i])
	}
	fmt.Println(timeline)
}

func printTable(processes []Process) {
	fmt.Println("\nCrew\tArrival\tBurst\tCompletion\tWaiting\tTurnaround")

	totalWT := 0
	totalTAT := 0

	for i := 0; i < len(processes); i++ {
		p := processes[i]
		fmt.Println(p.id, "\t", p.at, "\t", p.bt, "\t", p.ct, "\t\t", p.wt, "\t", p.tat)
		totalWT = totalWT + p.wt
		totalTAT = totalTAT + p.tat
	}

	avgWT := float64(totalWT) / float64(len(processes))
	avgTAT := float64(totalTAT) / float64(len(processes))

	fmt.Println("\nAverage Waiting Time:", avgWT)
	fmt.Println("Average Turnaround Time:", avgTAT)
}