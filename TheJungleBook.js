function getTreeHeigh(predators) {
  var dt = depthTracker(predators);
  return getHeigh(dt) + 1
}

function depthTracker(nodes) {
  var depth = new Array(nodes.length).fill(0);

  for (let i = 0; i < nodes.length; i++) {
    if (nodes[i] > -1 && depth[i] == 0) {
      depth[i] = checkDepth(i, depth, nodes)
    }
  }

  return depth;
}

function checkDepth(index, depth, nodes) {
  if (index == -1)
    return 0

  if(depth[index] > 0)
    return depth[index]
  
  var nextIdx = nodes[index]
  if(nextIdx == -1){
    depth[index] = 0
    return 0
  }

  var thisDepth = checkDepth(nextIdx, depth, nodes) + 1

  depth[index] = thisDepth;
  return thisDepth
}

function getHeigh(depth){
  return Math.max(...depth);
}

console.log(getTreeHeigh([1, - 1, 3, -1]))