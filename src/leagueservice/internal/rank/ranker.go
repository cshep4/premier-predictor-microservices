package rank

type (
	Ranker struct {
		root *rankNode
		size int
	}

	rankNode struct {
		leftSize int
		left     *rankNode
		right    *rankNode
		val      int
	}
)

func (r *Ranker) Insert(val int) {
	r.size++

	if r.root == nil {
		r.root = &rankNode{val: val}
		return
	}

	r.root.insert(val)
}

func (r *Ranker) Rank(val int) (int, bool) {
	if r.root == nil {
		return 0, false
	}

	rank, ok := r.root.rank(val)
	if !ok {
		return 0, false
	}

	// r.size - rank as root.rank(val) returns in reverse order
	return r.size - rank, true
}

func (r *Ranker) Clear() {
	r.root = nil
	r.size = 0
}

func (r *rankNode) insert(val int) {
	if val <= r.val {
		r.insertLeft(val)
		return
	}

	r.insertRight(val)
}

func (r *rankNode) insertLeft(val int) {
	r.leftSize++

	if r.left == nil {
		r.left = &rankNode{val: val}
		return
	}

	r.left.insert(val)
}

func (r *rankNode) insertRight(val int) {
	if r.right == nil {
		r.right = &rankNode{val: val}
		return
	}

	r.right.insert(val)
}

func (r *rankNode) rank(val int) (int, bool) {
	if r.val == val {
		return r.leftSize, true
	}

	if val < r.val {
		if r.left == nil {
			return 0, false
		}

		return r.left.rank(val)
	}

	if r.right == nil {
		return 0, false
	}

	rightRank, ok := r.right.rank(val)
	if !ok {
		return 0, false
	}

	return r.leftSize + 1 + rightRank, true
}
