func main() {
  x := make([]int, 0, 10)
  for i := 0; i < 10; i++ {
    x = append(x, i)
  }
  fmt.Println(x)
  x = removeIndex(x, 5)
  fmt.Println(x)
}

func removeIndex(s []int, index int) []int {
  return append(s[:index], s[index+1:]...)
} 
