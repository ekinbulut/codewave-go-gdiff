package internal

import (
	"github.com/sergi/go-diff/diffmatchpatch"
)

type DiffChecker struct {
	diffmatchpatch *diffmatchpatch.DiffMatchPatch
	diffs          []diffmatchpatch.Diff
}

func NewDiffChecker() *DiffChecker {
	return &DiffChecker{
		diffmatchpatch: diffmatchpatch.New(),
	}
}

func (d *DiffChecker) Check(old string, new string) (bool, error) {
	d.diffs = d.diffmatchpatch.DiffMain(old, new, true)
	return len(d.diffs) == 0, nil
}

func (d *DiffChecker) Diffs() []diffmatchpatch.Diff {
	return d.diffs
}

// print diffs to html
func (d *DiffChecker) PrintDiffsToHtml(old string, new string) (string, error) {
	return d.diffmatchpatch.DiffPrettyHtml(d.diffs), nil
}
