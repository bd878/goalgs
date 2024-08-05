package skiplist

import (
  "errors"
  "math/rand"
)

var ErrNotFound = errors.New("item not found")

const RAND_MAX int = 10e3

var lgNmax int = 10 /* max levels amount */

var nullIntItem IntItem

type Skiplist struct {
  Head *SkiplistNode  /* head */
  LgN   int           /* how many levels */
}

type SkiplistNode struct {
  Item  Item
  Level int
  Next  []*SkiplistNode /* slice of level links */
}

/**
 * We start from mock header, because
 * we need to remove next nodes somehow
 */
func NewSkiplist() *Skiplist {
  return &Skiplist{
    Head: NewSkiplistNode(&nullIntItem, lgNmax),
    LgN: 0,
  }
}

func NewSkiplistNode(v Item, level int) *SkiplistNode {
  return &SkiplistNode{
    Item: v,
    Level: level,
    Next: make([](*SkiplistNode), level),
  }
}

func (s *Skiplist) Insert(v Item) {
  nextLevel := randLevel()
  if nextLevel > s.LgN {
    s.LgN = nextLevel
  }
  insertR(s.Head, NewSkiplistNode(v, nextLevel), s.LgN)
}

func (s *Skiplist) Search(k int) (Item, error) {
  return searchR(s.Head, k, s.LgN)
}

func (s *Skiplist) Remove(k int) {
  removeR(s.Head, k, s.LgN)
}

/**
 * Generate new level randomly.
 * Generate new j-linked node with
 * probability of 1/2^j
 */
func randLevel() int {
  i, j := 1, 2
  t := rand.Intn(RAND_MAX)
  for i < (lgNmax-1) && t > int(RAND_MAX/j) {
    i += 1
    j = j*2
  }
  return i
}

/**
 * Connects new node
 * on each level
 */
func insertR(t, x *SkiplistNode, level int) {
  key := x.Item.Key()
  tk := t.Next[level] /* next node on level k */
  /* we might have non-nil links on lower levels */
  if tk == nil || key < tk.Item.Key() {
    if level < x.Level { /* below the level */
      x.Next[level] = tk // next
      t.Next[level] = x // previous
    }
    if level == 0 {
      return
    }
    insertR(t, x, level-1) // link next levels below
    return
  }
  insertR(tk, x, level) /* insert next on list */
}

func searchR(t *SkiplistNode, key, level int) (Item, error) {
  if t == nil {
    return nil, ErrNotFound
  }
  if key == t.Item.Key() {
    return t.Item, nil
  }
  x := t.Next[level]
  if x == nil || key < x.Item.Key() {
    if level == 0 {
      return nil, ErrNotFound
    }
    return searchR(t, key, level-1)
  }
  return searchR(x, key, level)
}

/**
 * Unlink given node on each level
 */
func removeR(t *SkiplistNode, key, level int) {
  x := t.Next[level]
  if x == nil {
    removeR(t, key, level-1)
    return
  }

  if key <= x.Item.Key() { /* move one level below on same node */
    if key == x.Item.Key() {
      t.Next[level] = x.Next[level] // relink
    }
    if level == 0 { // on the lowest level
      return
    }
    removeR(t, key, level-1)
    return
  }

  removeR(x, key, level) /* move to next node on same level */
}