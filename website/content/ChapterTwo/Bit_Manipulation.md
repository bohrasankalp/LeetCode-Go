---
title: 2.15 ✅ Bit Manipulation
type: docs
weight: 15
---

# Bit Manipulation

![](https://img.halfrost.com/Leetcode/Bit_Manipulation.png)

-XOR property. Question 136, Question 268, Question 389, Question 421,

```go
x ^ 0 = x
x ^ 11111……1111 = ~x
x ^ (~x) = 11111……1111
x ^ x = 0
a ^ b = c  => a ^ c = b  => b ^ c = a (commutative law)
a ^ b ^ c = a ^ (b ^ c) = (a ^ b）^ c (associative law)
```

-Construct a special Mask, put 0 or 1 in the special position.

```go
Clear the rightmost n bits of x to zero, x & ( ~0 << n )
Get the nth bit value of x (0 or 1)，(x >> n) & 1
Get the power of the nth bit of x, x & (1 << (n -1))
Set only the nth position to 1, x | (1 << n)
Set only the nth position to 0, x & (~(1 << n))
Clear the most significant bit of x to the nth bit (inclusive), x & ((1 << n) -1)
Clear bits n through 0, inclusive, x & (~((1 << (n + 1)) -1))
```
- The & bitwise operation has a special meaning. Question 260, Question 201, Question 318, Question 371, Question 397, Question 461, Question 693,

```go
X & 1 == 1 to determine whether it is an odd number (even number)
X & = (X -1) clears the least significant bit (LSB) of 1
X & -X get the least significant bit (LSB) of 1
X & ~X = 0
```



| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0029|Divide Two Integers|[Go]({{< relref "/ChapterFour/0001~0099/0029.Divide-Two-Integers.md" >}})|Medium||||17.2%|
|0067|Add Binary|[Go]({{< relref "/ChapterFour/0001~0099/0067.Add-Binary.md" >}})|Easy||||52.4%|
|0078|Subsets|[Go]({{< relref "/ChapterFour/0001~0099/0078.Subsets.md" >}})|Medium| O(n^2)| O(n)|❤️|74.8%|
|0089|Gray Code|[Go]({{< relref "/ChapterFour/0001~0099/0089.Gray-Code.md" >}})|Medium||||57.1%|
|0090|Subsets II|[Go]({{< relref "/ChapterFour/0001~0099/0090.Subsets-II.md" >}})|Medium||||55.8%|
|0136|Single Number|[Go]({{< relref "/ChapterFour/0100~0199/0136.Single-Number.md" >}})|Easy| O(n)| O(1)||70.6%|
|0137|Single Number II|[Go]({{< relref "/ChapterFour/0100~0199/0137.Single-Number-II.md" >}})|Medium| O(n)| O(1)|❤️|58.4%|
|0187|Repeated DNA Sequences|[Go]({{< relref "/ChapterFour/0100~0199/0187.Repeated-DNA-Sequences.md" >}})|Medium| O(n)| O(1)||46.9%|
|0190|Reverse Bits|[Go]({{< relref "/ChapterFour/0100~0199/0190.Reverse-Bits.md" >}})|Easy| O(n)| O(1)|❤️|53.8%|
|0191|Number of 1 Bits|[Go]({{< relref "/ChapterFour/0100~0199/0191.Number-of-1-Bits.md" >}})|Easy| O(n)| O(1)||66.4%|
|0201|Bitwise AND of Numbers Range|[Go]({{< relref "/ChapterFour/0200~0299/0201.Bitwise-AND-of-Numbers-Range.md" >}})|Medium| O(n)| O(1)|❤️|42.5%|
|0231|Power of Two|[Go]({{< relref "/ChapterFour/0200~0299/0231.Power-of-Two.md" >}})|Easy| O(1)| O(1)||46.0%|
|0260|Single Number III|[Go]({{< relref "/ChapterFour/0200~0299/0260.Single-Number-III.md" >}})|Medium| O(n)| O(1)|❤️|67.6%|
|0268|Missing Number|[Go]({{< relref "/ChapterFour/0200~0299/0268.Missing-Number.md" >}})|Easy| O(n)| O(1)||62.4%|
|0287|Find the Duplicate Number|[Go]({{< relref "/ChapterFour/0200~0299/0287.Find-the-Duplicate-Number.md" >}})|Medium||||59.1%|
|0318|Maximum Product of Word Lengths|[Go]({{< relref "/ChapterFour/0300~0399/0318.Maximum-Product-of-Word-Lengths.md" >}})|Medium| O(n)| O(1)||59.9%|
|0338|Counting Bits|[Go]({{< relref "/ChapterFour/0300~0399/0338.Counting-Bits.md" >}})|Easy| O(n)| O(n)||75.7%|
|0342|Power of Four|[Go]({{< relref "/ChapterFour/0300~0399/0342.Power-of-Four.md" >}})|Easy| O(n)| O(1)||46.1%|
|0371|Sum of Two Integers|[Go]({{< relref "/ChapterFour/0300~0399/0371.Sum-of-Two-Integers.md" >}})|Medium| O(n)| O(1)||50.7%|
|0389|Find the Difference|[Go]({{< relref "/ChapterFour/0300~0399/0389.Find-the-Difference.md" >}})|Easy| O(n)| O(1)||60.0%|
|0393|UTF-8 Validation|[Go]({{< relref "/ChapterFour/0300~0399/0393.UTF-8-Validation.md" >}})|Medium| O(n)| O(1)||45.1%|
|0397|Integer Replacement|[Go]({{< relref "/ChapterFour/0300~0399/0397.Integer-Replacement.md" >}})|Medium| O(n)| O(1)||35.2%|
|0401|Binary Watch|[Go]({{< relref "/ChapterFour/0400~0499/0401.Binary-Watch.md" >}})|Easy| O(1)| O(1)||52.2%|
|0405|Convert a Number to Hexadecimal|[Go]({{< relref "/ChapterFour/0400~0499/0405.Convert-a-Number-to-Hexadecimal.md" >}})|Easy| O(n)| O(1)||46.7%|
|0421|Maximum XOR of Two Numbers in an Array|[Go]({{< relref "/ChapterFour/0400~0499/0421.Maximum-XOR-of-Two-Numbers-in-an-Array.md" >}})|Medium| O(n)| O(1)|❤️|54.1%|
|0461|Hamming Distance|[Go]({{< relref "/ChapterFour/0400~0499/0461.Hamming-Distance.md" >}})|Easy| O(n)| O(1)||75.0%|
|0473|Matchsticks to Square|[Go]({{< relref "/ChapterFour/0400~0499/0473.Matchsticks-to-Square.md" >}})|Medium||||40.2%|
|0476|Number Complement|[Go]({{< relref "/ChapterFour/0400~0499/0476.Number-Complement.md" >}})|Easy| O(n)| O(1)||67.3%|
|0477|Total Hamming Distance|[Go]({{< relref "/ChapterFour/0400~0499/0477.Total-Hamming-Distance.md" >}})|Medium| O(n)| O(1)||52.1%|
|0491|Non-decreasing Subsequences|[Go]({{< relref "/ChapterFour/0400~0499/0491.Non-decreasing-Subsequences.md" >}})|Medium||||60.1%|
|0526|Beautiful Arrangement|[Go]({{< relref "/ChapterFour/0500~0599/0526.Beautiful-Arrangement.md" >}})|Medium||||64.5%|
|0638|Shopping Offers|[Go]({{< relref "/ChapterFour/0600~0699/0638.Shopping-Offers.md" >}})|Medium||||53.3%|
|0645|Set Mismatch|[Go]({{< relref "/ChapterFour/0600~0699/0645.Set-Mismatch.md" >}})|Easy||||42.8%|
|0693|Binary Number with Alternating Bits|[Go]({{< relref "/ChapterFour/0600~0699/0693.Binary-Number-with-Alternating-Bits.md" >}})|Easy| O(n)| O(1)|❤️|61.6%|
|0756|Pyramid Transition Matrix|[Go]({{< relref "/ChapterFour/0700~0799/0756.Pyramid-Transition-Matrix.md" >}})|Medium| O(n log n)| O(n)||52.7%|
|0762|Prime Number of Set Bits in Binary Representation|[Go]({{< relref "/ChapterFour/0700~0799/0762.Prime-Number-of-Set-Bits-in-Binary-Representation.md" >}})|Easy| O(n)| O(1)||68.0%|
|0784|Letter Case Permutation|[Go]({{< relref "/ChapterFour/0700~0799/0784.Letter-Case-Permutation.md" >}})|Medium| O(n)| O(1)||73.8%|
|0810|Chalkboard XOR Game|[Go]({{< relref "/ChapterFour/0800~0899/0810.Chalkboard-XOR-Game.md" >}})|Hard||||56.2%|
|0864|Shortest Path to Get All Keys|[Go]({{< relref "/ChapterFour/0800~0899/0864.Shortest-Path-to-Get-All-Keys.md" >}})|Hard||||45.5%|
|0898|Bitwise ORs of Subarrays|[Go]({{< relref "/ChapterFour/0800~0899/0898.Bitwise-ORs-of-Subarrays.md" >}})|Medium| O(n)| O(1)||37.2%|
|0980|Unique Paths III|[Go]({{< relref "/ChapterFour/0900~0999/0980.Unique-Paths-III.md" >}})|Hard||||81.9%|
|0995|Minimum Number of K Consecutive Bit Flips|[Go]({{< relref "/ChapterFour/0900~0999/0995.Minimum-Number-of-K-Consecutive-Bit-Flips.md" >}})|Hard||||51.2%|
|0996|Number of Squareful Arrays|[Go]({{< relref "/ChapterFour/0900~0999/0996.Number-of-Squareful-Arrays.md" >}})|Hard||||49.2%|
|1009|Complement of Base 10 Integer|[Go]({{< relref "/ChapterFour/1000~1099/1009.Complement-of-Base-10-Integer.md" >}})|Easy||||61.6%|
|1178|Number of Valid Words for Each Puzzle|[Go]({{< relref "/ChapterFour/1100~1199/1178.Number-of-Valid-Words-for-Each-Puzzle.md" >}})|Hard||||46.3%|
|1239|Maximum Length of a Concatenated String with Unique Characters|[Go]({{< relref "/ChapterFour/1200~1299/1239.Maximum-Length-of-a-Concatenated-String-with-Unique-Characters.md" >}})|Medium||||52.2%|
|1310|XOR Queries of a Subarray|[Go]({{< relref "/ChapterFour/1300~1399/1310.XOR-Queries-of-a-Subarray.md" >}})|Medium||||72.3%|
|1442|Count Triplets That Can Form Two Arrays of Equal XOR|[Go]({{< relref "/ChapterFour/1400~1499/1442.Count-Triplets-That-Can-Form-Two-Arrays-of-Equal-XOR.md" >}})|Medium||||76.0%|
|1461|Check If a String Contains All Binary Codes of Size K|[Go]({{< relref "/ChapterFour/1400~1499/1461.Check-If-a-String-Contains-All-Binary-Codes-of-Size-K.md" >}})|Medium||||56.6%|
|1486|XOR Operation in an Array|[Go]({{< relref "/ChapterFour/1400~1499/1486.XOR-Operation-in-an-Array.md" >}})|Easy||||84.6%|
|1655|Distribute Repeating Integers|[Go]({{< relref "/ChapterFour/1600~1699/1655.Distribute-Repeating-Integers.md" >}})|Hard||||38.9%|
|1659|Maximize Grid Happiness|[Go]({{< relref "/ChapterFour/1600~1699/1659.Maximize-Grid-Happiness.md" >}})|Hard||||38.2%|
|1680|Concatenation of Consecutive Binary Numbers|[Go]({{< relref "/ChapterFour/1600~1699/1680.Concatenation-of-Consecutive-Binary-Numbers.md" >}})|Medium||||57.0%|
|1681|Minimum Incompatibility|[Go]({{< relref "/ChapterFour/1600~1699/1681.Minimum-Incompatibility.md" >}})|Hard||||37.4%|
|1684|Count the Number of Consistent Strings|[Go]({{< relref "/ChapterFour/1600~1699/1684.Count-the-Number-of-Consistent-Strings.md" >}})|Easy||||82.2%|
|1720|Decode XORed Array|[Go]({{< relref "/ChapterFour/1700~1799/1720.Decode-XORed-Array.md" >}})|Easy||||85.8%|
|1734|Decode XORed Permutation|[Go]({{< relref "/ChapterFour/1700~1799/1734.Decode-XORed-Permutation.md" >}})|Medium||||63.0%|
|1738|Find Kth Largest XOR Coordinate Value|[Go]({{< relref "/ChapterFour/1700~1799/1738.Find-Kth-Largest-XOR-Coordinate-Value.md" >}})|Medium||||61.0%|
|1763|Longest Nice Substring|[Go]({{< relref "/ChapterFour/1700~1799/1763.Longest-Nice-Substring.md" >}})|Easy||||61.6%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|
-----|-------| ----------------| ---------------|-------------|-------------|-------------|



----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Sorting/">⬅️previous page</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/">next page➡️</a></p>
</div>
