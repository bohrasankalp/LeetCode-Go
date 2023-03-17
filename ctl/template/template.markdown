# LeetCode in Go - English

### This repository is english version of [Leetcode-go](https://github.com/halfrost/LeetCode-Go).
Thanks to original creator @halfrost for all the hard work and dedication. 
- Milestone
  - [x] Translate english of main page
  - [ ] Check and validate all internal pages with pdf generation
  - [ ] Validate all links
  - [ ] Multi-language support in single repository for PDF
  - [ ] Adding more questions along with diagrams & pictures

___


[LeetCode Online Judge](https://leetcode.com/) is a website containing many **algorithm questions**. Most of them are
real interview questions of **Google, Facebook, LinkedIn, Apple**, etc. and it always help to sharp our algorithm
Skills. Level up your coding skills and quickly land a job. This is the best place to expand your knowledge and get
prepared for your next interview. This repo shows my solutions in Go with the code style strictly follows
the [Google Golang Style Guide](https://github.com/golang/go/wiki/CodeReviewComments). Please feel free to reference
and **STAR** to support this repo, thank you!

<p align='center'>
<img src='./logo.png'>
</p>

![](./website/static/wechat-qr-code.png)

<p align='center'>
<a href="https://github.com/halfrost/leetcode-go/releases/" rel="nofollow"><img alt="GitHub All Releases" src="https://img.shields.io/github/downloads/halfrost/LeetCode-Go/total?label=PDF%20downloads"></a>
<img src="https://img.shields.io/badge/Total%20Word%20Count-738884-success">
<a href="https://github.com/halfrost/leetcode-go/actions" rel="nofollow"><img src="https://github.com/halfrost/leetcode-go/workflows/Deploy%20leetcode-cookbook/badge.svg?branch=master"></a>
<a href="https://travis-ci.org/github/halfrost/LeetCode-Go" rel="nofollow"><img src="https://travis-ci.org/halfrost/LeetCode-Go.svg?branch=master"></a>
<a href="https://goreportcard.com/report/github.com/halfrost/LeetCode-Go" rel="nofollow"><img src="https://goreportcard.com/badge/github.com/halfrost/LeetCode-Go"></a>
<img src="https://img.shields.io/badge/runtime%20beats-100%25-success">
<a href="https://codecov.io/gh/halfrost/LeetCode-Go"><img src="https://codecov.io/gh/halfrost/LeetCode-Go/branch/master/graph/badge.svg" /></a>
<!--<img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/halfrost/LeetCode-Go?color=26C2F0">-->
<img alt="Support Go version" src="https://img.shields.io/badge/Go-v1.15-26C2F0">
<img src="https://visitor-badge.laobi.icu/badge?page_id=halfrost.LeetCode-Go">
</p>

<p align='center'>
<a href="https://github.com/halfrost/leetcode-go/blob/master/LICENSE"><img alt="GitHub" src="https://img.shields.io/github/license/halfrost/LeetCode-Go?label=License"></a>
<img src="https://img.shields.io/badge/License-CC-000000.svg">
<a href="https://leetcode.com/halfrost/"><img src="https://img.shields.io/badge/@halfrost-8751-yellow.svg">
<img src="https://img.shields.io/badge/language-Golang-26C2F0.svg">
<a href="https://halfrost.com"><img src="https://img.shields.io/badge/Blog-Halfrost--Field-80d4f9.svg?style=flat"></a>
<a href="http://weibo.com/halfrost"><img src="https://img.shields.io/badge/weibo-@halfrost-f974ce.svg?style=flat&colorA=f4292e"></a>
<a href="https://twitter.com/halffrost"><img src="https://img.shields.io/badge/twitter-@halffrost-F8E81C.svg?style=flat&colorA=009df2"></a>
<a href="https://www.zhihu.com/people/halfrost/activities"><img src="https://img.shields.io/badge/%E7%9F%A5%E4%B9%8E-@halfrost-fd6f32.svg?style=flat&colorA=0083ea"></a>
<img src="https://img.shields.io/badge/made%20with-=1-blue.svg">
<a href="https://github.com/halfrost/leetcode-go/pulls"><img src="https://img.shields.io/badge/PR-Welcome-brightgreen.svg"></a>
</p>

E-books that support Progressive Web Apps and Dark Mode《LeetCode Cookbook》 <a href="https://books.halfrost.com/leetcode/" rel="nofollow">Online Reading</a>

<p align='center'>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_Safari_0.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_Chrome_PWA.png"></a>
</p>

Offline version of the e-book "LeetCode Cookbook" PDF <a href="https://github.com/halfrost/leetcode-go/releases/" rel="nofollow">
Download here</a>

<p align='center'>
<a href="https://github.com/halfrost/leetcode-go/releases/"><img src="https://img.halfrost.com/Leetcode/Cookbook.png"></a>
</p>

Install the PWA version of "LeetCode Cookbook" to the desktop of the device through the iOS /Android browser to learn at any time

<p align='center'>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad_example1__.png"></a>
<a href="https://books.halfrost.com/leetcode/"><img src="https://img.halfrost.com/Leetcode/Cookbook_PWA_iPad_example2__.png"></a>
</p>

## Data Structures

> The topics marked with ✅ have completed all the topics, and those without the mark have not completed all the topics

<a href="https://books.halfrost.com/leetcode/"><img src="./website/static/logo.png" alt="logo" height="550" align="right" /></a>

- [LeetCode in Go - English](#leetcode-in-go---english)
    - [This repository is english version of Leetcode-go.](#this-repository-is-english-version-of-leetcode-go)
  - [Data Structures](#data-structures)
  - [Algorithm](#algorithm)
  - [LeetCode Problems](#leetcode-problems)
  - [1. Personal data](#1-personal-data)
  - [2. Directory](#2-directory)
  - [3. Classification](#3-classification)
  - [Array](#array)
  - [String](#string)
  - [Two Pointers](#two-pointers)
  - [Linked List](#linked-list)
  - [Stack](#stack)
  - [Tree](#tree)
  - [Dynamic Programming](#dynamic-programming)
  - [Backtracking](#backtracking)
  - [Depth First Search](#depth-first-search)
  - [Breadth First Search](#breadth-first-search)
  - [Binary Search](#binary-search)
  - [Math](#math)
  - [Hash Table](#hash-table)
  - [Sort](#sort)
  - [Bit Manipulation](#bit-manipulation)
  - [Union Find](#union-find)
  - [Sliding Window](#sliding-window)
  - [Segment Tree](#segment-tree)
  - [Binary Indexed Tree](#binary-indexed-tree)
  - [♥️ Thanks](#️-thanks)

<br>
<br>

|          data structure           | variant                                                                                                                                                                                                                                   | related topics | Explanatory article |
| :-------------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------- | :------------------ |
| Sequential Linear Tables: Vectors |                                                                                                                                                                                                                                           |                |                     |
|            Single list            | 1. Doubly linked list<br>2. Static linked list<br>3. Symmetric matrix<br>4. Sparse matrix                                                                                                                                                 |                |                     |
|            hash table             | 1. Hash functions<br>2. Solve collision/fill factor<br>                                                                                                                                                                                   |                |                     |
|          Stack and Queue          | 1. Generalized stack<br>2. Double-ended queue<br>                                                                                                                                                                                         |                |                     |
|               Queue               | 1. Linked list implementation<br>2. Circular array implementation<br>3. Double-ended queue                                                                                                                                                |                |                     |
|              String               | 1. KMP Algorithm<br>2. Finite State Automata<br>3. Pattern Matching Finite State Automata<br>4. BM Pattern Matching Algorithm<br>5. BM-KMP Algorithm<br>6 .BF Algorithm                                                                   |                |                     |
|               Tree                | 1. Binary tree<br>2. Union search<br>3. Huffman tree                                                                                                                                                                                      |                |                     |
|      Array implemented heap       | 1. Max heap and min heap<br>2. Max and min heap<br>3. Double-ended heap<br>4. d fork heap                                                                                                                                                 |                |                     |
|     Heap implemented by tree      | 1. Left heap<br>2. Flat heap<br>3. Binomial heap<br>4. Fibonacci heap<br>5. Pairing heap                                                                                                                                                  |                |                     |
|              Lookup               | 1. Hash table<br>2. Jump table<br>3. Sorted binary tree<br>4. AVL tree<br>5. B tree/B+ tree/B\*tree<br>6. AA tree <br>7. Red-black tree<br>8. Sorted binary heap<br>9. Splay tree<br>10. Double chain tree<br>11. Trie tree<br>12. R tree |

## Algorithm

|            Algorithms            | Concrete Types                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Related Topics                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Explained Articles   |
| :------------------------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :------------------- |
|        Sorting Algorithms        | 1. Bubble sort<br>2. Insertion sort<br>3. Selection sort<br>4. Shell sort<br>5. Quick sort<br>6. Merge sort<br>7 . Heap sort<br>8. Linear sorting algorithm<br>9. Introspective sorting<br>10. Indirect sorting<br>11. Counting sorting<br>12. Radix sorting<br>13. Bucket sorting<br>14 . External sorting -k-way merge loser tree<br>15. External sorting -best merge tree                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |                      |
| Recursion and divide and conquer |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | 1. Binary Search/Find<br>2. Multiplication of Large Integers<br>3. Strassen Matrix Multiplication<br>4. Chessboard Covering<br>5. Merge Sort<br>6. Quick Sort<br>7. Linear Time Selection<br>8. Closest Point Pair Problem<br>9. Round Robin Schedule<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |                      |
|       Dynamic Programming        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | 1. Matrix multiplication problem<br>2. Longest common subsequence<br>3. Maximum subsegment sum<br>4. Optimal triangulation of convex polygons<br>5. Polygon games<br>6. Image Compression<br>7. Circuit Routing<br>8. Pipeline Job Scheduling<br>9. 0-1 Knapsack Problem/Knapsack Nine Lectures<br>10. Optimal Binary Search Tree<br>11. Dynamic Programming Acceleration Principle<br>12. Tree DP<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |                      |
|              greedy              | 1. Activity scheduling problem<br>2. Optimal loading<br>3. Huffman coding<br>4. Single-source shortest path<br>5. Minimum spanning tree<br>6. Multi-machine scheduling problem<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|           Backtracking           | 1. Loading Problem<br>2. Batch Job Scheduling<br>3. Signed Triangle Problem<br>4. Post n Problem<br>5. 0-1 Knapsack Problem<br>6. Maximum Clique Problem<br> >7. The m-coloring problem of graphs<br>8. The traveling salesman problem<br>9. The circle permutation problem<br>10. The circuit board permutation problem<br>11. The continuous postage problem<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|              Search              | 1. Enumeration<br>2. DFS<br>3. BFS<br>4. Heuristic Search<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |                      |
|          Randomization           | 1. Random Number<br>2. Numerical Randomization Algorithm<br>3. Sherwood Algorithm<br>4. Las Vegas Algorithm<br>5. Monte Carlo Algorithm<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 1. Calculate the value of π<br>2. Calculate the definite integral<br>3. Solve Systems of nonlinear equations<br>4. Linear time selection algorithms<br>5. Skip lists<br>6. Post n problems<br>7. Integer factorization<br>8. Principal element problems<br>9. Prime numbers test<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |                      |
|                                  |
|           Graph Theory           | 1. Traverse DFS /BFS<br>2. AOV /AOE Network<br>3. Kruskal Algorithm (Minimum Spanning Tree)<br>4. Prim Algorithm (Minimum Spanning Tree)<br>5. Boruvka Algorithm ( Minimum spanning tree)<br>6. Dijkstra's algorithm (single-source shortest path)<br>7. Bellman-Ford algorithm (single-source shortest path)<br>8. SPFA algorithm (single-source shortest path)<br>9. Floyd's algorithm (multi-source shortest path)<br>10. Johnson's algorithm (multi-source shortest path)<br>11. Fleury's algorithm (Eulerian circuit)<br>12. Ford-Fulkerson algorithm (maximum network flow augmentation path) <br>13. Edmonds-Karp Algorithm (Maximum Network Flow)<br>14. Dinic Algorithm (Maximum Network Flow)<br>15. General Preflow Pushing Algorithm<br>16. Highest Label Preflow Pushing HLPP Algorithm<br> >17. Primal-Dual algorithm (minimum cost flow)18. Kosaraju algorithm (strongly connected components of directed graph)<br>19. Tarjan algorithm (strongly connected components of directed graph)<br>20. Gabow algorithm (with strong connected components to the graph)<br>21. Hungarian algorithm (bipartite graph matching)<br>22. Hopcroft－Karp algorithm (bipartite graph matching)<br>23. kuhn munkras algorithm (bipartite graph matching)<br> 24. Edmonds' Blossom-Contraction Algorithm (General Graph Matching)<br> | 1. Graph Traversal<br>2. Strong and Weak Connectivity of Directed and Undirected Graphs<br>3. Cutting Points/Edges<br> >3. AOV network and topological sorting<br>4. AOE network and critical path<br>5. Minimum cost spanning tree/second smallest spanning tree<br>6. Shortest path problem/Kth short circuit problem<br>7. Maximum Network Flow Problem<br>8. Minimum Cost Flow Problem<br>9. Graph Coloring Problem<br>10. Differential Constraint Systems<br>11. Euler Circuit<br>12. Chinese Postman Problem<br>13. Hamiltonian circuit<br>14. Best edge cut set/best point cut set/minimum edge cut set/minimum point cut set/minimum path cover/minimum point set cover<br>15. Edge cover set<br>16. Bipartite Graph Perfect Matching and Maximum Matching Problem<br>17. Cactus Graph<br>18. Chord Graph<br>19. Stable Marriage Problem<br>20. Maximum Clique Problem<br>                                                                                                                                                                                                                                    |                      |
|          Number Theory           | 1. Greatest common divisor<br> 2. Least common multiple<br>3. Decomposition of prime factors<br>4. Prime number determination<br>5. Base conversion<br>6. High precision calculation<br>7. Divisibility Problems<br>8. Congruence Problems<br>9. Euler Functions<br>10. Extended Euclidean<br>11. Permutation Groups<br>12. Generating Functions<br>13. Discrete Transformations<br> >14. Cantor Expansion<br>15. Matrices<br>16. Vectors<br>17. Linear Equations<br>18. Linear Programming<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|             Geometry             | 1. Convex Hull -Gift wrapping<br>2. Convex Hull -Graham scan<br>3. Line Segment Problems<br> 4. Polygon and Polyhedron Related Problems<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|           NP-Complete            | 1. Computational Models<br>2. P-type and NP-type Problems<br>3. NP-Complete Problems<br>4. Approximation Algorithms for NP-Complete Problems<br>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | 1. Random Access Machine RAM<br>2. Random Access Stored Program Machine RASP<br>3. Turing Machine<br>4. Nondeterministic Turing Machine<br>5. P and NP Class Languages<br> 6. Polynomial Time Verification<br>7. Polynomial Time Transformation<br>8. Cook's Theorem<br>9. The Satisfiability Problem of Conjunctive Normal Form CNF-SAT<br>10. The Satisfiability of 3-Variable Conjunctive Normal Form Problem 3-SAT<br>11. Clique Problem CLIQUE<br>12. Vertex Cover Problem VERTEX-COVER<br>13. Subset Sum Problem SUBSET-SUM<br>14. Hamiltonian Circuit Problem HAM-CYCLE<br> 15. Traveling Salesman Problem TSP<br>16. Approximate Algorithms for Vertex Covering Problems<br>17. Approximate Algorithms for Traveling Salesman Problems<br>18. Traveling Salesman Problems with Triangle Inequality Properties<br>19. General Traveling Salesman Problems <br>20. Approximation Algorithms for Set Covering Problems<br>21. Approximation Algorithms for Subset Sum Problems<br>22. Exponential Time Algorithms for Subset Sum Problems<br>23. Polynomial Time Approximation Form for Subset Sum Problems <br> |                      |
|           ------------           | ------------------------------------------------------------------                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | -----------------------------------------------------------------                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | -------------------- |

## LeetCode Problems

## 1. Personal data

{{.PersonalData}}

## 2. Directory

{{.TotalNum}}

{{.AvailableTable}}

---

The following are free algorithm questions, but they cannot be solved using Go yet:

no yet

---

## 3. Classification

## Array

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Array/)

## String

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/String/)

## Two Pointers

![](./topic/Two_pointers.png)

- The classic way of writing a two-pointer sliding window. The right pointer keeps moving to the right until it cannot move to the right (the specific conditions depend on the topic). When the right pointer reaches the far right, start to move the left pointer to release the left boundary of the window. No. Question 3, Question 76, Question 209, Question 424, Question 438, Question 567, Question 713, Question 763, Question 845, Question 881, Question 904, Question 978, Question Question 992, Question 1004, Question 1040, Question 1052.

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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Two_Pointers/)

## Linked List

![](./topic/Linked_List.png)

- Clever construction of virtual head nodes. The traversal processing logic can be made more unified.
- Flexible use of recursion. Construct recursive conditions, and use recursion to solve problems cleverly. However, it should be noted that some topics cannot use recursion, because too deep recursion will cause timeout and stack overflow.
- The reverse order of the range of the linked list. Question 92.
- The linked list finds the middle node. Question 876. The linked list finds the nth last node. Question 19. It only takes one traversal to get the answer.
- Merge K ordered linked lists. Question 21, question 23.
- Linked list sorting. Question 86, question 328.
- Linked list sorting requires O(n \*log n) time complexity and O(1) space complexity. There is only one method, merge sort, and merge from top to bottom. Question 148.
- Judging whether there is a ring in the linked list, if there is a ring, output the subscript of the intersection point of the ring; judge whether there is an intersection point in the two linked lists, and output the intersection point if there is an intersection point. Question 141, Question 142 question, question 160.

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Linked_List/)

## Stack

![](./topic/Stack.png)

-Bracket matching issues and similar. Question 20, question 921, question 1021.
-Basic pop and push operations on the stack. Question 71, Question 150, Question 155, Question 224, Question 225, Question 232, Question 946, Question 1047.
-Coding problems using the stack. Question 394, Question 682, Question 856, Question 880. -**Monotonic stack**. **Use the stack to maintain a monotonically increasing or decreasing subscript array**. Question 84, Question 456, Question 496, Question 503, Question 739, Question 901, Question
Question 907, question 1019.
Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Stack/)

## Tree

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Tree/)

## Dynamic Programming

- Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Dynamic_Programming/)

## Backtracking

![](./topic/Backtracking.png)

- Permutations of permutation questions. Question 46, question 47. Question 60, Question 526, Question 996.
- Combination question Combination. Question 39, Question 40, Question 77, Question 216.
- Permutation and combination hybridization problems. Question 1079.
- N queens ultimate solution (binary solution). Question 51, question 52.
- Sudoku problems. Question 37.
- Search in four directions. Question 79, Question 212, Question 980.
- Subcollection issues. Question 78, question 90.
- Trie. Question 208, question 211.
- BFS optimization. Question 126, question 127.
- DFS templates. (just an example, does not correspond to any question)

```go
func combinationSum2(candidates []int, target int) [][]int {
if len(candidates) == 0 {
return [][]int{}
}
c, res := []int{}, [][]int{}
sort.Ints(candidates)
findcombinationSum2(candidates, target, 0, c, &res)
return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
if target == 0 {
b := make([]int, len(c))
copy(b, c)
*res = append(*res, b)
return
}
for i := index; i < len(nums); i++ {
if i > index && nums[i] == nums[i-1] { //Here is the key logic for deduplication
continue
}
if target >= nums[i] {
c = append(c, nums[i])
findcombinationSum2(nums, target-nums[i], i+1, c, res)
c = c[:len(c)-1]
}
}
}
```

- BFS template. (just an example, does not correspond to any question)

```go
package main
func main(){

}
func updateMatrix_BFS(matrix [][]int) [][]int {
res := make([][]int, len(matrix))
if len(matrix) == 0 || len(matrix[0]) == 0 {
return res
}
queue := make([][]int, 0)
for i, _ := range matrix {
res[i] = make([]int, len(matrix[0]))
for j, _ := range res[i] {
if matrix[i][j] == 0 {
res[i][j] = -1
queue = append(queue, []int{i, j})
}
}
}
level := 1
for len(queue) > 0 {
size := len(queue)
for size > 0 {
size -= 1
node := queue[0]
queue = queue[1:]
i, j := node[0], node[1]
for _, direction := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
x := i + direction[0]
y := j + direction[1]
if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || res[x][y] < 0 || res[x][y] > 0 {
continue
}
res[x][y] = level
queue = append(queue, []int{x, y})
}
}
level++
}
for i, row := range res {
for j, cell := range row {
if cell == -1 {
res[i][j] = 0
}
}
}
return res
}
```

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Backtracking/)

## Depth First Search

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Depth_First_Search/)

## Breadth First Search

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Breadth_First_Search/)

## Binary Search

![]()

-The classic way of writing binary search. Three points to note:

1. Loop exit condition, note that low <= high, not low < high.
2. The value of mid, mid := low + (high-low)>>1
3. Update of low and high. low = mid + 1, high = mid -1.

```go
func binarySearchMatrix(nums []int, target int) int {
low, high := 0, len(nums)-1
for low <= high {
mid := low + (high-low)>>1
if nums[mid] == target {
return mid
} else if nums[mid] > target {
high = mid - 1
} else {
low = mid + 1
}
}
return -1
}
```

-A variant of binary search. There are 4 basic variants:

1. Find the first element equal to target, time complexity O(logn)
2. Find the last element equal to target, time complexity O(logn)
3. Find the first element greater than or equal to target, time complexity O(logn)
4. Find the last element less than or equal to target, time complexity O(logn)

```go
//binary search for the first element equal to target, time complexity O(logn)
func searchFirstEqualElement(nums []int, target int) int {
low, high := 0, len(nums)-1
for low <= high {
mid := low + ((high - low) >> 1)
if nums[mid] > target {
high = mid - 1
} else if nums[mid] < target {
low = mid + 1
} else {
if (mid == 0) || (nums[mid-1] != target) { //Find the first element equal to target
return mid
}
high = mid - 1
}
}
return -1
}

//binary search for the last element equal to target, time complexity O(logn)
func searchLastEqualElement(nums []int, target int) int {
low, high := 0, len(nums)-1
for low <= high {
mid := low + ((high - low) >> 1)
if nums[mid] > target {
high = mid - 1
} else if nums[mid] < target {
low = mid + 1
} else {
if (mid == len(nums)-1) || (nums[mid+1] != target) { //Find the last element equal to target
return mid
}
low = mid + 1
}
}
return -1
}

//binary search for the first element greater than or equal to target, time complexity O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
low, high := 0, len(nums)-1
for low <= high {
mid := low + ((high - low) >> 1)
if nums[mid] >= target {
if (mid == 0) || (nums[mid-1] < target) { //Find the first element greater than or equal to target
return mid
}
high = mid - 1
} else {
low = mid + 1
}
}
return -1
}

//binary search for the last element less than or equal to target, time complexity O(logn)
func searchLastLessElement(nums []int, target int) int {
low, high := 0, len(nums)-1
for low <= high {
mid := low + ((high - low) >> 1)
if nums[mid] <= target {
if (mid == len(nums)-1) || (nums[mid+1] > target) { //Find the last element less than or equal to target
return mid
}
low = mid + 1
} else {
high = mid - 1
}
}
return -1
}
```

- Binary search in basically sorted arrays. The classic solution can be solved, and the variant writing method can also be written. The common question type is to find the peak in the peak array and find the dividing point in the rotated ordered array. No.
  Question 33, Question 81, Question 153, Question 154, Question 162, Question 852

```go
package main

func peakIndexInMountainArray(A []int) int {
low, high := 0, len(A)-1
for low < high {
mid := low + (high-low)>>1
//If mid is large, there is a peak on the left, high = m, if mid + 1 is large, there is a peak on the right, low = mid + 1
if A[mid] > A[mid+1] {
high = mid
} else {
low = mid + 1
}
}
return low
}
```

- max-min maximum minimization problem. Find the maximum value if the minimum condition is met. Question 410, Question 875, Question 1011, Question 1283.
  Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Binary_Search/)

## Math

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Math/)

## Hash Table

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Hash_Table/)

## Sort

![](./topic/Sort.png)

- Deep understanding of multi-way quick sort. Question 75.
- Sorting of linked lists, insertion sort (question 147) and merge sort (question 148)
- Bucket sort and radix sort. Question 164.
- "Swing Sort". Question 324.
- Sorting that is not adjacent to each other. Question 767, question 1054.
- "Sorting pancakes". Question 969.

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Sort/)

## Bit Manipulation

![](./topic/Bit_Manipulation.png)

-XOR property. Question 136, Question 268, Question 389, Question 421,

```go
x ^ 0 = x
x ^ 11111……1111 = ~x
x ^ (~x) = 11111……1111
x ^ x = 0
a ^ b = c = > a ^ c = b = > b ^ c = a (commutative law)
a ^ b ^ c = a ^ (b ^ c) = (a ^ b）^ c (associative law)
```

- Construct a special Mask, put 0 or 1 in the special position.

```go
Clear the rightmost n bits of x , x & ( ~0 << n )
Get the nth bit value (0 or 1) of x, (x >> n) & 1
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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Bit_Manipulation/)

## Union Find

![](./topic/Union_Find.png)

- Flexibly use the idea of ​​union find, master the [template] of union find (https://github.com/halfrost/leetcode-go/blob/master/template/UnionFind.go)
  , there are two implementations of union search in the template, one is the version of path compression + rank optimization, and the other is to calculate the number of elements in each set +
  The version with the maximum number of collection elements, both versions have their respective uses. The questions that can use the first type of union search template are: No. 128, No. 130, No. 547, No.
  Question 684, Question 721, Question 765, Question 778, Question 839, Question 924, Question 928, Question 947, Question 952, Question 959, Question 990
  question. The questions that can use the second type of union search template are: question 803 and question 952. Question 803 Rank optimization and the number of statistical collections will be time-consuming. If not optimized, it will be TLE.
- Union search is a kind of thinking, some questions need to use this kind of thinking flexibly, instead of a dead set template, such as question 399, this question is
  stringUnionFind, realized by using the idea of ​​union search. Each node here is based on strings and maps, rather than simply using int node numbers.
- Some questions can’t be done with a dead set template, such as question 685. This question cannot be path compressed and rank optimized, because the question involves a directed graph, and it is necessary to know the predecessor node of the node. If the path is compressed, this question There is no way to do it. This question does not require path compression and rank optimization.
- The information given by the flexible abstract topic, number the given information reasonably, use and search to solve the problem, and use map to reduce the time complexity, such as the 721st question and the 959th question.
- Regarding the topic of maps, bricks, and grids, you can create a special node, and union() the bricks or grids around the edges to this special node. Question 130, question 803.
- Questions that can be used for union search can generally be solved with DFS and BFS, but the time complexity will be higher.
  Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Union_Find/)

## Sliding Window

![](./topic/Sliding_Window.png)

- The classic way of writing a two-pointer sliding window. The right pointer keeps moving to the right until it cannot move to the right (the specific conditions depend on the topic). When the right pointer reaches the far right, start to move the left pointer to release the left boundary of the window. No.
  Question 3, Question 76, Question 209, Question 424, Question 438, Question 567, Question 713, Question 763, Question 845, Question 881, Question 904, Question 978, Question
  Question 992, Question 1004, Question 1040, Question 1052.

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

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Sliding_Window/)

## Segment Tree

![](./topic/Segment_Tree.png)

- The classic array implementation of the line segment tree. The pushUp logic of merging two nodes is abstracted, and any operation can be realized (common operations include: addition, max, min, etc.). No.
  Question 218, Question 303, Question 307, Question 699.
- The classic way of counting line segment trees. Question 315, Question 327, Question 493.
- The implementation of the tree of the line segment tree. Question 715, question 732.
- Range lazy update. Question 218, question 699.
- Discretization. Discretization needs to pay attention to a special case: if the three intervals are [1,10] [1,4] [6,10], after discretization x[1]=1,x[2]=4,x[3] =6,x[4]
  =10. The first interval is [1,4], the second interval is [1,2], and the third interval is [3,4], so that interval one = interval two +
  Interval three, which is inconsistent with the model before discretization. Before discretization, it is obvious that interval one > interval two + interval three. The correct way is: add a number between the numbers with a difference greater than 1, such as the above
  1 4 6 10 Add 5 in the middle, then x[1]=1,x[2]=4,x[3]=5,x[4]=6,x[5]=10. After this processing, interval one is 1-5, interval two is 1-2, and interval three is 4-5.
- Flexible construction of segment trees. Line segment tree nodes can store multiple pieces of information, and the pushUp operation for merging two nodes can also be diverse. Question 850, question 1157.

Line segment tree [question type] (https://blog.csdn.net/xuechelingxiao/article/details/38313105) from simple to difficult:

1. Single point update:
   - [HDU 1166 Enemy Soldier Arrangement](http://acm.hdu.edu.cn/showproblem.php?pid=1166) update: single-point increase and decrease query: interval summation
   - [HDU 1754 I Hate It](http://acm.hdu.edu.cn/showproblem.php?pid=1754) update: single point replacement query: maximum value of interval
   - [HDU 1394 Minimum Inversion Number](http://acm.hdu.edu.cn/showproblem.php?pid=1394) update: single-point increase and decrease query: interval summation
   - [HDU 2795 Billboard](http://acm.hdu.edu.cn/showproblem.php?pid=2795) query: Find the maximum value of the interval (
     Do the update operation directly in the query)
2. Interval update:
   - [HDU 1698 Just a Hook](http://acm.hdu.edu.cn/showproblem.php?pid=1698) update: segment replacement (since the total interval is only queried once, it can be output directly
     1 node information)
   - [POJ 3468 A Simple Problem with Integers](http://poj.org/problem?id=3468) update: Segment increase and decrease query: Interval summation
   - [POJ 2528 Mayor’s posters](http://poj.org/problem?id=2528) discretization + update: segment replacement query: simple hash
   - [POJ 3225 Help with Intervals](http://poj.org/problem?id=3225) update: segment replacement, interval XOR query: simple hash
3. Interval merging (this type of question will ask the longest continuous interval that meets the conditions in the interval, so when PushUp needs to merge the intervals of the left and right sons):
   [POJ 3667 Hotel](http://poj.org/problem?id=3667) update: interval replacement query: query the leftmost endpoint that meets the conditions
4. Scanning line (this kind of problem needs to sort some operations, and then use a scanning line to scan from left to right. The most typical one is rectangular area union, perimeter length union and other questions):
   - [HDU 1542 Atlantis](http://acm.hdu.edu.cn/showproblem.php?pid=1542) update: Interval increase and decrease query: Take the value of the root node directly
   - [HDU 1828 Picture](http://acm.hdu.edu.cn/showproblem.php?pid=1828) update: interval increase and decrease query: take the value of the root node directly

Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Segment_Tree/)

## Binary Indexed Tree

![](./topic/Binary_Indexed_Tree.png)
Problems List in [there](https://books.halfrost.com/leetcode/ChapterTwo/Binary_Indexed_Tree/)

---

<p align='center'>
<a href="https://github.com/halfrost/leetcode-go/releases/tag/Special"><img src="https://img.halfrost.com/Leetcode/ACM-ICPC_Algorithm_Template.png"></a>
</p>

Thank you for reading here. This is bonus. You can download my [《ACM-ICPC Algorithm Template》](https://github.com/halfrost/leetcode-go/releases/tag/Special/)

## ♥️ Thanks

Thanks for your Star！

[![Stargazers over time](https://starchart.cc/halfrost/LeetCode-Go.svg)](https://starchart.cc/halfrost/LeetCode-Go)
