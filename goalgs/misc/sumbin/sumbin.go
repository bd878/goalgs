package sumbin

func test(c, b, a bool) int {
  var res int
  if c { res += 4 }
  if b { res += 2 }
  if a { res += 1 }
  return res
}

func Sum(num1, num2 string) string {
  var res string

  num1 = reverse(num1)
  num2 = reverse(num2)

  var shift bool
  var a, b byte
  for i := 0; i < max(len(num1), len(num2)); i++ {
    if i > len(num1)-1 {
      a = '0'
    } else {
      a = num1[i]
    }

    if i > len(num2)-1 {
      b = '0'
    } else {
      b = num2[i]
    }

    t := test(shift, a != '0', b != '0')
    switch t {
    case 0: // 000
      res += "0"
    case 1, 2, 4: // 100, 001, 010
      res += "1"
    case 3, 5, 6: // 011, 101, 110
      res += "0"
      shift = true
    case 7: // 111
      res += "0"
      res += "1"
      shift = true
    }
  }

  if shift {
    res += "1"
  }

  return reverse(res)
}

func reverse(str string) string {
  var res string
  for i := len(str)-1; i >= 0; i-- {
    res += string(str[i])
  }
  return res
}
