package datastruct

type Set interface{
	Search(s Set, key int) *Element
	Insert(s Set, x *Element)
	Delete(s Set, x *Element)
	Minimum(s Set) *Element
	Maximum(s Set) *Element
	Successor(s Set, x *Element) *Element
	Predecessor(s Set, x *Element) *Element
}