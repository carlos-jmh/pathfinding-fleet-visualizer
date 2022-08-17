# %%
import matplotlib.pyplot as plt
import networkx as nx
import random as r
# %%
def create(col, row):
    G = nx.grid_2d_graph(col,row)

    # 1 = Plains; 2-3 = Wheat; 4 = Ore; 5-6 = Mountains;
    terrain_choices = [1, 2, 3, 4, 5, 6]
    terrain_weights = [60, 8, 8, 8, 8, 8]

    w = r.choices(terrain_choices, terrain_weights, k = col * row)
    weights = {node: we for node, we in zip(G.nodes(), w)}

    return G, weights
# %%
def plot(G, weights, dimStart = 10, dimEnd = 10, nodeSize = 2000):
    plt.figure(figsize=(dimStart,dimEnd))
    pos = {(x, y):(x, y) for x, y in G.nodes()}
    labelsdict = {x:f'{x} {weights[y]}' for x, y in zip(G.nodes, weights)}

    nx.draw(G, pos, 
            node_color='lightgreen',
            labels=labelsdict, 
            with_labels=True,
            node_size=nodeSize)

    print(f"{len(G.nodes())} nodes.")

# %%
if __name__ == "__main__":
    G, weights = create(5, 5)
    print(weights)
    plot(G, weights)

# %%
