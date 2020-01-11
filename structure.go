package main
import "fmt"
type student struct
{
  name string,
  age int,
  gender string 
 }
  func main()
  {
  student1:=student("Anjana",23,"Female");
  fmt.println(student1);
  var student2 student; //declaring variable and it is initialised all structure elements as 0
   fmt.println(student2);
  }
