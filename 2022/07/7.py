from pathlib import Path

class Node:
  def __init__(self, name, is_dir, size=None) -> None:
    self.name = name
    self.is_dir = is_dir
    self.size = size
    self.children = []
  def __str__(self) -> str:
    return "{}; dir: {}; size: {}; children: {}".format(self.name, self.is_dir, self.size, self.children)

def add_node_to_tree(root: Node, path: Path, name, size):
  print(root, path)
  parts = path.parts
  current_node = root

  # Iterate over the parts of the path and add nodes
  # to the tree as needed
  for part in parts:
    found = False
    # Check if the current node has a child with the
    # given name
    for child in current_node.children:
      if child.name == part:
        current_node = child
        found = True
        break
    # If no child with the given name was found, create
    # a new node and add it as a child of the current node
    if not found:
      new_node = Node(path.joinpath(name), is_dir=False, size=size)
      current_node.children.append(new_node)
      current_node = new_node

#root = Node('/', is_dir=True)

def cd(cur: Path, nxt: str):
  if nxt == "..":
    return cur.parent
  return cur.joinpath(nxt)


dir = Path('/')
root = Node(dir, True)
for line in open('inp').readlines()[1:]:
  parts = [p.strip() for p in line.split(' ')]
  if parts[0] == '$':
    if parts[1] == "cd":
      #print("cd", parts[2])
      dir = cd(dir, parts[2])
      print(dir)
  else:
    if parts[0] != "dir":
      add_node_to_tree(root, dir, parts[1], int(parts[0]))



  #print(line)

#for line in open('inp').readlines():

#  print(line)
  #print(dir)
  #add_node_to_tree(root, line)

#print(root)
