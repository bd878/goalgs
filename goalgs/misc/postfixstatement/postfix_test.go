package postfixstatement_test

import (
  "testing"

  ps "github.com/bd878/goalgs/misc/postfixstatement"
)

func TestPostfixEvaluate(t *testing.T) {
  expression := "598+46**7+*"
  result := ps.Evaluate(expression)
  t.Log(result)
  if result != 2075 {
    t.Error("result != 2075", result)
  }
}

func TestTranslateToPostfix(t *testing.T) {
  expression := "(5*(((9+8)*(4*6))+7))"
  postfix := ps.TranslateInfixToPostfix(expression)
  result := ps.Evaluate(postfix)

  if result != 2075 {
    t.Error("postfix form is wrong")
  }
}