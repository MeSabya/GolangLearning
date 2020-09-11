//Queue implementation in Golang 
/*
A Queue in golang can be implemented using two options:
1. using container/list
2. using slice
https://flaviocopes.com/golang-data-structure-queue/
*/

package main
import ( "fmt"
        "container/list"
    )

type customQueue struct {
    Queue *list.List
}

func (q *customQueue) EnQueue(s string){
    q.Queue.PushBack(s)
} 

func (q *customQueue) DeQueue() error {
    if q.Queue.Len() > 0 {
        ele := q.Queue.Front()
        q.Queue.Remove(ele)
        return fmt.Errorf("Success")
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}
func main() {
    queue := &customQueue{
        Queue : list.New(),
    }
    queue.EnQueue("A")
    fmt.Println("Dequeuing ", queue.DeQueue())
}