package internal

import (
	"errors"
	"strings"

	"github.com/fujianbang/summer"
)

type Tree struct {
	root *node
}

func (tree *Tree) AddRouter(uri string, handler summer.ControllerHandler) error {
	n := tree.root
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node

		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0 {
			for _, cnode := range childNodes {
				if cnode.segment == segment {
					objNode = cnode
					break
				}
			}
		}

		if objNode == nil {
			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.children = append(n.children, cnode)

			objNode = cnode
		}

		n = objNode
	}

	return nil
}

func (tree *Tree) FindHandler(uri string) summer.ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}

	return matchNode.handler
}
