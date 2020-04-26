package main
import (
   "encoding/json"
  "fmt"
  "os"
)
func main(){

var s1,s2 string
fmt.Print(" Enter the name and address: ")
fmt.Scan(&s1, &s2)
idMap:= make(map[string] string)
idMap["name"]=s1
idMap["adress"]=s2
barr,_:= json.Marshal(idMap)
os.Stdout.Write(barr)
}