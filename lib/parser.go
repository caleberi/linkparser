package lib

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href='...'>...</a>) node
type Link struct {
	Href string
	Text string
}

// Check exits on error but provides an error message
func Check(msg string, err error) {
	if err != nil {
		Exit(msg)
	}
}

// Exit causes the program to quit
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	for currentNode := node.FirstChild; currentNode != nil; currentNode = currentNode.NextSibling {
		ret = append(ret, linkNodes(currentNode)...)
	}
	return ret
}

// Parse will parse an html reader interface support source
// and returns a []Link struct or error
func Parse(r io.Reader) ([]Link, error) {
	document, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(document)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}
func extractTextFromNode(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var ret string
	for currentNode := node.FirstChild; currentNode != nil; currentNode = currentNode.NextSibling {
		ret += extractTextFromNode(currentNode) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}
func buildLink(node *html.Node) Link {
	var ret Link
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = extractTextFromNode(node)
	return ret
}
