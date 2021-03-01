package treex_test

import (
	"encoding/json"
	"fmt"
	"krolus/treex"
	"krolus/treex/models"
	"testing"

	. "github.com/franela/goblin"
)

func ResetRoot(root *models.Node) {
	root = models.NewNode("Root", "My root").
		AddNode(
			models.NewNode("Dev", "All about DEV").
				AddNode(models.NewNode("Rust", "Everything Rust")).
				AddNode(models.NewNode("PHP", "PHP is  good").AddLeaf(
					models.NewLeaf("Symfony", "Great resource"),
				)).
				AddNode(models.NewNode("Go", "Go development")).
				AddNode(models.NewNode("Python", "Python must know"))).
		AddNode(models.NewNode("Music", "My music").
			AddNode(
				models.NewNode("Rock", "Rock Music")))

}

// Pretty ...
func Pretty(n interface{}) {
	prettyJSON, err := json.MarshalIndent(n, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", string(prettyJSON))
}

func Test(t *testing.T) {

	root := models.NewNode("", "")
	g := Goblin(t)
	g.Describe("Test Find", func() {

		ResetRoot(root)
		var phpFound *models.Node

		g.It("Should find php node ", func() {
			phpFound = root.FindNode(func(node *models.Node) bool {
				return node.Label == "PHP"
			})
			g.Assert(phpFound).IsNotZero("Node not found")
		})

		g.It("Should find symfony sub ", func() {
			symFound := root.FindLeaf(func(leaf *models.Leaf) bool {
				return leaf.Label == "Symfony"
			})
			g.Assert(symFound).IsNotZero("Leaf not found")

		})
	})

	g.Describe("Test descendents", func() {
		ResetRoot(root)
		g.It("Rust should be descendent of Root ", func() {
			rust := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Rust"
			})
			g.Assert(rust).IsNotNil("Not found")
			flag := root.IsDescendent(rust)
			g.Assert(flag).IsTrue("Not descendent")
		})
		g.It("Rock should not be descendent of Dev ", func() {

			dev := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Dev"
			})
			g.Assert(dev).IsNotNil("Not found")

			rock := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Rock"
			})
			g.Assert(rock).IsNotNil("Not found")

			flag := dev.IsDescendent(rock)
			g.Assert(flag).IsFalse("Is descendent")
		})

	})

	g.Describe("Test move", func() {
		ResetRoot(root)
		rootState := treex.State{Root: root}
		g.It("Move Rust to Rock", func() {
			dev := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Dev"
			})
			g.Assert(dev).IsNotNil("Dev not found")

			rust := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Rust"
			})
			g.Assert(rust).IsNotNil("Rust not found")

			rock := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Rock"
			})
			g.Assert(rock).IsNotNil("Rock not found")

			if err := rootState.MoveNode(rust.ID, rock.ID); err != nil {
				g.Errorf("Error moving: %v", err)
			}
			g.Assert(rock.Nodes.Get(rust.ID)).IsNotNil("Rust not in Rock children")
			g.Assert(dev.Nodes.Get(rust.ID)).IsZero("Rust still in Dev children")
		})

		g.It("Move symfony to Rust", func() {
			php := root.FindNode(func(node *models.Node) bool {
				return node.Label == "PHP"
			})
			g.Assert(php).IsNotNil("PHP not found")

			rust := root.FindNode(func(node *models.Node) bool {
				return node.Label == "Rust"
			})
			g.Assert(rust).IsNotNil("Rust not found")

			symp := root.FindLeaf(func(leaf *models.Leaf) bool {
				return leaf.Label == "Symfony"
			})
			g.Assert(symp).IsNotNil("Symfony not found")

			if err := rootState.MoveLeaf(symp.ID, rust.ID); err != nil {
				g.Errorf("Error moving: %v", err)
			}

			g.Assert(rust.Leaves.Get(symp.ID)).IsNotNil("Symfony not in Rust leaves")
			g.Assert(php.Leaves.Get(symp.ID)).IsZero("Symfony still in PHP leaves")

			//Pretty(root)

		})

	})

	g.Describe("Test persistance", func() {

	})
}
