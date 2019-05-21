package main
//dude wtf man
//fml bitches
import "fmt"

type node struct {
	element int
	next    *node
}

var head *node = nil
var tail *node = nil

func ins_node(e int) {
	if head == nil {
		head = new(node)
		head.element = e
		head.next = nil
		tail = head
	} else {
		var newnode *node = new(node)
		newnode.element = e
		newnode.next = nil
		tail.next = newnode
		tail = newnode

	}

}
func disp() {
	var temp *node = head

	if head == nil {
		fmt.Println("The Linked List is empty bitch")
	} else {
		fmt.Println("The elements present int hte linked list are")
		for {
			if temp == nil {
				break
			}
			fmt.Println(temp.element)
			temp = temp.next
		}
	}
}
func del(pos int) {
	var curr int = 2
	var temp *node = head
	if pos == 1 {
		temp = head.next
		head = temp
	} else {
		var tempf *node = temp.next
		for {
			if curr == pos {
				if tempf.next == nil {
					tail = temp
				}
				temp.next = tempf.next
				tempf.next = nil
				break
			}
			temp = temp.next
			tempf = tempf.next
			curr++
		}
	}
}
func main() {
	var ch, ter, val int
	ter = 0
	for {
		fmt.Println("Enter ur choice 1 to insert 2 to display 3 to delete a node 4 to quit")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			fmt.Println("Enter the value to be inserted")
			fmt.Scanln(&val)
			ins_node(val)
		case 2:
			disp()
		case 3:
			fmt.Println("Enter teh position of the node to be deleted")
			fmt.Scanln(&val)
			del(val)
		case 4:
			ter = 1
		}
		if ter == 1 {
			break
		}
	}
}
