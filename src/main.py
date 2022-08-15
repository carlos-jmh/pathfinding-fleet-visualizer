# %%
from collections import deque
from networkx.algorithms.shortest_paths import weighted
import graph
import dijkstra

# %% Initializing Graph
G, weights = graph.create(10, 10)

# %% Plotting Graph
graph.plot(G, weights, 12, 12, 4000)

# %% Running Dijkstra's
source = (0, 0)
target = (9, 9)
parentsMap, nodeCosts = dijkstra.dijkstra(G, weights, source, target)

# %% Going through parentsMap to find path
path = deque()
pathW = deque()
temp = target

while temp != source:
    path.appendleft(temp)
    temp = parentsMap[temp]
    pathW.appendleft(weights[temp])
path.appendleft(source)

# %% Print Result
print(list(path), "Total cost: ", nodeCosts[target])
print("Cost per Node: ", list(pathW))
# %%
