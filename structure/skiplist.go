package structure

import "math/rand"

const (
	maxLevel int8    = 16
	p        float32 = 0.25
)

type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element //前进节点
}

type SkipList struct {
	Head  *Element //head 是空节点，只保留前进指针
	len   uint64
	level int8
}

func New() *SkipList {
	return &SkipList{
		Head: &Element{forward: make([]*Element, maxLevel)},
	}
}
func (skiplist *SkipList) Get(score float64) *Element {
	prev := skiplist.Head

	for i := skiplist.level - 1; i >= 0; i-- {
		for prev.forward[i] != nil && score > prev.forward[i].Score {
			prev = prev.forward[i]
		}
	}
	if prev.forward[0] != nil && prev.forward[0].Score == score {
		return prev.forward[0]
	}
	return nil
}
func (skiplist *SkipList) Put(score float64, value interface{}) *Element {
	prev := skiplist.Head
	update := make([]*Element, maxLevel)
	for i := skiplist.level - 1; i >= 0; i-- {
		for prev.forward[i] != nil && score > prev.forward[i].Score {
			prev = prev.forward[i]
		}
		update[i] = prev //保存前一个节点，方便之后添加新节点
	}
	//找到了相同score，直接更新
	if prev.forward[0] != nil && prev.forward[0].Score == score {
		prev.forward[0].Value = value
		return prev.forward[0]
	}
	level := randomLevel()
	// 加一层索引
	if level > skiplist.level {
		skiplist.level++
		level = skiplist.level
		update[level-1] = skiplist.Head
	}
	e := &Element{
		Score:   score,
		Value:   value,
		forward: make([]*Element, level),
	}
	for i := level - 1; i >= 0; i-- {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	skiplist.len++
	return e
}

func (skiplist *SkipList) Delete(score float64) *Element {
	prev := skiplist.Head
	update := make([]*Element, maxLevel)
	for i := skiplist.level - 1; i >= 0; i-- {
		for prev.forward[i] != nil && score > prev.forward[i].Score {
			prev = prev.forward[i]
		}
		update[i] = prev //保存前一个节点，方便之后删除
	}
	//找到了相同score，然后删除
	cur := prev.forward[0]
	if cur != nil && cur.Score == score {
		for i, v := range cur.forward {
			update[i].forward[i] = v
		}
		return cur
	}
	return nil
}

func randomLevel() int8 {
	var level int8 = 1
	if rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}
