package main

import "fmt"

type LinkNode struct {
	key   int
	value int
	prev  *LinkNode
	next  *LinkNode
}

type LRUCache struct {
	hash     map[int]*LinkNode
	capacity int
	head     *LinkNode
	tail     *LinkNode
	count    int
}

func Constructor(capacity int) LRUCache {
	vhead := &LinkNode{}
	return LRUCache{
		hash:     make(map[int]*LinkNode),
		capacity: capacity,
		head:     vhead,
		tail:     vhead,
		count:    0,
	}
}

func (this *LRUCache) Get(key int) int {
	n, has := this.hash[key]
	if has {
		this.moveLast(n)
		return (*n).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, has := this.hash[key]; has {
		this.moveLast(n)
		n.value = value
		return
	}

	if this.count < this.capacity {
		n := &LinkNode{key: key, value: value}
		this.append(n)
		this.hash[key] = n
		this.count++
		return
	}

	f := this.pop()
	if f != nil {
		delete(this.hash, f.key)
		this.count--
	}

	if this.count < this.capacity {
		n := &LinkNode{key: key, value: value}
		this.append(n)
		this.hash[key] = n
		this.count++
	}
}

func (this *LRUCache) moveLast(n *LinkNode) {
	if n == this.tail {
		return
	}

	n.prev.next = n.next
	n.next.prev = n.prev
	this.tail.next = n
	n.prev = this.tail
	n.next = nil
	this.tail = n
}

func (this *LRUCache) append(n *LinkNode) {
	this.tail.next = n
	n.prev = this.tail
	n.next = nil
	this.tail = n
}

func (this *LRUCache) pop() *LinkNode {
	first := this.head.next
	if first != nil {
		second := first.next
		this.head.next = second
		if second != nil {
			second.prev = this.head
		}
	}
	return first
}

func (this *LRUCache) Print() {
	fmt.Print("<", this.capacity, this.count, "> ")
	for p := this.head; p != nil; p = p.next {
		fmt.Print(p.key, ":", p.value, "->")
	}
	fmt.Print("\n")
}

func (this *LRUCache) PrintRev() {
	fmt.Print("<", this.capacity, this.count, "> ")
	for p := this.tail; p != nil; p = p.prev {
		fmt.Print(p.key, ":", p.value, "->")
	}
	fmt.Print("\n")
}

func main() {
	// cache := Constructor(2)
	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// cache.Print()
	// fmt.Println(cache.Get(1)) // returns 1
	// cache.Print()
	// cache.Put(3, 3) // evicts key 2
	// cache.Print()
	// fmt.Println(cache.Get(2)) // returns -1 (not found)
	// cache.Print()
	// cache.Put(4, 4) // evicts key 1
	// cache.Print()
	// fmt.Println(cache.Get(1)) // returns -1 (not found)
	// fmt.Println(cache.Get(3)) // returns 3
	// fmt.Println(cache.Get(4)) // returns 4

	// cache := Constructor(1)
	// cache.Put(2, 1)
	// fmt.Println(cache.Get(2)) // 1
	// cache.Put(3, 2)
	// fmt.Println(cache.Get(2)) // -1
	// fmt.Println(cache.Get(3)) // 2

	// ["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	// [[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	// [null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]

	// cache := Constructor(3)
	// cache.Put(1, 1)
	// cache.Put(2, 2)
	// cache.Put(3, 3)
	// cache.Put(4, 4)
	// cache.Print()
	// fmt.Println(cache.Get(4)) // 4
	// cache.Print()
	// fmt.Println(cache.Get(3)) // 3
	// cache.Print()
	// fmt.Println(cache.Get(2)) // 2
	// cache.Print()
	// fmt.Println(cache.Get(1)) // -1
	// cache.Print()
	// cache.Put(5, 5)
	// cache.Print()
	// fmt.Println(cache.Get(1)) // -1
	// fmt.Println(cache.Get(2)) // 2
	// fmt.Println(cache.Get(3)) // 3
	// fmt.Println(cache.Get(4)) // -1
	// fmt.Println(cache.Get(5)) // 5

	// ["LRUCache","put","put","put","put","put","get","put","get","get","put","get","put","put","put","get","put","get","get","get","get","put","put","get","get","get","put","put","get","put","get","put","get","get","get","put","put","put","get","put","get","get","put","put","get","put","put","put","put","get","put","put","get","put","put","get","put","put","put","put","put","get","put","put","get","put","get","get","get","put","get","get","put","put","put","put","get","put","put","put","put","get","get","get","put","put","put","get","put","put","put","get","put","put","put","get","get","get","put","put","put","put","get","put","put","put","put","put","put","put"]
	// [[10],[10,13],[3,17],[6,11],[10,5],[9,10],[13],[2,19],[2],[3],[5,25],[8],[9,22],[5,5],[1,30],[11],[9,12],[7],[5],[8],[9],[4,30],[9,3],[9],[10],[10],[6,14],[3,1],[3],[10,11],[8],[2,14],[1],[5],[4],[11,4],[12,24],[5,18],[13],[7,23],[8],[12],[3,27],[2,12],[5],[2,9],[13,4],[8,18],[1,7],[6],[9,29],[8,21],[5],[6,30],[1,12],[10],[4,15],[7,22],[11,26],[8,17],[9,29],[5],[3,4],[11,30],[12],[4,29],[3],[9],[6],[3,4],[1],[10],[3,29],[10,28],[1,20],[11,13],[3],[3,12],[3,8],[10,9],[3,26],[8],[7],[5],[13,17],[2,27],[11,15],[12],[9,19],[2,15],[3,16],[1],[12,17],[9,1],[6,19],[4],[5],[5],[8,1],[11,7],[5,2],[9,28],[1],[2,2],[7,4],[4,22],[7,24],[9,26],[13,28],[11,26]]
	cache := Constructor(10)
	cache.Put(10, 13)
	cache.Put(3, 17)
	cache.Put(6, 11)
	cache.Put(10, 5)
	cache.Put(9, 10)
	// cache.Print()
	// cache.PrintRev()
	fmt.Println(cache.Get(13)) //-1
	// cache.Print()
	// cache.PrintRev()
	cache.Put(2, 19)
	// cache.Print()
	// cache.PrintRev()
	fmt.Println(cache.Get(2)) //19
	// cache.Print()
	// cache.PrintRev()
	fmt.Println(cache.Get(3)) //17
	// cache.Print()
	// cache.PrintRev()
	cache.Put(5, 25)
	// cache.Print()
	fmt.Println(cache.Get(8)) //-1
	// cache.Print()
	cache.Put(9, 22)
	// cache.Print()
	cache.Put(5, 5)
	// cache.Print()
	cache.Put(1, 30)
	// cache.Print()
	fmt.Println(cache.Get(11)) //-1
	cache.Put(9, 12)
	fmt.Println(cache.Get(7)) //-1
	fmt.Println(cache.Get(5)) //5
	// cache.Print()
	fmt.Println(cache.Get(8)) //-1
	fmt.Println(cache.Get(9)) //12
	// cache.Print()
	cache.Put(4, 30)
	// cache.Print()
	cache.Put(9, 3)
	// cache.Print()
	fmt.Println(cache.Get(9)) //3
	// cache.Print()
	fmt.Println(cache.Get(10)) //5
	// cache.Print()
	fmt.Println(cache.Get(10)) //5
	// cache.Print()
	cache.Put(6, 14)
	// cache.Print()
	cache.Put(3, 1)
	// cache.Print()
	fmt.Println(cache.Get(3)) //1
	// cache.Print()
	cache.Put(10, 11)
	// cache.Print()
	fmt.Println(cache.Get(8)) //-1
	cache.Put(2, 14)
	// cache.Print()
	fmt.Println(cache.Get(1)) //30
	// cache.Print()
	fmt.Println(cache.Get(5)) //5
	// cache.Print()
	fmt.Println(cache.Get(4)) //30
	// cache.Print()
	cache.Put(11, 4)
	// cache.Print()
	cache.Put(12, 24)
	// cache.Print()
	cache.Put(5, 18)
	// cache.Print()
	fmt.Println(cache.Get(13)) //-1
	cache.Put(7, 23)
	// cache.Print()
	fmt.Println(cache.Get(8))  //-1
	fmt.Println(cache.Get(12)) //24
	// cache.Print()
	cache.Put(3, 27)
	// cache.Print()
	cache.Put(2, 12)
	// cache.Print()
	fmt.Println(cache.Get(5)) //18
	// cache.Print()
	cache.Put(2, 9)
	// cache.Print()
	cache.Put(13, 4)
	// cache.Print()
	cache.Put(8, 18)
	// cache.Print()
	cache.Put(1, 7)
	// cache.Print()
	fmt.Println(cache.Get(6)) //-1
	cache.Put(9, 29)
	// cache.Print()
	cache.Put(8, 21)
	// cache.Print()
	fmt.Println(cache.Get(5)) //18
	// cache.Print()
	cache.Put(6, 30)
	// cache.Print()
	cache.Put(1, 12)
	// cache.Print()
	fmt.Println(cache.Get(10)) //-1
	cache.Put(4, 15)
	// cache.Print()
	cache.Put(7, 22)
	// cache.Print()
	cache.Put(11, 26)
	// cache.Print()
	cache.Put(8, 17)
	// cache.Print()
	cache.Put(9, 29)
	// cache.Print()
	fmt.Println(cache.Get(5)) //18
	// cache.Print()
	cache.Put(3, 4)
	// cache.Print()
	cache.Put(11, 30)
	// cache.Print()
	fmt.Println(cache.Get(12)) //-1
}
