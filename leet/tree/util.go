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
	Next  *TreeNode
}

func (tree *TreeNode) Print() {
	g :=getGraphLink(tree,len(strconv.Itoa(tree.MaxVal())))
	fmt.Printf("%v", g)
}
type Graph struct {
	graph []string
	offSet int
}

func stringLen(s string) (l int) {
	for _, x := range s {
		if x > 0x4E00 {
			l += 2
		} else {
			l += 1
		}
	}
	return
}

func (g *Graph) RPush(lines...string) *Graph {
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
	return stringLen(graph[0])
}
func layElemS(i, sw, ew int, e string)  string{
	if i + ew > sw || stringLen(e) > ew{
		panic("err")
	}
	return strings.Repeat(" ", i) + fmt.Sprintf("%-[2]*[1]s",e, ew) + strings.Repeat(" ", sw - i - ew)
}
func layElem(i, sw, ew, e int)  string{
		if i + ew > sw || stringLen(strconv.Itoa(e)) > ew{
			panic("err")
		}
		return strings.Repeat(" ", i) + fmt.Sprintf("%-[2]*[1]d",e, ew) + strings.Repeat(" ", sw - i - ew)
}

func getOneValString(t *TreeNode, elemWidth int) (graph Graph) {
	line:= layElem(0, elemWidth, elemWidth,  t.Val)
	o := 0  // 左对齐
	return Graph{graph:[]string{line}, offSet: o}
}

func expandGraph(offset int, graph Graph, left, right, up, down int) (g Graph) {
	if left < 0 || right < 0 || up < 0 || down < 0 || graph.graph == nil ||  len(graph.graph) == 0 {
		panic("err")
	}
	g = Graph{graph: make([]string, len(graph.graph) + up + down), offSet: offset + left}
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
	return g
}

func mergeTwoGraph(lg Graph, rg Graph, gap int) (lOffset, rOffset int, graph Graph) {
	if lg.Height() == 0 || rg.Height() == 0 || gap < 0{
		panic("err")
	}
	rOffset = rg.offSet + lg.Width() + gap
	lOffset = lg.offSet
	graph = Graph{graph: make([]string, 0)}
	if lg.Height() < rg.Height() {
		lg = expandGraph(lg.offSet, lg, 0, 0, 0, rg.Height() - lg.Height())
	} else if lg.Height() > rg.Height() {
		rg = expandGraph(rg.offSet, rg, 0, 0, 0, lg.Height() - rg.Height())
	}
	for i:=0; i < lg.Height(); i++{
		graph.graph = append(graph.graph, lg.graph[i] + strings.Repeat(" ", gap) + rg.graph[i])
	}
	return
}
func genLinkLine(lo, to, ro, w int) string {
	var x string
	var o int
	if ro == -1 {
		x = "┌" + strings.Repeat("─", to - lo - 1) + "┘"
		o = lo
	} else if lo == -1 {
		x = "└" + strings.Repeat("─", ro - to - 1) + "┐"
		o = to
	} else {
		x = "┌" + strings.Repeat("─", to - lo - 1) + "┴" + strings.Repeat("─", ro - to - 1) + "┐"
		o = lo
	}
	if o + stringLen(x) > w {
		panic("err")
	}
	return layElemS(o, w, stringLen(x), x)
}

func getGraphLink(t *TreeNode, elemWidth int) (graph Graph) {
	if elemWidth == -1 {
		elemWidth = stringLen(strconv.Itoa(t.MaxVal()))
	}
	if t == nil || elemWidth < 1 {
		return
	}
	tg := getOneValString(t, elemWidth)
	if t.Left == nil && t.Right == nil {
		return tg
	}
	lg := getGraphLink(t.Left, elemWidth)
	rg := getGraphLink(t.Right, elemWidth)
	if t.Right == nil {
		rootLine := layElem(lg.offSet + 2, lg.Width() + 2, elemWidth, t.Val)
		linkline := genLinkLine(lg.offSet, lg.offSet + 2, -1, lg.Width() + 2)
		lg = expandGraph(lg.offSet, lg, 0, 2, 0, 0)
		g := Graph{graph: make([]string, 0), offSet: lg.offSet + 2}
		g.RPush(rootLine, linkline)
		g.RPush(lg.graph...)
		return g
	}
	if t.Left == nil {
		rootLine := layElem(0, rg.Width() + 2, elemWidth, t.Val)
		linkline := genLinkLine(-1, 0, rg.offSet + 2, rg.Width() + 2)
		rg = expandGraph(rg.offSet, rg, 2, 0, 0, 0)
		g := Graph{graph: make([]string, 0), offSet: 0}
		g.RPush(rootLine, linkline)
		g.RPush(rg.graph...)
		return g
	}
	lo, ro, gc := mergeTwoGraph(lg, rg, 1)
	tg.offSet = (lo + ro) / 2
	rootLine := layElem(tg.offSet, gc.Width(), elemWidth, t.Val)
	linkLine := genLinkLine(lo, tg.offSet, ro, gc.Width())
	g := Graph{graph: make([]string, 0), offSet: tg.offSet}
	g.RPush(rootLine, linkLine)
	g.RPush(gc.graph...)
	return g
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

func (t TreeNode) MaxVal() (res int) {
	res = t.Val
	if t.Left != nil {
		res = max(res, t.Left.MaxVal())
	}
	if t.Right != nil {
		res = max(res, t.Right.MaxVal())
	}
	return res
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

func exactElem(s string) (res []string) {
	res = make([]string, 0)
	s = strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	for _, x := range strings.Split(s, ",") {
		y := strings.TrimFunc(x, func(r rune) bool {
			return unicode.IsSpace(r)
		})
		res = append(res, y)
	}
	return res
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
