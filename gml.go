/*
Package gml provides utilities for working with XML documents using a simple node-based structure.
*/
package gml

import (
	"encoding/xml"
	"strings"

	"github.com/ovila98/ers"
)

// Node represents a node in an XML document.
type Node struct {
	// Tag is the name of the XML tag.
	Tag string
	// Parent is a pointer to the parent Node.
	Parent *Node
	// Children holds the child nodes of this Node.
	Children []*Node
	// InnerText contains the text within the XML node.
	InnerText string
	// Attributes holds the attributes of the XML node.
	Attributes map[string]string
}

// UnmarshalXML implements the xml.Unmarshaler interface for Node.
func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Tag = start.Name.Local
	n.Attributes = make(map[string]string)
	for _, attr := range start.Attr {
		n.Attributes[attr.Name.Local] = attr.Value
	}
	for {
		token, err := d.Token()
		if err != nil {
			return ers.Trace(err)
		}
		switch elem := token.(type) {
		case xml.StartElement:
			child := &Node{Parent: n}
			err := child.UnmarshalXML(d, elem)
			if err != nil {
				return ers.Trace(err)
			}
			n.Children = append(n.Children, child)
		case xml.CharData:
			n.InnerText = strings.TrimSpace(string(elem))
		case xml.EndElement:
			if elem.Name.Local == n.Tag {
				return nil
			}
		}
	}
}

// MarshalXML implements the xml.Marshaler interface for Node.
func (n *Node) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = n.Tag
	for key, value := range n.Attributes {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: key}, Value: value})
	}
	if len(n.Children) == 0 && n.InnerText == "" {
		err := e.EncodeToken(start)
		if err != nil {
			return ers.Trace(err)
		}
		return e.EncodeToken(xml.EndElement{Name: start.Name})
	}
	err := e.EncodeToken(start)
	if err != nil {
		return ers.Trace(err)
	}
	if n.InnerText != "" {
		err = e.EncodeToken(xml.CharData([]byte(n.InnerText)))
		if err != nil {
			return ers.Trace(err)
		}
	}
	for _, child := range n.Children {
		err = e.Encode(child)
		if err != nil {
			return ers.Trace(err)
		}
	}
	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return ers.Trace(err)
	}
	return e.Flush()
}

// FindChild recursively searches for the first occurrence of a node with the given tag name within the XML node tree.
func (n *Node) FindChild(tag string) *Node {
	if n == nil {
		return nil
	}
	if n.Tag == tag {
		return n
	}
	for _, child := range n.Children {
		if found := child.FindChild(tag); found != nil {
			return found
		}
	}
	return nil
}

// AppendChild appends a child node to the Node and returns the child node.
func (n *Node) AppendChild(node *Node) *Node {
	node.Parent = n
	n.Children = append(n.Children, node)
	return node
}

// ChainAppendChildren appends multiple child nodes to the Node. Returns the parent node.
func (n *Node) ChainAppendChildren(nodes ...*Node) *Node {
	for _, node := range nodes {
		n.AppendChild(node)
	}
	return n
}

// ChainAppendChild appends a child node to the Node and returns the parent node.
func (n *Node) ChainAppendChild(node *Node) *Node {
	n.AppendChild(node)
	return n
}

// SetAttribute sets the given attribute to the XML node.
func (n *Node) SetAttribute(name, value string) {
	if n.Attributes == nil {
		n.Attributes = make(map[string]string)
	}
	n.Attributes[name] = value
}

// GetAttribute gets the given attribute from the XML node.
func (n *Node) GetAttribute(name string) string {
	if n.Attributes == nil {
		return ""
	}
	return n.Attributes[name]
}

// HasAttribute checks if the XML node has the given attribute.
func (n *Node) HasAttribute(name string) bool {
	if n.Attributes == nil {
		return false
	}
	_, has := n.Attributes[name]
	return has
}

// RemoveAttribute removes the given attribute from the XML node.
func (n *Node) RemoveAttribute(name string) {
	if n.Attributes == nil {
		return
	}
	delete(n.Attributes, name)
}

// RemoveChildrenWithTag removes all children with the specified tag.
func (n *Node) RemoveChildrenWithTag(tag string) {
	var newChildren []*Node
	for _, child := range n.Children {
		if child.Tag != tag {
			newChildren = append(newChildren, child)
		}
	}
	n.Children = newChildren
}

// CheckPath checks if a specified path exists from the current node using DFS.
func (n *Node) CheckPath(path ...string) bool {
	if len(path) == 0 {
		return true
	}
	return n.dfsCheckPath(path)
}

// dfsCheckPath is a helper function that performs DFS to check if a specified path exists.
func (n *Node) dfsCheckPath(path []string) bool {
	if len(path) == 0 {
		return true
	}
	for _, child := range n.Children {
		if child.Tag == path[0] {
			if child.dfsCheckPath(path[1:]) {
				return true
			}
		}
	}
	return false
}

// CreatePath creates a path of nested XML nodes based on the provided tags.
func (n *Node) CreatePath(path ...string) *Node {
	currentNode := n
	for _, tag := range path {
		currentNode = currentNode.AppendChild(&Node{Tag: tag})
	}
	return currentNode
}

// CreateUniquePath creates a unique path of nested XML nodes based on the provided tags.
func (n *Node) CreateUniquePath(path ...string) *Node {
	if len(path) == 0 {
		return n
	}
	n.RemoveChildrenWithTag(path[0])
	return n.CreatePath(path...)
}

// EnsurePath checks if a specified path exists from the current node and creates any missing nodes.
func (n *Node) EnsurePath(path ...string) *Node {
	currentNode := n
	for i, tag := range path {
		found := false
		for _, child := range currentNode.Children {
			if child.Tag == tag {
				currentNode = child
				found = true
				break
			}
		}
		if !found {
			return currentNode.CreatePath(path[i:]...)
		}
	}
	return currentNode
}

// Bytes returns a byte representation of the XML node and its children.
func (n *Node) Bytes() []byte {
	output, err := xml.MarshalIndent(n, "", "  ")
	if err != nil {
		return nil
	}
	return output
}

// String returns a string representation of the XML node and its children.
func (n *Node) String() string {
	return string(n.Bytes())
}
