package gee

import "strings"

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang
	part     string  // 路由中的一部分，例如 :lang
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true
}

// 查找第一个匹配的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 查找全部匹配的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	// 到达高度了就停止
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]       // 获取当前的规则
	child := n.matchChild(part) // 尝试用当前的规则进行匹配
	// 如果没有匹配成功，就新建一个节点，并加入到当前节点的孩子们中去
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	// 递归进行插入
	child.insert(pattern, parts, height+1)
}

// 查询节点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]             // 获取当前的规则
	children := n.matchChildren(part) // 尝试用当前的规则进行匹配
	// 遍历所有当前匹配的节点进行递归匹配
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
