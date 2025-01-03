package main

// https://leetcode.com/problems/maximum-frequency-stack/

// 2nd
type FreqStack struct {
    FrequencyMap    map[int]int 
    MaxFrequency    int         
    FrequencyStacks [][]int     
}

func Constructor() FreqStack {
    return FreqStack{
        FrequencyMap:    make(map[int]int, 10), 
        MaxFrequency:    0,                    
        FrequencyStacks: make([][]int, 1, 10), 
    }
}

func (this *FreqStack) Push(value int) {
    frequency := this.FrequencyMap[value] + 1
    this.FrequencyMap[value] = frequency

    if frequency >= len(this.FrequencyStacks) {
        this.FrequencyStacks = append(this.FrequencyStacks, make([]int, 0, 10))
    }

    this.FrequencyStacks[frequency] = append(this.FrequencyStacks[frequency], value)

    if frequency > this.MaxFrequency {
        this.MaxFrequency = frequency
    }
}

func (this *FreqStack) Pop() int {
    stackWithMaxFrequency := this.FrequencyStacks[this.MaxFrequency]

    lastIndex := len(stackWithMaxFrequency) - 1
    valueToPop := stackWithMaxFrequency[lastIndex]

    this.FrequencyStacks[this.MaxFrequency] = stackWithMaxFrequency[:lastIndex]
    this.FrequencyMap[valueToPop]--

    if len(this.FrequencyStacks[this.MaxFrequency]) == 0 {
        this.MaxFrequency--
    }

    return valueToPop
}


// my 1st shit solution
// type Bucket struct {
// 	bucketElements []BuckerElement
// }

// func (Bucket) Constructor() Bucket {
// 	return Bucket{
// 		bucketElements: make([]BuckerElement, 1),
// 	}
// }

// type BuckerElement struct {
// 	elementsMap map[int]int
// }

// func (BuckerElement) Constructor() BuckerElement {
// 	return BuckerElement{
// 		elementsMap: make(map[int]int),
// 	}
// }

// type FreqStack struct {
// 	freqMap map[int]int
// 	data    []int
// 	bucket  Bucket
// }

// func Constructor() FreqStack {
// 	return FreqStack{
// 		freqMap: make(map[int]int),
// 		data:    make([]int, 0),
// 		bucket:  Bucket{}.Constructor(),
// 	}
// }

// func (this *FreqStack) Push(val int) {
// 	var elemCount int
// 	value, exists := this.freqMap[val]
// 	prevElemCount := value

// 	if exists {
// 		value++
// 		this.freqMap[val] = value
// 		elemCount = value

// 		prevElemBucket := this.bucket.bucketElements[prevElemCount].elementsMap
// 		delete(prevElemBucket, val)		

// 		if elemCount >= len(this.bucket.bucketElements) {
// 			newBucketElement := BuckerElement{
// 				elementsMap: map[int]int{val: elemCount},
// 			}
	
// 			this.bucket.bucketElements = append(this.bucket.bucketElements, newBucketElement)
// 		} else {
// 			existsBucketElement := this.bucket.bucketElements[elemCount]
// 			existsBucketElement.elementsMap[val] = value
// 		}
// 	} else {
// 		this.freqMap[val] = 1
// 		elemCount = 1
	
// 		if elemCount >= len(this.bucket.bucketElements) {
// 			newBucketElement := BuckerElement{
// 				elementsMap: map[int]int{val: elemCount},
// 			}

// 			this.bucket.bucketElements = append(this.bucket.bucketElements, newBucketElement)
// 		} else {
// 			this.bucket.bucketElements[1].elementsMap[val] = elemCount
// 		}
// 	}

// 	this.data = append(this.data, val)
// }

// func (this *FreqStack) Pop() int {
// 	if len(this.data) <= 0 {
// 		return -1
// 	}

// 	var result int
//     var lastBucketElements map[int]int
// 	var elementKey int

// 	for i := len(this.bucket.bucketElements) - 1; i >= 0; i-- {
// 		if len(this.bucket.bucketElements[i].elementsMap) > 0 {
// 			lastBucketElements = this.bucket.bucketElements[i].elementsMap
// 			break
// 		}
// 	}

// 	for i := len(this.data) - 1; i >= 0; i-- {
// 		elementKey = this.data[i]
// 		_, exists := lastBucketElements[elementKey]
// 		if exists {
// 			oldFreqValue := this.freqMap[elementKey]
			
// 			if oldFreqValue != 1 {
// 				this.freqMap[elementKey] = oldFreqValue - 1
// 				this.bucket.bucketElements[oldFreqValue - 1].elementsMap[elementKey] = oldFreqValue - 1
// 			} else {
// 				delete(this.freqMap, elementKey)
// 			}

// 			delete(lastBucketElements, elementKey)
// 			this.removeByIndex(i)

// 			result = elementKey
// 			break
// 		}
// 	}

// 	return result
// }

// func (this *FreqStack) removeByIndex(index int) {
// 	if index < 0 || index >= len(this.data) {
// 		return
// 	}

// 	this.data = append(this.data[:index], this.data[index+1:]...)
// }


// 	// 1. получаем список элементов из последнего бакета
// 	// 2. идем по списку стэка и проверяем есть ли в мапе бакет элемента этот эелемент
// 	// 3. удаляем элемент из текущего бакета и переводим в бакет - 1
// 	// 4. удаляем из массива стэка этот элемент
// 	// 5. чет забыли но потом поймем)

// 	return result
// }



// конечно можно и по мапе итерироваться но мне кажется это херней тк полюбому будут огромные пуши и O(n)
// похоже нужна какая то динамическая структура с ребалансировкой

// удаляется ближайший к топу самый часто встречающийся элемент
