package main

import (
	"context"
	"sync"
	"time"
)

// Loader is a minimal DataLoader (the idea from github.com/graphql/dataloader):
//   - Load(key) calls that arrive within `wait` of each other are collected
//     into ONE call to fetch(keys)  -> batching
//   - a key loaded once in this request is answered from memory -> per-request cache
//
// Production libraries (Go: graph-gophers/dataloader, Node: dataloader) add
// max batch size, error per key, and cache control. This is the core idea only.
type Loader[K comparable, V any] struct {
	wait  time.Duration
	fetch func(ctx context.Context, keys []K) (map[K]V, error)

	mu    sync.Mutex
	cache map[K]V
	cur   *batch[K, V]
}

type batch[K comparable, V any] struct {
	keys []K
	seen map[K]bool
	done chan struct{}
	res  map[K]V
	err  error
}

func NewLoader[K comparable, V any](wait time.Duration, fetch func(context.Context, []K) (map[K]V, error)) *Loader[K, V] {
	return &Loader[K, V]{wait: wait, fetch: fetch, cache: map[K]V{}}
}

func (l *Loader[K, V]) Load(ctx context.Context, key K) (V, error) {
	l.mu.Lock()
	if v, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return v, nil
	}
	if l.cur == nil {
		b := &batch[K, V]{seen: map[K]bool{}, done: make(chan struct{})}
		l.cur = b
		time.AfterFunc(l.wait, func() {
			l.mu.Lock()
			l.cur = nil // new Load calls start the next batch
			l.mu.Unlock()

			b.res, b.err = l.fetch(ctx, b.keys) // ONE query for all collected keys

			if b.err == nil {
				l.mu.Lock()
				for k, v := range b.res {
					l.cache[k] = v
				}
				l.mu.Unlock()
			}
			close(b.done)
		})
	}
	b := l.cur
	if !b.seen[key] {
		b.seen[key] = true
		b.keys = append(b.keys, key)
	}
	l.mu.Unlock()

	<-b.done
	var zero V
	if b.err != nil {
		return zero, b.err
	}
	return b.res[key], nil
}
