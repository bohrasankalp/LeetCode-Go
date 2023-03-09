---
title: 2.17 ✅ Sliding Window
type: docs
weight: 17
---

# Sliding Window

![](https://img.halfrost.com/Leetcode/Sliding_Window.png)

- The classic way of writing a two-pointer sliding window. The right pointer keeps moving to the right until it cannot move to the right (the specific conditions depend on the topic). When the right pointer reaches the far right, start to move the left pointer to release the left boundary of the window. Question 3, Question 76, Question 209, Question 424, Question 438, Question 567, Question 713, Question 763, Question 845, Question 881, Question 904, Question 978, Question 992 Question 1004, Question 1040, Question 1052.

```c
	left, right := 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			freq[s[right+1]-'a']++
			right++
		} else {
			freq[s[left]-'a']--
			left++
		}
		result = max(result, right-left+1)
	}
```
- Sliding window classic questions. Question 239, question 480.


{{.AvailableTagTable}}