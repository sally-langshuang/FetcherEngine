package tree

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Graph struct {
	graph []string
}
func (g Graph) RPush(lines...string) Graph {
	g.graph = append(g.graph, lines...)
	return g
}

func (g Graph) LPush(lines...string) Graph{
	g.graph = append(lines, g.graph...)
	return g
}
func (g Graph) Width() int {
	if g.graph == nil || len(g.graph) == 0 {
		return 0
	}
	return len(g.graph[0])
}

func (g Graph) Height() int {
	if g.graph == nil{
		return 0
	}
	return len(g.graph)
}

func (g Graph)String() string {
	var b strings.Builder
	fmt.Fprint(&b, fmt.Sprintf("+%s+\n",strings.Repeat("-", g.Width())))
	for _, l := range g.graph {
		fmt.Fprintf(&b, fmt.Sprintf("|%s|\n", l))
	}
	fmt.Fprint(&b, fmt.Sprintf("+%s+\n",strings.Repeat("-", g.Width())))
	return b.String()
}

func getWidth(graph []string) int {
	if graph == nil || len(graph) == 0 {
		return 0
	}
	return len(graph[0])
}
func layElemS(i, sw, ew int, e string)  string{
	if i + ew > sw || len(e) > ew{
		panic("err")
	}
	return strings.Repeat(" ", i) + fmt.Sprintf("%-[2]*[1]s",e, ew) + strings.Repeat(" ", sw - i - ew)
}
func layElem(i, sw, ew, e int)  string{
		if i + ew > sw || len(strconv.Itoa(e)) > ew{
			panic("err")
		}
		return strings.Repeat(" ", i) + fmt.Sprintf("%-[2]*[1]d",e, ew) + strings.Repeat(" ", sw - i - ew)
}

func getOneValString(t *TreeNode, elemWidth int) (offset int, graph Graph) {
	line:= layElem(0, elemWidth, elemWidth,  t.Val)
	o := (elemWidth - 1) / 2
	return o, Graph{[]string{line}}
}

func expandGraph(offset int, graph Graph, left, right, up, down int) (o int, g Graph) {
	if left < 0 || right < 0 || up < 0 || down < 0 || graph.graph == nil ||  len(graph.graph) == 0 {
		panic("err")
	}
	o = offset + left
	g = Graph{graph: make([]string, len(graph.graph) + up + down)}
	blandLine := strings.Repeat(" ", left + graph.Width() + right)
	for i := 0; i < up; i++ {
		g.graph[i] = blandLine
	}
	for i, _ := range graph.graph {
		g.graph[up + i] = strings.Repeat(" ", left) + graph.graph[i] + strings.Repeat(" ", right)
	}
	for i := up + graph.Height(); i < up + graph.Height() + down ; i++ {
		g.graph[i] = blandLine
	}
	return o, g
}

func mergeTwoGraph(lo int, lg Graph, ro int, rg Graph, gap int) (lOffset, rOffset int, graph Graph) {
	if lg.Height() == 0 || rg.Height() == 0 || gap < 0{
		panic("err")
	}
	rOffset = ro + lg.Width() + gap
	lOffset = lo
	graph = Graph{graph: make([]string, 0)}
	if lg.Height() < rg.Height() {
		lo, lg = expandGraph(lo, lg, 0, 0, 0, rg.Height() - lg.Height())
	} else if lg.Height() > rg.Height() {
		ro, rg = expandGraph(ro, rg, 0, 0, 0, lg.Height() - rg.Height())
	}
	for i:=0; i < lg.Height(); i++{
		graph.graph = append(graph.graph, lg.graph[i] + strings.Repeat(" ", gap) + rg.graph[i])
	}
	return
}
func genLinkLine(lo, to, ro, w int) string {
	var x string
	if ro == -1 {
		x = "┌" + strings.Repeat("─", to - lo - 1) + "┘"
	} else if lo == -1 {
		x = "└" + strings.Repeat("─", ro - to - 1) + "┐"
	} else {
		x = "┌" + strings.Repeat("─", to - lo - 1) + "┴" + strings.Repeat("─", ro - to - 1) + "┐"
	}
	return layElemS(lo, w, len(x), x)
}

func getGraphLink(t *TreeNode, elemWidth int) (offset int, graph Graph) {
	if t == nil || elemWidth < 1 {
		return
	}
	to, tg := getOneValString(t, elemWidth)
	if t.Left == nil && t.Right == nil {
		return to, tg
	}
	lo, lg := getGraphLink(t.Left, elemWidth)
	ro, rg := getGraphLink(t.Right, elemWidth)
	if t.Right == nil {
		rootLine := layElem(lo + 2, lg.Width() + 2, elemWidth, t.Val)
		linkline := genLinkLine(lo, lo + 2, -1, lg.Width() + 2)
		_, lg = expandGraph(lo, lg, 0, 2, 0, 0)
		g := Graph{graph: make([]string, 0)}
		g.RPush(rootLine, linkline)
		g.RPush(lg.graph...)
		return lo + 2, lg
	}
	if t.Left == nil {
		rootLine := layElem(0, rg.Width() + 2, elemWidth, t.Val)
		linkline := genLinkLine(-1, 0, ro + 2, rg.Width() + 2)
		_, rg = expandGraph(ro, rg, 2, 0, 0, 0)
		g := Graph{graph: make([]string, 0)}
		g.RPush(rootLine, linkline)
		g.RPush(rg.graph...)
		return 0, rg
	}
	lo, ro, gc := mergeTwoGraph(lo, lg, ro, rg, 1)
	to = (lo + ro) / 2
	rootLine := layElem(to, gc.Width(), elemWidth, t.Val)
	linkLine := genLinkLine(lo, to, ro, gc.Width())
	g := Graph{graph: make([]string, 0)}
	g.RPush(rootLine, linkLine)
	g.RPush(gc.graph...)
	return to, g
}

func mergeGraph(leftGraph, rightGraph[]string)  {

}

func (t *TreeNode) LevelsNode() [][]*TreeNode {
	if t == nil {return nil}
	res := [][]*TreeNode{}
	currentLevel := []*TreeNode{t}
	nextLevel := make([]*TreeNode, 0)
	for ;len(currentLevel) > 0; {
		for _, n := range currentLevel {
			if n != nil {
				nextLevel = append(nextLevel, n.Left, n.Right)
			}
		}
		if newLevel := RightStripNil(currentLevel); len(newLevel) > 0 {
			res = append(res, newLevel)
		}
		currentLevel, nextLevel = nextLevel, make([]*TreeNode, 0)
	}
	return res
}

func RightStripNil(arr []*TreeNode) []*TreeNode {
	if arr == nil || len(arr) == 0 {
		return make([]*TreeNode, 0)
	}
	endIdx := len(arr) - 1
	for ; endIdx >=0; endIdx-- {
		if arr[endIdx] != nil {
			break
		}
	}
	return arr[:endIdx + 1]
}

func (t TreeNode) MaxVal() int {
	return max(t.Val, t.Left.MaxVal(), t.Right.MaxVal() )
}

func MaxDepth(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return max(MaxDepth(t.Left), MaxDepth(t.Right)) + 1
}

func maxInt(num int, nums...int) int {
	res := num
	for _, i:= range nums{
		if i > res {
			res = i
		}
	}
	return res
}

//
//func (TreeNode) Init(s string) *TreeNode {
//
//}
type NumArr struct {
	arr []*int
}

func (a NumArr)String() string {
	var b strings.Builder
	fmt.Fprint(&b, "[")
	for idx, i:= range a.arr {
		if i == nil {
			fmt.Fprint(&b, "null" )
		} else {
			fmt.Fprintf(&b, "%d",*i)
		}
		if idx != len(a.arr) - 1 {
			fmt.Fprint(&b, ",")
		}

	}
	fmt.Fprint(&b, "]")
	return b.String()
}

func exactElem(s string) []string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	return strings.Split(s, ",")
}

func InitTree(s string) *TreeNode {
	numArr := exactInt(s)
	if len(numArr.arr) == 0 {
		return nil
	}
	root := &TreeNode{Val: *numArr.arr[0]}
	i := 1
	for parentLevel := []*TreeNode{root}; len(parentLevel) > 0; {
		nextLevel := make([]*TreeNode, 0)
		for _, n := range parentLevel {
			if i < len(numArr.arr) {
				if numArr.arr[i] != nil {
					n.Left = &TreeNode{Val: *numArr.arr[i]}
					nextLevel = append(nextLevel, n.Left)
				}
				i++
			} else{
				break
			}
			if i < len(numArr.arr) {
				if numArr.arr[i] != nil{
					n.Right = &TreeNode{Val: *numArr.arr[i]}
					nextLevel = append(nextLevel, n.Right)
				}
				i++
			} else {
				break
			}
		}
		parentLevel, nextLevel = nextLevel, make([]*TreeNode, 0)
	}
	return root
}

func exactInt(s string) NumArr {
	res := make([]*int, 0)
	elems := exactElem(s)
	for _, i:= range elems {
		if i == "null" {
			res = append(res, nil)
		} else{
			num, err := strconv.Atoi(i)
			if err == nil {
				res = append(res, &num)
			} else {
				panic(err)
			}
		}
	}
	return NumArr{arr: res}
}
