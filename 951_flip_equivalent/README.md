# Solution

## Thanks to srinivas_bodduru: https://leetcode.com/problems/cousins-in-binary-tree-ii/solutions/5955106/beats-100-00-explained-with-example/?envType=daily-question&envId=2024-10-23

Intuition
The code is designed to replace the value of each node in a binary tree with a new value that depends on its sibling nodes. Specifically, for each node, the new value is calculated as the sum of the values of all child nodes across the current level, excluding the current node's children.

Here's the intuition broken down step-by-step:

Current Node's Children Influence: For each node at a certain level, we look at the sum of all child nodes at that level (across all nodes at the same level).

Replacement Calculation: The node's left and right child values are replaced by the sum of all child nodes at that level minus the sum of their own children. In essence, each child node receives the sum of its "cousins" (children of its parent's sibling nodes).

Level-by-Level Traversal: The replacement process happens one level at a time, similar to a breadth-first traversal (BFS), but using DFS recursion to traverse down the tree while calculating and setting new values for each child node based on its cousins.

Root Initialization: The root node is initialized to 0, as it has no cousins at its level.

Approach
The execution can be broken down into these key steps:

Initialize Root Value:
The root of the tree is set to 0 since there are no sibling nodes for it at the top level.
Start DFS from the Root:
The dfs function is initially called with an array containing only the root node. This array represents the nodes at the current level of the tree.
Calculate Sum of Children:
For the current level of nodes (held in the arr array), you calculate the sum of all the child node values. This sum will be used to update the values of nodes at the next level.
You iterate over each node in the arr, checking if they have left and/or right children, and if so, you add their values to the sum.

Update Child Node Values:
For each node in the current level, you calculate the sum of its own children. Then, for each child (left and right), you subtract the sum of its own children from the total sum of all children at this level. This gives the new value for the child (which is essentially the sum of "cousins").
After updating each child’s value, you add the child node to the childArr array to be processed in the next recursive call.

Recursive Call for the Next Level:
After processing all nodes at the current level, you recursively call dfs with the childArr, which now contains all the nodes at the next level.
The recursion continues level by level until all nodes are processed.

End Condition:
The recursion stops when arr (the array of nodes at the current level) is empty, meaning no more child nodes are left to process.
Consider a simple tree:
```batch
  10
 /  \
5    15
/ \ /
2 6 12 20
```
Initial root value: Set the root node 10 to 0.

First level (Nodes: 5, 15):

Calculate the sum of children: 2 + 6 + 12 + 20 = 40.

Replace 5's children:

5 has children 2 and 6. The sum of their values is 8. Replace 2 and 6 with 40 - 8 = 32.
Replace 15's children:

15 has children 12 and 20. The sum of their values is 32. Replace 12 and 20 with 40 - 32 = 8.
After this step, the tree becomes:
```batch
  0
 /  \
5    15
/ \ /
32 32 8 8
```
Second level (Nodes: 32, 32, 8, 8):

Since these nodes have no children, the recursion will stop at this point.

disclaimer: i have written the code in js and converted it to other langs via gpt please resolve the conversion bugs if any thank you

```batch
Complexity
Time complexity:O(N)
Space complexity:O(H)
upvote.jpeg
```

```javascript
var replaceValueInTree = function (root) {
    function dfs(arr) {
        if (arr.length == 0) return

        let sum = 0

        for (let node of arr) {
            if (!node) continue
            if (node.left) sum += node.left.val
            if (node.right) sum += node.right.val
        }

        let childArr = []

        for (let node of arr) {
            let curSum = 0

            if (node.left) curSum += node.left.val
            if (node.right) curSum += node.right.val

            if (node.left) {
                node.left.val = sum - curSum
                childArr.push(node.left)
            }
            if (node.right) {
                node.right.val = sum - curSum
                childArr.push(node.right)
            }
        }

        dfs(childArr)
    }

    root.val = 0
    dfs([root])

    return root
};
```