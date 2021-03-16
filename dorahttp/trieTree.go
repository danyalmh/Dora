package dorahttp

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

//Tree is trie tree
type Tree struct {
	method map[string]*Node
}

//Node is node of tree
type Node struct {
	label    string
	handler  DoraHandler
	children map[string]*Node
}

//Params is all Parameter.
type Params map[string]string

//Result is search result.
type Result struct {
	handler DoraHandler
	params  *Params
}

const (
	pathDelimiter     = "/"
	paramDelimiter    = ":"
	leftPtnDelimiter  = "["
	rightPtnDelimiter = "]"
	ptnWildcard       = "(.+)"
)

//NewTree creates new trie tree.
func NewTree() *Tree {

	return &Tree{
		method: map[string]*Node{

			MethodGet: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			MethodPost: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			MethodPut: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			MethodPatch: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			MethodDelete: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
			MethodOptions: {
				label:    "",
				handler:  nil,
				children: make(map[string]*Node),
			},
		},
	}
}

//Insert inserts a route definition to tree
func (t *Tree) Insert(method string, path string, handler DoraHandler) error {
	currNode := t.method[method]

	if path == pathDelimiter {
		if len(currNode.label) != 0 && currNode.handler == nil {
			return errors.New("Root node already exists")
		}

		currNode.label = path
		currNode.handler = handler
	}

	for _, l := range deleteEmpty(strings.Split(path, pathDelimiter)) {
		if nextNode, ok := currNode.children[l]; ok {
			currNode = nextNode
		} else {
			currNode.children[l] = &Node{
				label:    l,
				handler:  handler,
				children: make(map[string]*Node),
			}

			currNode = currNode.children[l]
		}
	}

	return nil
}

type regCache struct {
	s sync.Map
}

//Get gets a compiled regexp from cache or create it.
func (rc *regCache) Get(ptn string) (*regexp.Regexp, error) {

	v, ok := rc.s.Load(ptn)
	if ok {
		reg, ok := v.(*regexp.Regexp)
		if !ok {
			return nil, fmt.Errorf("the value of %q is wrong", ptn)
		}

		return reg, nil
	}

	reg, err := regexp.Compile(ptn)
	if err != nil {
		return nil, err
	}

	rc.s.Store(ptn, reg)
	return reg, nil
}

var regC = &regCache{}

//Search searches a path from a tree
func (t *Tree) Search(method string, path string) (*Result, error) {

	var params Params = make(Params)
	n := t.method[method]

	if len(n.label) == 0 && len(n.children) == 0 {
		return nil, errors.New("tree is empty")
	}

	label := deleteEmpty(strings.Split(path, pathDelimiter))
	curNode := n

	for _, l := range label {
		if nextNode, ok := curNode.children[l]; ok {
			curNode = nextNode
		} else {
			// pattern matching priority depends on an order of routing definition
			// ex.
			// 1 /foo/:id
			// 2 /foo/:id[^\d+$]
			// 3 /foo/:id[^\w+$]
			// priority is 1, 2, 3
			if len(curNode.children) == 0 {
				return &Result{}, errors.New("handler is not registered")
			}

			count := 0
			for c := range curNode.children {
				if string([]rune(c)[0]) == paramDelimiter {
					ptn := getPattern(c)

					reg, err := regC.Get(ptn)
					if err != nil {
						return nil, err
					}
					if reg.Match([]byte(l)) {
						params[getParameter(c)] = l
						curNode = curNode.children[c]
						count++
						break
					} else {
						return &Result{}, errors.New("param does not match")
					}
				}

				count++

				// If no match is found until the last loop.
				if count == len(curNode.children) {
					return &Result{}, errors.New("handler is not registered")
				}
			}
		}
	}

	if curNode.handler == nil {
		return &Result{}, errors.New("handler is not registered")
	}

	return &Result{
		handler: curNode.handler,
		params:  &params,
	}, nil
}

// getPattern gets a pattern from a label.
// ex.
// :id[^\d+$] → ^\d+$
// :id        → (.+)
func getPattern(label string) string {
	leftI := strings.Index(label, leftPtnDelimiter)
	rightI := strings.Index(label, rightPtnDelimiter)

	// if label has not pattern, return wild card pattern as default.
	if leftI == -1 || rightI == -1 {
		return ptnWildcard
	}

	return label[leftI+1 : rightI]
}

// getParameter gets a parameter from a label.
// ex.
// :id[^\d+$] → id
// :id        → id
func getParameter(label string) string {
	leftI := strings.Index(label, paramDelimiter)
	rightI := func(l string) int {
		r := []rune(l)

		var n int

		for i := 0; i < len(r); i++ {
			n = i
			if string(r[i]) == leftPtnDelimiter {
				n = i
				break
			} else if i == len(r)-1 {
				n = i + 1
				break
			}
		}

		return n
	}(label)

	return label[leftI+1 : rightI]
}

// deleteEmpty removes an empty value in slice.
func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
