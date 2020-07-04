package test

import (
	"container/list"
	"strings"
)

// 字段匹配的树
type matchTree struct {
	key      string
	value    interface{}
	setted   bool
	children map[string]*matchTree
}

func newMatchTree() *matchTree {
	res := new(matchTree)
	res.key = "root"
	res.children = map[string]*matchTree{}
	return res
}

func (root *matchTree) setNodeValue(fieldName string, value interface{}) {
	names := strings.Split(fieldName, ".")
	for _, name := range names {
		if _, ok := root.children[name]; !ok {
			root.children[name] = newMatchTree()
		}
		root = root.children[name]
		root.key = name
	}
	root.setted = true
	root.value = value
}

func (root *matchTree) getNodeValue(fieldName string) interface{} {
	queue := list.New()
	queue.PushBack(root)
	names := strings.Split(fieldName, ".")
	for _, name := range names {
		for i := queue.Len(); i > 0; i-- {
			val := queue.Remove(queue.Front())
			node := val.(*matchTree)
			if child, ok := node.children[name]; ok {
				queue.PushBack(child)
			}
			if child, ok := node.children["*"]; ok {
				queue.PushBack(child)
			}
			if child, ok := node.children["**"]; ok {
				queue.PushBack(child)
			}
			if node.key == "**" {
				queue.PushBack(node)
			}
		}
	}
	for ele := queue.Front(); ele != nil; ele = ele.Next() {
		node := ele.Value.(*matchTree)
		if node.setted {
			return node.value
		}
	}
	return nil
}
