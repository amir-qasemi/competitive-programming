// https://leetcode.com/problems/insert-delete-getrandom-o1

type RandomizedSet struct {
	valToIndex map[int]int
	vals       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{valToIndex: map[int]int{}, vals: nil}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.valToIndex[val]
	if ok {
		return false
	}

	this.vals = append(this.vals, val)
	this.valToIndex[val] = len(this.vals) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.valToIndex[val]

	if !ok {
		return false
	}

	delete(this.valToIndex, val)

	if index != len(this.vals)-1 {
		lastElementInVals := this.vals[len(this.vals)-1]
		this.valToIndex[lastElementInVals] = index
		this.vals[index] = lastElementInVals
	}
	this.vals = this.vals[:len(this.vals)-1]

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */