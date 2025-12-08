package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const CONNECTION_COUNT = 1000

type Point struct {
	x, y, z int
	id      int
}

func (p1 Point) Distance(p2 Point) int {
	d := math.Pow(float64(p1.x-p2.x), float64(2)) + math.Pow(float64(p1.y-p2.y), float64(2)) + math.Pow(float64(p1.z-p2.z), float64(2))
	return int(math.Sqrt(d))
}

type Connection struct {
	Points   []int
	Distance int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func makeConnection(p1, p2 Point) Connection {
	return Connection{
		Points:   []int{p1.id, p2.id},
		Distance: p1.Distance(p2),
	}
}

type Network struct {
	containedNodeIds []int
}

// Next two functions are very slow
func (n1 Network) hasOverlap(n2 Network) bool {
	for _, v := range n1.containedNodeIds {
		if slices.Contains(n2.containedNodeIds, v) {
			return true
		}
	}
	return false
}

func (n1 *Network) swallow(n2 Network) {
	n1.add(n2.containedNodeIds)
}

func (n1 *Network) add(points []int) {
	for _, v := range points {
		if !slices.Contains(n1.containedNodeIds, v) {
			n1.containedNodeIds = append(n1.containedNodeIds, v)
		}
	}

}

type PriorityQueue []*Connection

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Distance > pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Connection)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getPoint(l string, id int) Point {
	parts := strings.Split(l, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	z, _ := strconv.Atoi(parts[2])

	return Point{
		x:  x,
		y:  y,
		z:  z,
		id: id,
	}
}

func fillPq(pq PriorityQueue, field []string) PriorityQueue {
	ready := false
	cnt := 0
	for i := range field {
		curr := getPoint(field[i], i)
		for j := i + 1; j < len(field); j++ {
			comp := getPoint(field[j], j)
			c := makeConnection(curr, comp)
			if !ready {
				pq[cnt] = &c
				cnt++
				if cnt == CONNECTION_COUNT {
					ready = true
					heap.Init(&pq)
				}
			} else {
				heap.Push(&pq, &c)
				if pq.Len() > CONNECTION_COUNT {
					heap.Pop(&pq)
				}
			}
		}
	}
	return pq
}

// This uses an ungodly amount of slices.contains
// And most of them are very likely to be solveable without too much hastle
func buildNetworks(pq *PriorityQueue) []*Network {
	overview := make(map[int]*Network)
	for pq.Len() > 0 {
		c := heap.Pop(pq).(*Connection)
		//fmt.Printf("Conn: %v\n", c.Points)
		n1, _ := overview[c.Points[0]]
		n2, _ := overview[c.Points[1]]
		// Both not in network
		if n1 == nil && n2 == nil {
			//fmt.Println("Added network")
			n := Network{
				containedNodeIds: c.Points,
			}
			overview[c.Points[0]] = &n
			overview[c.Points[1]] = &n
		} else if n1 != nil && n2 == nil {
			//fmt.Println("Appended to network A")
			n1.add(c.Points)
			overview[c.Points[1]] = n1
		} else if n1 == nil && n2 != nil {
			//fmt.Println("Appended to network B")
			n2.add(c.Points)
			overview[c.Points[0]] = n2

		} else {
			//fmt.Println("Swallowed network")
			//fmt.Printf("Old: %v + %v\n", n1.containedNodeIds, n2.containedNodeIds)
			n1.swallow(*n2)
			//fmt.Printf("New: %v\n", n1.containedNodeIds)
			for _, v := range n2.containedNodeIds {
				overview[v] = n1
			}
		}

	}
	//fmt.Printf("%v\n", overview)
	res := []*Network{}
	for _, v := range overview {
		if !slices.Contains(res, v) {
			res = append(res, v)
		}
	}
	return res
}

func solve(input []string) []string {
	pq := make(PriorityQueue, CONNECTION_COUNT)
	pq = fillPq(pq, input)
	networks := buildNetworks(&pq)
	lengths := []int{}
	for _, n := range networks {
		//	fmt.Printf("%v\n", n.containedNodeIds)
		lengths = append(lengths, len(n.containedNodeIds))
	}
	slices.Sort(lengths)
	res := lengths[len(lengths)-1] * lengths[len(lengths)-2] * lengths[len(lengths)-3]
	return []string{fmt.Sprintf("%d", res)}
}

func main() {
	content, _ := readLines("./in.txt")
	solution := solve(content)
	for _, line := range solution {
		fmt.Println(line)
	}
}
