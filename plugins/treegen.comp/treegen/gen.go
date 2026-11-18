package treegen

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/kkgams/sdk/go/tree"
)

type Random interface {
	Intn(n int) int
	Float64() float64
}

type GenerateTreeConfig struct {
	NodeCount    int `json:"nodeCount"`    // target node count (0 = unlimited)
	MaxDepth     int `json:"maxDepth"`     // maximum tree depth (0 = unlimited)
	MaxBranching int `json:"maxBranching"` // maximum children per node (0 = unlimited)
	RootBranches int `json:"rootBranches"` // specific number of root children (0 = auto)
}

func GenerateTree(rng Random, cfg *GenerateTreeConfig) (tree.Tree, error) {

	if cfg.NodeCount == 0 && cfg.MaxDepth == 0 {
		return nil, fmt.Errorf("MaxDepth or NodeCount should be not zero ")
	}

	type Victim struct {
		Index      int
		ChildCount int
		Depth      int
	}

	t := tree.Tree{}
	t.Add(-1, nil)

	canHaveChildren := make([]*Victim, 0)
	if cfg.RootBranches < 1 {
		canHaveChildren = append(canHaveChildren, &Victim{
			Index:      0,
			ChildCount: 0,
			Depth:      1,
		})
	}

	if len(t) >= cfg.NodeCount {
		return t, nil
	}

	for range cfg.RootBranches {
		if len(t) >= cfg.NodeCount {
			return t, nil
		}

		childId := len(t) // The current tree length is the index of the new node.
		t.Add(0, nil)

		depth := calculateNodeDepth(&t, childId)
		if cfg.MaxDepth > 0 && depth >= cfg.MaxDepth {
			continue
		}

		canHaveChildren = append(canHaveChildren, &Victim{
			Index:      childId,
			ChildCount: 0,
			Depth:      depth,
		})
	}

	if len(canHaveChildren) < 1 {
		return t, nil
	}

	for range cfg.NodeCount - len(t) {
		if len(canHaveChildren) < 1 {
			break
		}

		parentIndex := 0
		if len(canHaveChildren) > 1 {
			parentIndex = rng.Intn(len(canHaveChildren))
		}

		childId := len(t) // The current tree length is the index of the new node.
		t.Add(canHaveChildren[parentIndex].Index, nil)
		canHaveChildren[parentIndex].ChildCount += 1
		if cfg.MaxBranching > 0 && canHaveChildren[parentIndex].ChildCount >= cfg.MaxBranching {
			canHaveChildren = slices.Delete(canHaveChildren, parentIndex, parentIndex+1)
		}

		depth := calculateNodeDepth(&t, childId)
		if cfg.MaxDepth > 0 && depth >= cfg.MaxDepth {
			continue
		}

		canHaveChildren = append(canHaveChildren, &Victim{
			Index:      childId,
			ChildCount: 0,
			Depth:      depth,
		})
	}

	return t, nil
}

// Calculate depth of a node by traversing up to root
func calculateNodeDepth(nodes *tree.Tree, nodeIdx int) int {
	depth := 1
	current := (*nodes)[nodeIdx]
	for current.ParentId >= 0 {
		current = (*nodes)[current.ParentId]
		depth++
	}
	return depth
}

func PrettyJson(input any) string {
	return string(Must(json.MarshalIndent(input, "", "  ")))
}

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}
