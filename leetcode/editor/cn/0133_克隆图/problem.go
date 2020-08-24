package p0133

type Node struct {
  Val int
  Neighbors []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
  if node == nil {
    return nil
  }

  visited := make(map[*Node]*Node)
  return dfs(node, visited)
}

func dfs(src *Node, visited map[*Node]*Node) *Node {
  if src == nil {
    return nil
  }

  if dst, ok := visited[src]; ok {
    return dst
  }

  dst := &Node{Val: src.Val}
  visited[src] = dst
  for _, n := range src.Neighbors {
    dst.Neighbors = append(dst.Neighbors, dfs(n, visited))
  }

  return dst
}
//leetcode submit region end(Prohibit modification and deletion)
