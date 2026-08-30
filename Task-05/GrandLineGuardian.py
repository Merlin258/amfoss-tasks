import curses
import os
import time
import signal

REFRESH = 0.8
CLK_TCK = os.sysconf('SC_CLK_TCK')

prev_times = {}


def get_mem_total():
    for line in open('/proc/meminfo'):
        if line.startswith('MemTotal'):
            return int(line.split()[1])
    return 1


MEM_TOTAL = get_mem_total()


def get_processes():
    procs = []
    for pid in os.listdir('/proc'):
        if not pid.isdigit():
            continue
        pid = int(pid)
        try:
            name = open(f'/proc/{pid}/comm').read().strip()

            stat = open(f'/proc/{pid}/stat').read().split()
            utime = int(stat[13])
            stime = int(stat[14])
            cpu_ticks = utime + stime

            mem_kb = 0
            for line in open(f'/proc/{pid}/status'):
                if line.startswith('VmRSS'):
                    mem_kb = int(line.split()[1])
                    break

            old_ticks = prev_times.get(pid, cpu_ticks)
            cpu_percent = (cpu_ticks - old_ticks) / CLK_TCK / REFRESH * 100
            prev_times[pid] = cpu_ticks

            procs.append({
                'pid': pid,
                'name': name,
                'cpu': round(cpu_percent, 1),
                'mem': round(mem_kb / MEM_TOTAL * 100, 1),
            })
        except:
            continue
    return procs


def draw(stdscr, procs, selected, scroll, message):
    stdscr.erase()
    height, width = stdscr.getmaxyx()

    stdscr.addstr(0, 0, f"Simple Process Monitor - {len(procs)} processes running")
    stdscr.addstr(1, 0, "UP/DOWN = move   k = kill selected   q = quit")
    stdscr.addstr(2, 0, f"{'PID':>7} {'CPU%':>7} {'MEM%':>7}  NAME")
    stdscr.addstr(3, 0, "-" * (width - 1))

    list_height = height - 5
    visible = procs[scroll:scroll + list_height]

    for i, p in enumerate(visible):
        row = 4 + i
        line = f"{p['pid']:>7} {p['cpu']:>7} {p['mem']:>7}  {p['name']}"
        line = line[:width - 1]
        if scroll + i == selected:
            stdscr.attron(curses.A_REVERSE)
            stdscr.addstr(row, 0, line)
            stdscr.attroff(curses.A_REVERSE)
        else:
            stdscr.addstr(row, 0, line)

    if message:
        stdscr.addstr(height - 1, 0, message[:width - 1])

    stdscr.refresh()


def main(stdscr):
    curses.curs_set(0)
    stdscr.timeout(int(REFRESH * 1000))

    selected = 0
    scroll = 0
    message = ""

    while True:
        procs = get_processes()
        procs.sort(key=lambda p: p['cpu'], reverse=True)

        # keep selected pointer inside the list
        if selected >= len(procs):
            selected = len(procs) - 1
        if selected < 0:
            selected = 0

        height, _ = stdscr.getmaxyx()
        list_height = height - 5
        if selected < scroll:
            scroll = selected
        if selected >= scroll + list_height:
            scroll = selected - list_height + 1

        draw(stdscr, procs, selected, scroll, message)
        message = ""

        key = stdscr.getch()

        if key == -1:
            continue  # nothing pressed, just loop again (this is our refresh)
        elif key == ord('q'):
            break
        elif key == curses.KEY_UP:
            selected -= 1
        elif key == curses.KEY_DOWN:
            selected += 1
        elif key == ord('k'):
            if procs:
                target = procs[selected]
                stdscr.addstr(height - 1, 0,
                               f"Kill {target['name']} (pid {target['pid']})? y/n ")
                stdscr.refresh()
                stdscr.timeout(-1)  # wait for an actual answer this time
                confirm = stdscr.getch()
                stdscr.timeout(int(REFRESH * 1000))
                if confirm == ord('y'):
                    try:
                        os.kill(target['pid'], signal.SIGTERM)
                        message = f"Killed pid {target['pid']}"
                    except Exception as e:
                        message = f"Couldn't kill it: {e}"
                else:
                    message = "cancelled"


if __name__ == '__main__':
    curses.wrapper(main)