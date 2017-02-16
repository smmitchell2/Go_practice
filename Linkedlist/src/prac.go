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

func newList() List{
  var list List
  list.head = nil
  list.tail = nil
  return list
}

func printList(list List){
  var temp *Node
  temp = list.head
  for(temp != nil){
    fmt.Printf("%d \n",temp.data)
    temp = temp.next
  }
}

func insert(list List,data int) List{
  var a Node = newNode(data)
  if(list.head == nil){
    list.head = &a
  } else if(list.head.next == nil){
    list.head.next = &a
    list.tail = &a
  } else {
    var temp *Node = list.tail
    temp.next = &a
    list.tail = &a
  }
  return list
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

  var list List = newList()
  var i int = 0
  for(i<100){
    list = insert(list,i)
    i = i + 1
  }

  printList(list)
}
