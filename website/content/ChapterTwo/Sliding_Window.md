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


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0003|Longest Substring Without Repeating Characters|[Go]({{< relref "/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters.md" >}})|Medium| O(n)| O(1)|❤️|33.8%|
|0030|Substring with Concatenation of All Words|[Go]({{< relref "/ChapterFour/0001~0099/0030.Substring-with-Concatenation-of-All-Words.md" >}})|Hard||||31.2%|
|0076|Minimum Window Substring|[Go]({{< relref "/ChapterFour/0001~0099/0076.Minimum-Window-Substring.md" >}})|Hard| O(n)| O(n)|❤️|40.9%|
|0187|Repeated DNA Sequences|[Go]({{< relref "/ChapterFour/0100~0199/0187.Repeated-DNA-Sequences.md" >}})|Medium||||46.9%|
|0209|Minimum Size Subarray Sum|[Go]({{< relref "/ChapterFour/0200~0299/0209.Minimum-Size-Subarray-Sum.md" >}})|Medium||||44.9%|
|0219|Contains Duplicate II|[Go]({{< relref "/ChapterFour/0200~0299/0219.Contains-Duplicate-II.md" >}})|Easy||||42.5%|
|0220|Contains Duplicate III|[Go]({{< relref "/ChapterFour/0200~0299/0220.Contains-Duplicate-III.md" >}})|Hard||||22.1%|
|0239|Sliding Window Maximum|[Go]({{< relref "/ChapterFour/0200~0299/0239.Sliding-Window-Maximum.md" >}})|Hard| O(n * k)| O(n)|❤️|46.3%|
|0395|Longest Substring with At Least K Repeating Characters|[Go]({{< relref "/ChapterFour/0300~0399/0395.Longest-Substring-with-At-Least-K-Repeating-Characters.md" >}})|Medium||||44.8%|
|0424|Longest Repeating Character Replacement|[Go]({{< relref "/ChapterFour/0400~0499/0424.Longest-Repeating-Character-Replacement.md" >}})|Medium| O(n)| O(1) ||51.9%|
|0438|Find All Anagrams in a String|[Go]({{< relref "/ChapterFour/0400~0499/0438.Find-All-Anagrams-in-a-String.md" >}})|Medium||||50.2%|
|0480|Sliding Window Median|[Go]({{< relref "/ChapterFour/0400~0499/0480.Sliding-Window-Median.md" >}})|Hard| O(n * log k)| O(k)|❤️|41.3%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.4%|
|0632|Smallest Range Covering Elements from K Lists|[Go]({{< relref "/ChapterFour/0600~0699/0632.Smallest-Range-Covering-Elements-from-K-Lists.md" >}})|Hard||||60.9%|
|0643|Maximum Average Subarray I|[Go]({{< relref "/ChapterFour/0600~0699/0643.Maximum-Average-Subarray-I.md" >}})|Easy||||43.7%|
|0658|Find K Closest Elements|[Go]({{< relref "/ChapterFour/0600~0699/0658.Find-K-Closest-Elements.md" >}})|Medium||||46.8%|
|0713|Subarray Product Less Than K|[Go]({{< relref "/ChapterFour/0700~0799/0713.Subarray-Product-Less-Than-K.md" >}})|Medium||||45.7%|
|0718|Maximum Length of Repeated Subarray|[Go]({{< relref "/ChapterFour/0700~0799/0718.Maximum-Length-of-Repeated-Subarray.md" >}})|Medium||||51.3%|
|0862|Shortest Subarray with Sum at Least K|[Go]({{< relref "/ChapterFour/0800~0899/0862.Shortest-Subarray-with-Sum-at-Least-K.md" >}})|Hard||||26.1%|
|0904|Fruit Into Baskets|[Go]({{< relref "/ChapterFour/0900~0999/0904.Fruit-Into-Baskets.md" >}})|Medium||||43.7%|
|0930|Binary Subarrays With Sum|[Go]({{< relref "/ChapterFour/0900~0999/0930.Binary-Subarrays-With-Sum.md" >}})|Medium||||52.0%|
|0978|Longest Turbulent Subarray|[Go]({{< relref "/ChapterFour/0900~0999/0978.Longest-Turbulent-Subarray.md" >}})|Medium| O(n)| O(1)|❤️|47.2%|
|0992|Subarrays with K Different Integers|[Go]({{< relref "/ChapterFour/0900~0999/0992.Subarrays-with-K-Different-Integers.md" >}})|Hard| O(n)| O(n)|❤️|54.7%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0900~0999/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard| O(n)| O(1)|❤️|51.2%|
|1004|Max Consecutive Ones III|[Go]({{< relref "/ChapterFour/1000~1099/1004.Max-Consecutive-Ones-III.md" >}})|Medium| O(n)| O(1) ||63.2%|
|1052|Grumpy Bookstore Owner|[Go]({{< relref "/ChapterFour/1000~1099/1052.Grumpy-Bookstore-Owner.md" >}})|Medium| O(n log n)| O(1) ||57.1%|
|1208|Get Equal Substrings Within Budget|[Go]({{< relref "/ChapterFour/1200~1299/1208.Get-Equal-Substrings-Within-Budget.md" >}})|Medium||||48.5%|
|1234|Replace the Substring for Balanced String|[Go]({{< relref "/ChapterFour/1200~1299/1234.Replace-the-Substring-for-Balanced-String.md" >}})|Medium||||37.1%|
|1423|Maximum Points You Can Obtain from Cards|[Go]({{< relref "/ChapterFour/1400~1499/1423.Maximum-Points-You-Can-Obtain-from-Cards.md" >}})|Medium||||52.3%|
|1438|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit|[Go]({{< relref "/ChapterFour/1400~1499/1438.Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit.md" >}})|Medium||||48.3%|
|1658|Minimum Operations to Reduce X to Zero|[Go]({{< relref "/ChapterFour/1600~1699/1658.Minimum-Operations-to-Reduce-X-to-Zero.md" >}})|Medium||||37.6%|
|1695|Maximum Erasure Value|[Go]({{< relref "/ChapterFour/1600~1699/1695.Maximum-Erasure-Value.md" >}})|Medium||||57.6%|
|1696|Jump Game VI|[Go]({{< relref "/ChapterFour/1600~1699/1696.Jump-Game-VI.md" >}})|Medium||||46.1%|
|1763|Longest Nice Substring|[Go]({{< relref "/ChapterFour/1700~1799/1763.Longest-Nice-Substring.md" >}})|Easy||||61.6%|
|1984|Minimum Difference Between Highest and Lowest of K Scores|[Go]({{< relref "/ChapterFour/1900~1999/1984.Minimum-Difference-Between-Highest-and-Lowest-of-K-Scores.md" >}})|Easy||||54.4%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">⬅️previous page</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Segment_Tree/">next page➡️</a></p>
</div>
