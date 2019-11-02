type ItemID uint64
type ItemSlice []*Item

// GroupByID returns a map and a slice of given ItemSlice
func (ss ItemSlice) GroupByID() (grouped map[ItemID]*Item, ids []ItemID) {
	if len(ss) == 0 {
		return
	}
	grouped = make(map[ItemID]*Item)
	ids = make([]ItemID, len(ss))
	for idx, i := range ss {
		grouped[i.Id] = i
		ids[idx] = ItemID(i.Id)
	}
	return
}
