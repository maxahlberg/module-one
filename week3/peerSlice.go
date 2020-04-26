package main
import(

   "fmt"
   "sort"
   "strings"
   "strconv"
)
func main(){

sli := make([]int,0, 3)

for {

   var input string

   fmt.Print("Enter the number or x to exit: ")
   fmt.Scan(&input)

   if strings.ToLower(input) == "x" {

    fmt.Print("Bye..!!")
    break
   }
  
  num, err := strconv.Atoi(input)

  if err != nil {
   
   fmt.Print("Bad input")
   break
  }

  sli = append(sli,num)
  sort.Ints(sli)
  fmt.Print(sli)
}
}


   