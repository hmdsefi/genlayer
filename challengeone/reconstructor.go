package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrDataIntegrityFailed = errors.New("error: Data integrity verification failed")
)

type Fragments map[int]*DataHash

type DataHash struct {
	Data string
	Hash string
}

type hasher interface {
	SimpleHash(input string) string
}

type DataReconstructor struct {
	hash hasher
}

func NewDataReconstructor(hasher hasher) *DataReconstructor {
	return &DataReconstructor{hash: hasher}
}

func (r *DataReconstructor) ReconstructData(fragments Fragments) (string, error) {
	queue := InsertMapIntoHeap(fragments)
	var builder strings.Builder
	for range queue.Len() {
		fragment, ok := queue.Pop().(*Item)
		if !ok {
			return "", errors.New("data reconstructor: invalid type")
		}

		normalizedData := r.NormalizeData(fragment.key, fragment.value.Data)
		hash := r.hash.SimpleHash(normalizedData)
		if hash != fragment.value.Hash {
			return "", ErrDataIntegrityFailed
		}

		builder.WriteString(fragment.value.Data)
		builder.WriteString(" ")
	}

	return builder.String(), nil
}

func (r *DataReconstructor) NormalizeData(index int, data string) string {
	return fmt.Sprintf("%d:%s", index, data)
}
