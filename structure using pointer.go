package main
import "fmt"
type Employee struct
{
  name string,
  age int
  }
  func main()
  {
  employee1:=&Employee("Anjana",23);
  fmt.println(employee1.age);//to access a single elemnt
  }
