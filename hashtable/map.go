package hashtable

type MyHashMap struct {
	m []int
}

/** Initialize your data structure here. */
func Constructor2() *MyHashMap {
	hm := &MyHashMap{
		make([]int, 1000000),
	}
	for k, _ := range hm.m {
		hm.m[k] = -1
	}
	return hm
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.m[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.m[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.m[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = true
	}
	return false
}

func singleNumber(nums []int) int {
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a ^= nums[i]
	}
	return a
}

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	res := make([]int, 0)
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			res = append(res, k)
		}
	}
	return res
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var diff int
	var res []int
	for i := 0; i < len(nums); i++ {
		diff = target - nums[i]
		if k, ok := m[diff]; ok {
			res = append(res, k, i)
		}
		m[nums[i]] = i
	}
	return res
}
