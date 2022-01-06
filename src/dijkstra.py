from collections import defaultdict
import heapq as heap

def dijkstra(G, weights, startingNode, target):
	visited = set()
	parentsMap = {}
	pq = []
	nodeCosts = defaultdict(lambda: float('inf'))
	nodeCosts[startingNode] = 0
	heap.heappush(pq, (0, startingNode))
 
	while pq:

		_, node = heap.heappop(pq)
		if node == target:
			return parentsMap, nodeCosts
		visited.add(node)
 
		for adjNode in G.neighbors(node):
			if adjNode in visited:	continue
				
			newCost = nodeCosts[node] + weights[node]
			if nodeCosts[adjNode] > newCost:
				parentsMap[adjNode] = node
				nodeCosts[adjNode] = newCost
				heap.heappush(pq, (newCost, adjNode))
        
	return parentsMap, nodeCosts