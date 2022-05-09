package internal

import (
	"strings"

	"github.com/fujianbang/summer"
)

type node struct {
	isLast   bool
	segment  string
	handler  summer.ControllerHandler
	children []*node
}

func newNode() *node {
	return &node{
		isLast:   false,
		segment:  "",
		handler:  nil,
		children: nil,
	}
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.children
	}

	nodes := make([]*node, 0, len(n.children))
	for _, cnode := range n.children {
		if isWildSegment(cnode.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}

	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildNodes(segment)
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}
