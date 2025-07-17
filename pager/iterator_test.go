package pager_test

import (
	"context"
	"iter"

	"github.com/hypershadow-io/contract/pager"
)

func ExampleNewIterator() {
	var in pager.Pager
	type Model struct{}
	var service interface {
		FindList(c context.Context, page pager.Pager) iter.Seq2[Model, error]
	}
	var result = struct {
		Items   []any
		HasNext bool
	}{
		Items: make([]any, 0, in.GetCount()),
	}

	iterator, err := pager.NewIterator[Model](in)
	if err != nil {
		// error handling
	}
	for itemModel, errModel := range iterator.Scan(service.FindList(context.Background(), iterator)) {
		if errModel != nil {
			// error handling
		}
		result.Items = append(result.Items, itemModel)
	}
	result.HasNext = iterator.HasNext()
}

func ExampleNewIteratorRelaxed() {
	var in pager.Pager
	type Model struct{}
	var service interface {
		FindList(c context.Context, page pager.Pager) iter.Seq2[Model, error]
	}
	var result = struct {
		Items   []any
		HasNext bool
	}{
		Items: make([]any, 0, in.GetCount()),
	}

	iterator := pager.NewIteratorRelaxed[Model](in)
	for itemModel, errModel := range iterator.Scan(service.FindList(context.Background(), iterator)) {
		if errModel != nil {
			// error handling
		}
		result.Items = append(result.Items, itemModel)
	}
	result.HasNext = iterator.HasNext()
}
