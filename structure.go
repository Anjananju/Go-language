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
    fmt.println("Name is",student1.name);//to access single element
    //to set a new value to age
    student1.age=24;
    //to print new age value only
    fmt.println("age is",student1.age);
    //to print structure with name-anjana,age=24,gender female
    fmt.println(student1);
  var student2 student; //declaring variable and it is initialised all structure elements as 0
   fmt.println(student2);
  }
