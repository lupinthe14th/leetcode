package serializeanddeserializebinarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	add := func(c *[]string, node *TreeNode) {
		if node == nil {
			*c = append(*c, "nil")
		} else {
			*c = append(*c, strconv.Itoa(node.Val))
		}
	}
	q := []*TreeNode{root}
	c := []string{strconv.Itoa(root.Val)}
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
			add(&c, q[i].Left)
			add(&c, q[i].Right)
		}
		q = q[l:]
	}
	out := strings.Join(c, ",")
	return out
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	atoi := func(s string) int {
		a, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return a
	}
	c := strings.Split(data, ",")

	if len(data) == 0 {
		return nil
	}

	t := &TreeNode{Val: atoi(c[0])}
	q := []*TreeNode{t}

	i := 1
	for len(q) > 0 {
		l := len(q)
		for j := 0; j < l; j++ {
			if c[i] == "nil" {
				q[j].Left = nil
			} else {
				q[j].Left = &TreeNode{Val: atoi(c[i])}
				q = append(q, q[j].Left)
			}
			i++
			if c[i] == "nil" {
				q[j].Right = nil
			} else {
				q[j].Right = &TreeNode{Val: atoi(c[i])}
				q = append(q, q[j].Right)
			}
			i++
		}
		q = q[l:]
	}
	return t
}
