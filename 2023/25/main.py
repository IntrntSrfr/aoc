import networkx as nx
from math import prod

G = nx.Graph()
for entry in open('inp').readlines():
    src, tgts = entry.split(':')
    G.add_edges_from([(src, t) for t in tgts.split()])
_, parts = nx.stoer_wagner(G)
print(prod([len(p) for p in parts]))
