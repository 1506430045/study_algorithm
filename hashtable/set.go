package hashtable

type MyHashSet struct {
	m [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		m: make([][]int, 10000),
	}
}

func hash(key int) int {
	return key % 10000
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}
	for k, v := range this.m[hash(key)] {
		if v == -1 {
			this.m[hash(key)][k] = key
			return
		}
	}
	this.m[hash(key)] = append(this.m[hash(key)], key)
}

func (this *MyHashSet) Remove(key int) {
	for k, v := range this.m[hash(key)] {
		if v == key {
			this.m[hash(key)][k] = -1
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, v := range this.m[hash(key)] {
		if v == key {
			return true
		}
	}
	return false
}
