package path

import (
	"errors"
	"strings"
)

const sep string = "/"
const pathParamPrefix = ":"

type pathNode struct {
	segment string // path segment.
	isParam bool   // if this segment is path param.
}

type pathTree map[pathNode][]pathNode

var globalRouter pathTree = make(map[pathNode][]pathNode)

func RegisterPath(path string) error {
	if len(path) == 0 || !strings.HasPrefix(path, sep) {
		// if the path's length is 0 or it doesn't start with '/', return error.
		return errors.New("path is illegal, it must not be blank and start with '/'")
	}
	segments := strings.Split(path, sep)
	root := pathNode{sep, false}
	if _, ok := globalRouter[root]; ok {
		globalRouter[root] = make([]pathNode, 0)
	}
	var preNode = root
	for _, seg := range segments {
		if len(seg) == 0 {
			continue
		}
		node := pathNode{segment: seg, isParam: strings.HasPrefix(seg, pathParamPrefix)}
		if subNodes, ok := globalRouter[preNode]; !ok {
			subNodes = make([]pathNode, 0)
			globalRouter[preNode] = subNodes
		}
		globalRouter[preNode] = append(globalRouter[preNode], node)
		preNode = node

	}
	return nil
}

func GetAllPaths() []string {
	return nil
}
