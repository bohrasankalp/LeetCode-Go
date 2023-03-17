---
title: 2.16 ✅ Union Find
type: docs
weight: 16
---

# Union Find

![](https://img.halfrost.com/Leetcode/Union_Find.png)
- Use the idea of ​​union search flexibly, master the [template]({{< relref "/ChapterThree/UnionFind.md" >}}) of union search, there are two ways to realize union search in the template, one It is the version of path compression + rank optimization, and the other is the version of calculating the number of elements in each set + the maximum number of elements in the set. These two versions have their respective uses. The questions that can use the first type of union search template are: question 128, question 130, question 547, question 684, question 721, question 765, question 778, question 839, question 924, question Question 928, Question 947, Question 952, Question 959, Question 990. The questions that can use the second type of union search template are: question 803 and question 952. Question 803 Rank optimization and the number of statistical collections will be time-consuming. If not optimized, it will be TLE.
- Union search is an idea, some questions need to use this idea flexibly, instead of a dead set template, such as question 399, this question is stringUnionFind, which is realized by using the idea of ​​union search. Each node here is based on strings and maps, rather than simply using int node numbers.
- Some questions can’t be done with a dead set template, such as question 685. This question cannot be path compressed and rank optimized, because the question involves a directed graph, and it is necessary to know the predecessor node of the node. If the path is compressed, this question There is no way to do it. This question does not require path compression and rank optimization.
- The information given by the flexible abstract topic, number the given information reasonably, use and search to solve the problem, and use map to reduce the time complexity, such as the 721st question and the 959th question.
- Regarding the topic of maps, bricks, and grids, you can create a special node, and union() the bricks or grids around the edges to this special node. Question 130, question 803.
- Questions that can be used for union search can generally be solved with DFS and BFS, but the time complexity will be higher.


| No.      | Title | Solution | Difficulty | TimeComplexity | SpaceComplexity |Favorite| Acceptance |
|:--------:|:------- | :--------: | :----------: | :----: | :-----: | :-----: |:-----: |
|0128|Longest Consecutive Sequence|[Go]({{< relref "/ChapterFour/0100~0199/0128.Longest-Consecutive-Sequence.md" >}})|Medium| O(n)| O(n)|❤️|48.6%|
|0130|Surrounded Regions|[Go]({{< relref "/ChapterFour/0100~0199/0130.Surrounded-Regions.md" >}})|Medium| O(m\*n)| O(m\*n)||36.6%|
|0200|Number of Islands|[Go]({{< relref "/ChapterFour/0200~0299/0200.Number-of-Islands.md" >}})|Medium| O(m\*n)| O(m\*n)||56.9%|
|0399|Evaluate Division|[Go]({{< relref "/ChapterFour/0300~0399/0399.Evaluate-Division.md" >}})|Medium| O(n)| O(n)||59.6%|
|0547|Number of Provinces|[Go]({{< relref "/ChapterFour/0500~0599/0547.Number-of-Provinces.md" >}})|Medium| O(n^2)| O(n)||63.7%|
|0684|Redundant Connection|[Go]({{< relref "/ChapterFour/0600~0699/0684.Redundant-Connection.md" >}})|Medium| O(n)| O(n)||62.2%|
|0685|Redundant Connection II|[Go]({{< relref "/ChapterFour/0600~0699/0685.Redundant-Connection-II.md" >}})|Hard| O(n)| O(n)||34.1%|
|0695|Max Area of Island|[Go]({{< relref "/ChapterFour/0600~0699/0695.Max-Area-of-Island.md" >}})|Medium||||71.8%|
|0721|Accounts Merge|[Go]({{< relref "/ChapterFour/0700~0799/0721.Accounts-Merge.md" >}})|Medium| O(n)| O(n)|❤️|56.3%|
|0765|Couples Holding Hands|[Go]({{< relref "/ChapterFour/0700~0799/0765.Couples-Holding-Hands.md" >}})|Hard| O(n)| O(n)|❤️|56.8%|
|0778|Swim in Rising Water|[Go]({{< relref "/ChapterFour/0700~0799/0778.Swim-in-Rising-Water.md" >}})|Hard| O(n^2)| O(n)|❤️|59.8%|
|0785|Is Graph Bipartite?|[Go]({{< relref "/ChapterFour/0700~0799/0785.Is-Graph-Bipartite.md" >}})|Medium||||53.1%|
|0803|Bricks Falling When Hit|[Go]({{< relref "/ChapterFour/0800~0899/0803.Bricks-Falling-When-Hit.md" >}})|Hard| O(n^2)| O(n)|❤️|34.3%|
|0839|Similar String Groups|[Go]({{< relref "/ChapterFour/0800~0899/0839.Similar-String-Groups.md" >}})|Hard| O(n^2)| O(n)||48.0%|
|0924|Minimize Malware Spread|[Go]({{< relref "/ChapterFour/0900~0999/0924.Minimize-Malware-Spread.md" >}})|Hard| O(m\*n)| O(n)||42.1%|
|0928|Minimize Malware Spread II|[Go]({{< relref "/ChapterFour/0900~0999/0928.Minimize-Malware-Spread-II.md" >}})|Hard| O(m\*n)| O(n)|❤️|42.7%|
|0947|Most Stones Removed with Same Row or Column|[Go]({{< relref "/ChapterFour/0900~0999/0947.Most-Stones-Removed-with-Same-Row-or-Column.md" >}})|Medium| O(n)| O(n)||58.9%|
|0952|Largest Component Size by Common Factor|[Go]({{< relref "/ChapterFour/0900~0999/0952.Largest-Component-Size-by-Common-Factor.md" >}})|Hard| O(n)| O(n)|❤️|40.1%|
|0959|Regions Cut By Slashes|[Go]({{< relref "/ChapterFour/0900~0999/0959.Regions-Cut-By-Slashes.md" >}})|Medium| O(n^2)| O(n^2)|❤️|69.2%|
|0990|Satisfiability of Equality Equations|[Go]({{< relref "/ChapterFour/0900~0999/0990.Satisfiability-of-Equality-Equations.md" >}})|Medium| O(n)| O(n)||50.5%|
|1020|Number of Enclaves|[Go]({{< relref "/ChapterFour/1000~1099/1020.Number-of-Enclaves.md" >}})|Medium||||65.5%|
|1202|Smallest String With Swaps|[Go]({{< relref "/ChapterFour/1200~1299/1202.Smallest-String-With-Swaps.md" >}})|Medium||||57.6%|
|1254|Number of Closed Islands|[Go]({{< relref "/ChapterFour/1200~1299/1254.Number-of-Closed-Islands.md" >}})|Medium||||64.1%|
|1319|Number of Operations to Make Network Connected|[Go]({{< relref "/ChapterFour/1300~1399/1319.Number-of-Operations-to-Make-Network-Connected.md" >}})|Medium||||58.9%|
|1579|Remove Max Number of Edges to Keep Graph Fully Traversable|[Go]({{< relref "/ChapterFour/1500~1599/1579.Remove-Max-Number-of-Edges-to-Keep-Graph-Fully-Traversable.md" >}})|Hard||||53.2%|
|1631|Path With Minimum Effort|[Go]({{< relref "/ChapterFour/1600~1699/1631.Path-With-Minimum-Effort.md" >}})|Medium||||55.6%|
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|-------------|-------------|


----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Bit_Manipulation/">⬅️previous page</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterTwo/Sliding_Window/">next page➡️</a></p>
</div>
