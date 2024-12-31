package main

type FreqStack struct {
    freqMap map[int]int
	data []int
}


func Constructor() FreqStack {
    return FreqStack{
		freqMap: make(map[int]int),
		data : make([]int, 0),
	}
}


func (this *FreqStack) Push(val int)  {
	value, exists := this.freqMap[val]

	if exists {
		value++
		this.freqMap[val] = value
	} else {
		this.freqMap[val] = 1
	}

	this.data = append(this.data, val)
}


func (this *FreqStack) Pop() int {
    // конечно можно и по мапе итерироваться но мне кажется это херней тк полюбому будут огромные пуши и O(n)
	// похоже нужна какая то динамическая структура с ребалансировкой

	// удаляется ближайший к топу самый часто встречающийся элемент
}
