package lru

import (
	"fmt"
	"testing"
)

func TestNum146(t *testing.T)  {
//输入
	//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	//输出
	//[null, null, null, 1, null, -1, null, -1, 3, 4]
	//
	//解释
	//LRUCache lRUCache = new LRUCache(2);
	//lRUCache.put(1, 1); // 缓存是 {1=1}
	//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
	//lRUCache.get(1);    // 返回 1
	//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//lRUCache.get(2);    // 返回 -1 (未找到)
	//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//lRUCache.get(1);    // 返回 -1 (未找到)
	//lRUCache.get(3);    // 返回 3
	//lRUCache.get(4);    // 返回 4

	//输入：
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//输出：
	//[null,null,null,1,null,-1,null,1,3,4]
	//预期结果：
	//[null,null,null,1,null,-1,null,-1,3,4]
	lru := Constructor(2)
	fmt.Println(&lru)
	datas := [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	for i:= range datas {
		if len(datas[i]) == 2{
			lru.Put(datas[i][0], datas[i][1])
		} else {
			x:=lru.Get(datas[i][0])
			fmt.Println(x)
		}

		//fmt.Println(&lru)
	}
}
