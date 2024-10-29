This was asked in LinkedIn Interview
Given a list of child->parent relationships, build a binary tree out of it. All the element Ids inside the tree are unique.


Example:


Given the following relationships:


```
Child Parent IsLeft
15 20 true
19 80 true
17 20 false
16 80 false
80 50 false
50 null false
20 50 true
```

You should return the following tree:
```

    50
   /  \
  20   80
 / \   / \
15 17 19 16
```

Structure

```go
type Relation struct {
	Child  int
	Parent int // -1 if don't have
	IsLeft bool
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
```