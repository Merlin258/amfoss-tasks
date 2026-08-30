# Task-06 CPU Scheduler
## Approach
Most of the info I got about scheduling algorithms was from ai. AI was also the reason I understood that we were supposed to implement these algorithms seperately in a menu format and not do it all together. What i got from my ai (gemini pro):
Basic Definitions
Arrival Time: The time when a process is ready to run.

Burst Time: The time that the CPU needs to complete the task.

Completion Time: The time when the CPU finishes the task.

Time Quantum: A fixed limit of time given to a process.

1. First Come First Serve (FCFS)
The CPU executes the processes in the sequence that they arrive.

Sort all the processes by their arrival time.

Select the first process in the list.

The CPU executes this process until it completes.

The CPU selects the next process in the list.

2. Shortest Job First (SJF - Non-Preemptive)
The CPU always selects the available process that has the smallest burst time.

Make a list of all processes that are available at the current time.

Compare the burst times of these available processes.

Select the process with the smallest burst time.

The CPU executes this process until it completes.

Update the current time.

Do steps 1 through 5 again until all processes complete.

3. Round Robin (RR)
The CPU uses a queue to share time equally among all processes.

Put all available processes into a ready queue.

Remove the first process from the front of the queue.

The CPU executes this process for the duration of the Time Quantum.

Stop the process if it does not complete in this Time Quantum.

Put the stopped process at the back of the ready queue.

Do steps 2 through 5 again until the queue is empty.

4. Metric Calculations
The code uses these mathematical formulas to calculate the performance:

Turnaround Time = Completion Time - Arrival Time

Waiting Time = Turnaround Time - Burst Time

Average Time = The sum of all times divided by the total number of processes

(What I thought from all this was like implement fcfs first, then if there's a queue implement sjf. I didn't know where round robin would come in this so i just uploaded the whole task to it and *then* it told me that i was supposed to do it in a menu driven program.)