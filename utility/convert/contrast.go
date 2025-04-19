package convert

import "github.com/gogf/gf/v2/container/gset"

func Contrast(source, target []int) (added, removed []int) {
	sourceSet := gset.NewIntSetFrom(source)
	targetSet := gset.NewIntSetFrom(target)

	added = targetSet.Diff(sourceSet).Slice()
	removed = sourceSet.Diff(targetSet).Slice()

	return
}
