import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

class Solution {
  public static void main(String[] args) {
    TreeNode root = new TreeNode(3);
    root.left = new TreeNode(5);
    root.right = new TreeNode(1);
    root.left.left = new TreeNode(6);
    root.left.right = new TreeNode(2);
    root.right.left = new TreeNode(0);
    root.right.right = new TreeNode(8);
    root.left.right.left = new TreeNode(7);
    root.left.right.right = new TreeNode(4);
    TreeNode target = root.left;
    System.out.println(distanceK(root, target, 2));

    TreeNode parent = new TreeNode(20);
    parent.left = new TreeNode(8);
    parent.right = new TreeNode(22);
    parent.left.left = new TreeNode(4);
    parent.left.right = new TreeNode(12);
    parent.left.right.left = new TreeNode(10);
    parent.left.right.right = new TreeNode(14);
    TreeNode nodeTarget = root.left.right.right;
    System.out.println(distanceK(parent, nodeTarget, 3));
  }

  public static List<Integer> distanceK(TreeNode root, TreeNode target, int K) {
    List<Integer> ans = new ArrayList<>();
    if (K == 0) {
      ans.add(target.val);
      return ans;
    }
    dfs(root, 0, target, ans, K);

    return ans;
  }

  private static int dfs(TreeNode root, int dept, TreeNode target, List<Integer> ans, int K) {
    if (root == null) {
      return 0;
    }
    if (dept == K) {
      ans.add(root.val);
      return 0;
    }

    // add nodes which are in the subtree rooted at the target
    if (target == root || dept > 0) {
      dfs(root.left, dept + 1, target, ans, K);
      dfs(root.right, dept + 1, target, ans, K);
    } else {
      // add nodes which are in the rest of the tree
      int left = dfs(root.left, dept, target, ans, K);
      int right = dfs(root.right, dept, target, ans, K);
      if (left == K || right == K) {
        ans.add(root.val);
        return 0;
      } else if (left > 0) {
        dfs(root.right, left + 1, target, ans, K);
        return left + 1;
      } else if (right > 0) {
        dfs(root.left, right + 1, target, ans, K);
        return right + 1;
      }
    }
    if (target == root) {
      return 1;
    }

    return 0;
  }
}

class TreeNode {
  public int val;
  public TreeNode left;
  public TreeNode right;

  public TreeNode() {
  }

  public TreeNode(int val) {
    this.val = val;
  }
}