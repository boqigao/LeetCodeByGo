package LC1647

func minDeletions(s string) int {
	var bucket [26]int
	for _, v := range s {
		bucket[v-'a']++
	}
	freq := make(map[int]int)
	maxKey := 0
	// freq存储了每个出现次数的freq，比如aabbcc那就是，2（key）出现了3（value）次
	for _, count := range bucket {
		if count == 0 {
			continue
		}
		freq[count]++
		if count > maxKey {
			maxKey = count
		}
	}
	deleted := 0

	for key := maxKey; key > 0; key-- {
		if freq[key] < 2 {
			continue
		}
		toDelete := freq[key] - 1
		freq[key-1] += toDelete
		deleted += toDelete
	}
	return deleted

}
