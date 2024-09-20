package gml

import (
	"encoding/xml"
	"testing"
)

var xmlDataToUnmarshal = `
    <library name="City Library" location="Downtown" rating="4.3">
        <section name="Fiction">
            <book>
                <title>The Great Gatsby</title>
                <author>F. Scott Fitzgerald</author>
                <published>1925</published>
            </book>
            <book>
                <title>1984</title>
                <author>George Orwell</author>
                <published>1949</published>
            </book>
        </section>
        <section name="Non-Fiction">
            <book>
                <title>Sapiens: A Brief History of Humankind</title>
                <author>Yuval Noah Harari</author>
                <published>2011</published>
            </book>
            <book>
                <title>Educated</title>
                <author>Tara Westover</author>
                <published>2018</published>
            </book>
        </section>
    </library>`

var xmlDataMarshaled = "<library name=\"City Library\" location=\"Downtown\" rating=\"4.3\"><section name=\"Fiction\"><book><title>1984</title><author>George Orwell</author><published>1949</published></book></section></library>"

// TestUnmarshalXML tests the UnmarshalXML function
func TestUnmarshalXML(t *testing.T) {
	var root Node
	err := xml.Unmarshal([]byte(xmlDataToUnmarshal), &root)
	if err != nil {
		t.Fatalf("Error during XML unmarshaling: %v", err)
	}

	// Check root node
	if root.Tag != "library" {
		t.Errorf("Expected root tag to be 'library', got '%s'", root.Tag)
	}

	if root.GetAttribute("name") != "City Library" {
		t.Errorf("Expected attribute 'name' to be 'City Library', got '%s'", root.GetAttribute("name"))
	}

	// Check first section
	section := root.FindChild("section")
	if section == nil || section.GetAttribute("name") != "Fiction" {
		t.Fatalf("Expected first section 'Fiction' not found")
	}

	// Check first book in Fiction
	book := section.FindChild("book")
	if book == nil {
		t.Fatalf("Expected first book in Fiction section not found")
	}

	title := book.FindChild("title")
	if title == nil || title.InnerText != "The Great Gatsby" {
		t.Errorf("Expected first book title to be 'The Great Gatsby', got '%s'", title.InnerText)
	}
}

// TestMarshalXML tests the MarshalXML function
func TestMarshalXML(t *testing.T) {
	// Construct the same node as in the marshaling example
	root := &Node{
		Tag: "library",
		Attributes: map[string]string{
			"name":     "City Library",
			"location": "Downtown",
			"rating":   "4.3",
		},
		Children: []*Node{
			{
				Tag: "section",
				Attributes: map[string]string{
					"name": "Fiction",
				},
				Children: []*Node{
					{
						Tag: "book",
						Children: []*Node{
							{
								Tag:       "title",
								InnerText: "1984",
							},
							{
								Tag:       "author",
								InnerText: "George Orwell",
							},
							{
								Tag:       "published",
								InnerText: "1949",
							},
						},
					},
				},
			},
		},
	}

	outputBytes, err := xml.Marshal(root)
	if err != nil {
		t.Fatalf("Error during XML marshaling: %v", err)
	}

	output := string(outputBytes)
	if output != xmlDataMarshaled {
		t.Errorf("Expected marshaled XML:\n%s\nGot:\n%s", xmlDataMarshaled, output)
	}
}

// TestAppendChild tests the AppendChild and ChainAppendChild methods
func TestAppendChild(t *testing.T) {
	root := &Node{Tag: "root"}
	child1 := &Node{Tag: "child1"}
	child2 := &Node{Tag: "child2"}

	root.AppendChild(child1)
	root.AppendChild(child2)

	if len(root.Children) != 2 {
		t.Fatalf("Expected 2 children, got %d", len(root.Children))
	}

	if root.Children[0].Tag != "child1" || root.Children[1].Tag != "child2" {
		t.Errorf("Expected children tags to be 'child1' and 'child2', got '%s' and '%s'", root.Children[0].Tag, root.Children[1].Tag)
	}
}

// TestChainAppendChildren tests appending multiple children in a chain
func TestChainAppendChildren(t *testing.T) {
	root := &Node{Tag: "root"}
	child1 := &Node{Tag: "child1"}
	child2 := &Node{Tag: "child2"}

	root.ChainAppendChildren(child1, child2)

	if len(root.Children) != 2 {
		t.Fatalf("Expected 2 children after chaining, got %d", len(root.Children))
	}

	if root.Children[0].Tag != "child1" || root.Children[1].Tag != "child2" {
		t.Errorf("Expected children tags to be 'child1' and 'child2', got '%s' and '%s'", root.Children[0].Tag, root.Children[1].Tag)
	}
}

// TestEnsurePath tests the EnsurePath method
func TestEnsurePath(t *testing.T) {
	root := &Node{Tag: "root"}
	node := root.EnsurePath("level1", "level2", "level3")

	if node == nil || node.Tag != "level3" {
		t.Fatalf("Expected node with tag 'level3', got '%s'", node.Tag)
	}

	if len(root.Children) != 1 || root.Children[0].Tag != "level1" {
		t.Fatalf("Expected first child to be 'level1', got '%s'", root.Children[0].Tag)
	}

	if root.Children[0].Children[0].Tag != "level2" {
		t.Errorf("Expected second level child to be 'level2', got '%s'", root.Children[0].Children[0].Tag)
	}
}

// TestRemoveChildrenWithTag tests removing children with a specific tag
func TestRemoveChildrenWithTag(t *testing.T) {
	root := &Node{
		Tag: "root",
		Children: []*Node{
			{Tag: "child1"},
			{Tag: "child2"},
			{Tag: "child1"},
		},
	}

	root.RemoveChildrenWithTag("child1")

	if len(root.Children) != 1 || root.Children[0].Tag != "child2" {
		t.Fatalf("Expected remaining child tag to be 'child2', got '%s'", root.Children[0].Tag)
	}
}
