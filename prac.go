package main
import "fmt"

type Person struct{
    name string
    age int
}

func main()  {
  var person1 Person
  var person2 Person

  person1.name = "Tony Stark"
  person1.age = 45
  person2.name = "Ned Flanders"
  person2.age = 39
  fmt.Printf("Person 1 name: %s age: %d\n",person1.name,person1.age)
  fmt.Printf("Person 2 name: %s age: %d\n",person2.name,person2.age)
}