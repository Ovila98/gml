# gml - Go Markup Language

[![GoDoc](https://godoc.org/github.com/ovila98/gml?status.svg)](https://godoc.org/github.com/ovila98/gml) [![Go Report Card](https://goreportcard.com/badge/github.com/ovila98/gml)](https://goreportcard.com/report/github.com/ovila98/gml)

`gml` (Go Markup Language) is a Go package that simplifies working with XML data by providing an intuitive and flexible abstraction over XML structures. The package revolves around the `Node` struct, which models an XML element with support for attributes, child elements, and text content. You can easily marshal/unmarshal XML data, navigate or manipulate XML node trees, and build XML queries with ease.

It's a subtle play on words—'gml' sounds like 'XML' but refers to a Go-based approach for working with markup!

## Features

- **Unmarshaling**: Convert XML into an easy-to-use `Node` structure with attributes and children.
- **Marshaling**: Serialize `Node` structures back into XML format.
- **Node Manipulation**: Add, remove, or search child nodes and attributes.
- **Path Creation**: Easily ensure or create paths of nested XML elements.
- **Custom XML Building**: Use `Node` as a base to create complex XML queries programmatically.
- **Method chaining**: Build your XML more easily by using method chaining.

## Installation

To install the package, simply run:

```bash
go get github.com/ovila98/gml
```

Then, import it into your Go code:

```go
import "github.com/ovila98/gml"
```

## Basic Usage

Here is a basic usage example showing how to create, manipulate, and marshal/unmarshal XML data with `gml`.

### Unmarshaling XML

The `gml` package allows you to easily convert an XML string into a `Node` tree structure.

```go
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/ovila98/gml"
)

func main() {
    xmlData := `
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

    // Unmarshal the XML data
    var root gml.Node
    if err := xml.Unmarshal([]byte(xmlData), &root); err != nil {
        panic(err)
    }

    // Print root node and first book title in Fiction section
    fmt.Printf("Library name: %s\n", root.GetAttribute("name"))
    fictionSection := root.FindChild("section")
    firstBook := fictionSection.FindChild("book")
    fmt.Printf("First Fiction Book: %s\n", firstBook.FindChild("title").InnerText)
}
```

### Marshaling XML

You can construct `Node` elements in your Go code and then serialize them back to XML.

```go
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/ovila98/gml"
)

func main() {
	// Create root node
	school := &gml.Node{
		Tag: "school",
		Attributes: map[string]string{
			"name":        "Green Valley High School",
			"location":    "North Side",
			"established": "1992",
		},
	}

	// Create Science classroom and add students
	scienceClass := school.AppendChild(&gml.Node{Tag: "classroom"})
	scienceClass.SetAttribute("name", "Science")

	student1 := scienceClass.AppendChild(&gml.Node{Tag: "student"})
	student1.AppendChild(&gml.Node{Tag: "name", InnerText: "John Doe"})
	student1.AppendChild(&gml.Node{Tag: "age", InnerText: "16"})
	student1.AppendChild(&gml.Node{Tag: "grade", InnerText: "A"})

	student2 := scienceClass.AppendChild(&gml.Node{Tag: "student"})
	student2.AppendChild(&gml.Node{Tag: "name", InnerText: "Jane Smith"})
	student2.AppendChild(&gml.Node{Tag: "age", InnerText: "15"})
	student2.AppendChild(&gml.Node{Tag: "grade", InnerText: "B+"})

	// Create Mathematics classroom and add students
	mathClass := school.AppendChild(&gml.Node{Tag: "classroom"})
	mathClass.SetAttribute("name", "Mathematics")

	student3 := mathClass.AppendChild(&gml.Node{Tag: "student"})
	student3.AppendChild(&gml.Node{Tag: "name", InnerText: "Alice Johnson"})
	student3.AppendChild(&gml.Node{Tag: "age", InnerText: "16"})
	student3.AppendChild(&gml.Node{Tag: "grade", InnerText: "A-"})

	student4 := mathClass.AppendChild(&gml.Node{Tag: "student"})
	student4.AppendChild(&gml.Node{Tag: "name", InnerText: "Bob Brown"})
	student4.AppendChild(&gml.Node{Tag: "age", InnerText: "17"})
	student4.AppendChild(&gml.Node{Tag: "grade", InnerText: "B"})

	// Marshal the XML data
	output, err := xml.MarshalIndent(school, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))
}
```

### Manipulating XML Nodes

The `Node` struct allows you to add, remove, and find child nodes easily.

#### Adding a Child Node

```go
root := &gml.Node{Tag: "root"}
child := &gml.Node{Tag: "child", InnerText: "This is a child node"}
root.AppendChild(child)
```

#### Removing a Child Node

```go
root.RemoveChildrenWithTag("child")
```

#### Finding a Child Node

```go
foundNode := root.FindChild("child")
if foundNode != nil {
    fmt.Println("Found child node with tag:", foundNode.Tag)
}
```

### Attribute Manipulation

Attributes are stored in a map, and `gml` provides methods to easily manage them.

#### Setting an Attribute

```go
node := &gml.Node{Tag: "node"}
node.SetAttribute("key", "value")
```

#### Getting an Attribute

```go
value := node.GetAttribute("key")
fmt.Println("Value of 'key':", value)
```

#### Checking for an Attribute

```go
if node.HasAttribute("key") {
    fmt.Println("'key' attribute exists")
}
```

#### Removing an Attribute

```go
node.RemoveAttribute("key")
```

### Path Handling

`gml` provides path-based methods to ensure or create paths of XML nodes, which is useful for complex XML structures.

#### Creating a Path

```go
root := &gml.Node{Tag: "root"}
deepNode := root.CreatePath("level1", "level2", "level3")
deepNode.InnerText = "Deep Node Content"
```

#### Ensuring a Path Exists

```go
root := &gml.Node{Tag: "root"}
ensuredNode := root.EnsurePath("level1", "level2", "level3")
ensuredNode.InnerText = "This node was created if it didn't exist"
```

### Method Chaining Example

This example demonstrates a concise way to construct an XML document with chained methods:

```go
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/yourusername/gml"
)

func main() {
    // Build the XML structure using method chaining
    school := (&gml.Node{Tag: "school"}).
        SetAttribute("name", "Green Valley High School").
        SetAttribute("location", "North Side").
        AppendChild(&gml.Node{Tag: "classroom"}).
        SetAttribute("name", "Science").
        AppendChild(&gml.Node{Tag: "student"}).
        AppendChild(&gml.Node{Tag: "name", InnerText: "John Doe"}).Parent.
        AppendChild(&gml.Node{Tag: "age", InnerText: "16"}).Parent.Parent.
        AppendChild(&gml.Node{Tag: "student"}).
        AppendChild(&gml.Node{Tag: "name", InnerText: "Jane Smith"}).Parent.
        AppendChild(&gml.Node{Tag: "age", InnerText: "15"})

    // Marshal the XML structure
    output, err := xml.MarshalIndent(school, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(output))
}
```

## API Reference

The complete API reference can be found on [GoDoc](https://godoc.org/github.com/ovila98/gml).

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests. Please follow the [contribution guidelines](CONTRIBUTING.md).

## License

This package is distributed under the MIT License. See the [LICENSE](LICENSE) file for more details.
