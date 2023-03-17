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


{{.AvailableTagTable}}