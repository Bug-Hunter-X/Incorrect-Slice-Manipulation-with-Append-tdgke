func main() {
  x := make([]int, 0, 10)
  for i := 0; i < 10; i++ {
    x = append(x, i)
  }
  fmt.Println(x)
  x = append(x[:5], x[6:]...)
  fmt.Println(x)
}