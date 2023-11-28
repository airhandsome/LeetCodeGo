type FrontMiddleBackQueue struct {
    left *list.List
    right *list.List

}


func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{
        left: list.New(),
        right: list.New(),
    }
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
    this.left.PushFront(val)
    if this.left.Len() == this.right.Len() + 2{
        this.right.PushFront(this.left.Back().Value.(int))
        this.left.Remove(this.left.Back())
    }
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
    if this.left.Len() == this.right.Len() + 1{
        this.right.PushFront(this.left.Back().Value.(int))
        this.left.Remove(this.left.Back())
    }
    this.left.PushBack(val)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
    this.right.PushBack(val)
    if this.left.Len() + 1 == this.right.Len(){
        this.left.PushBack(this.right.Front().Value.(int))
        this.right.Remove(this.right.Front())
    }
}

func (this *FrontMiddleBackQueue) PopFront() int {
    if this.left.Len() == 0 {
        return -1
    }
    val := this.left.Front().Value.(int)
    this.left.Remove(this.left.Front())
    if this.left.Len() + 1 == this.right.Len() {
        this.left.PushBack(this.right.Front().Value.(int))
        this.right.Remove(this.right.Front())
    }
    return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
    if this.left.Len() == 0 {
        return -1
    }
    val := this.left.Back().Value.(int)
    this.left.Remove(this.left.Back())
    if this.left.Len() + 1 == this.right.Len() {
        this.left.PushBack(this.right.Front().Value.(int))
        this.right.Remove(this.right.Front())
    }
    return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
    if this.left.Len() == 0 {
        return -1
    }
    if this.right.Len() == 0 {
        val := this.left.Back().Value.(int)
        this.left.Remove(this.left.Back())
        return val
    } else {
        val := this.right.Back().Value.(int)
        this.right.Remove(this.right.Back())
        if this.left.Len() == this.right.Len() + 2 {
            this.right.PushFront(this.left.Back().Value.(int))
            this.left.Remove(this.left.Back())
        }
        return val
    }
}



/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */