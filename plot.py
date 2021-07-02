#! /usr/bin/python

import matplotlib.pyplot as plt
import string
import pandas as pd
import sys
import math

class Event:

    def __init__(self, kind, owner, source, destination, timestamp):
        self.kind = kind
        self.owner = owner
        self.source = source
        self.destination = destination
        self.timestamp = timestamp

    def maxTimestamp(self):
        return max(self.timestamp)

    def ownerTimestamp(self):
        if len(self.timestamp) == 1:
            return self.timestamp[0]

        return self.timestamp[self.owner]

events = []
if len(sys.argv) > 1:
    path = sys.argv[1]
    df = pd.read_csv(path, sep=" ", header=None)
    df.columns = ['Kind', 'Owner', 'Source', 'Destination', 'T1', 'T2', 'T3']

    events = [Event(row.Kind, row.Owner, row.Source, row.Destination, [row.T1, row.T2, row.T3]) for index, row in df.iterrows() ]
else:
    for line in sys.stdin:
        splittedLine = line.split(' ')
        kind        = splittedLine[0]
        owner       = int(splittedLine[1])
        source      = splittedLine[2] if splittedLine[2] == 'NaN' else int(splittedLine[2])
        destination = splittedLine[3] if splittedLine[3] == 'NaN' else int(splittedLine[3])
        
        timestamp   = list(map(lambda x : int(x), splittedLine[4:]))

        events.append(Event(kind, owner, source, destination, timestamp))

if len(events) == 0:
    sys.exit(0)



events.sort(key = lambda e : e.maxTimestamp())
maxTimestamp = events[-1].maxTimestamp()
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
        plt.text(e.owner - 0.275 * math.log(1 + len(e.timestamp), 2) + 1, e.maxTimestamp() + 0.1, "$({},{})$".format(",".join(map(str, e.timestamp)), string.ascii_uppercase[i]))

# events
while len(events) > 0:
    e = events.pop(0)

    if e.kind == 'local':
        plt.scatter([e.owner + 1], [e.ownerTimestamp()], c = "k")
    elif e.kind == 'sent':
        s = e
        (x0, y0) = (s.owner + 1, s.ownerTimestamp())
        r = next(e for e in events if e.kind == 'received' and e.source == s.source and e.destination == s.destination)
        (x1, y1)  = (r.owner + 1, r.maxTimestamp())
        
        if y0 != y1:
            plt.scatter([x0, x1], [y0, y1], c = "k")
            plt.plot([x0, x1], [y0, y1])
        else:
            plt.scatter([x0, x1], [y0, y1 + 0.2], c = "k")
            plt.plot([x0, x1], [y0, y1 + 0.2])
        events.remove(r)
    else:
        pass

plt.axis("off")
plt.show()