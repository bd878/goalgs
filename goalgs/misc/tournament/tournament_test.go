package tournament_test

import (
  "testing"
  "math/rand"

  tree "github.com/bd878/goalgs/ds/tree"
  "github.com/bd878/goalgs/misc/tournament"
)

func TestTournament(t *testing.T) {
  size := 5
  perm := rand.Perm(size)

  root := tournament.Max(perm, 0, len(perm)-1)
  if root.CountTotal() != size {
    t.Error("total != size", root.CountTotal(), size)
  }

  root.Print(tree.PrintInt)
}