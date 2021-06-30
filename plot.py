import matplotlib.pyplot as plt
import string
import pandas as pd
import sys

class Event:

    def __init__(self, kind, owner, source, destination, timestamp):
        self.kind = kind
        self.owner = owner
        self.source = source
        self.destination = destination
        self.timestamp = timestamp

path = sys.argv[1]
df = pd.read_csv(path, sep=" ", header=None)
df.columns = ['Kind', 'Owner', 'Source', 'Destination', 'Timestamp']

events = [(Event(row.Kind, row.Owner, row.Source, row.Destination, row.Timestamp)) for index, row in df.iterrows() ]
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
    elif e.kind == 'sent':
        s = e
        (x0, y0) = (s.owner + 1, s.timestamp)
        r = next(e for e in events if e.kind == 'received' and e.source == s.source)
        (x1, y1)  = (r.owner + 1, r.timestamp)
        plt.scatter([x0, x1], [y0, y1], c = "k")
        plt.plot([x0, x1], [y0, y1])
        events.remove(r)
    else:
        pass

plt.axis("off")
plt.savefig("clock.png", bbox_inches="tight")