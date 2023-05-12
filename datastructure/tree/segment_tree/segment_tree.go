package segment_tree

type node struct {
	l, r int
	data int
}

var ele []int
var tree []node

func Init(nums []int) {
	tree = make([]node, 4*len(nums))
	ele = make([]int, len(nums))
	copy(ele, nums)
	build(0, len(nums), 0)
}

// build 建树
// 每个结点所掌管的区间为 [l, r)
func build(l, r, treeIndex int) {
	tree[treeIndex].l, tree[treeIndex].r = l, r
	if l+1 >= r {
		tree[treeIndex].data = ele[l]
		return
	}

	mid := l + (r-l)>>1
	build(l, mid, leftChild(treeIndex))
	build(mid, r, rightChild(treeIndex))

	// 更新当前树结点的值
	updateNode(treeIndex)
}

// Get 单点查询 返回索引为idx的值
func Get(idx int) int {
	return ele[idx]
}

// Query 区间查询
// 查询[l, r)区间中的特定值
func Query(l, r, treeIndex int) int {
	// [l, r)区间是否被当前结点命中
	if l == tree[treeIndex].l && r == tree[treeIndex].r {
		return tree[treeIndex].data
	}

	mid := tree[treeIndex].l + (tree[treeIndex].r-tree[treeIndex].l)>>1
	// 剩下三种情况
	// 待查询的区间在结点区间的左边（即[l, r) 是 [tree[idx].l, mid) 的子集）
	// 问我的左结点要值
	if r <= mid {
		return Query(l, r, leftChild(treeIndex))
	}
	// 待查询的区间在结点区间的右边（即[l, r) 是 [mid, tree[idx].r) 的子集）
	// 问我的右结点要值
	if l >= mid {
		return Query(l, r, rightChild(treeIndex))
	}
	// 待查询的区间在结点区间的中间（即 tree[idx].l <= l < mid，mid < r <= tree[idx].r）
	return Query(l, mid, leftChild(treeIndex)) + Query(mid, r, rightChild(treeIndex))
}

// updateNode 树结点的更新逻辑
// 这里以记录区间和为例
func updateNode(treeIndex int) {
	tree[treeIndex].data = tree[leftChild(treeIndex)].data + tree[rightChild(treeIndex)].data
}

// leftChild 返回结点treeIndex的左结点的索引
func leftChild(treeIndex int) int {
	return 2*treeIndex + 1
}

// rightChild 返回结点treeIndex的右结点的索引
func rightChild(treeIndex int) int {
	return 2*treeIndex + 2
}
