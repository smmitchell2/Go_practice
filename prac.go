package main
import "fmt"

type Person struct{
    name string
    age int
}

type Node struct{
  data int
  next *Node
}

type List struct{
  head *Node
  tail *Node
}

func newNode(data int) Node{
  var n Node
  n.data = data
  n.next = nil
  return n
}

func printList(list List){
  var temp *Node
  temp = list.head
  for(temp != nil){
    fmt.Printf("%d \n",temp.data)
    temp = temp.next
  }
}

func insert(list List,data int){
  if(list.head == nil){
    var a Node = newNode(data)
    list.head = &a
  }else if(list.head.next == nil){
    var b Node = newNode(data)
    list.head.next = &b
    list.tail = &b
  }else{
    var c Node = newNode(data)
    temp *Node = list.tail
    temp.next = &c
    list.tail = &c
  }

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
  //var a Node = newNode(10)
//  var b Node = newNode(20)
//var c Node = newNode(30)
  //a.next = &b
  //b.next = &c
  var list List
  insert(list,10)
  //list.head = &a
  //list.tail = &c
  printList(list)
}
