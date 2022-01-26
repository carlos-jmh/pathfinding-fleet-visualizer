# pathfinding-fleet-visualizer
An early version of a pathfinding and fleet management algorithm visualizer I'm working on.

Currently: When running main.py cell by cell, it'll first draw a grid with a group of weights, within a specific random distribution* (more below).
Format: (x, y) weight

![nodes_example](https://user-images.githubusercontent.com/75096034/148436760-db55afbd-b53c-42d4-8d28-a5e44a428d4a.png)

Then, using Djikstra's pathfinding algorithm, it will return a list which represents a path from source to target; along with the costs per node to get to target.  

Note: Djikstra's guarantees the shortest path but be aware, multiple shortest paths might exist.  
Note2: The weight of the source node is counted, but not of the target node, this is because I want to count the nodes we have *traveled across* not the ones we have only reached.

![path_example](https://user-images.githubusercontent.com/75096034/148437806-030b3e0f-5e77-4215-8d86-e72794fa4a47.png)

specific random distribution* : 60% chance of a 1, 16% change of a 2-3, 8% of a 4, 16% chance of a 5-6. These weight distributions represent future features of the map (1 = Plains, 2-3 = Wheat, 4 = Ore, 5-6 = Mountains).
