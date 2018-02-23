package stringutil

import "bytes"

func Duckalize(s string) string {
  var buffer bytes.Buffer

  buffer.WriteString("quac!")

  return buffer.String()
}
