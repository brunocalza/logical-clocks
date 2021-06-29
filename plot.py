import matplotlib.pyplot as plt
import string

data = """send 0 0 1 1
send 1 1 0 1
send 1 1 2 2
receive 1 0 1 3
local 1 -1 -1 4
send 1 1 2 5
send 1 1 0 6
local 1 -1 -1 7
local 2 -1 -1 1
send 2 2 1 2
receive 1 2 1 8
receive 0 1 0 2
local 0 -1 -1 3
receive 0 1 0 7
receive 2 1 2 3
receive 2 1 2 6"""

events = list(map(lambda x : x.split(' '),  data.split("\n")))   
events = list(map(lambda x : [x[0], int(x[1]), int(x[2]), int(x[3]),  int(x[4])], events))
events.sort(key = lambda x : x[4])
max_timestamp = events[-1][4]

num_threads = 3

# horizontal lines
for i in range(max_timestamp):
    plt.plot([0, num_threads - 1], [i + 0.5, i + 0.5], "k--")

# vertical lines
for i in range(num_threads):
    plt.plot([i, i], [-0.5, max_timestamp + 0.5], "k")

# thread labels
for i in range(num_threads):
    plt.text(i, max_timestamp + 2, "process ${}$".format(string.ascii_letters[i]), rotation="vertical", horizontalalignment="center")

# event labels
for i in range(num_threads):
    for (j, e) in enumerate(filter(lambda e: e[1] == i, events)):
        plt.text(e[1] - 0.1, e[4], "${}_{}$".format(string.ascii_letters[i], j))

# events
while len(events) > 0:
    e = events.pop(0)

    if e[0] == 'local':
        plt.scatter([e[1]], [e[4]], c="k")
    elif e[0] == 'send':
        s = e
        (x0, y0) = (s[1], s[4])
        r = next(e for e in events if e[0] == 'receive'and e[2] == s[2])
        (x1, y1)  = (r[1], r[4])
        plt.scatter([x0, x1], [y0, y1], c="k")
        plt.plot([x0, x1], [y0, y1])
        events.remove(r)
    else:
        pass

plt.axis("off")
plt.savefig("clock.png", bbox_inches="tight")