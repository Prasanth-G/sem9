import networkx as nx
import matplotlib.pyplot as plt

p = 0.4
k = 6

#rg = nx.watts_strogatz_graph(n, k, p)

# for i in range(10, 20):
n = 10000
rg = nx.erdos_renyi_graph(n, p, seed=None, directed=False)
print("Average Clustering: ", nx.average_clustering(rg))
# plt.figure(str(n))
# nx.draw(rg, with_labels=True)
# plt.show()