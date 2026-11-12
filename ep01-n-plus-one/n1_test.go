package main

// Step 10: a test that FAILS if someone brings the N+1 back.
// It needs the database running (docker compose up -d); otherwise it is skipped.
//
//   go test -run TestQueryBudget -v

import (
	"context"
	"testing"
)

func TestQueryBudget(t *testing.T) {
	ctx := context.Background()
	db, err := newPool(ctx)
	if err == nil {
		err = db.Ping(ctx)
	}
	if err != nil {
		t.Skipf("database not reachable: %v", err)
	}
	defer db.Close()

	cases := []struct {
		name   string
		fn     strategy
		budget int64 // max queries allowed, whatever the page size
	}{
		{"batch", batchIn, 2},
		{"join", joinAll, 1},
		{"json", jsonAgg, 1},
		{"nested-loader", nestedLoader, 4},
		// Uncomment to see the test catch an N+1:
		// {"naive", naive, 2},
	}

	for _, c := range cases {
		for _, limit := range []int{10, 100} { // the count must NOT grow with the page size
			st := &stats{}
			qctx := context.WithValue(ctx, ctxKey{}, st)
			if _, err := c.fn(qctx, db, limit); err != nil {
				t.Fatalf("%s: %v", c.name, err)
			}
			if got := st.queries.Load(); got > c.budget {
				t.Errorf("%s limit=%d: %d queries, budget is %d (N+1?)", c.name, limit, got, c.budget)
			}
		}
	}
}
