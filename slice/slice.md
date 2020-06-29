#### 1. Slice define
(point to array) + length + capacity

#### 2. Init slice
1. var s []int = make([]int, 5, 8) //zero slice
2. var s []int = []int{2,7,5...} //init slice
3. var s []int //nil slice

#### 3. Copy slice
slice copy is shallow copy, when changed, both influence

#### 4. Copy function is deep copy
  func copy(dst, src []T) int //return length

#### 5. Operate slice
assign and echo slice just like array

#### 6. Slice expansion
> increase 50% when capacity less 1024
> increase 25% when capacity greater 1024
