import matplotlib.pyplot as plt
import string

class Event:

    def __init__(self, kind, owner, src, dst, timestamp):
        self.kind = kind
        self.owner = owner
        self.src = src
        self.dst = dst
        self.timestamp = timestamp

data = """local 2 -1 -1 1
send 2 2 1 2
send 0 0 1 1
send 1 1 0 1
send 1 1 2 2
receive 2 1 2 3
receive 0 1 0 2
local 0 -1 -1 3
receive 1 0 1 3
local 1 -1 -1 4
send 1 1 2 5
send 1 1 0 6
receive 0 1 0 7
receive 2 1 2 6
local 1 -1 -1 7
receive 1 2 1 8"""

events = list(map(lambda x : x.split(' '),  data.split("\n")))   
events = list(map(lambda x : Event(x[0], int(x[1]), int(x[2]), int(x[3]),  int(x[4])), events))
events.sort(key = lambda e : e.timestamp)
maxTimestamp = events[-1].timestamp

numThreads = len({e.owner for e in events}) # numThreads is the number of distinct owners (processes)

# horizontal lines
for i in range(maxTimestamp):
    plt.plot([0, numThreads], [i + 0.5, i + 0.5], "k--")

# vertical lines
for i in range(numThreads + 1):
    plt.plot([i, i], [-0.5, maxTimestamp + 0.5], "k")

# thread labels
labels = ["time"] + ["${}$".format(string.ascii_uppercase[i]) for i in range(numThreads)]
for i, label in enumerate(labels):
    plt.text(i, maxTimestamp + 1, label, rotation="vertical", horizontalalignment="center")

# time labels
for t in range(1, maxTimestamp + 1):
    plt.text(0.075, t - 0.1, "$t={}$".format(t))

# event labels
for i in range(numThreads):
    for (j, e) in enumerate(filter(lambda e: e.owner == i, events)):
        plt.text(e.owner - 0.275 + 1, e.timestamp + 0.1, "$({},{})$".format(e.timestamp, string.ascii_uppercase[i]))

# events
while len(events) > 0:
    e = events.pop(0)

    if e.kind == 'local':
        plt.scatter([e.owner + 1], [e.timestamp], c = "k")
    elif e.kind == 'send':
        s = e
        (x0, y0) = (s.owner + 1, s.timestamp)
        r = next(e for e in events if e.kind == 'receive' and e.src == s.src)
        (x1, y1)  = (r.owner + 1, r.timestamp)
        plt.scatter([x0, x1], [y0, y1], c = "k")
        plt.plot([x0, x1], [y0, y1])
        events.remove(r)
    else:
        pass

plt.axis("off")
plt.savefig("clock.png", bbox_inches="tight")