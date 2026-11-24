# Minimum Number of Coins (Greedy)

> Striver SDE Sheet 047 · Day 8 Greedy — repeatedly take the largest Indian denomination that fits; optimal only because the system is canonical. Go, stdlib only, tested.

## Problem

Make a given amount with the fewest notes/coins of Indian denominations (₹1, 2, 5, 10, 20, 50, 100, 200, 500, 2000), returning the actual denominations used.

## Approach

Walk the denominations largest-first, taking each as many times as it fits. The honest caveat this problem exists to teach: greedy coin change is correct **only for canonical coin systems** like this one. With denominations {1, 3, 4} and amount 6, greedy gives 4+1+1 (three coins) while 3+3 (two) is optimal — that counterexample is why the general problem (Day 26, Coin Change) is dynamic programming. Greedy here is a property of the currency, not of coin change.

**Complexity:** O(amount/largest + k) effectively; output-sized space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
