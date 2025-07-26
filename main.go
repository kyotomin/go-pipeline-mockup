package main

import (
	"fmt"
	"strconv"
	"strings"
)


type Pipeline[T any] struct {
	data []T
}

// Конструктор пайплайна

func NewPipeline[T any](data []T) Pipeline[T] {
	return Pipeline[T]{
		data: data,
	}
}

func From[T any](items ...T) Pipeline[T] {
	var data []T
	for _, item := range items {
		data = append(data, item)
	}

	return Pipeline[T]{
		data: data,
	}
}

func (p Pipeline[T]) Add(item T) Pipeline[T] {
	p.data = append(p.data, item)

	return p
}

func (p Pipeline[T]) Filter(param func(T) bool) Pipeline[T] {
	var filtered []T
	for _, item := range p.data {
		if param(item) {
			filtered = append(filtered, item)
		}
	}

	return Pipeline[T]{
		data: filtered,
	}
}

func Map[T, U any](p Pipeline[T], transfrom func(T) U) Pipeline[U] {
   var result []U
	for _, item := range p.data {
		result = append(result, transfrom(item))
	}
   return Pipeline[U]{
      data: result,
   }
}  

func Reduce[T, U any](p Pipeline[T], reducer func(T) U) U {
   var result U

   for _, item := range p.data {
      result = reducer(item)
   }

   return result
}

func (p Pipeline[T]) ForEach(action func(T)) {
   for _, item := range p.data {
      action(item)
   }
}

func (p Pipeline[T]) Count(predicate func(T) bool) int {
   count := 0

   for _, item := range p.data {
      if predicate(item) {
         count++
      }
   }

   return count
}

func (p Pipeline[T]) ToSlice() []T {
   return p.data
}

func (p Pipeline[T]) Chunk(size int) [][]T {
   result := make([][]T, 0, len(p.data)/size+1)

   for i := 0; i < len(p.data); i += size {
      end := i + size

      if end > len(p.data) {
         end = len(p.data)
      }
   }

   return result
}

func Unique[T comparable](p Pipeline[T]) Pipeline[T] {
   seen := make(map[T]struct{})
   result := make([]T, 0, len(p.data))

   for _, item := range p.data {
      if _, ok := seen[item]; !ok {
         seen[item] = struct{}{}
         result = append(result, item)
      }
   }

   return Pipeline[T]{data: result}
}


func main() {
   
}
