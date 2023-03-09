---
title: 2.03 ✅ Two Pointers
type: docs
weight: 3
---

# Two Pointers

![](https://img.halfrost.com/Leetcode/Two_pointers.png)

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

- Fast and slow pointers can find repeated numbers, time complexity O(n), question 287.
- After replacing letters, the same letter can appear for the longest length in a row. Question 424.
- SUM problem sets. Question 1, Question 15, Question 16, Question 18, Question 167, Question 923, Question 1074.


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0011|Container With Most Water|[Go]({{< relref "/ChapterFour/0001~0099/0011.Container-With-Most-Water.md" >}})|Medium| O(n)| O(1)||54.0%|
|0015|3Sum|[Go]({{< relref "/ChapterFour/0001~0099/0015.3Sum.md" >}})|Medium| O(n^2)| O(n)|❤️|32.5%|
|0016|3Sum Closest|[Go]({{< relref "/ChapterFour/0001~0099/0016.3Sum-Closest.md" >}})|Medium| O(n^2)| O(1)|❤️|45.8%|
|0018|4Sum|[Go]({{< relref "/ChapterFour/0001~0099/0018.4Sum.md" >}})|Medium| O(n^3)| O(n^2)|❤️|36.0%|
|0019|Remove Nth Node From End of List|[Go]({{< relref "/ChapterFour/0001~0099/0019.Remove-Nth-Node-From-End-of-List.md" >}})|Medium| O(n)| O(1)||40.9%|
|0026|Remove Duplicates from Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0026.Remove-Duplicates-from-Sorted-Array.md" >}})|Easy| O(n)| O(1)||51.4%|
|0027|Remove Element|[Go]({{< relref "/ChapterFour/0001~0099/0027.Remove-Element.md" >}})|Easy| O(n)| O(1)||52.9%|
|0028|Find the Index of the First Occurrence in a String|[Go]({{< relref "/ChapterFour/0001~0099/0028.Find-the-Index-of-the-First-Occurrence-in-a-String.md" >}})|Easy| O(n)| O(1)||38.8%|
|0031|Next Permutation|[Go]({{< relref "/ChapterFour/0001~0099/0031.Next-Permutation.md" >}})|Medium||||37.5%|
|0042|Trapping Rain Water|[Go]({{< relref "/ChapterFour/0001~0099/0042.Trapping-Rain-Water.md" >}})|Hard| O(n)| O(1)|❤️|59.2%|
|0061|Rotate List|[Go]({{< relref "/ChapterFour/0001~0099/0061.Rotate-List.md" >}})|Medium| O(n)| O(1)||36.0%|
|0075|Sort Colors|[Go]({{< relref "/ChapterFour/0001~0099/0075.Sort-Colors.md" >}})|Medium| O(n)| O(1)|❤️|58.3%|
|0080|Remove Duplicates from Sorted Array II|[Go]({{< relref "/ChapterFour/0001~0099/0080.Remove-Duplicates-from-Sorted-Array-II.md" >}})|Medium| O(n)| O(1||52.2%|
|0082|Remove Duplicates from Sorted List II|[Go]({{< relref "/ChapterFour/0001~0099/0082.Remove-Duplicates-from-Sorted-List-II.md" >}})|Medium||||45.8%|
|0086|Partition List|[Go]({{< relref "/ChapterFour/0001~0099/0086.Partition-List.md" >}})|Medium| O(n)| O(1)|❤️|51.9%|
|0088|Merge Sorted Array|[Go]({{< relref "/ChapterFour/0001~0099/0088.Merge-Sorted-Array.md" >}})|Easy| O(n)| O(1)|❤️|46.5%|
|0125|Valid Palindrome|[Go]({{< relref "/ChapterFour/0100~0199/0125.Valid-Palindrome.md" >}})|Easy| O(n)| O(1)||44.3%|
|0141|Linked List Cycle|[Go]({{< relref "/ChapterFour/0100~0199/0141.Linked-List-Cycle.md" >}})|Easy| O(n)| O(1)|❤️|47.4%|
|0142|Linked List Cycle II|[Go]({{< relref "/ChapterFour/0100~0199/0142.Linked-List-Cycle-II.md" >}})|Medium| O(n)| O(1)|❤️|48.4%|
|0143|Reorder List|[Go]({{< relref "/ChapterFour/0100~0199/0143.Reorder-List.md" >}})|Medium||||52.3%|
|0148|Sort List|[Go]({{< relref "/ChapterFour/0100~0199/0148.Sort-List.md" >}})|Medium||||55.0%|
|0151|Reverse Words in a String|[Go]({{< relref "/ChapterFour/0100~0199/0151.Reverse-Words-in-a-String.md" >}})|Medium||||32.6%|
|0160|Intersection of Two Linked Lists|[Go]({{< relref "/ChapterFour/0100~0199/0160.Intersection-of-Two-Linked-Lists.md" >}})|Easy||||54.2%|
|0167|Two Sum II - Input Array Is Sorted|[Go]({{< relref "/ChapterFour/0100~0199/0167.Two-Sum-II-Input-Array-Is-Sorted.md" >}})|Medium| O(n)| O(1)||60.0%|
|0189|Rotate Array|[Go]({{< relref "/ChapterFour/0100~0199/0189.Rotate-Array.md" >}})|Medium||||39.4%|
|0202|Happy Number|[Go]({{< relref "/ChapterFour/0200~0299/0202.Happy-Number.md" >}})|Easy||||54.7%|
|0234|Palindrome Linked List|[Go]({{< relref "/ChapterFour/0200~0299/0234.Palindrome-Linked-List.md" >}})|Easy| O(n)| O(1)||50.1%|
|0283|Move Zeroes|[Go]({{< relref "/ChapterFour/0200~0299/0283.Move-Zeroes.md" >}})|Easy| O(n)| O(1)||61.4%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium| O(n)| O(1)|❤️|59.1%|
|0344|Reverse String|[Go]({{< relref "/ChapterFour/0300~0399/0344.Reverse-String.md" >}})|Easy| O(n)| O(1)||76.7%|
|0345|Reverse Vowels of a String|[Go]({{< relref "/ChapterFour/0300~0399/0345.Reverse-Vowels-of-a-String.md" >}})|Easy| O(n)| O(1)||50.0%|
|0349|Intersection of Two Arrays|[Go]({{< relref "/ChapterFour/0300~0399/0349.Intersection-of-Two-Arrays.md" >}})|Easy| O(n)| O(n) ||70.8%|
|0350|Intersection of Two Arrays II|[Go]({{< relref "/ChapterFour/0300~0399/0350.Intersection-of-Two-Arrays-II.md" >}})|Easy| O(n)| O(n) ||55.9%|
|0392|Is Subsequence|[Go]({{< relref "/ChapterFour/0300~0399/0392.Is-Subsequence.md" >}})|Easy||||47.8%|
|0455|Assign Cookies|[Go]({{< relref "/ChapterFour/0400~0499/0455.Assign-Cookies.md" >}})|Easy||||50.0%|
|0457|Circular Array Loop|[Go]({{< relref "/ChapterFour/0400~0499/0457.Circular-Array-Loop.md" >}})|Medium||||32.6%|
|0475|Heaters|[Go]({{< relref "/ChapterFour/0400~0499/0475.Heaters.md" >}})|Medium||||36.4%|
|0524|Longest Word in Dictionary through Deleting|[Go]({{< relref "/ChapterFour/0500~0599/0524.Longest-Word-in-Dictionary-through-Deleting.md" >}})|Medium| O(n)| O(1) ||51.0%|
|0532|K-diff Pairs in an Array|[Go]({{< relref "/ChapterFour/0500~0599/0532.K-diff-Pairs-in-an-Array.md" >}})|Medium| O(n)| O(n)||41.1%|
|0541|Reverse String II|[Go]({{< relref "/ChapterFour/0500~0599/0541.Reverse-String-II.md" >}})|Easy||||50.5%|
|0557|Reverse Words in a String III|[Go]({{< relref "/ChapterFour/0500~0599/0557.Reverse-Words-in-a-String-III.md" >}})|Easy||||81.9%|
|0567|Permutation in String|[Go]({{< relref "/ChapterFour/0500~0599/0567.Permutation-in-String.md" >}})|Medium| O(n)| O(1)|❤️|44.4%|
|0581|Shortest Unsorted Continuous Subarray|[Go]({{< relref "/ChapterFour/0500~0599/0581.Shortest-Unsorted-Continuous-Subarray.md" >}})|Medium||||36.4%|
|0611|Valid Triangle Number|[Go]({{< relref "/ChapterFour/0600~0699/0611.Valid-Triangle-Number.md" >}})|Medium||||50.5%|
|0633|Sum of Square Numbers|[Go]({{< relref "/ChapterFour/0600~0699/0633.Sum-of-Square-Numbers.md" >}})|Medium||||34.4%|
|0653|Two Sum IV - Input is a BST|[Go]({{< relref "/ChapterFour/0600~0699/0653.Two-Sum-IV-Input-is-a-BST.md" >}})|Easy||||61.0%|
|0658|Find K Closest Elements|[Go]({{< relref "/ChapterFour/0600~0699/0658.Find-K-Closest-Elements.md" >}})|Medium||||46.8%|
|0696|Count Binary Substrings|[Go]({{< relref "/ChapterFour/0600~0699/0696.Count-Binary-Substrings.md" >}})|Easy||||65.6%|
|0719|Find K-th Smallest Pair Distance|[Go]({{< relref "/ChapterFour/0700~0799/0719.Find-K-th-Smallest-Pair-Distance.md" >}})|Hard||||36.7%|
|0763|Partition Labels|[Go]({{< relref "/ChapterFour/0700~0799/0763.Partition-Labels.md" >}})|Medium| O(n)| O(1)|❤️|79.8%|
|0795|Number of Subarrays with Bounded Maximum|[Go]({{< relref "/ChapterFour/0700~0799/0795.Number-of-Subarrays-with-Bounded-Maximum.md" >}})|Medium||||52.8%|
|0821|Shortest Distance to a Character|[Go]({{< relref "/ChapterFour/0800~0899/0821.Shortest-Distance-to-a-Character.md" >}})|Easy||||71.3%|
|0825|Friends Of Appropriate Ages|[Go]({{< relref "/ChapterFour/0800~0899/0825.Friends-Of-Appropriate-Ages.md" >}})|Medium||||46.3%|
|0826|Most Profit Assigning Work|[Go]({{< relref "/ChapterFour/0800~0899/0826.Most-Profit-Assigning-Work.md" >}})|Medium| O(n log n)| O(n)||44.8%|
|0832|Flipping an Image|[Go]({{< relref "/ChapterFour/0800~0899/0832.Flipping-an-Image.md" >}})|Easy||||80.7%|
|0838|Push Dominoes|[Go]({{< relref "/ChapterFour/0800~0899/0838.Push-Dominoes.md" >}})|Medium| O(n)| O(n)||57.0%|
|0844|Backspace String Compare|[Go]({{< relref "/ChapterFour/0800~0899/0844.Backspace-String-Compare.md" >}})|Easy| O(n)| O(n) ||48.0%|
|0845|Longest Mountain in Array|[Go]({{< relref "/ChapterFour/0800~0899/0845.Longest-Mountain-in-Array.md" >}})|Medium| O(n)| O(1) ||40.2%|
|0870|Advantage Shuffle|[Go]({{< relref "/ChapterFour/0800~0899/0870.Advantage-Shuffle.md" >}})|Medium||||51.9%|
|0876|Middle of the Linked List|[Go]({{< relref "/ChapterFour/0800~0899/0876.Middle-of-the-Linked-List.md" >}})|Easy||||75.4%|
|0881|Boats to Save People|[Go]({{< relref "/ChapterFour/0800~0899/0881.Boats-to-Save-People.md" >}})|Medium| O(n log n)| O(1) ||53.0%|
|0922|Sort Array By Parity II|[Go]({{< relref "/ChapterFour/0900~0999/0922.Sort-Array-By-Parity-II.md" >}})|Easy||||70.7%|
|0923|3Sum With Multiplicity|[Go]({{< relref "/ChapterFour/0900~0999/0923.3Sum-With-Multiplicity.md" >}})|Medium| O(n^2)| O(n) ||45.3%|
|0925|Long Pressed Name|[Go]({{< relref "/ChapterFour/0900~0999/0925.Long-Pressed-Name.md" >}})|Easy| O(n)| O(1)||33.2%|
|0942|DI String Match|[Go]({{< relref "/ChapterFour/0900~0999/0942.DI-String-Match.md" >}})|Easy||||77.2%|
|0969|Pancake Sorting|[Go]({{< relref "/ChapterFour/0900~0999/0969.Pancake-Sorting.md" >}})|Medium||||70.1%|
|0977|Squares of a Sorted Array|[Go]({{< relref "/ChapterFour/0900~0999/0977.Squares-of-a-Sorted-Array.md" >}})|Easy| O(n)| O(1)||71.9%|
|0986|Interval List Intersections|[Go]({{< relref "/ChapterFour/0900~0999/0986.Interval-List-Intersections.md" >}})|Medium| O(n)| O(1)||71.4%|
|1040|Moving Stones Until Consecutive II|[Go]({{< relref "/ChapterFour/1000~1099/1040.Moving-Stones-Until-Consecutive-II.md" >}})|Medium||||55.8%|
|1048|Longest String Chain|[Go]({{< relref "/ChapterFour/1000~1099/1048.Longest-String-Chain.md" >}})|Medium||||59.2%|
|1089|Duplicate Zeros|[Go]({{< relref "/ChapterFour/1000~1099/1089.Duplicate-Zeros.md" >}})|Easy||||51.5%|
|1332|Remove Palindromic Subsequences|[Go]({{< relref "/ChapterFour/1300~1399/1332.Remove-Palindromic-Subsequences.md" >}})|Easy||||76.2%|
|1385|Find the Distance Value Between Two Arrays|[Go]({{< relref "/ChapterFour/1300~1399/1385.Find-the-Distance-Value-Between-Two-Arrays.md" >}})|Easy||||66.4%|
|1679|Max Number of K-Sum Pairs|[Go]({{< relref "/ChapterFour/1600~1699/1679.Max-Number-of-K-Sum-Pairs.md" >}})|Medium||||57.3%|
|1721|Swapping Nodes in a Linked List|[Go]({{< relref "/ChapterFour/1700~1799/1721.Swapping-Nodes-in-a-Linked-List.md" >}})|Medium||||67.3%|
|1877|Minimize Maximum Pair Sum in Array|[Go]({{< relref "/ChapterFour/1800~1899/1877.Minimize-Maximum-Pair-Sum-in-Array.md" >}})|Medium||||79.9%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|
------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/String/">⬅️上一页</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/">下一页➡️</a></p>
</div>
