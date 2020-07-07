package hashtable

import (
	"fmt"
	"testing"
)

func TestConstructor2(t *testing.T) {
	hashMap := Constructor2()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // 返回 1
	fmt.Println(hashMap.Get(3)) // 返回 -1 (未找到)
	hashMap.Put(2, 1)           // 更新已有的值
	fmt.Println(hashMap.Get(2)) // 返回 1
	hashMap.Remove(2)           // 删除键为2的数据
	fmt.Println(hashMap.Get(2)) // 返回 -1 (未找到)

}

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{3, 2, 4}, 7))
}
