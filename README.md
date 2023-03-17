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

|    |  Easy  |  Medium  |  Hard |  Total |
|:--------:|:--------:|:--------:|:--------:|:--------:|
|Optimizing|0|0|0|0|
|Accepted|**0**|**0**|**0**|**0**|
|Total|634|1377|579|2590|
|Perfection Rate|NaN%|NaN%|NaN%|NaN%|
|Completion Rate|0.0%|0.0%|0.0%|0.0%|
|------------|----------------------------|----------------------------|----------------------------|----------------------------|

## 2. Directory

The solutions of 787 questions have been collected below, and there are still 8 questions trying to optimize to beats 100%

| No.    |  Title  |  Solution  |  Acceptance |  Difficulty |  Frequency |
|:--------:|:--------------------------------------------------------------|:--------:|:--------:|:--------:|:--------:|
|0001|Two Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0001.Two-Sum)|49.6%|Easy||
|0002|Add Two Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0002.Add-Two-Numbers)|40.3%|Medium||
|0003|Longest Substring Without Repeating Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0003.Longest-Substring-Without-Repeating-Characters)|33.8%|Medium||
|0004|Median of Two Sorted Arrays|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0004.Median-of-Two-Sorted-Arrays)|36.1%|Hard||
|0005|Longest Palindromic Substring|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0005.Longest-Palindromic-Substring)|32.4%|Medium||
|0006|Zigzag Conversion|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0006.Zigzag-Conversion)|44.8%|Medium||
|0007|Reverse Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0007.Reverse-Integer)|27.4%|Medium||
|0008|String to Integer (atoi)|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0008.String-to-Integer-atoi)|16.6%|Medium||
|0009|Palindrome Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0009.Palindrome-Number)|53.5%|Easy||
|0010|Regular Expression Matching||28.0%|Hard||
|0011|Container With Most Water|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0011.Container-With-Most-Water)|54.0%|Medium||
|0012|Integer to Roman|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0012.Integer-to-Roman)|61.9%|Medium||
|0013|Roman to Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0013.Roman-to-Integer)|58.5%|Easy||
|0014|Longest Common Prefix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0014.Longest-Common-Prefix)|40.8%|Easy||
|0015|3Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0015.3Sum)|32.5%|Medium||
|0016|3Sum Closest|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0016.3Sum-Closest)|45.8%|Medium||
|0017|Letter Combinations of a Phone Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0017.Letter-Combinations-of-a-Phone-Number)|56.4%|Medium||
|0018|4Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0018.4Sum)|36.0%|Medium||
|0019|Remove Nth Node From End of List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0019.Remove-Nth-Node-From-End-of-List)|40.9%|Medium||
|0020|Valid Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0020.Valid-Parentheses)|40.3%|Easy||
|0021|Merge Two Sorted Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0021.Merge-Two-Sorted-Lists)|62.4%|Easy||
|0022|Generate Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0022.Generate-Parentheses)|72.4%|Medium||
|0023|Merge k Sorted Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0023.Merge-k-Sorted-Lists)|49.7%|Hard||
|0024|Swap Nodes in Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0024.Swap-Nodes-in-Pairs)|61.2%|Medium||
|0025|Reverse Nodes in k-Group|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0025.Reverse-Nodes-in-k-Group)|54.5%|Hard||
|0026|Remove Duplicates from Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0026.Remove-Duplicates-from-Sorted-Array)|51.5%|Easy||
|0027|Remove Element|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0027.Remove-Element)|52.9%|Easy||
|0028|Find the Index of the First Occurrence in a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0028.Find-the-Index-of-the-First-Occurrence-in-a-String)|38.9%|Easy||
|0029|Divide Two Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0029.Divide-Two-Integers)|17.2%|Medium||
|0030|Substring with Concatenation of All Words|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0030.Substring-with-Concatenation-of-All-Words)|31.2%|Hard||
|0031|Next Permutation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0031.Next-Permutation)|37.5%|Medium||
|0032|Longest Valid Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0032.Longest-Valid-Parentheses)|32.8%|Hard||
|0033|Search in Rotated Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0033.Search-in-Rotated-Sorted-Array)|38.9%|Medium||
|0034|Find First and Last Position of Element in Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0034.Find-First-and-Last-Position-of-Element-in-Sorted-Array)|41.8%|Medium||
|0035|Search Insert Position|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0035.Search-Insert-Position)|43.3%|Easy||
|0036|Valid Sudoku|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0036.Valid-Sudoku)|58.0%|Medium||
|0037|Sudoku Solver|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0037.Sudoku-Solver)|57.6%|Hard||
|0038|Count and Say||52.0%|Medium||
|0039|Combination Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0039.Combination-Sum)|68.5%|Medium||
|0040|Combination Sum II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0040.Combination-Sum-II)|53.4%|Medium||
|0041|First Missing Positive|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0041.First-Missing-Positive)|36.7%|Hard||
|0042|Trapping Rain Water|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0042.Trapping-Rain-Water)|59.2%|Hard||
|0043|Multiply Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0043.Multiply-Strings)|39.1%|Medium||
|0044|Wildcard Matching||26.9%|Hard||
|0045|Jump Game II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0045.Jump-Game-II)|39.8%|Medium||
|0046|Permutations|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0046.Permutations)|75.6%|Medium||
|0047|Permutations II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0047.Permutations-II)|57.3%|Medium||
|0048|Rotate Image|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0048.Rotate-Image)|70.9%|Medium||
|0049|Group Anagrams|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0049.Group-Anagrams)|66.7%|Medium||
|0050|Pow(x, n)|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0050.Powx-n)|33.0%|Medium||
|0051|N-Queens|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0051.N-Queens)|64.0%|Hard||
|0052|N-Queens II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0052.N-Queens-II)|71.6%|Hard||
|0053|Maximum Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0053.Maximum-Subarray)|50.2%|Medium||
|0054|Spiral Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0054.Spiral-Matrix)|44.8%|Medium||
|0055|Jump Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0055.Jump-Game)|38.9%|Medium||
|0056|Merge Intervals|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0056.Merge-Intervals)|46.2%|Medium||
|0057|Insert Interval|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0057.Insert-Interval)|39.0%|Medium||
|0058|Length of Last Word|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0058.Length-of-Last-Word)|42.6%|Easy||
|0059|Spiral Matrix II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0059.Spiral-Matrix-II)|67.3%|Medium||
|0060|Permutation Sequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0060.Permutation-Sequence)|44.3%|Hard||
|0061|Rotate List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0061.Rotate-List)|36.0%|Medium||
|0062|Unique Paths|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0062.Unique-Paths)|62.6%|Medium||
|0063|Unique Paths II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0063.Unique-Paths-II)|39.4%|Medium||
|0064|Minimum Path Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0064.Minimum-Path-Sum)|61.1%|Medium||
|0065|Valid Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0065.Valid-Number)|18.7%|Hard||
|0066|Plus One|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0066.Plus-One)|43.6%|Easy||
|0067|Add Binary|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0067.Add-Binary)|52.4%|Easy||
|0068|Text Justification||37.4%|Hard||
|0069|Sqrt(x)|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0069.Sqrtx)|37.4%|Easy||
|0070|Climbing Stairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0070.Climbing-Stairs)|52.2%|Easy||
|0071|Simplify Path|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0071.Simplify-Path)|39.3%|Medium||
|0072|Edit Distance||54.3%|Hard||
|0073|Set Matrix Zeroes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0073.Set-Matrix-Zeroes)|51.1%|Medium||
|0074|Search a 2D Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0074.Search-a-2D-Matrix)|47.6%|Medium||
|0075|Sort Colors|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0075.Sort-Colors)|58.4%|Medium||
|0076|Minimum Window Substring|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0076.Minimum-Window-Substring)|40.9%|Hard||
|0077|Combinations|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0077.Combinations)|66.9%|Medium||
|0078|Subsets|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0078.Subsets)|74.8%|Medium||
|0079|Word Search|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0079.Word-Search)|40.2%|Medium||
|0080|Remove Duplicates from Sorted Array II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0080.Remove-Duplicates-from-Sorted-Array-II)|52.2%|Medium||
|0081|Search in Rotated Sorted Array II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0081.Search-in-Rotated-Sorted-Array-II)|35.7%|Medium||
|0082|Remove Duplicates from Sorted List II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0082.Remove-Duplicates-from-Sorted-List-II)|45.9%|Medium||
|0083|Remove Duplicates from Sorted List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0083.Remove-Duplicates-from-Sorted-List)|50.5%|Easy||
|0084|Largest Rectangle in Histogram|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0084.Largest-Rectangle-in-Histogram)|42.5%|Hard||
|0085|Maximal Rectangle||44.7%|Hard||
|0086|Partition List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0086.Partition-List)|51.9%|Medium||
|0087|Scramble String||36.1%|Hard||
|0088|Merge Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0088.Merge-Sorted-Array)|46.5%|Easy||
|0089|Gray Code|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0089.Gray-Code)|57.1%|Medium||
|0090|Subsets II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0090.Subsets-II)|55.8%|Medium||
|0091|Decode Ways|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0091.Decode-Ways)|32.7%|Medium||
|0092|Reverse Linked List II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0092.Reverse-Linked-List-II)|45.4%|Medium||
|0093|Restore IP Addresses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0093.Restore-IP-Addresses)|47.3%|Medium||
|0094|Binary Tree Inorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0094.Binary-Tree-Inorder-Traversal)|73.7%|Easy||
|0095|Unique Binary Search Trees II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0095.Unique-Binary-Search-Trees-II)|52.3%|Medium||
|0096|Unique Binary Search Trees|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0096.Unique-Binary-Search-Trees)|59.6%|Medium||
|0097|Interleaving String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0097.Interleaving-String)|37.3%|Medium||
|0098|Validate Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0098.Validate-Binary-Search-Tree)|32.0%|Medium||
|0099|Recover Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0099.Recover-Binary-Search-Tree)|50.9%|Medium||
|0100|Same Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0100.Same-Tree)|58.0%|Easy||
|0101|Symmetric Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0101.Symmetric-Tree)|54.2%|Easy||
|0102|Binary Tree Level Order Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0102.Binary-Tree-Level-Order-Traversal)|64.2%|Medium||
|0103|Binary Tree Zigzag Level Order Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0103.Binary-Tree-Zigzag-Level-Order-Traversal)|56.8%|Medium||
|0104|Maximum Depth of Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0104.Maximum-Depth-of-Binary-Tree)|73.9%|Easy||
|0105|Construct Binary Tree from Preorder and Inorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0105.Construct-Binary-Tree-from-Preorder-and-Inorder-Traversal)|61.4%|Medium||
|0106|Construct Binary Tree from Inorder and Postorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0106.Construct-Binary-Tree-from-Inorder-and-Postorder-Traversal)|58.1%|Medium||
|0107|Binary Tree Level Order Traversal II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0107.Binary-Tree-Level-Order-Traversal-II)|61.0%|Medium||
|0108|Convert Sorted Array to Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0108.Convert-Sorted-Array-to-Binary-Search-Tree)|69.7%|Easy||
|0109|Convert Sorted List to Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0109.Convert-Sorted-List-to-Binary-Search-Tree)|60.1%|Medium||
|0110|Balanced Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0110.Balanced-Binary-Tree)|48.9%|Easy||
|0111|Minimum Depth of Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0111.Minimum-Depth-of-Binary-Tree)|44.3%|Easy||
|0112|Path Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0112.Path-Sum)|48.2%|Easy||
|0113|Path Sum II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0113.Path-Sum-II)|57.0%|Medium||
|0114|Flatten Binary Tree to Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0114.Flatten-Binary-Tree-to-Linked-List)|61.7%|Medium||
|0115|Distinct Subsequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0115.Distinct-Subsequences)|44.4%|Hard||
|0116|Populating Next Right Pointers in Each Node|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0116.Populating-Next-Right-Pointers-in-Each-Node)|60.3%|Medium||
|0117|Populating Next Right Pointers in Each Node II||50.2%|Medium||
|0118|Pascal's Triangle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0118.Pascals-Triangle)|70.6%|Easy||
|0119|Pascal's Triangle II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0119.Pascals-Triangle-II)|60.6%|Easy||
|0120|Triangle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0120.Triangle)|54.4%|Medium||
|0121|Best Time to Buy and Sell Stock|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0121.Best-Time-to-Buy-and-Sell-Stock)|54.3%|Easy||
|0122|Best Time to Buy and Sell Stock II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0122.Best-Time-to-Buy-and-Sell-Stock-II)|63.8%|Medium||
|0123|Best Time to Buy and Sell Stock III||45.4%|Hard||
|0124|Binary Tree Maximum Path Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0124.Binary-Tree-Maximum-Path-Sum)|39.2%|Hard||
|0125|Valid Palindrome|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0125.Valid-Palindrome)|44.3%|Easy||
|0126|Word Ladder II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0126.Word-Ladder-II)|27.5%|Hard||
|0127|Word Ladder|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0127.Word-Ladder)|37.1%|Hard||
|0128|Longest Consecutive Sequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0128.Longest-Consecutive-Sequence)|48.6%|Medium||
|0129|Sum Root to Leaf Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0129.Sum-Root-to-Leaf-Numbers)|60.9%|Medium||
|0130|Surrounded Regions|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0130.Surrounded-Regions)|36.6%|Medium||
|0131|Palindrome Partitioning|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0131.Palindrome-Partitioning)|64.7%|Medium||
|0132|Palindrome Partitioning II||33.8%|Hard||
|0133|Clone Graph||51.8%|Medium||
|0134|Gas Station||46.1%|Medium||
|0135|Candy|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0135.Candy)|41.0%|Hard||
|0136|Single Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0136.Single-Number)|70.6%|Easy||
|0137|Single Number II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0137.Single-Number-II)|58.4%|Medium||
|0138|Copy List with Random Pointer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0138.Copy-List-with-Random-Pointer)|51.3%|Medium||
|0139|Word Break||45.6%|Medium||
|0140|Word Break II||45.2%|Hard||
|0141|Linked List Cycle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0141.Linked-List-Cycle)|47.4%|Easy||
|0142|Linked List Cycle II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0142.Linked-List-Cycle-II)|48.6%|Medium||
|0143|Reorder List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0143.Reorder-List)|52.4%|Medium||
|0144|Binary Tree Preorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0144.Binary-Tree-Preorder-Traversal)|66.7%|Easy||
|0145|Binary Tree Postorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0145.Binary-Tree-Postorder-Traversal)|67.8%|Easy||
|0146|LRU Cache|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0146.LRU-Cache)|40.6%|Medium||
|0147|Insertion Sort List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0147.Insertion-Sort-List)|50.9%|Medium||
|0148|Sort List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0148.Sort-List)|55.0%|Medium||
|0149|Max Points on a Line||25.1%|Hard||
|0150|Evaluate Reverse Polish Notation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0150.Evaluate-Reverse-Polish-Notation)|45.7%|Medium||
|0151|Reverse Words in a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0151.Reverse-Words-in-a-String)|32.6%|Medium||
|0152|Maximum Product Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0152.Maximum-Product-Subarray)|34.9%|Medium||
|0153|Find Minimum in Rotated Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0153.Find-Minimum-in-Rotated-Sorted-Array)|48.8%|Medium||
|0154|Find Minimum in Rotated Sorted Array II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0154.Find-Minimum-in-Rotated-Sorted-Array-II)|43.5%|Hard||
|0155|Min Stack|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0155.Min-Stack)|52.3%|Medium||
|0156|Binary Tree Upside Down||61.9%|Medium||
|0157|Read N Characters Given Read4||40.9%|Easy||
|0158|Read N Characters Given read4 II - Call Multiple Times||41.6%|Hard||
|0159|Longest Substring with At Most Two Distinct Characters||53.8%|Medium||
|0160|Intersection of Two Linked Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0160.Intersection-of-Two-Linked-Lists)|54.2%|Easy||
|0161|One Edit Distance||34.1%|Medium||
|0162|Find Peak Element|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0162.Find-Peak-Element)|46.0%|Medium||
|0163|Missing Ranges||32.2%|Easy||
|0164|Maximum Gap|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0164.Maximum-Gap)|43.3%|Hard||
|0165|Compare Version Numbers||35.7%|Medium||
|0166|Fraction to Recurring Decimal||24.3%|Medium||
|0167|Two Sum II - Input Array Is Sorted|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0167.Two-Sum-II-Input-Array-Is-Sorted)|60.0%|Medium||
|0168|Excel Sheet Column Title|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0168.Excel-Sheet-Column-Title)|35.4%|Easy||
|0169|Majority Element|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0169.Majority-Element)|63.9%|Easy||
|0170|Two Sum III - Data structure design||37.4%|Easy||
|0171|Excel Sheet Column Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0171.Excel-Sheet-Column-Number)|62.0%|Easy||
|0172|Factorial Trailing Zeroes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0172.Factorial-Trailing-Zeroes)|42.1%|Medium||
|0173|Binary Search Tree Iterator|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0173.Binary-Search-Tree-Iterator)|69.7%|Medium||
|0174|Dungeon Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0174.Dungeon-Game)|37.5%|Hard||
|0175|Combine Two Tables||74.0%|Easy||
|0176|Second Highest Salary||37.5%|Medium||
|0177|Nth Highest Salary||37.6%|Medium||
|0178|Rank Scores||60.4%|Medium||
|0179|Largest Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0179.Largest-Number)|34.5%|Medium||
|0180|Consecutive Numbers||46.6%|Medium||
|0181|Employees Earning More Than Their Managers||69.0%|Easy||
|0182|Duplicate Emails||70.8%|Easy||
|0183|Customers Who Never Order||68.5%|Easy||
|0184|Department Highest Salary||50.0%|Medium||
|0185|Department Top Three Salaries||50.4%|Hard||
|0186|Reverse Words in a String II||52.8%|Medium||
|0187|Repeated DNA Sequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0187.Repeated-DNA-Sequences)|46.9%|Medium||
|0188|Best Time to Buy and Sell Stock IV||38.8%|Hard||
|0189|Rotate Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0189.Rotate-Array)|39.4%|Medium||
|0190|Reverse Bits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0190.Reverse-Bits)|53.8%|Easy||
|0191|Number of 1 Bits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0191.Number-of-1-Bits)|66.4%|Easy||
|0192|Word Frequency||25.6%|Medium||
|0193|Valid Phone Numbers||26.0%|Easy||
|0194|Transpose File||25.4%|Medium||
|0195|Tenth Line||32.9%|Easy||
|0196|Delete Duplicate Emails||60.5%|Easy||
|0197|Rising Temperature||44.7%|Easy||
|0198|House Robber|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0198.House-Robber)|49.4%|Medium||
|0199|Binary Tree Right Side View|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0199.Binary-Tree-Right-Side-View)|61.6%|Medium||
|0200|Number of Islands|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0200.Number-of-Islands)|56.9%|Medium||
|0201|Bitwise AND of Numbers Range|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0201.Bitwise-AND-of-Numbers-Range)|42.5%|Medium||
|0202|Happy Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0202.Happy-Number)|54.8%|Easy||
|0203|Remove Linked List Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0203.Remove-Linked-List-Elements)|45.8%|Easy||
|0204|Count Primes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0204.Count-Primes)|33.1%|Medium||
|0205|Isomorphic Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0205.Isomorphic-Strings)|42.8%|Easy||
|0206|Reverse Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0206.Reverse-Linked-List)|73.4%|Easy||
|0207|Course Schedule|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0207.Course-Schedule)|45.4%|Medium||
|0208|Implement Trie (Prefix Tree)|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0208.Implement-Trie-Prefix-Tree)|61.7%|Medium||
|0209|Minimum Size Subarray Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0209.Minimum-Size-Subarray-Sum)|44.9%|Medium||
|0210|Course Schedule II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0210.Course-Schedule-II)|48.4%|Medium||
|0211|Design Add and Search Words Data Structure|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0211.Design-Add-and-Search-Words-Data-Structure)|43.0%|Medium||
|0212|Word Search II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0212.Word-Search-II)|36.5%|Hard||
|0213|House Robber II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0213.House-Robber-II)|41.0%|Medium||
|0214|Shortest Palindrome||32.3%|Hard||
|0215|Kth Largest Element in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0215.Kth-Largest-Element-in-an-Array)|66.1%|Medium||
|0216|Combination Sum III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0216.Combination-Sum-III)|67.6%|Medium||
|0217|Contains Duplicate|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0217.Contains-Duplicate)|61.4%|Easy||
|0218|The Skyline Problem|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0218.The-Skyline-Problem)|41.8%|Hard||
|0219|Contains Duplicate II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0219.Contains-Duplicate-II)|42.5%|Easy||
|0220|Contains Duplicate III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0220.Contains-Duplicate-III)|22.1%|Hard||
|0221|Maximal Square||44.9%|Medium||
|0222|Count Complete Tree Nodes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0222.Count-Complete-Tree-Nodes)|60.4%|Medium||
|0223|Rectangle Area|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0223.Rectangle-Area)|45.1%|Medium||
|0224|Basic Calculator|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0224.Basic-Calculator)|42.4%|Hard||
|0225|Implement Stack using Queues|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0225.Implement-Stack-using-Queues)|58.5%|Easy||
|0226|Invert Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0226.Invert-Binary-Tree)|74.6%|Easy||
|0227|Basic Calculator II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0227.Basic-Calculator-II)|42.4%|Medium||
|0228|Summary Ranges|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0228.Summary-Ranges)|47.2%|Easy||
|0229|Majority Element II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0229.Majority-Element-II)|45.0%|Medium||
|0230|Kth Smallest Element in a BST|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0230.Kth-Smallest-Element-in-a-BST)|70.0%|Medium||
|0231|Power of Two|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0231.Power-of-Two)|46.0%|Easy||
|0232|Implement Queue using Stacks|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0232.Implement-Queue-using-Stacks)|63.1%|Easy||
|0233|Number of Digit One||34.0%|Hard||
|0234|Palindrome Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0234.Palindrome-Linked-List)|50.1%|Easy||
|0235|Lowest Common Ancestor of a Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0235.Lowest-Common-Ancestor-of-a-Binary-Search-Tree)|61.4%|Medium||
|0236|Lowest Common Ancestor of a Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0236.Lowest-Common-Ancestor-of-a-Binary-Tree)|58.7%|Medium||
|0237|Delete Node in a Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0237.Delete-Node-in-a-Linked-List)|75.9%|Medium||
|0238|Product of Array Except Self||65.0%|Medium||
|0239|Sliding Window Maximum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0239.Sliding-Window-Maximum)|46.3%|Hard||
|0240|Search a 2D Matrix II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0240.Search-a-2D-Matrix-II)|51.0%|Medium||
|0241|Different Ways to Add Parentheses||63.8%|Medium||
|0242|Valid Anagram|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0242.Valid-Anagram)|63.0%|Easy||
|0243|Shortest Word Distance||65.0%|Easy||
|0244|Shortest Word Distance II||60.8%|Medium||
|0245|Shortest Word Distance III||57.6%|Medium||
|0246|Strobogrammatic Number||47.7%|Easy||
|0247|Strobogrammatic Number II||51.5%|Medium||
|0248|Strobogrammatic Number III||41.9%|Hard||
|0249|Group Shifted Strings||64.3%|Medium||
|0250|Count Univalue Subtrees||55.4%|Medium||
|0251|Flatten 2D Vector||49.1%|Medium||
|0252|Meeting Rooms||57.2%|Easy||
|0253|Meeting Rooms II||50.6%|Medium||
|0254|Factor Combinations||49.1%|Medium||
|0255|Verify Preorder Sequence in Binary Search Tree||48.2%|Medium||
|0256|Paint House||61.0%|Medium||
|0257|Binary Tree Paths|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0257.Binary-Tree-Paths)|61.3%|Easy||
|0258|Add Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0258.Add-Digits)|63.9%|Easy||
|0259|3Sum Smaller||50.7%|Medium||
|0260|Single Number III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0260.Single-Number-III)|67.6%|Medium||
|0261|Graph Valid Tree||47.1%|Medium||
|0262|Trips and Users||37.5%|Hard||
|0263|Ugly Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0263.Ugly-Number)|42.3%|Easy||
|0264|Ugly Number II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0264.Ugly-Number-II)|46.2%|Medium||
|0265|Paint House II||52.9%|Hard||
|0266|Palindrome Permutation||66.0%|Easy||
|0267|Palindrome Permutation II||40.7%|Medium||
|0268|Missing Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0268.Missing-Number)|62.4%|Easy||
|0269|Alien Dictionary||35.3%|Hard||
|0270|Closest Binary Search Tree Value||54.7%|Easy||
|0271|Encode and Decode Strings||42.7%|Medium||
|0272|Closest Binary Search Tree Value II||58.4%|Hard||
|0273|Integer to English Words||30.0%|Hard||
|0274|H-Index|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0274.H-Index)|38.3%|Medium||
|0275|H-Index II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0275.H-Index-II)|37.5%|Medium||
|0276|Paint Fence||44.5%|Medium||
|0277|Find the Celebrity||46.6%|Medium||
|0278|First Bad Version|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0278.First-Bad-Version)|43.3%|Easy||
|0279|Perfect Squares|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0279.Perfect-Squares)|52.6%|Medium||
|0280|Wiggle Sort||67.0%|Medium||
|0281|Zigzag Iterator||62.7%|Medium||
|0282|Expression Add Operators||39.2%|Hard||
|0283|Move Zeroes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0283.Move-Zeroes)|61.4%|Easy||
|0284|Peeking Iterator|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0284.Peeking-Iterator)|58.6%|Medium||
|0285|Inorder Successor in BST||48.7%|Medium||
|0286|Walls and Gates||60.4%|Medium||
|0287|Find the Duplicate Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0287.Find-the-Duplicate-Number)|59.1%|Medium||
|0288|Unique Word Abbreviation||25.6%|Medium||
|0289|Game of Life||67.1%|Medium||
|0290|Word Pattern|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0290.Word-Pattern)|41.7%|Easy||
|0291|Word Pattern II||47.1%|Medium||
|0292|Nim Game||56.1%|Easy||
|0293|Flip Game||63.2%|Easy||
|0294|Flip Game II||51.9%|Medium||
|0295|Find Median from Data Stream||51.5%|Hard||
|0296|Best Meeting Point||60.1%|Hard||
|0297|Serialize and Deserialize Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0297.Serialize-and-Deserialize-Binary-Tree)|55.4%|Hard||
|0298|Binary Tree Longest Consecutive Sequence||52.7%|Medium||
|0299|Bulls and Cows|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0299.Bulls-and-Cows)|49.3%|Medium||
|0300|Longest Increasing Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0300.Longest-Increasing-Subsequence)|52.1%|Medium||
|0301|Remove Invalid Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0301.Remove-Invalid-Parentheses)|47.2%|Hard||
|0302|Smallest Rectangle Enclosing Black Pixels||58.4%|Hard||
|0303|Range Sum Query - Immutable|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0303.Range-Sum-Query-Immutable)|59.3%|Easy||
|0304|Range Sum Query 2D - Immutable|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0304.Range-Sum-Query-2D-Immutable)|52.7%|Medium||
|0305|Number of Islands II||39.6%|Hard||
|0306|Additive Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0306.Additive-Number)|31.0%|Medium||
|0307|Range Sum Query - Mutable|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0307.Range-Sum-Query-Mutable)|40.7%|Medium||
|0308|Range Sum Query 2D - Mutable||42.8%|Hard||
|0309|Best Time to Buy and Sell Stock with Cooldown|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0309.Best-Time-to-Buy-and-Sell-Stock-with-Cooldown)|56.2%|Medium||
|0310|Minimum Height Trees||38.5%|Medium||
|0311|Sparse Matrix Multiplication||67.3%|Medium||
|0312|Burst Balloons||57.0%|Hard||
|0313|Super Ugly Number||45.5%|Medium||
|0314|Binary Tree Vertical Order Traversal||52.3%|Medium||
|0315|Count of Smaller Numbers After Self|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0315.Count-of-Smaller-Numbers-After-Self)|42.6%|Hard||
|0316|Remove Duplicate Letters||44.9%|Medium||
|0317|Shortest Distance from All Buildings||42.6%|Hard||
|0318|Maximum Product of Word Lengths|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0318.Maximum-Product-of-Word-Lengths)|59.9%|Medium||
|0319|Bulb Switcher|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0319.Bulb-Switcher)|48.3%|Medium||
|0320|Generalized Abbreviation||57.6%|Medium||
|0321|Create Maximum Number||29.1%|Hard||
|0322|Coin Change|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0322.Coin-Change)|42.0%|Medium||
|0323|Number of Connected Components in an Undirected Graph||62.2%|Medium||
|0324|Wiggle Sort II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0324.Wiggle-Sort-II)|33.3%|Medium||
|0325|Maximum Size Subarray Sum Equals k||49.2%|Medium||
|0326|Power of Three|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0326.Power-of-Three)|45.5%|Easy||
|0327|Count of Range Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0327.Count-of-Range-Sum)|36.0%|Hard||
|0328|Odd Even Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0328.Odd-Even-Linked-List)|61.2%|Medium||
|0329|Longest Increasing Path in a Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0329.Longest-Increasing-Path-in-a-Matrix)|52.4%|Hard||
|0330|Patching Array||40.2%|Hard||
|0331|Verify Preorder Serialization of a Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0331.Verify-Preorder-Serialization-of-a-Binary-Tree)|44.6%|Medium||
|0332|Reconstruct Itinerary||41.2%|Hard||
|0333|Largest BST Subtree||42.8%|Medium||
|0334|Increasing Triplet Subsequence||42.7%|Medium||
|0335|Self Crossing||29.4%|Hard||
|0336|Palindrome Pairs||35.0%|Hard||
|0337|House Robber III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0337.House-Robber-III)|53.9%|Medium||
|0338|Counting Bits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0338.Counting-Bits)|75.7%|Easy||
|0339|Nested List Weight Sum||82.3%|Medium||
|0340|Longest Substring with At Most K Distinct Characters||48.0%|Medium||
|0341|Flatten Nested List Iterator|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0341.Flatten-Nested-List-Iterator)|61.8%|Medium||
|0342|Power of Four|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0342.Power-of-Four)|46.1%|Easy||
|0343|Integer Break|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0343.Integer-Break)|56.0%|Medium||
|0344|Reverse String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0344.Reverse-String)|76.7%|Easy||
|0345|Reverse Vowels of a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0345.Reverse-Vowels-of-a-String)|50.0%|Easy||
|0346|Moving Average from Data Stream||77.1%|Easy||
|0347|Top K Frequent Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0347.Top-K-Frequent-Elements)|64.3%|Medium||
|0348|Design Tic-Tac-Toe||57.6%|Medium||
|0349|Intersection of Two Arrays|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0349.Intersection-of-Two-Arrays)|70.9%|Easy||
|0350|Intersection of Two Arrays II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0350.Intersection-of-Two-Arrays-II)|55.9%|Easy||
|0351|Android Unlock Patterns||51.5%|Medium||
|0352|Data Stream as Disjoint Intervals|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0352.Data-Stream-as-Disjoint-Intervals)|59.7%|Hard||
|0353|Design Snake Game||39.2%|Medium||
|0354|Russian Doll Envelopes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0354.Russian-Doll-Envelopes)|38.0%|Hard||
|0355|Design Twitter||37.2%|Medium||
|0356|Line Reflection||34.8%|Medium||
|0357|Count Numbers with Unique Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0357.Count-Numbers-with-Unique-Digits)|51.8%|Medium||
|0358|Rearrange String k Distance Apart||37.7%|Hard||
|0359|Logger Rate Limiter||75.6%|Easy||
|0360|Sort Transformed Array||55.0%|Medium||
|0361|Bomb Enemy||51.3%|Medium||
|0362|Design Hit Counter||68.4%|Medium||
|0363|Max Sum of Rectangle No Larger Than K||44.0%|Hard||
|0364|Nested List Weight Sum II||66.6%|Medium||
|0365|Water and Jug Problem||37.4%|Medium||
|0366|Find Leaves of Binary Tree||80.3%|Medium||
|0367|Valid Perfect Square|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0367.Valid-Perfect-Square)|43.3%|Easy||
|0368|Largest Divisible Subset|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0368.Largest-Divisible-Subset)|41.5%|Medium||
|0369|Plus One Linked List||61.0%|Medium||
|0370|Range Addition||71.1%|Medium||
|0371|Sum of Two Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0371.Sum-of-Two-Integers)|50.7%|Medium||
|0372|Super Pow|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0372.Super-Pow)|36.4%|Medium||
|0373|Find K Pairs with Smallest Sums|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0373.Find-K-Pairs-with-Smallest-Sums)|38.3%|Medium||
|0374|Guess Number Higher or Lower|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0374.Guess-Number-Higher-or-Lower)|51.8%|Easy||
|0375|Guess Number Higher or Lower II||46.8%|Medium||
|0376|Wiggle Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0376.Wiggle-Subsequence)|48.3%|Medium||
|0377|Combination Sum IV|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0377.Combination-Sum-IV)|52.2%|Medium||
|0378|Kth Smallest Element in a Sorted Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0378.Kth-Smallest-Element-in-a-Sorted-Matrix)|61.8%|Medium||
|0379|Design Phone Directory||51.2%|Medium||
|0380|Insert Delete GetRandom O(1)||52.8%|Medium||
|0381|Insert Delete GetRandom O(1) - Duplicates allowed||35.5%|Hard||
|0382|Linked List Random Node|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0382.Linked-List-Random-Node)|62.8%|Medium||
|0383|Ransom Note|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0383.Ransom-Note)|58.2%|Easy||
|0384|Shuffle an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0384.Shuffle-an-Array)|57.8%|Medium||
|0385|Mini Parser|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0385.Mini-Parser)|36.9%|Medium||
|0386|Lexicographical Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0386.Lexicographical-Numbers)|61.5%|Medium||
|0387|First Unique Character in a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0387.First-Unique-Character-in-a-String)|59.5%|Easy||
|0388|Longest Absolute File Path||46.6%|Medium||
|0389|Find the Difference|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0389.Find-the-Difference)|60.0%|Easy||
|0390|Elimination Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0390.Elimination-Game)|46.2%|Medium||
|0391|Perfect Rectangle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0391.Perfect-Rectangle)|32.8%|Hard||
|0392|Is Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0392.Is-Subsequence)|47.7%|Easy||
|0393|UTF-8 Validation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0393.UTF-8-Validation)|45.1%|Medium||
|0394|Decode String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0394.Decode-String)|57.8%|Medium||
|0395|Longest Substring with At Least K Repeating Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0395.Longest-Substring-with-At-Least-K-Repeating-Characters)|44.8%|Medium||
|0396|Rotate Function|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0396.Rotate-Function)|41.1%|Medium||
|0397|Integer Replacement|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0397.Integer-Replacement)|35.2%|Medium||
|0398|Random Pick Index||62.5%|Medium||
|0399|Evaluate Division|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0399.Evaluate-Division)|59.6%|Medium||
|0400|Nth Digit|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0400.Nth-Digit)|34.1%|Medium||
|0401|Binary Watch|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0401.Binary-Watch)|52.2%|Easy||
|0402|Remove K Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0402.Remove-K-Digits)|30.6%|Medium||
|0403|Frog Jump||43.2%|Hard||
|0404|Sum of Left Leaves|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0404.Sum-of-Left-Leaves)|56.7%|Easy||
|0405|Convert a Number to Hexadecimal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0405.Convert-a-Number-to-Hexadecimal)|46.7%|Easy||
|0406|Queue Reconstruction by Height||72.9%|Medium||
|0407|Trapping Rain Water II||47.6%|Hard||
|0408|Valid Word Abbreviation||34.8%|Easy||
|0409|Longest Palindrome|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0409.Longest-Palindrome)|54.3%|Easy||
|0410|Split Array Largest Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0410.Split-Array-Largest-Sum)|53.5%|Hard||
|0411|Minimum Unique Word Abbreviation||39.4%|Hard||
|0412|Fizz Buzz|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0412.Fizz-Buzz)|69.8%|Easy||
|0413|Arithmetic Slices|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0413.Arithmetic-Slices)|65.1%|Medium||
|0414|Third Maximum Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0414.Third-Maximum-Number)|33.1%|Easy||
|0415|Add Strings||52.5%|Easy||
|0416|Partition Equal Subset Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0416.Partition-Equal-Subset-Sum)|46.3%|Medium||
|0417|Pacific Atlantic Water Flow|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0417.Pacific-Atlantic-Water-Flow)|54.4%|Medium||
|0418|Sentence Screen Fitting||35.6%|Medium||
|0419|Battleships in a Board|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0419.Battleships-in-a-Board)|74.8%|Medium||
|0420|Strong Password Checker||13.8%|Hard||
|0421|Maximum XOR of Two Numbers in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0421.Maximum-XOR-of-Two-Numbers-in-an-Array)|54.1%|Medium||
|0422|Valid Word Square||39.9%|Easy||
|0423|Reconstruct Original Digits from English|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0423.Reconstruct-Original-Digits-from-English)|51.3%|Medium||
|0424|Longest Repeating Character Replacement|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0424.Longest-Repeating-Character-Replacement)|51.9%|Medium||
|0425|Word Squares||52.8%|Hard||
|0426|Convert Binary Search Tree to Sorted Doubly Linked List||64.6%|Medium||
|0427|Construct Quad Tree||74.5%|Medium||
|0428|Serialize and Deserialize N-ary Tree||65.9%|Hard||
|0429|N-ary Tree Level Order Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0429.N-ary-Tree-Level-Order-Traversal)|70.7%|Medium||
|0430|Flatten a Multilevel Doubly Linked List||59.6%|Medium||
|0431|Encode N-ary Tree to Binary Tree||78.9%|Hard||
|0432|All O`one Data Structure||36.6%|Hard||
|0433|Minimum Genetic Mutation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0433.Minimum-Genetic-Mutation)|52.3%|Medium||
|0434|Number of Segments in a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0434.Number-of-Segments-in-a-String)|37.2%|Easy||
|0435|Non-overlapping Intervals|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0435.Non-overlapping-Intervals)|50.2%|Medium||
|0436|Find Right Interval|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0436.Find-Right-Interval)|50.7%|Medium||
|0437|Path Sum III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0437.Path-Sum-III)|48.1%|Medium||
|0438|Find All Anagrams in a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0438.Find-All-Anagrams-in-a-String)|50.2%|Medium||
|0439|Ternary Expression Parser||58.4%|Medium||
|0440|K-th Smallest in Lexicographical Order||31.1%|Hard||
|0441|Arranging Coins|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0441.Arranging-Coins)|46.2%|Easy||
|0442|Find All Duplicates in an Array||73.5%|Medium||
|0443|String Compression||52.1%|Medium||
|0444|Sequence Reconstruction||26.7%|Medium||
|0445|Add Two Numbers II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0445.Add-Two-Numbers-II)|59.6%|Medium||
|0446|Arithmetic Slices II - Subsequence||46.6%|Hard||
|0447|Number of Boomerangs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0447.Number-of-Boomerangs)|54.9%|Medium||
|0448|Find All Numbers Disappeared in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0448.Find-All-Numbers-Disappeared-in-an-Array)|59.9%|Easy||
|0449|Serialize and Deserialize BST||57.0%|Medium||
|0450|Delete Node in a BST||50.3%|Medium||
|0451|Sort Characters By Frequency|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0451.Sort-Characters-By-Frequency)|70.0%|Medium||
|0452|Minimum Number of Arrows to Burst Balloons||55.3%|Medium||
|0453|Minimum Moves to Equal Array Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0453.Minimum-Moves-to-Equal-Array-Elements)|56.0%|Medium||
|0454|4Sum II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0454.4Sum-II)|57.2%|Medium||
|0455|Assign Cookies|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0455.Assign-Cookies)|50.0%|Easy||
|0456|132 Pattern|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0456.132-Pattern)|32.4%|Medium||
|0457|Circular Array Loop|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0457.Circular-Array-Loop)|32.6%|Medium||
|0458|Poor Pigs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0458.Poor-Pigs)|63.0%|Hard||
|0459|Repeated Substring Pattern||43.7%|Easy||
|0460|LFU Cache|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0460.LFU-Cache)|43.0%|Hard||
|0461|Hamming Distance|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0461.Hamming-Distance)|75.0%|Easy||
|0462|Minimum Moves to Equal Array Elements II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0462.Minimum-Moves-to-Equal-Array-Elements-II)|60.0%|Medium||
|0463|Island Perimeter|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0463.Island-Perimeter)|69.7%|Easy||
|0464|Can I Win||29.7%|Medium||
|0465|Optimal Account Balancing||49.2%|Hard||
|0466|Count The Repetitions||29.4%|Hard||
|0467|Unique Substrings in Wraparound String||38.5%|Medium||
|0468|Validate IP Address||26.6%|Medium||
|0469|Convex Polygon||38.8%|Medium||
|0470|Implement Rand10() Using Rand7()|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0470.Implement-Rand10-Using-Rand7)|46.4%|Medium||
|0471|Encode String with Shortest Length||50.6%|Hard||
|0472|Concatenated Words||50.1%|Hard||
|0473|Matchsticks to Square|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0473.Matchsticks-to-Square)|40.2%|Medium||
|0474|Ones and Zeroes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0474.Ones-and-Zeroes)|46.7%|Medium||
|0475|Heaters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0475.Heaters)|36.4%|Medium||
|0476|Number Complement|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0476.Number-Complement)|67.3%|Easy||
|0477|Total Hamming Distance|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0477.Total-Hamming-Distance)|52.1%|Medium||
|0478|Generate Random Point in a Circle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0478.Generate-Random-Point-in-a-Circle)|39.6%|Medium||
|0479|Largest Palindrome Product||32.0%|Hard||
|0480|Sliding Window Median|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0480.Sliding-Window-Median)|41.3%|Hard||
|0481|Magical String||50.7%|Medium||
|0482|License Key Formatting||43.3%|Easy||
|0483|Smallest Good Base|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0483.Smallest-Good-Base)|38.7%|Hard||
|0484|Find Permutation||66.9%|Medium||
|0485|Max Consecutive Ones|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0485.Max-Consecutive-Ones)|56.5%|Easy||
|0486|Predict the Winner||51.1%|Medium||
|0487|Max Consecutive Ones II||49.6%|Medium||
|0488|Zuma Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0488.Zuma-Game)|34.0%|Hard||
|0489|Robot Room Cleaner||76.6%|Hard||
|0490|The Maze||55.8%|Medium||
|0491|Non-decreasing Subsequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0491.Non-decreasing-Subsequences)|60.1%|Medium||
|0492|Construct the Rectangle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0492.Construct-the-Rectangle)|54.6%|Easy||
|0493|Reverse Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0493.Reverse-Pairs)|30.9%|Hard||
|0494|Target Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0494.Target-Sum)|45.7%|Medium||
|0495|Teemo Attacking|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0495.Teemo-Attacking)|56.9%|Easy||
|0496|Next Greater Element I|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0496.Next-Greater-Element-I)|71.5%|Easy||
|0497|Random Point in Non-overlapping Rectangles|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0497.Random-Point-in-Non-overlapping-Rectangles)|39.4%|Medium||
|0498|Diagonal Traverse|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0498.Diagonal-Traverse)|58.3%|Medium||
|0499|The Maze III||47.3%|Hard||
|0500|Keyboard Row|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0500.Keyboard-Row)|69.5%|Easy||
|0501|Find Mode in Binary Search Tree||49.2%|Easy||
|0502|IPO||49.7%|Hard||
|0503|Next Greater Element II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0503.Next-Greater-Element-II)|63.2%|Medium||
|0504|Base 7|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0504.Base-7)|48.4%|Easy||
|0505|The Maze II||52.5%|Medium||
|0506|Relative Ranks|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0506.Relative-Ranks)|60.3%|Easy||
|0507|Perfect Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0507.Perfect-Number)|37.7%|Easy||
|0508|Most Frequent Subtree Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0508.Most-Frequent-Subtree-Sum)|64.8%|Medium||
|0509|Fibonacci Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0509.Fibonacci-Number)|69.8%|Easy||
|0510|Inorder Successor in BST II||61.0%|Medium||
|0511|Game Play Analysis I||76.1%|Easy||
|0512|Game Play Analysis II||52.8%|Easy||
|0513|Find Bottom Left Tree Value|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0513.Find-Bottom-Left-Tree-Value)|66.8%|Medium||
|0514|Freedom Trail||47.0%|Hard||
|0515|Find Largest Value in Each Tree Row|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0515.Find-Largest-Value-in-Each-Tree-Row)|64.6%|Medium||
|0516|Longest Palindromic Subsequence||61.0%|Medium||
|0517|Super Washing Machines||40.3%|Hard||
|0518|Coin Change II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0518.Coin-Change-II)|60.5%|Medium||
|0519|Random Flip Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0519.Random-Flip-Matrix)|40.0%|Medium||
|0520|Detect Capital|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0520.Detect-Capital)|57.0%|Easy||
|0521|Longest Uncommon Subsequence I||60.3%|Easy||
|0522|Longest Uncommon Subsequence II||40.5%|Medium||
|0523|Continuous Subarray Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0523.Continuous-Subarray-Sum)|28.5%|Medium||
|0524|Longest Word in Dictionary through Deleting|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0524.Longest-Word-in-Dictionary-through-Deleting)|51.0%|Medium||
|0525|Contiguous Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0525.Contiguous-Array)|46.8%|Medium||
|0526|Beautiful Arrangement|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0526.Beautiful-Arrangement)|64.5%|Medium||
|0527|Word Abbreviation||60.6%|Hard||
|0528|Random Pick with Weight|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0528.Random-Pick-with-Weight)|46.1%|Medium||
|0529|Minesweeper|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0529.Minesweeper)|65.7%|Medium||
|0530|Minimum Absolute Difference in BST|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0530.Minimum-Absolute-Difference-in-BST)|57.3%|Easy||
|0531|Lonely Pixel I||62.2%|Medium||
|0532|K-diff Pairs in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0532.K-diff-Pairs-in-an-Array)|41.1%|Medium||
|0533|Lonely Pixel II||48.5%|Medium||
|0534|Game Play Analysis III||81.5%|Medium||
|0535|Encode and Decode TinyURL|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0535.Encode-and-Decode-TinyURL)|85.9%|Medium||
|0536|Construct Binary Tree from String||56.2%|Medium||
|0537|Complex Number Multiplication|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0537.Complex-Number-Multiplication)|71.4%|Medium||
|0538|Convert BST to Greater Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0538.Convert-BST-to-Greater-Tree)|67.7%|Medium||
|0539|Minimum Time Difference||56.4%|Medium||
|0540|Single Element in a Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0540.Single-Element-in-a-Sorted-Array)|59.1%|Medium||
|0541|Reverse String II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0541.Reverse-String-II)|50.5%|Easy||
|0542|01 Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0542.01-Matrix)|44.7%|Medium||
|0543|Diameter of Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0543.Diameter-of-Binary-Tree)|56.7%|Easy||
|0544|Output Contest Matches||76.9%|Medium||
|0545|Boundary of Binary Tree||44.5%|Medium||
|0546|Remove Boxes||47.7%|Hard||
|0547|Number of Provinces|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0547.Number-of-Provinces)|63.7%|Medium||
|0548|Split Array with Equal Sum||50.1%|Hard||
|0549|Binary Tree Longest Consecutive Sequence II||49.5%|Medium||
|0550|Game Play Analysis IV||43.0%|Medium||
|0551|Student Attendance Record I|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0551.Student-Attendance-Record-I)|48.2%|Easy||
|0552|Student Attendance Record II||41.4%|Hard||
|0553|Optimal Division||59.9%|Medium||
|0554|Brick Wall|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0554.Brick-Wall)|53.5%|Medium||
|0555|Split Concatenated Strings||43.6%|Medium||
|0556|Next Greater Element III||34.0%|Medium||
|0557|Reverse Words in a String III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0557.Reverse-Words-in-a-String-III)|81.9%|Easy||
|0558|Logical OR of Two Binary Grids Represented as Quad-Trees||48.6%|Medium||
|0559|Maximum Depth of N-ary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0559.Maximum-Depth-of-N-ary-Tree)|71.7%|Easy||
|0560|Subarray Sum Equals K|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0560.Subarray-Sum-Equals-K)|43.7%|Medium||
|0561|Array Partition|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0561.Array-Partition)|77.1%|Easy||
|0562|Longest Line of Consecutive One in Matrix||50.2%|Medium||
|0563|Binary Tree Tilt|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0563.Binary-Tree-Tilt)|59.9%|Easy||
|0564|Find the Closest Palindrome||21.9%|Hard||
|0565|Array Nesting||56.4%|Medium||
|0566|Reshape the Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0566.Reshape-the-Matrix)|62.8%|Easy||
|0567|Permutation in String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0567.Permutation-in-String)|44.4%|Medium||
|0568|Maximum Vacation Days||45.0%|Hard||
|0569|Median Employee Salary||67.4%|Hard||
|0570|Managers with at Least 5 Direct Reports||66.6%|Medium||
|0571|Find Median Given Frequency of Numbers||43.8%|Hard||
|0572|Subtree of Another Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0572.Subtree-of-Another-Tree)|46.3%|Easy||
|0573|Squirrel Simulation||55.1%|Medium||
|0574|Winning Candidate||59.8%|Medium||
|0575|Distribute Candies|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0575.Distribute-Candies)|66.4%|Easy||
|0576|Out of Boundary Paths|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0576.Out-of-Boundary-Paths)|44.3%|Medium||
|0577|Employee Bonus||75.1%|Easy||
|0578|Get Highest Answer Rate Question||41.0%|Medium||
|0579|Find Cumulative Salary of an Employee||45.6%|Hard||
|0580|Count Student Number in Departments||58.6%|Medium||
|0581|Shortest Unsorted Continuous Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0581.Shortest-Unsorted-Continuous-Subarray)|36.4%|Medium||
|0582|Kill Process||68.7%|Medium||
|0583|Delete Operation for Two Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0583.Delete-Operation-for-Two-Strings)|59.7%|Medium||
|0584|Find Customer Referee||69.4%|Easy||
|0585|Investments in 2016||52.3%|Medium||
|0586|Customer Placing the Largest Number of Orders||69.4%|Easy||
|0587|Erect the Fence||52.2%|Hard||
|0588|Design In-Memory File System||48.6%|Hard||
|0589|N-ary Tree Preorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0589.N-ary-Tree-Preorder-Traversal)|75.9%|Easy||
|0590|N-ary Tree Postorder Traversal||77.4%|Easy||
|0591|Tag Validator||37.2%|Hard||
|0592|Fraction Addition and Subtraction||52.4%|Medium||
|0593|Valid Square||44.0%|Medium||
|0594|Longest Harmonious Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0594.Longest-Harmonious-Subsequence)|53.5%|Easy||
|0595|Big Countries||70.7%|Easy||
|0596|Classes More Than 5 Students||47.9%|Easy||
|0597|Friend Requests I: Overall Acceptance Rate||42.5%|Easy||
|0598|Range Addition II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0598.Range-Addition-II)|55.3%|Easy||
|0599|Minimum Index Sum of Two Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0599.Minimum-Index-Sum-of-Two-Lists)|53.3%|Easy||
|0600|Non-negative Integers without Consecutive Ones||39.2%|Hard||
|0601|Human Traffic of Stadium||50.0%|Hard||
|0602|Friend Requests II: Who Has the Most Friends||61.1%|Medium||
|0603|Consecutive Available Seats||67.6%|Easy||
|0604|Design Compressed String Iterator||39.5%|Easy||
|0605|Can Place Flowers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0605.Can-Place-Flowers)|32.5%|Easy||
|0606|Construct String from Binary Tree||64.0%|Easy||
|0607|Sales Person||69.6%|Easy||
|0608|Tree Node||71.6%|Medium||
|0609|Find Duplicate File in System|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0609.Find-Duplicate-File-in-System)|67.8%|Medium||
|0610|Triangle Judgement||70.9%|Easy||
|0611|Valid Triangle Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0611.Valid-Triangle-Number)|50.5%|Medium||
|0612|Shortest Distance in a Plane||62.7%|Medium||
|0613|Shortest Distance in a Line||81.0%|Easy||
|0614|Second Degree Follower||37.5%|Medium||
|0615|Average Salary: Departments VS Company||56.6%|Hard||
|0616|Add Bold Tag in String||48.8%|Medium||
|0617|Merge Two Binary Trees|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0617.Merge-Two-Binary-Trees)|78.6%|Easy||
|0618|Students Report By Geography||64.2%|Hard||
|0619|Biggest Single Number||49.8%|Easy||
|0620|Not Boring Movies||72.3%|Easy||
|0621|Task Scheduler||56.3%|Medium||
|0622|Design Circular Queue|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0622.Design-Circular-Queue)|51.6%|Medium||
|0623|Add One Row to Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0623.Add-One-Row-to-Tree)|59.5%|Medium||
|0624|Maximum Distance in Arrays||41.8%|Medium||
|0625|Minimum Factorization||33.5%|Medium||
|0626|Exchange Seats||70.0%|Medium||
|0627|Swap Salary||83.1%|Easy||
|0628|Maximum Product of Three Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0628.Maximum-Product-of-Three-Numbers)|46.0%|Easy||
|0629|K Inverse Pairs Array||42.7%|Hard||
|0630|Course Schedule III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0630.Course-Schedule-III)|40.2%|Hard||
|0631|Design Excel Sum Formula||43.5%|Hard||
|0632|Smallest Range Covering Elements from K Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0632.Smallest-Range-Covering-Elements-from-K-Lists)|61.0%|Hard||
|0633|Sum of Square Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0633.Sum-of-Square-Numbers)|34.4%|Medium||
|0634|Find the Derangement of An Array||42.0%|Medium||
|0635|Design Log Storage System||62.8%|Medium||
|0636|Exclusive Time of Functions|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0636.Exclusive-Time-of-Functions)|61.2%|Medium||
|0637|Average of Levels in Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0637.Average-of-Levels-in-Binary-Tree)|71.8%|Easy||
|0638|Shopping Offers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0638.Shopping-Offers)|53.3%|Medium||
|0639|Decode Ways II||30.4%|Hard||
|0640|Solve the Equation||43.4%|Medium||
|0641|Design Circular Deque||57.4%|Medium||
|0642|Design Search Autocomplete System||48.6%|Hard||
|0643|Maximum Average Subarray I|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0643.Maximum-Average-Subarray-I)|43.7%|Easy||
|0644|Maximum Average Subarray II||35.9%|Hard||
|0645|Set Mismatch|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0645.Set-Mismatch)|42.8%|Easy||
|0646|Maximum Length of Pair Chain||56.5%|Medium||
|0647|Palindromic Substrings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0647.Palindromic-Substrings)|66.8%|Medium||
|0648|Replace Words|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0648.Replace-Words)|62.7%|Medium||
|0649|Dota2 Senate||40.6%|Medium||
|0650|2 Keys Keyboard||53.4%|Medium||
|0651|4 Keys Keyboard||54.6%|Medium||
|0652|Find Duplicate Subtrees||59.0%|Medium||
|0653|Two Sum IV - Input is a BST|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0653.Two-Sum-IV-Input-is-a-BST)|61.0%|Easy||
|0654|Maximum Binary Tree||84.7%|Medium||
|0655|Print Binary Tree||61.9%|Medium||
|0656|Coin Path||31.8%|Hard||
|0657|Robot Return to Origin||75.3%|Easy||
|0658|Find K Closest Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0658.Find-K-Closest-Elements)|46.8%|Medium||
|0659|Split Array into Consecutive Subsequences||50.8%|Medium||
|0660|Remove 9||57.3%|Hard||
|0661|Image Smoother|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0661.Image-Smoother)|55.4%|Easy||
|0662|Maximum Width of Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0662.Maximum-Width-of-Binary-Tree)|40.7%|Medium||
|0663|Equal Tree Partition||41.4%|Medium||
|0664|Strange Printer||46.9%|Hard||
|0665|Non-decreasing Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0665.Non-decreasing-Array)|24.3%|Medium||
|0666|Path Sum IV||59.5%|Medium||
|0667|Beautiful Arrangement II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0667.Beautiful-Arrangement-II)|59.8%|Medium||
|0668|Kth Smallest Number in Multiplication Table|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0668.Kth-Smallest-Number-in-Multiplication-Table)|51.5%|Hard||
|0669|Trim a Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0669.Trim-a-Binary-Search-Tree)|66.4%|Medium||
|0670|Maximum Swap||47.9%|Medium||
|0671|Second Minimum Node In a Binary Tree||44.1%|Easy||
|0672|Bulb Switcher II||50.7%|Medium||
|0673|Number of Longest Increasing Subsequence||42.8%|Medium||
|0674|Longest Continuous Increasing Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0674.Longest-Continuous-Increasing-Subsequence)|49.3%|Easy||
|0675|Cut Off Trees for Golf Event||34.1%|Hard||
|0676|Implement Magic Dictionary|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0676.Implement-Magic-Dictionary)|56.9%|Medium||
|0677|Map Sum Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0677.Map-Sum-Pairs)|56.8%|Medium||
|0678|Valid Parenthesis String||34.1%|Medium||
|0679|24 Game||49.2%|Hard||
|0680|Valid Palindrome II||39.3%|Easy||
|0681|Next Closest Time||46.4%|Medium||
|0682|Baseball Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0682.Baseball-Game)|74.2%|Easy||
|0683|K Empty Slots||37.0%|Hard||
|0684|Redundant Connection|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0684.Redundant-Connection)|62.2%|Medium||
|0685|Redundant Connection II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0685.Redundant-Connection-II)|34.1%|Hard||
|0686|Repeated String Match||34.2%|Medium||
|0687|Longest Univalue Path||40.4%|Medium||
|0688|Knight Probability in Chessboard||52.1%|Medium||
|0689|Maximum Sum of 3 Non-Overlapping Subarrays||49.0%|Hard||
|0690|Employee Importance|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0690.Employee-Importance)|65.5%|Medium||
|0691|Stickers to Spell Word||46.3%|Hard||
|0692|Top K Frequent Words|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0692.Top-K-Frequent-Words)|57.1%|Medium||
|0693|Binary Number with Alternating Bits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0693.Binary-Number-with-Alternating-Bits)|61.6%|Easy||
|0694|Number of Distinct Islands||60.7%|Medium||
|0695|Max Area of Island|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0695.Max-Area-of-Island)|71.8%|Medium||
|0696|Count Binary Substrings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0696.Count-Binary-Substrings)|65.6%|Easy||
|0697|Degree of an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0697.Degree-of-an-Array)|55.9%|Easy||
|0698|Partition to K Equal Sum Subsets||40.1%|Medium||
|0699|Falling Squares|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0699.Falling-Squares)|44.5%|Hard||
|0700|Search in a Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0700.Search-in-a-Binary-Search-Tree)|77.6%|Easy||
|0701|Insert into a Binary Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0701.Insert-into-a-Binary-Search-Tree)|74.3%|Medium||
|0702|Search in a Sorted Array of Unknown Size||71.5%|Medium||
|0703|Kth Largest Element in a Stream|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0703.Kth-Largest-Element-in-a-Stream)|55.5%|Easy||
|0704|Binary Search|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0704.Binary-Search)|55.4%|Easy||
|0705|Design HashSet|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0705.Design-HashSet)|65.7%|Easy||
|0706|Design HashMap|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0706.Design-HashMap)|64.7%|Easy||
|0707|Design Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0707.Design-Linked-List)|27.6%|Medium||
|0708|Insert into a Sorted Circular Linked List||34.5%|Medium||
|0709|To Lower Case|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0709.To-Lower-Case)|82.3%|Easy||
|0710|Random Pick with Blacklist|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0710.Random-Pick-with-Blacklist)|33.5%|Hard||
|0711|Number of Distinct Islands II||51.8%|Hard||
|0712|Minimum ASCII Delete Sum for Two Strings||62.5%|Medium||
|0713|Subarray Product Less Than K|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0713.Subarray-Product-Less-Than-K)|45.7%|Medium||
|0714|Best Time to Buy and Sell Stock with Transaction Fee|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0714.Best-Time-to-Buy-and-Sell-Stock-with-Transaction-Fee)|65.0%|Medium||
|0715|Range Module|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0715.Range-Module)|44.6%|Hard||
|0716|Max Stack||45.2%|Hard||
|0717|1-bit and 2-bit Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0717.1-bit-and-2-bit-Characters)|45.7%|Easy||
|0718|Maximum Length of Repeated Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0718.Maximum-Length-of-Repeated-Subarray)|51.3%|Medium||
|0719|Find K-th Smallest Pair Distance|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0719.Find-K-th-Smallest-Pair-Distance)|36.7%|Hard||
|0720|Longest Word in Dictionary|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0720.Longest-Word-in-Dictionary)|51.9%|Medium||
|0721|Accounts Merge|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0721.Accounts-Merge)|56.3%|Medium||
|0722|Remove Comments||38.2%|Medium||
|0723|Candy Crush||76.6%|Medium||
|0724|Find Pivot Index|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0724.Find-Pivot-Index)|54.5%|Easy||
|0725|Split Linked List in Parts|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0725.Split-Linked-List-in-Parts)|57.2%|Medium||
|0726|Number of Atoms|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0726.Number-of-Atoms)|52.1%|Hard||
|0727|Minimum Window Subsequence||42.9%|Hard||
|0728|Self Dividing Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0728.Self-Dividing-Numbers)|77.8%|Easy||
|0729|My Calendar I|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0729.My-Calendar-I)|56.9%|Medium||
|0730|Count Different Palindromic Subsequences||44.8%|Hard||
|0731|My Calendar II||55.0%|Medium||
|0732|My Calendar III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0732.My-Calendar-III)|71.5%|Hard||
|0733|Flood Fill|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0733.Flood-Fill)|61.8%|Easy||
|0734|Sentence Similarity||44.0%|Easy||
|0735|Asteroid Collision|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0735.Asteroid-Collision)|44.4%|Medium||
|0736|Parse Lisp Expression||51.6%|Hard||
|0737|Sentence Similarity II||49.0%|Medium||
|0738|Monotone Increasing Digits||47.2%|Medium||
|0739|Daily Temperatures|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0739.Daily-Temperatures)|66.3%|Medium||
|0740|Delete and Earn||57.1%|Medium||
|0741|Cherry Pickup||36.3%|Hard||
|0742|Closest Leaf in a Binary Tree||45.9%|Medium||
|0743|Network Delay Time||51.8%|Medium||
|0744|Find Smallest Letter Greater Than Target|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0744.Find-Smallest-Letter-Greater-Than-Target)|45.6%|Easy||
|0745|Prefix and Suffix Search|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0745.Prefix-and-Suffix-Search)|41.3%|Hard||
|0746|Min Cost Climbing Stairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0746.Min-Cost-Climbing-Stairs)|63.1%|Easy||
|0747|Largest Number At Least Twice of Others|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0747.Largest-Number-At-Least-Twice-of-Others)|47.0%|Easy||
|0748|Shortest Completing Word|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0748.Shortest-Completing-Word)|59.2%|Easy||
|0749|Contain Virus||51.1%|Hard||
|0750|Number Of Corner Rectangles||67.5%|Medium||
|0751|IP to CIDR||54.5%|Medium||
|0752|Open the Lock|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0752.Open-the-Lock)|55.6%|Medium||
|0753|Cracking the Safe|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0753.Cracking-the-Safe)|55.8%|Hard||
|0754|Reach a Number||42.7%|Medium||
|0755|Pour Water||46.4%|Medium||
|0756|Pyramid Transition Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0756.Pyramid-Transition-Matrix)|52.7%|Medium||
|0757|Set Intersection Size At Least Two||43.8%|Hard||
|0758|Bold Words in String||50.8%|Medium||
|0759|Employee Free Time||71.8%|Hard||
|0760|Find Anagram Mappings||82.9%|Easy||
|0761|Special Binary String||60.2%|Hard||
|0762|Prime Number of Set Bits in Binary Representation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0762.Prime-Number-of-Set-Bits-in-Binary-Representation)|68.0%|Easy||
|0763|Partition Labels|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0763.Partition-Labels)|79.8%|Medium||
|0764|Largest Plus Sign||48.3%|Medium||
|0765|Couples Holding Hands|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0765.Couples-Holding-Hands)|56.8%|Hard||
|0766|Toeplitz Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0766.Toeplitz-Matrix)|68.7%|Easy||
|0767|Reorganize String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0767.Reorganize-String)|52.9%|Medium||
|0768|Max Chunks To Make Sorted II||52.8%|Hard||
|0769|Max Chunks To Make Sorted||58.2%|Medium||
|0770|Basic Calculator IV||55.9%|Hard||
|0771|Jewels and Stones|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0771.Jewels-and-Stones)|88.2%|Easy||
|0772|Basic Calculator III||48.8%|Hard||
|0773|Sliding Puzzle||64.1%|Hard||
|0774|Minimize Max Distance to Gas Station||51.7%|Hard||
|0775|Global and Local Inversions|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0775.Global-and-Local-Inversions)|43.3%|Medium||
|0776|Split BST||72.6%|Medium||
|0777|Swap Adjacent in LR String||36.8%|Medium||
|0778|Swim in Rising Water|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0778.Swim-in-Rising-Water)|59.8%|Hard||
|0779|K-th Symbol in Grammar||41.3%|Medium||
|0780|Reaching Points||32.6%|Hard||
|0781|Rabbits in Forest|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0781.Rabbits-in-Forest)|54.8%|Medium||
|0782|Transform to Chessboard||51.7%|Hard||
|0783|Minimum Distance Between BST Nodes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0783.Minimum-Distance-Between-BST-Nodes)|59.3%|Easy||
|0784|Letter Case Permutation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0784.Letter-Case-Permutation)|73.8%|Medium||
|0785|Is Graph Bipartite?|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0785.Is-Graph-Bipartite)|53.1%|Medium||
|0786|K-th Smallest Prime Fraction|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0786.K-th-Smallest-Prime-Fraction)|51.6%|Medium||
|0787|Cheapest Flights Within K Stops||37.0%|Medium||
|0788|Rotated Digits||56.7%|Medium||
|0789|Escape The Ghosts||60.8%|Medium||
|0790|Domino and Tromino Tiling||52.9%|Medium||
|0791|Custom Sort String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0791.Custom-Sort-String)|69.2%|Medium||
|0792|Number of Matching Subsequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0792.Number-of-Matching-Subsequences)|51.7%|Medium||
|0793|Preimage Size of Factorial Zeroes Function|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0793.Preimage-Size-of-Factorial-Zeroes-Function)|43.2%|Hard||
|0794|Valid Tic-Tac-Toe State|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0794.Valid-Tic-Tac-Toe-State)|35.1%|Medium||
|0795|Number of Subarrays with Bounded Maximum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0795.Number-of-Subarrays-with-Bounded-Maximum)|52.8%|Medium||
|0796|Rotate String||55.1%|Easy||
|0797|All Paths From Source to Target||82.3%|Medium||
|0798|Smallest Rotation with Highest Score||50.2%|Hard||
|0799|Champagne Tower||51.3%|Medium||
|0800|Similar RGB Color||67.0%|Easy||
|0801|Minimum Swaps To Make Sequences Increasing||39.4%|Hard||
|0802|Find Eventual Safe States|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0802.Find-Eventual-Safe-States)|56.4%|Medium||
|0803|Bricks Falling When Hit|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0803.Bricks-Falling-When-Hit)|34.3%|Hard||
|0804|Unique Morse Code Words||82.6%|Easy||
|0805|Split Array With Same Average||25.7%|Hard||
|0806|Number of Lines To Write String||66.6%|Easy||
|0807|Max Increase to Keep City Skyline|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0807.Max-Increase-to-Keep-City-Skyline)|85.9%|Medium||
|0808|Soup Servings||43.4%|Medium||
|0809|Expressive Words||46.2%|Medium||
|0810|Chalkboard XOR Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0810.Chalkboard-XOR-Game)|56.2%|Hard||
|0811|Subdomain Visit Count|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0811.Subdomain-Visit-Count)|75.5%|Medium||
|0812|Largest Triangle Area|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0812.Largest-Triangle-Area)|59.9%|Easy||
|0813|Largest Sum of Averages||53.0%|Medium||
|0814|Binary Tree Pruning||72.4%|Medium||
|0815|Bus Routes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0815.Bus-Routes)|45.6%|Hard||
|0816|Ambiguous Coordinates|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0816.Ambiguous-Coordinates)|56.3%|Medium||
|0817|Linked List Components|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0817.Linked-List-Components)|57.7%|Medium||
|0818|Race Car||43.3%|Hard||
|0819|Most Common Word|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0819.Most-Common-Word)|44.8%|Easy||
|0820|Short Encoding of Words|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0820.Short-Encoding-of-Words)|60.6%|Medium||
|0821|Shortest Distance to a Character|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0821.Shortest-Distance-to-a-Character)|71.3%|Easy||
|0822|Card Flipping Game||45.9%|Medium||
|0823|Binary Trees With Factors|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0823.Binary-Trees-With-Factors)|49.7%|Medium||
|0824|Goat Latin||67.9%|Easy||
|0825|Friends Of Appropriate Ages|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0825.Friends-Of-Appropriate-Ages)|46.3%|Medium||
|0826|Most Profit Assigning Work|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0826.Most-Profit-Assigning-Work)|44.8%|Medium||
|0827|Making A Large Island||44.9%|Hard||
|0828|Count Unique Characters of All Substrings of a Given String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0828.Count-Unique-Characters-of-All-Substrings-of-a-Given-String)|51.7%|Hard||
|0829|Consecutive Numbers Sum||41.6%|Hard||
|0830|Positions of Large Groups|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0830.Positions-of-Large-Groups)|51.8%|Easy||
|0831|Masking Personal Information||47.4%|Medium||
|0832|Flipping an Image|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0832.Flipping-an-Image)|80.7%|Easy||
|0833|Find And Replace in String||54.0%|Medium||
|0834|Sum of Distances in Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0834.Sum-of-Distances-in-Tree)|59.3%|Hard||
|0835|Image Overlap||63.9%|Medium||
|0836|Rectangle Overlap|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0836.Rectangle-Overlap)|43.9%|Easy||
|0837|New 21 Game||36.2%|Medium||
|0838|Push Dominoes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0838.Push-Dominoes)|57.0%|Medium||
|0839|Similar String Groups|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0839.Similar-String-Groups)|48.0%|Hard||
|0840|Magic Squares In Grid||38.6%|Medium||
|0841|Keys and Rooms|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0841.Keys-and-Rooms)|71.5%|Medium||
|0842|Split Array into Fibonacci Sequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0842.Split-Array-into-Fibonacci-Sequence)|38.4%|Medium||
|0843|Guess the Word||41.4%|Hard||
|0844|Backspace String Compare|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0844.Backspace-String-Compare)|48.0%|Easy||
|0845|Longest Mountain in Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0845.Longest-Mountain-in-Array)|40.2%|Medium||
|0846|Hand of Straights|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0846.Hand-of-Straights)|56.2%|Medium||
|0847|Shortest Path Visiting All Nodes||61.1%|Hard||
|0848|Shifting Letters||45.2%|Medium||
|0849|Maximize Distance to Closest Person||47.6%|Medium||
|0850|Rectangle Area II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0850.Rectangle-Area-II)|53.9%|Hard||
|0851|Loud and Rich|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0851.Loud-and-Rich)|58.4%|Medium||
|0852|Peak Index in a Mountain Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0852.Peak-Index-in-a-Mountain-Array)|69.1%|Medium||
|0853|Car Fleet|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0853.Car-Fleet)|50.3%|Medium||
|0854|K-Similar Strings||40.1%|Hard||
|0855|Exam Room||43.4%|Medium||
|0856|Score of Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0856.Score-of-Parentheses)|64.8%|Medium||
|0857|Minimum Cost to Hire K Workers||52.3%|Hard||
|0858|Mirror Reflection||63.1%|Medium||
|0859|Buddy Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0859.Buddy-Strings)|29.2%|Easy||
|0860|Lemonade Change||52.9%|Easy||
|0861|Score After Flipping Matrix||75.0%|Medium||
|0862|Shortest Subarray with Sum at Least K|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0862.Shortest-Subarray-with-Sum-at-Least-K)|26.1%|Hard||
|0863|All Nodes Distance K in Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0863.All-Nodes-Distance-K-in-Binary-Tree)|62.3%|Medium||
|0864|Shortest Path to Get All Keys|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0864.Shortest-Path-to-Get-All-Keys)|45.5%|Hard||
|0865|Smallest Subtree with all the Deepest Nodes||68.9%|Medium||
|0866|Prime Palindrome||25.7%|Medium||
|0867|Transpose Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0867.Transpose-Matrix)|64.1%|Easy||
|0868|Binary Gap||62.2%|Easy||
|0869|Reordered Power of 2|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0869.Reordered-Power-of-2)|63.6%|Medium||
|0870|Advantage Shuffle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0870.Advantage-Shuffle)|51.8%|Medium||
|0871|Minimum Number of Refueling Stops||39.8%|Hard||
|0872|Leaf-Similar Trees|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0872.Leaf-Similar-Trees)|67.6%|Easy||
|0873|Length of Longest Fibonacci Subsequence||48.3%|Medium||
|0874|Walking Robot Simulation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0874.Walking-Robot-Simulation)|38.9%|Medium||
|0875|Koko Eating Bananas|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0875.Koko-Eating-Bananas)|52.3%|Medium||
|0876|Middle of the Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0876.Middle-of-the-Linked-List)|75.5%|Easy||
|0877|Stone Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0877.Stone-Game)|69.7%|Medium||
|0878|Nth Magical Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0878.Nth-Magical-Number)|35.3%|Hard||
|0879|Profitable Schemes||40.6%|Hard||
|0880|Decoded String at Index|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0880.Decoded-String-at-Index)|28.3%|Medium||
|0881|Boats to Save People|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0881.Boats-to-Save-People)|53.1%|Medium||
|0882|Reachable Nodes In Subdivided Graph||50.2%|Hard||
|0883|Projection Area of 3D Shapes||71.2%|Easy||
|0884|Uncommon Words from Two Sentences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0884.Uncommon-Words-from-Two-Sentences)|66.3%|Easy||
|0885|Spiral Matrix III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0885.Spiral-Matrix-III)|73.4%|Medium||
|0886|Possible Bipartition||50.0%|Medium||
|0887|Super Egg Drop|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0887.Super-Egg-Drop)|27.2%|Hard||
|0888|Fair Candy Swap|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0888.Fair-Candy-Swap)|60.7%|Easy||
|0889|Construct Binary Tree from Preorder and Postorder Traversal||71.0%|Medium||
|0890|Find and Replace Pattern|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0890.Find-and-Replace-Pattern)|77.6%|Medium||
|0891|Sum of Subsequence Widths|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0891.Sum-of-Subsequence-Widths)|36.6%|Hard||
|0892|Surface Area of 3D Shapes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0892.Surface-Area-of-3D-Shapes)|64.0%|Easy||
|0893|Groups of Special-Equivalent Strings||71.1%|Medium||
|0894|All Possible Full Binary Trees||80.0%|Medium||
|0895|Maximum Frequency Stack|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0895.Maximum-Frequency-Stack)|66.6%|Hard||
|0896|Monotonic Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0896.Monotonic-Array)|58.4%|Easy||
|0897|Increasing Order Search Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0897.Increasing-Order-Search-Tree)|78.5%|Easy||
|0898|Bitwise ORs of Subarrays|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0898.Bitwise-ORs-of-Subarrays)|37.2%|Medium||
|0899|Orderly Queue||66.4%|Hard||
|0900|RLE Iterator||59.4%|Medium||
|0901|Online Stock Span|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0901.Online-Stock-Span)|65.2%|Medium||
|0902|Numbers At Most N Given Digit Set||41.4%|Hard||
|0903|Valid Permutations for DI Sequence||57.8%|Hard||
|0904|Fruit Into Baskets|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0904.Fruit-Into-Baskets)|43.7%|Medium||
|0905|Sort Array By Parity||75.6%|Easy||
|0906|Super Palindromes||38.9%|Hard||
|0907|Sum of Subarray Minimums|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0907.Sum-of-Subarray-Minimums)|35.8%|Medium||
|0908|Smallest Range I||68.2%|Easy||
|0909|Snakes and Ladders|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0909.Snakes-and-Ladders)|45.1%|Medium||
|0910|Smallest Range II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0910.Smallest-Range-II)|35.1%|Medium||
|0911|Online Election|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0911.Online-Election)|52.2%|Medium||
|0912|Sort an Array||59.7%|Medium||
|0913|Cat and Mouse||34.9%|Hard||
|0914|X of a Kind in a Deck of Cards|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0914.X-of-a-Kind-in-a-Deck-of-Cards)|31.3%|Easy||
|0915|Partition Array into Disjoint Intervals||48.5%|Medium||
|0916|Word Subsets|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0916.Word-Subsets)|53.7%|Medium||
|0917|Reverse Only Letters||61.9%|Easy||
|0918|Maximum Sum Circular Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0918.Maximum-Sum-Circular-Subarray)|42.9%|Medium||
|0919|Complete Binary Tree Inserter||65.1%|Medium||
|0920|Number of Music Playlists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0920.Number-of-Music-Playlists)|50.8%|Hard||
|0921|Minimum Add to Make Parentheses Valid|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0921.Minimum-Add-to-Make-Parentheses-Valid)|75.8%|Medium||
|0922|Sort Array By Parity II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0922.Sort-Array-By-Parity-II)|70.7%|Easy||
|0923|3Sum With Multiplicity|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0923.3Sum-With-Multiplicity)|45.3%|Medium||
|0924|Minimize Malware Spread|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0924.Minimize-Malware-Spread)|42.1%|Hard||
|0925|Long Pressed Name|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0925.Long-Pressed-Name)|33.2%|Easy||
|0926|Flip String to Monotone Increasing||61.5%|Medium||
|0927|Three Equal Parts|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0927.Three-Equal-Parts)|39.5%|Hard||
|0928|Minimize Malware Spread II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0928.Minimize-Malware-Spread-II)|42.7%|Hard||
|0929|Unique Email Addresses||67.1%|Easy||
|0930|Binary Subarrays With Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0930.Binary-Subarrays-With-Sum)|52.0%|Medium||
|0931|Minimum Falling Path Sum||69.1%|Medium||
|0932|Beautiful Array||65.2%|Medium||
|0933|Number of Recent Calls|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0933.Number-of-Recent-Calls)|73.2%|Easy||
|0934|Shortest Bridge||54.3%|Medium||
|0935|Knight Dialer||50.4%|Medium||
|0936|Stamping The Sequence||63.0%|Hard||
|0937|Reorder Data in Log Files||56.4%|Medium||
|0938|Range Sum of BST|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0938.Range-Sum-of-BST)|85.9%|Easy||
|0939|Minimum Area Rectangle||52.9%|Medium||
|0940|Distinct Subsequences II||44.0%|Hard||
|0941|Valid Mountain Array||33.4%|Easy||
|0942|DI String Match|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0942.DI-String-Match)|77.3%|Easy||
|0943|Find the Shortest Superstring||44.5%|Hard||
|0944|Delete Columns to Make Sorted||74.8%|Easy||
|0945|Minimum Increment to Make Array Unique||51.1%|Medium||
|0946|Validate Stack Sequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0946.Validate-Stack-Sequences)|67.7%|Medium||
|0947|Most Stones Removed with Same Row or Column|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0947.Most-Stones-Removed-with-Same-Row-or-Column)|58.9%|Medium||
|0948|Bag of Tokens||52.1%|Medium||
|0949|Largest Time for Given Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0949.Largest-Time-for-Given-Digits)|35.2%|Medium||
|0950|Reveal Cards In Increasing Order||77.7%|Medium||
|0951|Flip Equivalent Binary Trees||66.8%|Medium||
|0952|Largest Component Size by Common Factor|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0952.Largest-Component-Size-by-Common-Factor)|40.1%|Hard||
|0953|Verifying an Alien Dictionary|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0953.Verifying-an-Alien-Dictionary)|54.5%|Easy||
|0954|Array of Doubled Pairs||39.1%|Medium||
|0955|Delete Columns to Make Sorted II||34.7%|Medium||
|0956|Tallest Billboard||39.9%|Hard||
|0957|Prison Cells After N Days||39.1%|Medium||
|0958|Check Completeness of a Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0958.Check-Completeness-of-a-Binary-Tree)|56.0%|Medium||
|0959|Regions Cut By Slashes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0959.Regions-Cut-By-Slashes)|69.2%|Medium||
|0960|Delete Columns to Make Sorted III||57.2%|Hard||
|0961|N-Repeated Element in Size 2N Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0961.N-Repeated-Element-in-Size-2N-Array)|76.0%|Easy||
|0962|Maximum Width Ramp||49.0%|Medium||
|0963|Minimum Area Rectangle II||54.7%|Medium||
|0964|Least Operators to Express Number||47.9%|Hard||
|0965|Univalued Binary Tree||69.6%|Easy||
|0966|Vowel Spellchecker|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0966.Vowel-Spellchecker)|51.4%|Medium||
|0967|Numbers With Same Consecutive Differences||57.4%|Medium||
|0968|Binary Tree Cameras|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0968.Binary-Tree-Cameras)|46.7%|Hard||
|0969|Pancake Sorting|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0969.Pancake-Sorting)|70.1%|Medium||
|0970|Powerful Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0970.Powerful-Integers)|43.6%|Medium||
|0971|Flip Binary Tree To Match Preorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0971.Flip-Binary-Tree-To-Match-Preorder-Traversal)|50.1%|Medium||
|0972|Equal Rational Numbers||43.4%|Hard||
|0973|K Closest Points to Origin|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0973.K-Closest-Points-to-Origin)|65.7%|Medium||
|0974|Subarray Sums Divisible by K||54.4%|Medium||
|0975|Odd Even Jump||39.0%|Hard||
|0976|Largest Perimeter Triangle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0976.Largest-Perimeter-Triangle)|54.6%|Easy||
|0977|Squares of a Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0977.Squares-of-a-Sorted-Array)|71.9%|Easy||
|0978|Longest Turbulent Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0978.Longest-Turbulent-Subarray)|47.2%|Medium||
|0979|Distribute Coins in Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0979.Distribute-Coins-in-Binary-Tree)|72.2%|Medium||
|0980|Unique Paths III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0980.Unique-Paths-III)|81.9%|Hard||
|0981|Time Based Key-Value Store|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0981.Time-Based-Key-Value-Store)|52.4%|Medium||
|0982|Triples with Bitwise AND Equal To Zero||57.6%|Hard||
|0983|Minimum Cost For Tickets||64.3%|Medium||
|0984|String Without AAA or BBB|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0984.String-Without-AAA-or-BBB)|43.1%|Medium||
|0985|Sum of Even Numbers After Queries|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0985.Sum-of-Even-Numbers-After-Queries)|68.2%|Medium||
|0986|Interval List Intersections|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0986.Interval-List-Intersections)|71.3%|Medium||
|0987|Vertical Order Traversal of a Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0987.Vertical-Order-Traversal-of-a-Binary-Tree)|45.0%|Hard||
|0988|Smallest String Starting From Leaf||50.2%|Medium||
|0989|Add to Array-Form of Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0989.Add-to-Array-Form-of-Integer)|47.1%|Easy||
|0990|Satisfiability of Equality Equations|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0990.Satisfiability-of-Equality-Equations)|50.5%|Medium||
|0991|Broken Calculator|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0991.Broken-Calculator)|54.1%|Medium||
|0992|Subarrays with K Different Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0992.Subarrays-with-K-Different-Integers)|54.7%|Hard||
|0993|Cousins in Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0993.Cousins-in-Binary-Tree)|54.6%|Easy||
|0994|Rotting Oranges||52.9%|Medium||
|0995|Minimum Number of K Consecutive Bit Flips|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0995.Minimum-Number-of-K-Consecutive-Bit-Flips)|51.2%|Hard||
|0996|Number of Squareful Arrays|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0996.Number-of-Squareful-Arrays)|49.2%|Hard||
|0997|Find the Town Judge|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0997.Find-the-Town-Judge)|49.6%|Easy||
|0998|Maximum Binary Tree II||67.2%|Medium||
|0999|Available Captures for Rook|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/0999.Available-Captures-for-Rook)|68.1%|Easy||
|1000|Minimum Cost to Merge Stones||42.3%|Hard||
|1001|Grid Illumination||36.2%|Hard||
|1002|Find Common Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1002.Find-Common-Characters)|68.5%|Easy||
|1003|Check If Word Is Valid After Substitutions|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1003.Check-If-Word-Is-Valid-After-Substitutions)|58.2%|Medium||
|1004|Max Consecutive Ones III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1004.Max-Consecutive-Ones-III)|63.2%|Medium||
|1005|Maximize Sum Of Array After K Negations|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1005.Maximize-Sum-Of-Array-After-K-Negations)|50.9%|Easy||
|1006|Clumsy Factorial|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1006.Clumsy-Factorial)|55.4%|Medium||
|1007|Minimum Domino Rotations For Equal Row||52.3%|Medium||
|1008|Construct Binary Search Tree from Preorder Traversal||81.1%|Medium||
|1009|Complement of Base 10 Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1009.Complement-of-Base-10-Integer)|61.6%|Easy||
|1010|Pairs of Songs With Total Durations Divisible by 60|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1010.Pairs-of-Songs-With-Total-Durations-Divisible-by-60)|52.8%|Medium||
|1011|Capacity To Ship Packages Within D Days|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1011.Capacity-To-Ship-Packages-Within-D-Days)|67.7%|Medium||
|1012|Numbers With Repeated Digits||39.8%|Hard||
|1013|Partition Array Into Three Parts With Equal Sum||42.8%|Easy||
|1014|Best Sightseeing Pair||59.4%|Medium||
|1015|Smallest Integer Divisible by K||46.8%|Medium||
|1016|Binary String With Substrings Representing 1 To N||57.4%|Medium||
|1017|Convert to Base -2|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1017.Convert-to-Base-2)|60.8%|Medium||
|1018|Binary Prefix Divisible By 5|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1018.Binary-Prefix-Divisible-By-5)|46.9%|Easy||
|1019|Next Greater Node In Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1019.Next-Greater-Node-In-Linked-List)|59.9%|Medium||
|1020|Number of Enclaves|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1020.Number-of-Enclaves)|65.5%|Medium||
|1021|Remove Outermost Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1021.Remove-Outermost-Parentheses)|80.6%|Easy||
|1022|Sum of Root To Leaf Binary Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1022.Sum-of-Root-To-Leaf-Binary-Numbers)|73.6%|Easy||
|1023|Camelcase Matching||60.6%|Medium||
|1024|Video Stitching||50.5%|Medium||
|1025|Divisor Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1025.Divisor-Game)|67.5%|Easy||
|1026|Maximum Difference Between Node and Ancestor|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1026.Maximum-Difference-Between-Node-and-Ancestor)|75.7%|Medium||
|1027|Longest Arithmetic Subsequence||46.8%|Medium||
|1028|Recover a Tree From Preorder Traversal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1028.Recover-a-Tree-From-Preorder-Traversal)|73.2%|Hard||
|1029|Two City Scheduling||65.2%|Medium||
|1030|Matrix Cells in Distance Order|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1030.Matrix-Cells-in-Distance-Order)|69.6%|Easy||
|1031|Maximum Sum of Two Non-Overlapping Subarrays||59.6%|Medium||
|1032|Stream of Characters||51.6%|Hard||
|1033|Moving Stones Until Consecutive||46.2%|Medium||
|1034|Coloring A Border|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1034.Coloring-A-Border)|49.2%|Medium||
|1035|Uncrossed Lines||59.1%|Medium||
|1036|Escape a Large Maze||34.1%|Hard||
|1037|Valid Boomerang|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1037.Valid-Boomerang)|37.1%|Easy||
|1038|Binary Search Tree to Greater Sum Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1038.Binary-Search-Tree-to-Greater-Sum-Tree)|85.5%|Medium||
|1039|Minimum Score Triangulation of Polygon||55.3%|Medium||
|1040|Moving Stones Until Consecutive II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1040.Moving-Stones-Until-Consecutive-II)|55.8%|Medium||
|1041|Robot Bounded In Circle||55.3%|Medium||
|1042|Flower Planting With No Adjacent||50.6%|Medium||
|1043|Partition Array for Maximum Sum||71.5%|Medium||
|1044|Longest Duplicate Substring||30.6%|Hard||
|1045|Customers Who Bought All Products||66.8%|Medium||
|1046|Last Stone Weight||64.8%|Easy||
|1047|Remove All Adjacent Duplicates In String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1047.Remove-All-Adjacent-Duplicates-In-String)|69.7%|Easy||
|1048|Longest String Chain|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1048.Longest-String-Chain)|59.2%|Medium||
|1049|Last Stone Weight II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1049.Last-Stone-Weight-II)|53.1%|Medium||
|1050|Actors and Directors Who Cooperated At Least Three Times||70.8%|Easy||
|1051|Height Checker|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1051.Height-Checker)|75.5%|Easy||
|1052|Grumpy Bookstore Owner|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1052.Grumpy-Bookstore-Owner)|57.1%|Medium||
|1053|Previous Permutation With One Swap||50.5%|Medium||
|1054|Distant Barcodes|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1054.Distant-Barcodes)|45.8%|Medium||
|1055|Shortest Way to Form String||59.7%|Medium||
|1056|Confusing Number||48.2%|Easy||
|1057|Campus Bikes||57.6%|Medium||
|1058|Minimize Rounding Error to Meet Target||44.9%|Medium||
|1059|All Paths from Source Lead to Destination||39.3%|Medium||
|1060|Missing Element in Sorted Array||54.7%|Medium||
|1061|Lexicographically Smallest Equivalent String||76.5%|Medium||
|1062|Longest Repeating Substring||59.2%|Medium||
|1063|Number of Valid Subarrays||74.4%|Hard||
|1064|Fixed Point||64.2%|Easy||
|1065|Index Pairs of a String||63.8%|Easy||
|1066|Campus Bikes II||54.9%|Medium||
|1067|Digit Count in Range||45.1%|Hard||
|1068|Product Sales Analysis I||80.3%|Easy||
|1069|Product Sales Analysis II||81.9%|Easy||
|1070|Product Sales Analysis III||48.7%|Medium||
|1071|Greatest Common Divisor of Strings||56.6%|Easy||
|1072|Flip Columns For Maximum Number of Equal Rows||63.3%|Medium||
|1073|Adding Two Negabinary Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1073.Adding-Two-Negabinary-Numbers)|36.4%|Medium||
|1074|Number of Submatrices That Sum to Target|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1074.Number-of-Submatrices-That-Sum-to-Target)|69.7%|Hard||
|1075|Project Employees I||66.5%|Easy||
|1076|Project Employees II||50.3%|Easy||
|1077|Project Employees III||77.5%|Medium||
|1078|Occurrences After Bigram|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1078.Occurrences-After-Bigram)|63.6%|Easy||
|1079|Letter Tile Possibilities|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1079.Letter-Tile-Possibilities)|76.0%|Medium||
|1080|Insufficient Nodes in Root to Leaf Paths||53.4%|Medium||
|1081|Smallest Subsequence of Distinct Characters||57.8%|Medium||
|1082|Sales Analysis I||74.9%|Easy||
|1083|Sales Analysis II||50.0%|Easy||
|1084|Sales Analysis III||49.1%|Easy||
|1085|Sum of Digits in the Minimum Number||76.3%|Easy||
|1086|High Five||75.1%|Easy||
|1087|Brace Expansion||66.2%|Medium||
|1088|Confusing Number II||47.3%|Hard||
|1089|Duplicate Zeros|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1089.Duplicate-Zeros)|51.5%|Easy||
|1090|Largest Values From Labels||61.1%|Medium||
|1091|Shortest Path in Binary Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1091.Shortest-Path-in-Binary-Matrix)|44.6%|Medium||
|1092|Shortest Common Supersequence||58.0%|Hard||
|1093|Statistics from a Large Sample|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1093.Statistics-from-a-Large-Sample)|43.6%|Medium||
|1094|Car Pooling||56.9%|Medium||
|1095|Find in Mountain Array||35.6%|Hard||
|1096|Brace Expansion II||63.5%|Hard||
|1097|Game Play Analysis V||53.5%|Hard||
|1098|Unpopular Books||44.4%|Medium||
|1099|Two Sum Less Than K||61.0%|Easy||
|1100|Find K-Length Substrings With No Repeated Characters||74.7%|Medium||
|1101|The Earliest Moment When Everyone Become Friends||64.7%|Medium||
|1102|Path With Maximum Minimum Value||53.3%|Medium||
|1103|Distribute Candies to People||64.2%|Easy||
|1104|Path In Zigzag Labelled Binary Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1104.Path-In-Zigzag-Labelled-Binary-Tree)|75.1%|Medium||
|1105|Filling Bookcase Shelves|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1105.Filling-Bookcase-Shelves)|59.3%|Medium||
|1106|Parsing A Boolean Expression||58.4%|Hard||
|1107|New Users Daily Count||45.3%|Medium||
|1108|Defanging an IP Address|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1108.Defanging-an-IP-Address)|89.1%|Easy||
|1109|Corporate Flight Bookings||60.6%|Medium||
|1110|Delete Nodes And Return Forest|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1110.Delete-Nodes-And-Return-Forest)|69.3%|Medium||
|1111|Maximum Nesting Depth of Two Valid Parentheses Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1111.Maximum-Nesting-Depth-of-Two-Valid-Parentheses-Strings)|73.0%|Medium||
|1112|Highest Grade For Each Student||72.4%|Medium||
|1113|Reported Posts||65.4%|Easy||
|1114|Print in Order||68.4%|Easy||
|1115|Print FooBar Alternately||62.5%|Medium||
|1116|Print Zero Even Odd||60.5%|Medium||
|1117|Building H2O||55.9%|Medium||
|1118|Number of Days in a Month||57.1%|Easy||
|1119|Remove Vowels from a String||90.8%|Easy||
|1120|Maximum Average Subtree||65.6%|Medium||
|1121|Divide Array Into Increasing Sequences||60.5%|Hard||
|1122|Relative Sort Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1122.Relative-Sort-Array)|68.6%|Easy||
|1123|Lowest Common Ancestor of Deepest Leaves|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1123.Lowest-Common-Ancestor-of-Deepest-Leaves)|70.9%|Medium||
|1124|Longest Well-Performing Interval||34.6%|Medium||
|1125|Smallest Sufficient Team||46.8%|Hard||
|1126|Active Businesses||67.0%|Medium||
|1127|User Purchase Platform||50.1%|Hard||
|1128|Number of Equivalent Domino Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1128.Number-of-Equivalent-Domino-Pairs)|47.1%|Easy||
|1129|Shortest Path with Alternating Colors||48.5%|Medium||
|1130|Minimum Cost Tree From Leaf Values||68.3%|Medium||
|1131|Maximum of Absolute Value Expression||49.1%|Medium||
|1132|Reported Posts II||33.2%|Medium||
|1133|Largest Unique Number||67.6%|Easy||
|1134|Armstrong Number||77.8%|Easy||
|1135|Connecting Cities With Minimum Cost||61.3%|Medium||
|1136|Parallel Courses||61.6%|Medium||
|1137|N-th Tribonacci Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1137.N-th-Tribonacci-Number)|63.8%|Easy||
|1138|Alphabet Board Path||52.0%|Medium||
|1139|Largest 1-Bordered Square||50.2%|Medium||
|1140|Stone Game II||64.8%|Medium||
|1141|User Activity for the Past 30 Days I||48.7%|Easy||
|1142|User Activity for the Past 30 Days II||35.8%|Easy||
|1143|Longest Common Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1143.Longest-Common-Subsequence)|58.4%|Medium||
|1144|Decrease Elements To Make Array Zigzag||47.3%|Medium||
|1145|Binary Tree Coloring Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1145.Binary-Tree-Coloring-Game)|51.7%|Medium||
|1146|Snapshot Array||37.2%|Medium||
|1147|Longest Chunked Palindrome Decomposition||59.4%|Hard||
|1148|Article Views I||76.3%|Easy||
|1149|Article Views II||47.3%|Medium||
|1150|Check If a Number Is Majority Element in a Sorted Array||57.2%|Easy||
|1151|Minimum Swaps to Group All 1's Together||60.9%|Medium||
|1152|Analyze User Website Visit Pattern||43.1%|Medium||
|1153|String Transforms Into Another String||35.3%|Hard||
|1154|Day of the Year|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1154.Day-of-the-Year)|49.6%|Easy||
|1155|Number of Dice Rolls With Target Sum||53.8%|Medium||
|1156|Swap For Longest Repeated Character Substring||45.2%|Medium||
|1157|Online Majority Element In Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1157.Online-Majority-Element-In-Subarray)|41.7%|Hard||
|1158|Market Analysis I||60.5%|Medium||
|1159|Market Analysis II||58.1%|Hard||
|1160|Find Words That Can Be Formed by Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1160.Find-Words-That-Can-Be-Formed-by-Characters)|67.5%|Easy||
|1161|Maximum Level Sum of a Binary Tree||66.0%|Medium||
|1162|As Far from Land as Possible||51.9%|Medium||
|1163|Last Substring in Lexicographical Order||34.8%|Hard||
|1164|Product Price at a Given Date||67.3%|Medium||
|1165|Single-Row Keyboard||86.5%|Easy||
|1166|Design File System||62.1%|Medium||
|1167|Minimum Cost to Connect Sticks||68.4%|Medium||
|1168|Optimize Water Distribution in a Village||64.5%|Hard||
|1169|Invalid Transactions||31.2%|Medium||
|1170|Compare Strings by Frequency of the Smallest Character|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1170.Compare-Strings-by-Frequency-of-the-Smallest-Character)|61.5%|Medium||
|1171|Remove Zero Sum Consecutive Nodes from Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1171.Remove-Zero-Sum-Consecutive-Nodes-from-Linked-List)|43.2%|Medium||
|1172|Dinner Plate Stacks||33.1%|Hard||
|1173|Immediate Food Delivery I||82.6%|Easy||
|1174|Immediate Food Delivery II||62.9%|Medium||
|1175|Prime Arrangements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1175.Prime-Arrangements)|54.5%|Easy||
|1176|Diet Plan Performance||52.7%|Easy||
|1177|Can Make Palindrome from Substring||38.0%|Medium||
|1178|Number of Valid Words for Each Puzzle|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1178.Number-of-Valid-Words-for-Each-Puzzle)|46.3%|Hard||
|1179|Reformat Department Table||80.8%|Easy||
|1180|Count Substrings with Only One Distinct Letter||79.0%|Easy||
|1181|Before and After Puzzle||45.3%|Medium||
|1182|Shortest Distance to Target Color||55.4%|Medium||
|1183|Maximum Number of Ones||61.3%|Hard||
|1184|Distance Between Bus Stops|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1184.Distance-Between-Bus-Stops)|54.0%|Easy||
|1185|Day of the Week|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1185.Day-of-the-Week)|57.5%|Easy||
|1186|Maximum Subarray Sum with One Deletion||41.4%|Medium||
|1187|Make Array Strictly Increasing||45.3%|Hard||
|1188|Design Bounded Blocking Queue||72.8%|Medium||
|1189|Maximum Number of Balloons|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1189.Maximum-Number-of-Balloons)|61.1%|Easy||
|1190|Reverse Substrings Between Each Pair of Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1190.Reverse-Substrings-Between-Each-Pair-of-Parentheses)|65.9%|Medium||
|1191|K-Concatenation Maximum Sum||23.8%|Medium||
|1192|Critical Connections in a Network||54.6%|Hard||
|1193|Monthly Transactions I||66.0%|Medium||
|1194|Tournament Winners||50.8%|Hard||
|1195|Fizz Buzz Multithreaded||72.9%|Medium||
|1196|How Many Apples Can You Put into the Basket||66.9%|Easy||
|1197|Minimum Knight Moves||39.7%|Medium||
|1198|Find Smallest Common Element in All Rows||76.6%|Medium||
|1199|Minimum Time to Build Blocks||40.8%|Hard||
|1200|Minimum Absolute Difference|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1200.Minimum-Absolute-Difference)|69.6%|Easy||
|1201|Ugly Number III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1201.Ugly-Number-III)|28.8%|Medium||
|1202|Smallest String With Swaps|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1202.Smallest-String-With-Swaps)|57.6%|Medium||
|1203|Sort Items by Groups Respecting Dependencies|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1203.Sort-Items-by-Groups-Respecting-Dependencies)|51.1%|Hard||
|1204|Last Person to Fit in the Bus||73.4%|Medium||
|1205|Monthly Transactions II||43.0%|Medium||
|1206|Design Skiplist||60.4%|Hard||
|1207|Unique Number of Occurrences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1207.Unique-Number-of-Occurrences)|73.5%|Easy||
|1208|Get Equal Substrings Within Budget|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1208.Get-Equal-Substrings-Within-Budget)|48.5%|Medium||
|1209|Remove All Adjacent Duplicates in String II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1209.Remove-All-Adjacent-Duplicates-in-String-II)|56.2%|Medium||
|1210|Minimum Moves to Reach Target with Rotations||49.2%|Hard||
|1211|Queries Quality and Percentage||71.4%|Easy||
|1212|Team Scores in Football Tournament||57.1%|Medium||
|1213|Intersection of Three Sorted Arrays||79.9%|Easy||
|1214|Two Sum BSTs||66.1%|Medium||
|1215|Stepping Numbers||46.2%|Medium||
|1216|Valid Palindrome III||50.4%|Hard||
|1217|Minimum Cost to Move Chips to The Same Position|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1217.Minimum-Cost-to-Move-Chips-to-The-Same-Position)|71.9%|Easy||
|1218|Longest Arithmetic Subsequence of Given Difference||51.9%|Medium||
|1219|Path with Maximum Gold||63.7%|Medium||
|1220|Count Vowels Permutation||60.3%|Hard||
|1221|Split a String in Balanced Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1221.Split-a-String-in-Balanced-Strings)|85.1%|Easy||
|1222|Queens That Can Attack the King||71.8%|Medium||
|1223|Dice Roll Simulation||48.7%|Hard||
|1224|Maximum Equal Frequency||37.0%|Hard||
|1225|Report Contiguous Dates||62.0%|Hard||
|1226|The Dining Philosophers||56.0%|Medium||
|1227|Airplane Seat Assignment Probability||65.6%|Medium||
|1228|Missing Number In Arithmetic Progression||51.5%|Easy||
|1229|Meeting Scheduler||55.2%|Medium||
|1230|Toss Strange Coins||54.0%|Medium||
|1231|Divide Chocolate||57.4%|Hard||
|1232|Check If It Is a Straight Line|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1232.Check-If-It-Is-a-Straight-Line)|40.4%|Easy||
|1233|Remove Sub-Folders from the Filesystem||65.6%|Medium||
|1234|Replace the Substring for Balanced String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1234.Replace-the-Substring-for-Balanced-String)|37.1%|Medium||
|1235|Maximum Profit in Job Scheduling|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1235.Maximum-Profit-in-Job-Scheduling)|53.4%|Hard||
|1236|Web Crawler||66.5%|Medium||
|1237|Find Positive Integer Solution for a Given Equation||69.2%|Medium||
|1238|Circular Permutation in Binary Representation||68.8%|Medium||
|1239|Maximum Length of a Concatenated String with Unique Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1239.Maximum-Length-of-a-Concatenated-String-with-Unique-Characters)|52.2%|Medium||
|1240|Tiling a Rectangle with the Fewest Squares||54.0%|Hard||
|1241|Number of Comments per Post||66.5%|Easy||
|1242|Web Crawler Multithreaded||48.9%|Medium||
|1243|Array Transformation||51.1%|Easy||
|1244|Design A Leaderboard||68.7%|Medium||
|1245|Tree Diameter||61.5%|Medium||
|1246|Palindrome Removal||46.2%|Hard||
|1247|Minimum Swaps to Make Strings Equal||64.0%|Medium||
|1248|Count Number of Nice Subarrays||60.8%|Medium||
|1249|Minimum Remove to Make Valid Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1249.Minimum-Remove-to-Make-Valid-Parentheses)|65.8%|Medium||
|1250|Check If It Is a Good Array||58.9%|Hard||
|1251|Average Selling Price||81.6%|Easy||
|1252|Cells with Odd Values in a Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1252.Cells-with-Odd-Values-in-a-Matrix)|78.5%|Easy||
|1253|Reconstruct a 2-Row Binary Matrix||44.3%|Medium||
|1254|Number of Closed Islands|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1254.Number-of-Closed-Islands)|64.1%|Medium||
|1255|Maximum Score Words Formed by Letters||72.4%|Hard||
|1256|Encode Number||69.9%|Medium||
|1257|Smallest Common Region||64.3%|Medium||
|1258|Synonymous Sentences||56.7%|Medium||
|1259|Handshakes That Don't Cross||60.4%|Hard||
|1260|Shift 2D Grid|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1260.Shift-2D-Grid)|67.8%|Easy||
|1261|Find Elements in a Contaminated Binary Tree||76.3%|Medium||
|1262|Greatest Sum Divisible by Three||50.8%|Medium||
|1263|Minimum Moves to Move a Box to Their Target Location||49.0%|Hard||
|1264|Page Recommendations||66.7%|Medium||
|1265|Print Immutable Linked List in Reverse||94.2%|Medium||
|1266|Minimum Time Visiting All Points|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1266.Minimum-Time-Visiting-All-Points)|79.1%|Easy||
|1267|Count Servers that Communicate||59.4%|Medium||
|1268|Search Suggestions System|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1268.Search-Suggestions-System)|66.3%|Medium||
|1269|Number of Ways to Stay in the Same Place After Some Steps||43.5%|Hard||
|1270|All People Report to the Given Manager||87.0%|Medium||
|1271|Hexspeak||57.1%|Easy||
|1272|Remove Interval||63.6%|Medium||
|1273|Delete Tree Nodes||61.1%|Medium||
|1274|Number of Ships in a Rectangle||69.2%|Hard||
|1275|Find Winner on a Tic Tac Toe Game|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1275.Find-Winner-on-a-Tic-Tac-Toe-Game)|54.2%|Easy||
|1276|Number of Burgers with No Waste of Ingredients||50.7%|Medium||
|1277|Count Square Submatrices with All Ones||74.5%|Medium||
|1278|Palindrome Partitioning III||60.8%|Hard||
|1279|Traffic Light Controlled Intersection||74.4%|Easy||
|1280|Students and Examinations||71.9%|Easy||
|1281|Subtract the Product and Sum of Digits of an Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1281.Subtract-the-Product-and-Sum-of-Digits-of-an-Integer)|86.7%|Easy||
|1282|Group the People Given the Group Size They Belong To||85.6%|Medium||
|1283|Find the Smallest Divisor Given a Threshold|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1283.Find-the-Smallest-Divisor-Given-a-Threshold)|56.0%|Medium||
|1284|Minimum Number of Flips to Convert Binary Matrix to Zero Matrix||71.9%|Hard||
|1285|Find the Start and End Number of Continuous Ranges||86.1%|Medium||
|1286|Iterator for Combination||73.3%|Medium||
|1287|Element Appearing More Than 25% In Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1287.Element-Appearing-More-Than-25-In-Sorted-Array)|59.5%|Easy||
|1288|Remove Covered Intervals||57.0%|Medium||
|1289|Minimum Falling Path Sum II||58.5%|Hard||
|1290|Convert Binary Number in a Linked List to Integer|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1290.Convert-Binary-Number-in-a-Linked-List-to-Integer)|82.2%|Easy||
|1291|Sequential Digits||61.3%|Medium||
|1292|Maximum Side Length of a Square with Sum Less than or Equal to Threshold||53.3%|Medium||
|1293|Shortest Path in a Grid with Obstacles Elimination|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1293.Shortest-Path-in-a-Grid-with-Obstacles-Elimination)|45.4%|Hard||
|1294|Weather Type in Each Country||67.3%|Easy||
|1295|Find Numbers with Even Number of Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1295.Find-Numbers-with-Even-Number-of-Digits)|77.0%|Easy||
|1296|Divide Array in Sets of K Consecutive Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1296.Divide-Array-in-Sets-of-K-Consecutive-Numbers)|56.5%|Medium||
|1297|Maximum Number of Occurrences of a Substring||51.9%|Medium||
|1298|Maximum Candies You Can Get from Boxes||60.2%|Hard||
|1299|Replace Elements with Greatest Element on Right Side|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1299.Replace-Elements-with-Greatest-Element-on-Right-Side)|73.5%|Easy||
|1300|Sum of Mutated Array Closest to Target|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1300.Sum-of-Mutated-Array-Closest-to-Target)|43.6%|Medium||
|1301|Number of Paths with Max Score||38.8%|Hard||
|1302|Deepest Leaves Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1302.Deepest-Leaves-Sum)|86.7%|Medium||
|1303|Find the Team Size||90.1%|Easy||
|1304|Find N Unique Integers Sum up to Zero|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1304.Find-N-Unique-Integers-Sum-up-to-Zero)|77.0%|Easy||
|1305|All Elements in Two Binary Search Trees|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1305.All-Elements-in-Two-Binary-Search-Trees)|79.8%|Medium||
|1306|Jump Game III|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1306.Jump-Game-III)|63.5%|Medium||
|1307|Verbal Arithmetic Puzzle||34.2%|Hard||
|1308|Running Total for Different Genders||87.2%|Medium||
|1309|Decrypt String from Alphabet to Integer Mapping||79.7%|Easy||
|1310|XOR Queries of a Subarray|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1310.XOR-Queries-of-a-Subarray)|72.3%|Medium||
|1311|Get Watched Videos by Your Friends||45.9%|Medium||
|1312|Minimum Insertion Steps to Make a String Palindrome||66.3%|Hard||
|1313|Decompress Run-Length Encoded List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1313.Decompress-Run-Length-Encoded-List)|85.8%|Easy||
|1314|Matrix Block Sum||75.4%|Medium||
|1315|Sum of Nodes with Even-Valued Grandparent||85.5%|Medium||
|1316|Distinct Echo Substrings||49.6%|Hard||
|1317|Convert Integer to the Sum of Two No-Zero Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1317.Convert-Integer-to-the-Sum-of-Two-No-Zero-Integers)|55.5%|Easy||
|1318|Minimum Flips to Make a OR b Equal to c||66.0%|Medium||
|1319|Number of Operations to Make Network Connected|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1319.Number-of-Operations-to-Make-Network-Connected)|58.9%|Medium||
|1320|Minimum Distance to Type a Word Using Two Fingers||59.5%|Hard||
|1321|Restaurant Growth||70.7%|Medium||
|1322|Ads Performance||60.4%|Easy||
|1323|Maximum 69 Number||82.1%|Easy||
|1324|Print Words Vertically||61.4%|Medium||
|1325|Delete Leaves With a Given Value||74.7%|Medium||
|1326|Minimum Number of Taps to Open to Water a Garden||47.6%|Hard||
|1327|List the Products Ordered in a Period||76.3%|Easy||
|1328|Break a Palindrome||52.7%|Medium||
|1329|Sort the Matrix Diagonally|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1329.Sort-the-Matrix-Diagonally)|83.4%|Medium||
|1330|Reverse Subarray To Maximize Array Value||40.2%|Hard||
|1331|Rank Transform of an Array||59.2%|Easy||
|1332|Remove Palindromic Subsequences|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1332.Remove-Palindromic-Subsequences)|76.2%|Easy||
|1333|Filter Restaurants by Vegan-Friendly, Price and Distance||60.0%|Medium||
|1334|Find the City With the Smallest Number of Neighbors at a Threshold Distance||54.4%|Medium||
|1335|Minimum Difficulty of a Job Schedule||58.5%|Hard||
|1336|Number of Transactions per Visit||49.9%|Hard||
|1337|The K Weakest Rows in a Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1337.The-K-Weakest-Rows-in-a-Matrix)|72.1%|Easy||
|1338|Reduce Array Size to The Half||69.4%|Medium||
|1339|Maximum Product of Splitted Binary Tree||47.8%|Medium||
|1340|Jump Game V||62.5%|Hard||
|1341|Movie Rating||57.3%|Medium||
|1342|Number of Steps to Reduce a Number to Zero||85.2%|Easy||
|1343|Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold||67.6%|Medium||
|1344|Angle Between Hands of a Clock||63.3%|Medium||
|1345|Jump Game IV||47.0%|Hard||
|1346|Check If N and Its Double Exist||36.5%|Easy||
|1347|Minimum Number of Steps to Make Two Strings Anagram||77.8%|Medium||
|1348|Tweet Counts Per Frequency||43.9%|Medium||
|1349|Maximum Students Taking Exam||48.9%|Hard||
|1350|Students With Invalid Departments||90.3%|Easy||
|1351|Count Negative Numbers in a Sorted Matrix||75.5%|Easy||
|1352|Product of the Last K Numbers||50.0%|Medium||
|1353|Maximum Number of Events That Can Be Attended|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1353.Maximum-Number-of-Events-That-Can-Be-Attended)|32.6%|Medium||
|1354|Construct Target Array With Multiple Sums||36.3%|Hard||
|1355|Activity Participants||73.8%|Medium||
|1356|Sort Integers by The Number of 1 Bits||72.9%|Easy||
|1357|Apply Discount Every n Orders||69.9%|Medium||
|1358|Number of Substrings Containing All Three Characters||63.3%|Medium||
|1359|Count All Valid Pickup and Delivery Options||62.7%|Hard||
|1360|Number of Days Between Two Dates||48.1%|Easy||
|1361|Validate Binary Tree Nodes||39.9%|Medium||
|1362|Closest Divisors||59.8%|Medium||
|1363|Largest Multiple of Three||33.3%|Hard||
|1364|Number of Trusted Contacts of a Customer||77.7%|Medium||
|1365|How Many Numbers Are Smaller Than the Current Number||86.7%|Easy||
|1366|Rank Teams by Votes||57.9%|Medium||
|1367|Linked List in Binary Tree||43.7%|Medium||
|1368|Minimum Cost to Make at Least One Valid Path in a Grid||61.5%|Hard||
|1369|Get the Second Most Recent Activity||69.1%|Hard||
|1370|Increasing Decreasing String||77.2%|Easy||
|1371|Find the Longest Substring Containing Vowels in Even Counts||63.2%|Medium||
|1372|Longest ZigZag Path in a Binary Tree||60.2%|Medium||
|1373|Maximum Sum BST in Binary Tree||39.4%|Hard||
|1374|Generate a String With Characters That Have Odd Counts||77.6%|Easy||
|1375|Number of Times Binary String Is Prefix-Aligned||65.8%|Medium||
|1376|Time Needed to Inform All Employees||58.3%|Medium||
|1377|Frog Position After T Seconds||35.8%|Hard||
|1378|Replace Employee ID With The Unique Identifier||91.1%|Easy||
|1379|Find a Corresponding Node of a Binary Tree in a Clone of That Tree||86.7%|Easy||
|1380|Lucky Numbers in a Matrix|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1380.Lucky-Numbers-in-a-Matrix)|70.6%|Easy||
|1381|Design a Stack With Increment Operation||77.1%|Medium||
|1382|Balance a Binary Search Tree||80.7%|Medium||
|1383|Maximum Performance of a Team|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1383.Maximum-Performance-of-a-Team)|48.5%|Hard||
|1384|Total Sales Amount by Year||65.4%|Hard||
|1385|Find the Distance Value Between Two Arrays|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1385.Find-the-Distance-Value-Between-Two-Arrays)|66.4%|Easy||
|1386|Cinema Seat Allocation||40.9%|Medium||
|1387|Sort Integers by The Power Value||70.1%|Medium||
|1388|Pizza With 3n Slices||50.2%|Hard||
|1389|Create Target Array in the Given Order|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1389.Create-Target-Array-in-the-Given-Order)|85.8%|Easy||
|1390|Four Divisors||41.2%|Medium||
|1391|Check if There is a Valid Path in a Grid||47.3%|Medium||
|1392|Longest Happy Prefix||44.9%|Hard||
|1393|Capital Gain/Loss||87.7%|Medium||
|1394|Find Lucky Integer in an Array||64.4%|Easy||
|1395|Count Number of Teams||67.3%|Medium||
|1396|Design Underground System|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1396.Design-Underground-System)|73.6%|Medium||
|1397|Find All Good Strings||42.3%|Hard||
|1398|Customers Who Bought Products A and B but Not C||76.0%|Medium||
|1399|Count Largest Group||67.0%|Easy||
|1400|Construct K Palindrome Strings||62.9%|Medium||
|1401|Circle and Rectangle Overlapping||44.1%|Medium||
|1402|Reducing Dishes||72.0%|Hard||
|1403|Minimum Subsequence in Non-Increasing Order||72.3%|Easy||
|1404|Number of Steps to Reduce a Number in Binary Representation to One||52.6%|Medium||
|1405|Longest Happy String||57.4%|Medium||
|1406|Stone Game III||59.5%|Hard||
|1407|Top Travellers||62.2%|Easy||
|1408|String Matching in an Array||63.8%|Easy||
|1409|Queries on a Permutation With Key||83.4%|Medium||
|1410|HTML Entity Parser||51.7%|Medium||
|1411|Number of Ways to Paint N × 3 Grid||62.4%|Hard||
|1412|Find the Quiet Students in All Exams||62.1%|Hard||
|1413|Minimum Value to Get Positive Step by Step Sum||67.4%|Easy||
|1414|Find the Minimum Number of Fibonacci Numbers Whose Sum Is K||65.4%|Medium||
|1415|The k-th Lexicographical String of All Happy Strings of Length n||72.4%|Medium||
|1416|Restore The Array||38.7%|Hard||
|1417|Reformat The String||54.8%|Easy||
|1418|Display Table of Food Orders in a Restaurant||74.1%|Medium||
|1419|Minimum Number of Frogs Croaking||50.2%|Medium||
|1420|Build Array Where You Can Find The Maximum Exactly K Comparisons||63.1%|Hard||
|1421|NPV Queries||83.3%|Easy||
|1422|Maximum Score After Splitting a String||58.0%|Easy||
|1423|Maximum Points You Can Obtain from Cards|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1423.Maximum-Points-You-Can-Obtain-from-Cards)|52.3%|Medium||
|1424|Diagonal Traverse II||50.5%|Medium||
|1425|Constrained Subsequence Sum||47.5%|Hard||
|1426|Counting Elements||59.7%|Easy||
|1427|Perform String Shifts||54.3%|Easy||
|1428|Leftmost Column with at Least a One||53.4%|Medium||
|1429|First Unique Number||52.9%|Medium||
|1430|Check If a String Is a Valid Sequence from Root to Leaves Path in a Binary Tree||46.4%|Medium||
|1431|Kids With the Greatest Number of Candies||87.3%|Easy||
|1432|Max Difference You Can Get From Changing an Integer||42.3%|Medium||
|1433|Check If a String Can Break Another String||68.9%|Medium||
|1434|Number of Ways to Wear Different Hats to Each Other||43.0%|Hard||
|1435|Create a Session Bar Chart||77.1%|Easy||
|1436|Destination City||77.5%|Easy||
|1437|Check If All 1's Are at Least Length K Places Away|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1437.Check-If-All-1s-Are-at-Least-Length-K-Places-Away)|58.8%|Easy||
|1438|Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1438.Longest-Continuous-Subarray-With-Absolute-Diff-Less-Than-or-Equal-to-Limit)|48.3%|Medium||
|1439|Find the Kth Smallest Sum of a Matrix With Sorted Rows|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1439.Find-the-Kth-Smallest-Sum-of-a-Matrix-With-Sorted-Rows)|61.4%|Hard||
|1440|Evaluate Boolean Expression||75.2%|Medium||
|1441|Build an Array With Stack Operations||71.8%|Medium||
|1442|Count Triplets That Can Form Two Arrays of Equal XOR|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1442.Count-Triplets-That-Can-Form-Two-Arrays-of-Equal-XOR)|76.0%|Medium||
|1443|Minimum Time to Collect All Apples in a Tree||62.8%|Medium||
|1444|Number of Ways of Cutting a Pizza||56.1%|Hard||
|1445|Apples & Oranges||89.8%|Medium||
|1446|Consecutive Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1446.Consecutive-Characters)|61.3%|Easy||
|1447|Simplified Fractions||65.1%|Medium||
|1448|Count Good Nodes in Binary Tree||74.3%|Medium||
|1449|Form Largest Integer With Digits That Add up to Target||47.7%|Hard||
|1450|Number of Students Doing Homework at a Given Time||75.9%|Easy||
|1451|Rearrange Words in a Sentence||63.0%|Medium||
|1452|People Whose List of Favorite Companies Is Not a Subset of Another List||56.8%|Medium||
|1453|Maximum Number of Darts Inside of a Circular Dartboard||36.9%|Hard||
|1454|Active Users||37.8%|Medium||
|1455|Check If a Word Occurs As a Prefix of Any Word in a Sentence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1455.Check-If-a-Word-Occurs-As-a-Prefix-of-Any-Word-in-a-Sentence)|64.2%|Easy||
|1456|Maximum Number of Vowels in a Substring of Given Length||58.1%|Medium||
|1457|Pseudo-Palindromic Paths in a Binary Tree||67.8%|Medium||
|1458|Max Dot Product of Two Subsequences||46.6%|Hard||
|1459|Rectangles Area||69.4%|Medium||
|1460|Make Two Arrays Equal by Reversing Subarrays||72.2%|Easy||
|1461|Check If a String Contains All Binary Codes of Size K|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1461.Check-If-a-String-Contains-All-Binary-Codes-of-Size-K)|56.6%|Medium||
|1462|Course Schedule IV||49.0%|Medium||
|1463|Cherry Pickup II|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1463.Cherry-Pickup-II)|69.7%|Hard||
|1464|Maximum Product of Two Elements in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1464.Maximum-Product-of-Two-Elements-in-an-Array)|79.9%|Easy||
|1465|Maximum Area of a Piece of Cake After Horizontal and Vertical Cuts|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1465.Maximum-Area-of-a-Piece-of-Cake-After-Horizontal-and-Vertical-Cuts)|40.9%|Medium||
|1466|Reorder Routes to Make All Paths Lead to the City Zero||61.7%|Medium||
|1467|Probability of a Two Boxes Having The Same Number of Distinct Balls||60.8%|Hard||
|1468|Calculate Salaries||80.7%|Medium||
|1469|Find All The Lonely Nodes||82.1%|Easy||
|1470|Shuffle the Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1470.Shuffle-the-Array)|89.0%|Easy||
|1471|The k Strongest Values in an Array||60.4%|Medium||
|1472|Design Browser History||76.2%|Medium||
|1473|Paint House III||61.6%|Hard||
|1474|Delete N Nodes After M Nodes of a Linked List||73.4%|Easy||
|1475|Final Prices With a Special Discount in a Shop||76.0%|Easy||
|1476|Subrectangle Queries||88.3%|Medium||
|1477|Find Two Non-overlapping Sub-arrays Each With Target Sum||36.8%|Medium||
|1478|Allocate Mailboxes||55.6%|Hard||
|1479|Sales by Day of the Week||81.0%|Hard||
|1480|Running Sum of 1d Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1480.Running-Sum-of-1d-Array)|87.7%|Easy||
|1481|Least Number of Unique Integers after K Removals||55.7%|Medium||
|1482|Minimum Number of Days to Make m Bouquets|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1482.Minimum-Number-of-Days-to-Make-m-Bouquets)|54.2%|Medium||
|1483|Kth Ancestor of a Tree Node||33.8%|Hard||
|1484|Group Sold Products By The Date||80.0%|Easy||
|1485|Clone Binary Tree With Random Pointer||79.7%|Medium||
|1486|XOR Operation in an Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1486.XOR-Operation-in-an-Array)|84.6%|Easy||
|1487|Making File Names Unique||36.2%|Medium||
|1488|Avoid Flood in The City||26.4%|Medium||
|1489|Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree||52.5%|Hard||
|1490|Clone N-ary Tree||83.5%|Medium||
|1491|Average Salary Excluding the Minimum and Maximum Salary||61.4%|Easy||
|1492|The kth Factor of n||62.8%|Medium||
|1493|Longest Subarray of 1's After Deleting One Element||60.3%|Medium||
|1494|Parallel Courses II||30.4%|Hard||
|1495|Friendly Movies Streamed Last Month||49.3%|Easy||
|1496|Path Crossing||56.0%|Easy||
|1497|Check If Array Pairs Are Divisible by k||39.2%|Medium||
|1498|Number of Subsequences That Satisfy the Given Sum Condition||37.5%|Medium||
|1499|Max Value of Equation||45.8%|Hard||
|1500|Design a File Sharing System||44.6%|Medium||
|1501|Countries You Can Safely Invest In||54.3%|Medium||
|1502|Can Make Arithmetic Progression From Sequence||67.7%|Easy||
|1503|Last Moment Before All Ants Fall Out of a Plank||55.6%|Medium||
|1504|Count Submatrices With All Ones||57.5%|Medium||
|1505|Minimum Possible Integer After at Most K Adjacent Swaps On Digits||38.4%|Hard||
|1506|Find Root of N-Ary Tree||78.5%|Medium||
|1507|Reformat Date||63.2%|Easy||
|1508|Range Sum of Sorted Subarray Sums||59.0%|Medium||
|1509|Minimum Difference Between Largest and Smallest Value in Three Moves||54.6%|Medium||
|1510|Stone Game IV||60.4%|Hard||
|1511|Customer Order Frequency||71.5%|Easy||
|1512|Number of Good Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1512.Number-of-Good-Pairs)|88.2%|Easy||
|1513|Number of Substrings With Only 1s||45.4%|Medium||
|1514|Path with Maximum Probability||48.6%|Medium||
|1515|Best Position for a Service Centre||37.3%|Hard||
|1516|Move Sub-Tree of N-Ary Tree||63.4%|Hard||
|1517|Find Users With Valid E-Mails||54.1%|Easy||
|1518|Water Bottles|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1518.Water-Bottles)|60.4%|Easy||
|1519|Number of Nodes in the Sub-Tree With the Same Label||55.3%|Medium||
|1520|Maximum Number of Non-Overlapping Substrings||38.2%|Hard||
|1521|Find a Value of a Mysterious Function Closest to Target||43.5%|Hard||
|1522|Diameter of N-Ary Tree||73.5%|Medium||
|1523|Count Odd Numbers in an Interval Range||49.7%|Easy||
|1524|Number of Sub-arrays With Odd Sum||43.3%|Medium||
|1525|Number of Good Ways to Split a String||69.0%|Medium||
|1526|Minimum Number of Increments on Subarrays to Form a Target Array||68.5%|Hard||
|1527|Patients With a Condition||40.8%|Easy||
|1528|Shuffle String||85.2%|Easy||
|1529|Minimum Suffix Flips||72.6%|Medium||
|1530|Number of Good Leaf Nodes Pairs||61.1%|Medium||
|1531|String Compression II||49.5%|Hard||
|1532|The Most Recent Three Orders||70.0%|Medium||
|1533|Find the Index of the Large Integer||56.0%|Medium||
|1534|Count Good Triplets||80.9%|Easy||
|1535|Find the Winner of an Array Game||48.6%|Medium||
|1536|Minimum Swaps to Arrange a Binary Grid||46.5%|Medium||
|1537|Get the Maximum Score||39.3%|Hard||
|1538|Guess the Majority in a Hidden Array||63.5%|Medium||
|1539|Kth Missing Positive Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1539.Kth-Missing-Positive-Number)|58.5%|Easy||
|1540|Can Convert String in K Moves||33.4%|Medium||
|1541|Minimum Insertions to Balance a Parentheses String||49.8%|Medium||
|1542|Find Longest Awesome Substring||41.4%|Hard||
|1543|Fix Product Name Format||60.9%|Easy||
|1544|Make The String Great||63.3%|Easy||
|1545|Find Kth Bit in Nth Binary String||58.5%|Medium||
|1546|Maximum Number of Non-Overlapping Subarrays With Sum Equals Target||47.4%|Medium||
|1547|Minimum Cost to Cut a Stick||57.1%|Hard||
|1548|The Most Similar Path in a Graph||57.1%|Hard||
|1549|The Most Recent Orders for Each Product||67.4%|Medium||
|1550|Three Consecutive Odds||63.5%|Easy||
|1551|Minimum Operations to Make Array Equal|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1551.Minimum-Operations-to-Make-Array-Equal)|81.5%|Medium||
|1552|Magnetic Force Between Two Balls||58.0%|Medium||
|1553|Minimum Number of Days to Eat N Oranges||34.7%|Hard||
|1554|Strings Differ by One Character||42.6%|Medium||
|1555|Bank Account Summary||52.3%|Medium||
|1556|Thousand Separator||54.6%|Easy||
|1557|Minimum Number of Vertices to Reach All Nodes||79.4%|Medium||
|1558|Minimum Numbers of Function Calls to Make Target Array||63.9%|Medium||
|1559|Detect Cycles in 2D Grid||48.0%|Medium||
|1560|Most Visited Sector in  a Circular Track||58.6%|Easy||
|1561|Maximum Number of Coins You Can Get||79.1%|Medium||
|1562|Find Latest Group of Size M||42.8%|Medium||
|1563|Stone Game V||40.3%|Hard||
|1564|Put Boxes Into the Warehouse I||67.0%|Medium||
|1565|Unique Orders and Customers Per Month||82.9%|Easy||
|1566|Detect Pattern of Length M Repeated K or More Times||43.4%|Easy||
|1567|Maximum Length of Subarray With Positive Product||44.0%|Medium||
|1568|Minimum Number of Days to Disconnect Island||46.5%|Hard||
|1569|Number of Ways to Reorder Array to Get Same BST||47.8%|Hard||
|1570|Dot Product of Two Sparse Vectors||90.3%|Medium||
|1571|Warehouse Manager||89.0%|Easy||
|1572|Matrix Diagonal Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1572.Matrix-Diagonal-Sum)|80.2%|Easy||
|1573|Number of Ways to Split a String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1573.Number-of-Ways-to-Split-a-String)|32.5%|Medium||
|1574|Shortest Subarray to be Removed to Make Array Sorted||37.0%|Medium||
|1575|Count All Possible Routes||56.8%|Hard||
|1576|Replace All ?'s to Avoid Consecutive Repeating Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1576.Replace-All-s-to-Avoid-Consecutive-Repeating-Characters)|48.3%|Easy||
|1577|Number of Ways Where Square of Number Is Equal to Product of Two Numbers||40.1%|Medium||
|1578|Minimum Time to Make Rope Colorful||63.5%|Medium||
|1579|Remove Max Number of Edges to Keep Graph Fully Traversable|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1579.Remove-Max-Number-of-Edges-to-Keep-Graph-Fully-Traversable)|53.2%|Hard||
|1580|Put Boxes Into the Warehouse II||63.7%|Medium||
|1581|Customer Who Visited but Did Not Make Any Transactions||85.2%|Easy||
|1582|Special Positions in a Binary Matrix||65.4%|Easy||
|1583|Count Unhappy Friends||60.9%|Medium||
|1584|Min Cost to Connect All Points||64.0%|Medium||
|1585|Check If String Is Transformable With Substring Sort Operations||47.9%|Hard||
|1586|Binary Search Tree Iterator II||70.9%|Medium||
|1587|Bank Account Summary II||86.3%|Easy||
|1588|Sum of All Odd Length Subarrays||83.4%|Easy||
|1589|Maximum Sum Obtained of Any Permutation||37.2%|Medium||
|1590|Make Sum Divisible by P||28.1%|Medium||
|1591|Strange Printer II||58.6%|Hard||
|1592|Rearrange Spaces Between Words||43.5%|Easy||
|1593|Split a String Into the Max Number of Unique Substrings||55.3%|Medium||
|1594|Maximum Non Negative Product in a Matrix||33.1%|Medium||
|1595|Minimum Cost to Connect Two Groups of Points||46.5%|Hard||
|1596|The Most Frequently Ordered Products for Each Customer||83.5%|Medium||
|1597|Build Binary Expression Tree From Infix Expression||62.7%|Hard||
|1598|Crawler Log Folder||64.5%|Easy||
|1599|Maximum Profit of Operating a Centennial Wheel||43.8%|Medium||
|1600|Throne Inheritance|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1600.Throne-Inheritance)|63.6%|Medium||
|1601|Maximum Number of Achievable Transfer Requests||51.2%|Hard||
|1602|Find Nearest Right Node in Binary Tree||75.6%|Medium||
|1603|Design Parking System|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1603.Design-Parking-System)|87.9%|Easy||
|1604|Alert Using Same Key-Card Three or More Times in a One Hour Period||47.0%|Medium||
|1605|Find Valid Matrix Given Row and Column Sums||78.0%|Medium||
|1606|Find Servers That Handled Most Number of Requests||42.9%|Hard||
|1607|Sellers With No Sales||54.9%|Easy||
|1608|Special Array With X Elements Greater Than or Equal X|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1608.Special-Array-With-X-Elements-Greater-Than-or-Equal-X)|60.4%|Easy||
|1609|Even Odd Tree|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1609.Even-Odd-Tree)|54.3%|Medium||
|1610|Maximum Number of Visible Points||37.3%|Hard||
|1611|Minimum One Bit Operations to Make Integers Zero||63.3%|Hard||
|1612|Check If Two Expression Trees are Equivalent||70.2%|Medium||
|1613|Find the Missing IDs||75.1%|Medium||
|1614|Maximum Nesting Depth of the Parentheses|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1614.Maximum-Nesting-Depth-of-the-Parentheses)|82.3%|Easy||
|1615|Maximal Network Rank||58.5%|Medium||
|1616|Split Two Strings to Make Palindrome||31.2%|Medium||
|1617|Count Subtrees With Max Distance Between Cities||66.1%|Hard||
|1618|Maximum Font to Fit a Sentence in a Screen||59.7%|Medium||
|1619|Mean of Array After Removing Some Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1619.Mean-of-Array-After-Removing-Some-Elements)|65.6%|Easy||
|1620|Coordinate With Maximum Network Quality||37.7%|Medium||
|1621|Number of Sets of K Non-Overlapping Line Segments||42.3%|Medium||
|1622|Fancy Sequence||16.1%|Hard||
|1623|All Valid Triplets That Can Represent a Country||86.8%|Easy||
|1624|Largest Substring Between Two Equal Characters|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1624.Largest-Substring-Between-Two-Equal-Characters)|59.1%|Easy||
|1625|Lexicographically Smallest String After Applying Operations||65.5%|Medium||
|1626|Best Team With No Conflicts||51.1%|Medium||
|1627|Graph Connectivity With Threshold||46.2%|Hard||
|1628|Design an Expression Tree With Evaluate Function||82.8%|Medium||
|1629|Slowest Key|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1629.Slowest-Key)|59.2%|Easy||
|1630|Arithmetic Subarrays||80.3%|Medium||
|1631|Path With Minimum Effort|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1631.Path-With-Minimum-Effort)|55.6%|Medium||
|1632|Rank Transform of a Matrix||40.9%|Hard||
|1633|Percentage of Users Attended a Contest||66.9%|Easy||
|1634|Add Two Polynomials Represented as Linked Lists||54.4%|Medium||
|1635|Hopper Company Queries I||51.0%|Hard||
|1636|Sort Array by Increasing Frequency|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1636.Sort-Array-by-Increasing-Frequency)|69.5%|Easy||
|1637|Widest Vertical Area Between Two Points Containing No Points||84.3%|Medium||
|1638|Count Substrings That Differ by One Character||71.4%|Medium||
|1639|Number of Ways to Form a Target String Given a Dictionary||43.2%|Hard||
|1640|Check Array Formation Through Concatenation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1640.Check-Array-Formation-Through-Concatenation)|56.2%|Easy||
|1641|Count Sorted Vowel Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1641.Count-Sorted-Vowel-Strings)|77.4%|Medium||
|1642|Furthest Building You Can Reach|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1642.Furthest-Building-You-Can-Reach)|48.3%|Medium||
|1643|Kth Smallest Instructions||47.0%|Hard||
|1644|Lowest Common Ancestor of a Binary Tree II||59.9%|Medium||
|1645|Hopper Company Queries II||38.9%|Hard||
|1646|Get Maximum in Generated Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1646.Get-Maximum-in-Generated-Array)|50.2%|Easy||
|1647|Minimum Deletions to Make Character Frequencies Unique|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1647.Minimum-Deletions-to-Make-Character-Frequencies-Unique)|59.1%|Medium||
|1648|Sell Diminishing-Valued Colored Balls|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1648.Sell-Diminishing-Valued-Colored-Balls)|30.4%|Medium||
|1649|Create Sorted Array through Instructions|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1649.Create-Sorted-Array-through-Instructions)|37.3%|Hard||
|1650|Lowest Common Ancestor of a Binary Tree III||77.4%|Medium||
|1651|Hopper Company Queries III||68.1%|Hard||
|1652|Defuse the Bomb|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1652.Defuse-the-Bomb)|62.2%|Easy||
|1653|Minimum Deletions to Make String Balanced|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1653.Minimum-Deletions-to-Make-String-Balanced)|59.1%|Medium||
|1654|Minimum Jumps to Reach Home|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1654.Minimum-Jumps-to-Reach-Home)|29.0%|Medium||
|1655|Distribute Repeating Integers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1655.Distribute-Repeating-Integers)|38.9%|Hard||
|1656|Design an Ordered Stream|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1656.Design-an-Ordered-Stream)|85.3%|Easy||
|1657|Determine if Two Strings Are Close|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1657.Determine-if-Two-Strings-Are-Close)|56.3%|Medium||
|1658|Minimum Operations to Reduce X to Zero|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1658.Minimum-Operations-to-Reduce-X-to-Zero)|37.6%|Medium||
|1659|Maximize Grid Happiness|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1659.Maximize-Grid-Happiness)|38.2%|Hard||
|1660|Correct a Binary Tree||72.5%|Medium||
|1661|Average Time of Process per Machine||78.5%|Easy||
|1662|Check If Two String Arrays are Equivalent|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1662.Check-If-Two-String-Arrays-are-Equivalent)|83.4%|Easy||
|1663|Smallest String With A Given Numeric Value|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1663.Smallest-String-With-A-Given-Numeric-Value)|66.8%|Medium||
|1664|Ways to Make a Fair Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1664.Ways-to-Make-a-Fair-Array)|63.5%|Medium||
|1665|Minimum Initial Energy to Finish Tasks|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1665.Minimum-Initial-Energy-to-Finish-Tasks)|56.3%|Hard||
|1666|Change the Root of a Binary Tree||73.1%|Medium||
|1667|Fix Names in a Table||65.2%|Easy||
|1668|Maximum Repeating Substring|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1668.Maximum-Repeating-Substring)|39.5%|Easy||
|1669|Merge In Between Linked Lists|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1669.Merge-In-Between-Linked-Lists)|73.8%|Medium||
|1670|Design Front Middle Back Queue|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1670.Design-Front-Middle-Back-Queue)|57.3%|Medium||
|1671|Minimum Number of Removals to Make Mountain Array||42.5%|Hard||
|1672|Richest Customer Wealth|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1672.Richest-Customer-Wealth)|87.9%|Easy||
|1673|Find the Most Competitive Subsequence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1673.Find-the-Most-Competitive-Subsequence)|49.3%|Medium||
|1674|Minimum Moves to Make Array Complementary|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1674.Minimum-Moves-to-Make-Array-Complementary)|38.6%|Medium||
|1675|Minimize Deviation in Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1675.Minimize-Deviation-in-Array)|54.7%|Hard||
|1676|Lowest Common Ancestor of a Binary Tree IV||79.5%|Medium||
|1677|Product's Worth Over Invoices||38.4%|Easy||
|1678|Goal Parser Interpretation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1678.Goal-Parser-Interpretation)|86.6%|Easy||
|1679|Max Number of K-Sum Pairs|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1679.Max-Number-of-K-Sum-Pairs)|57.3%|Medium||
|1680|Concatenation of Consecutive Binary Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1680.Concatenation-of-Consecutive-Binary-Numbers)|57.0%|Medium||
|1681|Minimum Incompatibility|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1681.Minimum-Incompatibility)|37.4%|Hard||
|1682|Longest Palindromic Subsequence II||49.6%|Medium||
|1683|Invalid Tweets||90.7%|Easy||
|1684|Count the Number of Consistent Strings|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1684.Count-the-Number-of-Consistent-Strings)|82.2%|Easy||
|1685|Sum of Absolute Differences in a Sorted Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1685.Sum-of-Absolute-Differences-in-a-Sorted-Array)|63.8%|Medium||
|1686|Stone Game VI||54.8%|Medium||
|1687|Delivering Boxes from Storage to Ports||38.7%|Hard||
|1688|Count of Matches in Tournament|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1688.Count-of-Matches-in-Tournament)|83.2%|Easy||
|1689|Partitioning Into Minimum Number Of Deci-Binary Numbers|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1689.Partitioning-Into-Minimum-Number-Of-Deci-Binary-Numbers)|89.3%|Medium||
|1690|Stone Game VII|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1690.Stone-Game-VII)|58.2%|Medium||
|1691|Maximum Height by Stacking Cuboids|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1691.Maximum-Height-by-Stacking-Cuboids)|54.4%|Hard||
|1692|Count Ways to Distribute Candies||62.0%|Hard||
|1693|Daily Leads and Partners||88.4%|Easy||
|1694|Reformat Phone Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1694.Reformat-Phone-Number)|65.0%|Easy||
|1695|Maximum Erasure Value|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1695.Maximum-Erasure-Value)|57.6%|Medium||
|1696|Jump Game VI|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1696.Jump-Game-VI)|46.1%|Medium||
|1697|Checking Existence of Edge Length Limited Paths||50.7%|Hard||
|1698|Number of Distinct Substrings in a String||63.5%|Medium||
|1699|Number of Calls Between Two Persons||83.1%|Medium||
|1700|Number of Students Unable to Eat Lunch|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1700.Number-of-Students-Unable-to-Eat-Lunch)|68.6%|Easy||
|1701|Average Waiting Time||62.6%|Medium||
|1702|Maximum Binary String After Change||46.3%|Medium||
|1703|Minimum Adjacent Swaps for K Consecutive Ones||42.3%|Hard||
|1704|Determine if String Halves Are Alike|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1704.Determine-if-String-Halves-Are-Alike)|77.8%|Easy||
|1705|Maximum Number of Eaten Apples|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1705.Maximum-Number-of-Eaten-Apples)|38.0%|Medium||
|1706|Where Will the Ball Fall||71.5%|Medium||
|1707|Maximum XOR With an Element From Array||44.8%|Hard||
|1708|Largest Subarray Length K||63.9%|Easy||
|1709|Biggest Window Between Visits||75.8%|Medium||
|1710|Maximum Units on a Truck|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1710.Maximum-Units-on-a-Truck)|73.8%|Easy||
|1711|Count Good Meals||29.2%|Medium||
|1712|Ways to Split Array Into Three Subarrays||32.7%|Medium||
|1713|Minimum Operations to Make a Subsequence||48.9%|Hard||
|1714|Sum Of Special Evenly-Spaced Elements In Array||49.1%|Hard||
|1715|Count Apples and Oranges||76.3%|Medium||
|1716|Calculate Money in Leetcode Bank|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1716.Calculate-Money-in-Leetcode-Bank)|66.1%|Easy||
|1717|Maximum Score From Removing Substrings||46.2%|Medium||
|1718|Construct the Lexicographically Largest Valid Sequence||52.0%|Medium||
|1719|Number Of Ways To Reconstruct A Tree||43.6%|Hard||
|1720|Decode XORed Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1720.Decode-XORed-Array)|85.8%|Easy||
|1721|Swapping Nodes in a Linked List|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1721.Swapping-Nodes-in-a-Linked-List)|67.2%|Medium||
|1722|Minimize Hamming Distance After Swap Operations||48.7%|Medium||
|1723|Find Minimum Time to Finish All Jobs||42.5%|Hard||
|1724|Checking Existence of Edge Length Limited Paths II||52.5%|Hard||
|1725|Number Of Rectangles That Can Form The Largest Square|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1725.Number-Of-Rectangles-That-Can-Form-The-Largest-Square)|78.6%|Easy||
|1726|Tuple with Same Product||60.8%|Medium||
|1727|Largest Submatrix With Rearrangements||61.2%|Medium||
|1728|Cat and Mouse II||39.9%|Hard||
|1729|Find Followers Count||71.2%|Easy||
|1730|Shortest Path to Get Food||54.1%|Medium||
|1731|The Number of Employees Which Report to Each Employee||49.5%|Easy||
|1732|Find the Highest Altitude|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1732.Find-the-Highest-Altitude)|78.8%|Easy||
|1733|Minimum Number of People to Teach||41.9%|Medium||
|1734|Decode XORed Permutation|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1734.Decode-XORed-Permutation)|63.0%|Medium||
|1735|Count Ways to Make Array With Product||50.3%|Hard||
|1736|Latest Time by Replacing Hidden Digits|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1736.Latest-Time-by-Replacing-Hidden-Digits)|42.4%|Easy||
|1737|Change Minimum Characters to Satisfy One of Three Conditions||35.3%|Medium||
|1738|Find Kth Largest XOR Coordinate Value|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1738.Find-Kth-Largest-XOR-Coordinate-Value)|61.0%|Medium||
|1739|Building Boxes||52.0%|Hard||
|1740|Find Distance in a Binary Tree||69.1%|Medium||
|1741|Find Total Time Spent by Each Employee||89.8%|Easy||
|1742|Maximum Number of Balls in a Box|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1742.Maximum-Number-of-Balls-in-a-Box)|73.7%|Easy||
|1743|Restore the Array From Adjacent Pairs||68.6%|Medium||
|1744|Can You Eat Your Favorite Candy on Your Favorite Day?|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1744.Can-You-Eat-Your-Favorite-Candy-on-Your-Favorite-Day)|33.0%|Medium||
|1745|Palindrome Partitioning IV||45.4%|Hard||
|1746|Maximum Subarray Sum After One Operation||62.2%|Medium||
|1747|Leetflex Banned Accounts||67.3%|Medium||
|1748|Sum of Unique Elements|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1748.Sum-of-Unique-Elements)|76.2%|Easy||
|1749|Maximum Absolute Sum of Any Subarray||58.4%|Medium||
|1750|Minimum Length of String After Deleting Similar Ends||43.9%|Medium||
|1751|Maximum Number of Events That Can Be Attended II||56.9%|Hard||
|1752|Check if Array Is Sorted and Rotated|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1752.Check-if-Array-Is-Sorted-and-Rotated)|50.0%|Easy||
|1753|Maximum Score From Removing Stones||66.6%|Medium||
|1754|Largest Merge Of Two Strings||45.8%|Medium||
|1755|Closest Subsequence Sum||36.7%|Hard||
|1756|Design Most Recently Used Queue||78.5%|Medium||
|1757|Recyclable and Low Fat Products||92.7%|Easy||
|1758|Minimum Changes To Make Alternating Binary String|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1758.Minimum-Changes-To-Make-Alternating-Binary-String)|58.3%|Easy||
|1759|Count Number of Homogenous Substrings||48.4%|Medium||
|1760|Minimum Limit of Balls in a Bag||60.4%|Medium||
|1761|Minimum Degree of a Connected Trio in a Graph||41.6%|Hard||
|1762|Buildings With an Ocean View||79.1%|Medium||
|1763|Longest Nice Substring|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1763.Longest-Nice-Substring)|61.6%|Easy||
|1764|Form Array by Concatenating Subarrays of Another Array||52.9%|Medium||
|1765|Map of Highest Peak||60.2%|Medium||
|1766|Tree of Coprimes||39.3%|Hard||
|1767|Find the Subtasks That Did Not Execute||81.5%|Hard||
|1768|Merge Strings Alternately||77.3%|Easy||
|1769|Minimum Number of Operations to Move All Balls to Each Box||85.1%|Medium||
|1770|Maximum Score from Performing Multiplication Operations||37.4%|Hard||
|1771|Maximize Palindrome Length From Subsequences||35.3%|Hard||
|1772|Sort Features by Popularity||65.1%|Medium||
|1773|Count Items Matching a Rule||84.4%|Easy||
|1774|Closest Dessert Cost||47.4%|Medium||
|1775|Equal Sum Arrays With Minimum Number of Operations||53.0%|Medium||
|1776|Car Fleet II||53.5%|Hard||
|1777|Product's Price for Each Store||84.2%|Easy||
|1778|Shortest Path in a Hidden Grid||39.1%|Medium||
|1779|Find Nearest Point That Has the Same X or Y Coordinate||67.3%|Easy||
|1780|Check if Number is a Sum of Powers of Three||65.8%|Medium||
|1781|Sum of Beauty of All Substrings||60.8%|Medium||
|1782|Count Pairs Of Nodes||38.4%|Hard||
|1783|Grand Slam Titles||86.9%|Medium||
|1784|Check if Binary String Has at Most One Segment of Ones||40.2%|Easy||
|1785|Minimum Elements to Add to Form a Given Sum||42.6%|Medium||
|1786|Number of Restricted Paths From First to Last Node||39.3%|Medium||
|1787|Make the XOR of All Segments Equal to Zero||39.4%|Hard||
|1788|Maximize the Beauty of the Garden||66.6%|Hard||
|1789|Primary Department for Each Employee||79.3%|Easy||
|1790|Check if One String Swap Can Make Strings Equal||45.4%|Easy||
|1791|Find Center of Star Graph|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1791.Find-Center-of-Star-Graph)|83.5%|Easy||
|1792|Maximum Average Pass Ratio||52.3%|Medium||
|1793|Maximum Score of a Good Subarray||53.5%|Hard||
|1794|Count Pairs of Equal Substrings With Minimum Difference||65.1%|Medium||
|1795|Rearrange Products Table||86.9%|Easy||
|1796|Second Largest Digit in a String||49.3%|Easy||
|1797|Design Authentication Manager||56.9%|Medium||
|1798|Maximum Number of Consecutive Values You Can Make||55.7%|Medium||
|1799|Maximize Score After N Operations||46.0%|Hard||
|1800|Maximum Ascending Subarray Sum||63.2%|Easy||
|1801|Number of Orders in the Backlog||48.2%|Medium||
|1802|Maximum Value at a Given Index in a Bounded Array||32.4%|Medium||
|1803|Count Pairs With XOR in a Range||47.0%|Hard||
|1804|Implement Trie II (Prefix Tree)||59.7%|Medium||
|1805|Number of Different Integers in a String||36.6%|Easy||
|1806|Minimum Number of Operations to Reinitialize a Permutation||72.0%|Medium||
|1807|Evaluate the Bracket Pairs of a String||66.6%|Medium||
|1808|Maximize Number of Nice Divisors||31.7%|Hard||
|1809|Ad-Free Sessions||59.0%|Easy||
|1810|Minimum Path Cost in a Hidden Grid||54.5%|Medium||
|1811|Find Interview Candidates||63.5%|Medium||
|1812|Determine Color of a Chessboard Square||77.7%|Easy||
|1813|Sentence Similarity III||33.3%|Medium||
|1814|Count Nice Pairs in an Array||42.0%|Medium||
|1815|Maximum Number of Groups Getting Fresh Donuts||40.1%|Hard||
|1816|Truncate Sentence|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1816.Truncate-Sentence)|83.0%|Easy||
|1817|Finding the Users Active Minutes||80.4%|Medium||
|1818|Minimum Absolute Sum Difference|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1818.Minimum-Absolute-Sum-Difference)|30.3%|Medium||
|1819|Number of Different Subsequences GCDs||39.0%|Hard||
|1820|Maximum Number of Accepted Invitations||49.5%|Medium||
|1821|Find Customers With Positive Revenue this Year||89.1%|Easy||
|1822|Sign of the Product of an Array||65.3%|Easy||
|1823|Find the Winner of the Circular Game||78.2%|Medium||
|1824|Minimum Sideway Jumps||49.4%|Medium||
|1825|Finding MK Average||35.6%|Hard||
|1826|Faulty Sensor||49.8%|Easy||
|1827|Minimum Operations to Make the Array Increasing||78.3%|Easy||
|1828|Queries on Number of Points Inside a Circle||86.3%|Medium||
|1829|Maximum XOR for Each Query||76.8%|Medium||
|1830|Minimum Number of Operations to Make String Sorted||49.1%|Hard||
|1831|Maximum Transaction Each Day||83.1%|Medium||
|1832|Check if the Sentence Is Pangram||83.6%|Easy||
|1833|Maximum Ice Cream Bars||74.0%|Medium||
|1834|Single-Threaded CPU||45.8%|Medium||
|1835|Find XOR Sum of All Pairs Bitwise AND||60.8%|Hard||
|1836|Remove Duplicates From an Unsorted Linked List||70.5%|Medium||
|1837|Sum of Digits in Base K||77.0%|Easy||
|1838|Frequency of the Most Frequent Element||39.2%|Medium||
|1839|Longest Substring Of All Vowels in Order||48.6%|Medium||
|1840|Maximum Building Height||35.5%|Hard||
|1841|League Statistics||55.3%|Medium||
|1842|Next Palindrome Using Same Digits||53.1%|Hard||
|1843|Suspicious Bank Accounts||47.3%|Medium||
|1844|Replace All Digits with Characters||80.2%|Easy||
|1845|Seat Reservation Manager||65.6%|Medium||
|1846|Maximum Element After Decreasing and Rearranging|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1846.Maximum-Element-After-Decreasing-and-Rearranging)|58.9%|Medium||
|1847|Closest Room||35.6%|Hard||
|1848|Minimum Distance to the Target Element||57.9%|Easy||
|1849|Splitting a String Into Descending Consecutive Values||33.1%|Medium||
|1850|Minimum Adjacent Swaps to Reach the Kth Smallest Number||71.8%|Medium||
|1851|Minimum Interval to Include Each Query||48.2%|Hard||
|1852|Distinct Numbers in Each Subarray||71.2%|Medium||
|1853|Convert Date Format||86.7%|Easy||
|1854|Maximum Population Year||60.0%|Easy||
|1855|Maximum Distance Between a Pair of Values||53.1%|Medium||
|1856|Maximum Subarray Min-Product||37.8%|Medium||
|1857|Largest Color Value in a Directed Graph||40.8%|Hard||
|1858|Longest Word With All Prefixes||67.1%|Medium||
|1859|Sorting the Sentence||84.1%|Easy||
|1860|Incremental Memory Leak||71.6%|Medium||
|1861|Rotating the Box||65.3%|Medium||
|1862|Sum of Floored Pairs||28.1%|Hard||
|1863|Sum of All Subset XOR Totals||79.9%|Easy||
|1864|Minimum Number of Swaps to Make the Binary String Alternating||42.8%|Medium||
|1865|Finding Pairs With a Certain Sum||50.3%|Medium||
|1866|Number of Ways to Rearrange Sticks With K Sticks Visible||55.6%|Hard||
|1867|Orders With Maximum Quantity Above Average||73.9%|Medium||
|1868|Product of Two Run-Length Encoded Arrays||57.7%|Medium||
|1869|Longer Contiguous Segments of Ones than Zeros||60.4%|Easy||
|1870|Minimum Speed to Arrive on Time||38.4%|Medium||
|1871|Jump Game VII||25.1%|Medium||
|1872|Stone Game VIII||52.7%|Hard||
|1873|Calculate Special Bonus||59.9%|Easy||
|1874|Minimize Product Sum of Two Arrays||90.1%|Medium||
|1875|Group Employees of the Same Salary||73.4%|Medium||
|1876|Substrings of Size Three with Distinct Characters||71.1%|Easy||
|1877|Minimize Maximum Pair Sum in Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1877.Minimize-Maximum-Pair-Sum-in-Array)|79.9%|Medium||
|1878|Get Biggest Three Rhombus Sums in a Grid||47.1%|Medium||
|1879|Minimum XOR Sum of Two Arrays||45.0%|Hard||
|1880|Check if Word Equals Summation of Two Words||73.9%|Easy||
|1881|Maximum Value after Insertion||36.8%|Medium||
|1882|Process Tasks Using Servers||39.5%|Medium||
|1883|Minimum Skips to Arrive at Meeting On Time||38.2%|Hard||
|1884|Egg Drop With 2 Eggs and N Floors||70.8%|Medium||
|1885|Count Pairs in Two Arrays||59.4%|Medium||
|1886|Determine Whether Matrix Can Be Obtained By Rotation||55.8%|Easy||
|1887|Reduction Operations to Make the Array Elements Equal||62.5%|Medium||
|1888|Minimum Number of Flips to Make the Binary String Alternating||38.8%|Medium||
|1889|Minimum Space Wasted From Packaging||31.2%|Hard||
|1890|The Latest Login in 2020||79.0%|Easy||
|1891|Cutting Ribbons||48.3%|Medium||
|1892|Page Recommendations II||43.6%|Hard||
|1893|Check if All the Integers in a Range Are Covered||50.4%|Easy||
|1894|Find the Student that Will Replace the Chalk||44.3%|Medium||
|1895|Largest Magic Square||52.0%|Medium||
|1896|Minimum Cost to Change the Final Value of Expression||54.7%|Hard||
|1897|Redistribute Characters to Make All Strings Equal||59.6%|Easy||
|1898|Maximum Number of Removable Characters||40.5%|Medium||
|1899|Merge Triplets to Form Target Triplet||64.6%|Medium||
|1900|The Earliest and Latest Rounds Where Players Compete||51.8%|Hard||
|1901|Find a Peak Element II||52.4%|Medium||
|1902|Depth of BST Given Insertion Order||44.5%|Medium||
|1903|Largest Odd Number in String||56.0%|Easy||
|1904|The Number of Full Rounds You Have Played||44.9%|Medium||
|1905|Count Sub Islands||67.6%|Medium||
|1906|Minimum Absolute Difference Queries||43.9%|Medium||
|1907|Count Salary Categories||63.5%|Medium||
|1908|Game of Nim||64.5%|Medium||
|1909|Remove One Element to Make the Array Strictly Increasing||26.0%|Easy||
|1910|Remove All Occurrences of a Substring||74.5%|Medium||
|1911|Maximum Alternating Subsequence Sum||59.3%|Medium||
|1912|Design Movie Rental System||39.9%|Hard||
|1913|Maximum Product Difference Between Two Pairs||81.3%|Easy||
|1914|Cyclically Rotating a Grid||48.4%|Medium||
|1915|Number of Wonderful Substrings||45.2%|Medium||
|1916|Count Ways to Build Rooms in an Ant Colony||49.4%|Hard||
|1917|Leetcodify Friends Recommendations||28.2%|Hard||
|1918|Kth Smallest Subarray Sum||52.6%|Medium||
|1919|Leetcodify Similar Friends||43.0%|Hard||
|1920|Build Array from Permutation||90.3%|Easy||
|1921|Eliminate Maximum Number of Monsters||37.4%|Medium||
|1922|Count Good Numbers||39.8%|Medium||
|1923|Longest Common Subpath||27.6%|Hard||
|1924|Erect the Fence II||53.1%|Hard||
|1925|Count Square Sum Triples||68.1%|Easy||
|1926|Nearest Exit from Entrance in Maze||48.8%|Medium||
|1927|Sum Game||46.8%|Medium||
|1928|Minimum Cost to Reach Destination in Time||37.8%|Hard||
|1929|Concatenation of Array||90.3%|Easy||
|1930|Unique Length-3 Palindromic Subsequences||52.0%|Medium||
|1931|Painting a Grid With Three Different Colors||57.1%|Hard||
|1932|Merge BSTs to Create Single BST||35.5%|Hard||
|1933|Check if String Is Decomposable Into Value-Equal Substrings||50.5%|Easy||
|1934|Confirmation Rate||76.5%|Medium||
|1935|Maximum Number of Words You Can Type||71.5%|Easy||
|1936|Add Minimum Number of Rungs||42.7%|Medium||
|1937|Maximum Number of Points with Cost||36.1%|Medium||
|1938|Maximum Genetic Difference Query||39.6%|Hard||
|1939|Users That Actively Request Confirmation Messages||59.3%|Easy||
|1940|Longest Common Subsequence Between Sorted Arrays||79.1%|Medium||
|1941|Check if All Characters Have Equal Number of Occurrences||76.8%|Easy||
|1942|The Number of the Smallest Unoccupied Chair||40.7%|Medium||
|1943|Describe the Painting||48.2%|Medium||
|1944|Number of Visible People in a Queue||69.2%|Hard||
|1945|Sum of Digits of String After Convert||61.4%|Easy||
|1946|Largest Number After Mutating Substring||34.9%|Medium||
|1947|Maximum Compatibility Score Sum||61.0%|Medium||
|1948|Delete Duplicate Folders in System||56.6%|Hard||
|1949|Strong Friendship||56.5%|Medium||
|1950|Maximum of Minimum Values in All Subarrays||49.8%|Medium||
|1951|All the Pairs With the Maximum Number of Common Followers||71.0%|Medium||
|1952|Three Divisors||58.1%|Easy||
|1953|Maximum Number of Weeks for Which You Can Work||39.4%|Medium||
|1954|Minimum Garden Perimeter to Collect Enough Apples||53.3%|Medium||
|1955|Count Number of Special Subsequences||51.2%|Hard||
|1956|Minimum Time For K Virus Variants to Spread||48.0%|Hard||
|1957|Delete Characters to Make Fancy String||57.0%|Easy||
|1958|Check if Move is Legal||45.4%|Medium||
|1959|Minimum Total Space Wasted With K Resizing Operations||42.3%|Medium||
|1960|Maximum Product of the Length of Two Palindromic Substrings||29.8%|Hard||
|1961|Check If String Is a Prefix of Array||53.5%|Easy||
|1962|Remove Stones to Minimize the Total||60.2%|Medium||
|1963|Minimum Number of Swaps to Make the String Balanced||69.2%|Medium||
|1964|Find the Longest Valid Obstacle Course at Each Position||47.2%|Hard||
|1965|Employees With Missing Information||75.6%|Easy||
|1966|Binary Searchable Numbers in an Unsorted Array||65.4%|Medium||
|1967|Number of Strings That Appear as Substrings in Word||80.1%|Easy||
|1968|Array With Elements Not Equal to Average of Neighbors||49.5%|Medium||
|1969|Minimum Non-Zero Product of the Array Elements||34.2%|Medium||
|1970|Last Day Where You Can Still Cross||49.6%|Hard||
|1971|Find if Path Exists in Graph||52.1%|Easy||
|1972|First and Last Call On the Same Day||52.2%|Hard||
|1973|Count Nodes Equal to Sum of Descendants||75.1%|Medium||
|1974|Minimum Time to Type Word Using Special Typewriter||72.3%|Easy||
|1975|Maximum Matrix Sum||48.3%|Medium||
|1976|Number of Ways to Arrive at Destination||31.0%|Medium||
|1977|Number of Ways to Separate Numbers||21.0%|Hard||
|1978|Employees Whose Manager Left the Company||50.0%|Easy||
|1979|Find Greatest Common Divisor of Array||76.9%|Easy||
|1980|Find Unique Binary String||65.1%|Medium||
|1981|Minimize the Difference Between Target and Chosen Elements||32.4%|Medium||
|1982|Find Array Given Subset Sums||48.5%|Hard||
|1983|Widest Pair of Indices With Equal Range Sum||54.2%|Medium||
|1984|Minimum Difference Between Highest and Lowest of K Scores|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/1984.Minimum-Difference-Between-Highest-and-Lowest-of-K-Scores)|54.4%|Easy||
|1985|Find the Kth Largest Integer in the Array||44.5%|Medium||
|1986|Minimum Number of Work Sessions to Finish the Tasks||32.7%|Medium||
|1987|Number of Unique Good Subsequences||52.2%|Hard||
|1988|Find Cutoff Score for Each School||68.7%|Medium||
|1989|Maximum Number of People That Can Be Caught in Tag||52.3%|Medium||
|1990|Count the Number of Experiments||51.2%|Medium||
|1991|Find the Middle Index in Array||67.4%|Easy||
|1992|Find All Groups of Farmland||68.7%|Medium||
|1993|Operations on Tree||44.1%|Medium||
|1994|The Number of Good Subsets||34.8%|Hard||
|1995|Count Special Quadruplets||59.7%|Easy||
|1996|The Number of Weak Characters in the Game||44.0%|Medium||
|1997|First Day Where You Have Been in All the Rooms||37.0%|Medium||
|1998|GCD Sort of an Array||45.6%|Hard||
|1999|Smallest Greater Multiple Made of Two Digits||49.5%|Medium||
|2000|Reverse Prefix of Word||78.6%|Easy||
|2001|Number of Pairs of Interchangeable Rectangles||46.0%|Medium||
|2002|Maximum Product of the Length of Two Palindromic Subsequences||54.4%|Medium||
|2003|Smallest Missing Genetic Value in Each Subtree||44.7%|Hard||
|2004|The Number of Seniors and Juniors to Join the Company||40.2%|Hard||
|2005|Subtree Removal Game with Fibonacci Tree||63.5%|Hard||
|2006|Count Number of Pairs With Absolute Difference K||82.7%|Easy||
|2007|Find Original Array From Doubled Array||40.7%|Medium||
|2008|Maximum Earnings From Taxi||43.6%|Medium||
|2009|Minimum Number of Operations to Make Array Continuous||45.6%|Hard||
|2010|The Number of Seniors and Juniors to Join the Company II||59.7%|Hard||
|2011|Final Value of Variable After Performing Operations||88.7%|Easy||
|2012|Sum of Beauty in the Array||46.9%|Medium||
|2013|Detect Squares||50.3%|Medium||
|2014|Longest Subsequence Repeated k Times||55.6%|Hard||
|2015|Average Height of Buildings in Each Segment||58.9%|Medium||
|2016|Maximum Difference Between Increasing Elements||54.2%|Easy||
|2017|Grid Game||43.6%|Medium||
|2018|Check if Word Can Be Placed In Crossword||49.4%|Medium||
|2019|The Score of Students Solving Math Expression||33.8%|Hard||
|2020|Number of Accounts That Did Not Stream||73.3%|Medium||
|2021|Brightest Position on Street|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2021.Brightest-Position-on-Street)|62.3%|Medium||
|2022|Convert 1D Array Into 2D Array|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2022.Convert-1D-Array-Into-2D-Array)|59.0%|Easy||
|2023|Number of Pairs of Strings With Concatenation Equal to Target||73.3%|Medium||
|2024|Maximize the Confusion of an Exam||60.3%|Medium||
|2025|Maximum Number of Ways to Partition an Array||32.8%|Hard||
|2026|Low-Quality Problems||84.9%|Easy||
|2027|Minimum Moves to Convert String||53.9%|Easy||
|2028|Find Missing Observations||44.5%|Medium||
|2029|Stone Game IX||26.6%|Medium||
|2030|Smallest K-Length Subsequence With Occurrences of a Letter||38.8%|Hard||
|2031|Count Subarrays With More Ones Than Zeros||51.9%|Medium||
|2032|Two Out of Three||73.5%|Easy||
|2033|Minimum Operations to Make a Uni-Value Grid||52.4%|Medium||
|2034|Stock Price Fluctuation||48.9%|Medium||
|2035|Partition Array Into Two Arrays to Minimize Sum Difference||18.8%|Hard||
|2036|Maximum Alternating Subarray Sum||41.0%|Medium||
|2037|Minimum Number of Moves to Seat Everyone|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2037.Minimum-Number-of-Moves-to-Seat-Everyone)|82.0%|Easy||
|2038|Remove Colored Pieces if Both Neighbors are the Same Color|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2038.Remove-Colored-Pieces-if-Both-Neighbors-are-the-Same-Color)|58.0%|Medium||
|2039|The Time When the Network Becomes Idle||51.0%|Medium||
|2040|Kth Smallest Product of Two Sorted Arrays||28.9%|Hard||
|2041|Accepted Candidates From the Interviews||79.1%|Medium||
|2042|Check if Numbers Are Ascending in a Sentence||67.2%|Easy||
|2043|Simple Bank System|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2043.Simple-Bank-System)|65.3%|Medium||
|2044|Count Number of Maximum Bitwise-OR Subsets||75.6%|Medium||
|2045|Second Minimum Time to Reach Destination||39.2%|Hard||
|2046|Sort Linked List Already Sorted Using Absolute Values||68.6%|Medium||
|2047|Number of Valid Words in a Sentence||29.4%|Easy||
|2048|Next Greater Numerically Balanced Number||47.3%|Medium||
|2049|Count Nodes With the Highest Score||47.6%|Medium||
|2050|Parallel Courses III||59.3%|Hard||
|2051|The Category of Each Member in the Store||71.5%|Medium||
|2052|Minimum Cost to Separate Sentence Into Rows||50.1%|Medium||
|2053|Kth Distinct String in an Array||72.1%|Easy||
|2054|Two Best Non-Overlapping Events||45.2%|Medium||
|2055|Plates Between Candles||44.6%|Medium||
|2056|Number of Valid Move Combinations On Chessboard||58.1%|Hard||
|2057|Smallest Index With Equal Value||71.5%|Easy||
|2058|Find the Minimum and Maximum Number of Nodes Between Critical Points||57.1%|Medium||
|2059|Minimum Operations to Convert Number||47.7%|Medium||
|2060|Check if an Original String Exists Given Two Encoded Strings||40.9%|Hard||
|2061|Number of Spaces Cleaning Robot Cleaned||54.6%|Medium||
|2062|Count Vowel Substrings of a String||66.0%|Easy||
|2063|Vowels of All Substrings||54.7%|Medium||
|2064|Minimized Maximum of Products Distributed to Any Store||50.7%|Medium||
|2065|Maximum Path Quality of a Graph||57.3%|Hard||
|2066|Account Balance||84.6%|Medium||
|2067|Number of Equal Count Substrings||48.3%|Medium||
|2068|Check Whether Two Strings are Almost Equivalent||64.4%|Easy||
|2069|Walking Robot Simulation II||23.1%|Medium||
|2070|Most Beautiful Item for Each Query||49.6%|Medium||
|2071|Maximum Number of Tasks You Can Assign||34.0%|Hard||
|2072|The Winner University||72.3%|Easy||
|2073|Time Needed to Buy Tickets||62.3%|Easy||
|2074|Reverse Nodes in Even Length Groups||53.4%|Medium||
|2075|Decode the Slanted Ciphertext||50.3%|Medium||
|2076|Process Restricted Friend Requests||53.3%|Hard||
|2077|Paths in Maze That Lead to Same Room||55.4%|Medium||
|2078|Two Furthest Houses With Different Colors||67.1%|Easy||
|2079|Watering Plants||80.0%|Medium||
|2080|Range Frequency Queries||38.6%|Medium||
|2081|Sum of k-Mirror Numbers||42.0%|Hard||
|2082|The Number of Rich Customers||79.9%|Easy||
|2083|Substrings That Begin and End With the Same Letter||67.8%|Medium||
|2084|Drop Type 1 Orders for Customers With Type 0 Orders||88.7%|Medium||
|2085|Count Common Words With One Occurrence||70.0%|Easy||
|2086|Minimum Number of Food Buckets to Feed the Hamsters||44.9%|Medium||
|2087|Minimum Cost Homecoming of a Robot in a Grid||51.1%|Medium||
|2088|Count Fertile Pyramids in a Land||63.7%|Hard||
|2089|Find Target Indices After Sorting Array||76.4%|Easy||
|2090|K Radius Subarray Averages||42.7%|Medium||
|2091|Removing Minimum and Maximum From Array||55.9%|Medium||
|2092|Find All People With Secret||34.2%|Hard||
|2093|Minimum Cost to Reach City With Discounts||56.1%|Medium||
|2094|Finding 3-Digit Even Numbers||57.7%|Easy||
|2095|Delete the Middle Node of a Linked List||59.6%|Medium||
|2096|Step-By-Step Directions From a Binary Tree Node to Another|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2096.Step-By-Step-Directions-From-a-Binary-Tree-Node-to-Another)|48.5%|Medium||
|2097|Valid Arrangement of Pairs||41.2%|Hard||
|2098|Subsequence of Size K With the Largest Even Sum||38.3%|Medium||
|2099|Find Subsequence of Length K With the Largest Sum||42.9%|Easy||
|2100|Find Good Days to Rob the Bank||49.2%|Medium||
|2101|Detonate the Maximum Bombs||42.0%|Medium||
|2102|Sequentially Ordinal Rank Tracker||65.9%|Hard||
|2103|Rings and Rods||81.3%|Easy||
|2104|Sum of Subarray Ranges||60.3%|Medium||
|2105|Watering Plants II||49.5%|Medium||
|2106|Maximum Fruits Harvested After at Most K Steps||35.0%|Hard||
|2107|Number of Unique Flavors After Sharing K Candies||56.2%|Medium||
|2108|Find First Palindromic String in the Array||78.7%|Easy||
|2109|Adding Spaces to a String||56.5%|Medium||
|2110|Number of Smooth Descent Periods of a Stock||57.7%|Medium||
|2111|Minimum Operations to Make the Array K-Increasing||37.6%|Hard||
|2112|The Airport With the Most Traffic||70.9%|Medium||
|2113|Elements in Array After Removing and Replacing Elements||73.2%|Medium||
|2114|Maximum Number of Words Found in Sentences||87.2%|Easy||
|2115|Find All Possible Recipes from Given Supplies||48.5%|Medium||
|2116|Check if a Parentheses String Can Be Valid||31.1%|Medium||
|2117|Abbreviating the Product of a Range||27.2%|Hard||
|2118|Build the Equation||57.2%|Hard||
|2119|A Number After a Double Reversal||76.9%|Easy||
|2120|Execution of All Suffix Instructions Staying in a Grid||83.3%|Medium||
|2121|Intervals Between Identical Elements||43.2%|Medium||
|2122|Recover the Original Array||38.5%|Hard||
|2123|Minimum Operations to Remove Adjacent Ones in Matrix||41.1%|Hard||
|2124|Check if All A's Appears Before All B's||71.4%|Easy||
|2125|Number of Laser Beams in a Bank||82.2%|Medium||
|2126|Destroying Asteroids||49.9%|Medium||
|2127|Maximum Employees to Be Invited to a Meeting||34.6%|Hard||
|2128|Remove All Ones With Row and Column Flips||76.3%|Medium||
|2129|Capitalize the Title||62.0%|Easy||
|2130|Maximum Twin Sum of a Linked List||80.7%|Medium||
|2131|Longest Palindrome by Concatenating Two Letter Words||48.5%|Medium||
|2132|Stamping the Grid||31.7%|Hard||
|2133|Check if Every Row and Column Contains All Numbers||52.5%|Easy||
|2134|Minimum Swaps to Group All 1's Together II||51.3%|Medium||
|2135|Count Words Obtained After Adding a Letter||42.8%|Medium||
|2136|Earliest Possible Day of Full Bloom||73.6%|Hard||
|2137|Pour Water Between Buckets to Make Water Levels Equal||67.1%|Medium||
|2138|Divide a String Into Groups of Size k||65.4%|Easy||
|2139|Minimum Moves to Reach Target Score||48.8%|Medium||
|2140|Solving Questions With Brainpower||46.4%|Medium||
|2141|Maximum Running Time of N Computers||39.2%|Hard||
|2142|The Number of Passengers in Each Bus I||50.0%|Medium||
|2143|Choose Numbers From Two Arrays in Range||52.3%|Hard||
|2144|Minimum Cost of Buying Candies With Discount||61.0%|Easy||
|2145|Count the Hidden Sequences||36.8%|Medium||
|2146|K Highest Ranked Items Within a Price Range||41.4%|Medium||
|2147|Number of Ways to Divide a Long Corridor||39.9%|Hard||
|2148|Count Elements With Strictly Smaller and Greater Elements||59.7%|Easy||
|2149|Rearrange Array Elements by Sign||80.8%|Medium||
|2150|Find All Lonely Numbers in the Array||60.6%|Medium||
|2151|Maximum Good People Based on Statements||49.2%|Hard||
|2152|Minimum Number of Lines to Cover Points||46.7%|Medium||
|2153|The Number of Passengers in Each Bus II||46.5%|Hard||
|2154|Keep Multiplying Found Values by Two||72.7%|Easy||
|2155|All Divisions With the Highest Score of a Binary Array||63.5%|Medium||
|2156|Find Substring With Given Hash Value||22.3%|Hard||
|2157|Groups of Strings||25.7%|Hard||
|2158|Amount of New Area Painted Each Day||54.6%|Hard||
|2159|Order Two Columns Independently||61.2%|Medium||
|2160|Minimum Sum of Four Digit Number After Splitting Digits||86.9%|Easy||
|2161|Partition Array According to Given Pivot||84.8%|Medium||
|2162|Minimum Cost to Set Cooking Time||39.9%|Medium||
|2163|Minimum Difference in Sums After Removal of Elements||47.6%|Hard||
|2164|Sort Even and Odd Indices Independently|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2164.Sort-Even-and-Odd-Indices-Independently)|65.1%|Easy||
|2165|Smallest Value of the Rearranged Number|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2165.Smallest-Value-of-the-Rearranged-Number)|51.4%|Medium||
|2166|Design Bitset|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2166.Design-Bitset)|31.7%|Medium||
|2167|Minimum Time to Remove All Cars Containing Illegal Goods|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2167.Minimum-Time-to-Remove-All-Cars-Containing-Illegal-Goods)|40.7%|Hard||
|2168|Unique Substrings With Equal Digit Frequency||59.3%|Medium||
|2169|Count Operations to Obtain Zero|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2169.Count-Operations-to-Obtain-Zero)|75.1%|Easy||
|2170|Minimum Operations to Make the Array Alternating|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2170.Minimum-Operations-to-Make-the-Array-Alternating)|33.2%|Medium||
|2171|Removing Minimum Number of Magic Beans|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2171.Removing-Minimum-Number-of-Magic-Beans)|42.0%|Medium||
|2172|Maximum AND Sum of Array||48.3%|Hard||
|2173|Longest Winning Streak||54.1%|Hard||
|2174|Remove All Ones With Row and Column Flips II||68.2%|Medium||
|2175|The Change in Global Rankings||65.1%|Medium||
|2176|Count Equal and Divisible Pairs in an Array||79.7%|Easy||
|2177|Find Three Consecutive Integers That Sum to a Given Number||63.9%|Medium||
|2178|Maximum Split of Positive Even Integers||59.3%|Medium||
|2179|Count Good Triplets in an Array||38.4%|Hard||
|2180|Count Integers With Even Digit Sum|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2180.Count-Integers-With-Even-Digit-Sum)|65.4%|Easy||
|2181|Merge Nodes in Between Zeros|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2181.Merge-Nodes-in-Between-Zeros)|86.4%|Medium||
|2182|Construct String With Repeat Limit|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2182.Construct-String-With-Repeat-Limit)|52.4%|Medium||
|2183|Count Array Pairs Divisible by K|[Go](https://github.com/halfrost/leetcode-go/tree/master/leetcode/2183.Count-Array-Pairs-Divisible-by-K)|28.4%|Hard||
|2184|Number of Ways to Build Sturdy Brick Wall||50.2%|Medium||
|2185|Counting Words With a Given Prefix||77.3%|Easy||
|2186|Minimum Number of Steps to Make Two Strings Anagram II||72.0%|Medium||
|2187|Minimum Time to Complete Trips||39.2%|Medium||
|2188|Minimum Time to Finish the Race||41.7%|Hard||
|2189|Number of Ways to Build House of Cards||62.7%|Medium||
|2190|Most Frequent Number Following Key In an Array||60.0%|Easy||
|2191|Sort the Jumbled Numbers||45.8%|Medium||
|2192|All Ancestors of a Node in a Directed Acyclic Graph||50.7%|Medium||
|2193|Minimum Number of Moves to Make Palindrome||51.3%|Hard||
|2194|Cells in a Range on an Excel Sheet||85.1%|Easy||
|2195|Append K Integers With Minimal Sum||25.0%|Medium||
|2196|Create Binary Tree From Descriptions||72.3%|Medium||
|2197|Replace Non-Coprime Numbers in Array||38.6%|Hard||
|2198|Number of Single Divisor Triplets||55.9%|Medium||
|2199|Finding the Topic of Each Post||49.1%|Hard||
|2200|Find All K-Distant Indices in an Array||65.0%|Easy||
|2201|Count Artifacts That Can Be Extracted||55.3%|Medium||
|2202|Maximize the Topmost Element After K Moves||22.8%|Medium||
|2203|Minimum Weighted Subgraph With the Required Paths||35.9%|Hard||
|2204|Distance to a Cycle in Undirected Graph||70.3%|Hard||
|2205|The Number of Users That Are Eligible for Discount||50.2%|Easy||
|2206|Divide Array Into Equal Pairs||73.9%|Easy||
|2207|Maximize Number of Subsequences in a String||33.3%|Medium||
|2208|Minimum Operations to Halve Array Sum||45.4%|Medium||
|2209|Minimum White Tiles After Covering With Carpets||34.2%|Hard||
|2210|Count Hills and Valleys in an Array||58.3%|Easy||
|2211|Count Collisions on a Road||42.3%|Medium||
|2212|Maximum Points in an Archery Competition||49.5%|Medium||
|2213|Longest Substring of One Repeating Character||30.9%|Hard||
|2214|Minimum Health to Beat Game||58.0%|Medium||
|2215|Find the Difference of Two Arrays||70.5%|Easy||
|2216|Minimum Deletions to Make Array Beautiful||46.9%|Medium||
|2217|Find Palindrome With Fixed Length||34.0%|Medium||
|2218|Maximum Value of K Coins From Piles||48.2%|Hard||
|2219|Maximum Sum Score of Array||60.8%|Medium||
|2220|Minimum Bit Flips to Convert Number||82.3%|Easy||
|2221|Find Triangular Sum of an Array||78.5%|Medium||
|2222|Number of Ways to Select Buildings||51.3%|Medium||
|2223|Sum of Scores of Built Strings||37.4%|Hard||
|2224|Minimum Number of Operations to Convert Time||65.4%|Easy||
|2225|Find Players With Zero or One Losses||72.7%|Medium||
|2226|Maximum Candies Allocated to K Children||36.6%|Medium||
|2227|Encrypt and Decrypt Strings||39.4%|Hard||
|2228|Users With Two Purchases Within Seven Days||44.0%|Medium||
|2229|Check if an Array Is Consecutive||61.1%|Easy||
|2230|The Users That Are Eligible for Discount||49.5%|Easy||
|2231|Largest Number After Digit Swaps by Parity||61.0%|Easy||
|2232|Minimize Result by Adding Parentheses to Expression||65.8%|Medium||
|2233|Maximum Product After K Increments||41.1%|Medium||
|2234|Maximum Total Beauty of the Gardens||28.5%|Hard||
|2235|Add Two Integers||88.6%|Easy||
|2236|Root Equals Sum of Children||84.8%|Easy||
|2237|Count Positions on Street With Required Brightness||64.7%|Medium||
|2238|Number of Times a Driver Was a Passenger||73.7%|Medium||
|2239|Find Closest Number to Zero||45.5%|Easy||
|2240|Number of Ways to Buy Pens and Pencils||56.9%|Medium||
|2241|Design an ATM Machine||38.9%|Medium||
|2242|Maximum Score of a Node Sequence||37.9%|Hard||
|2243|Calculate Digit Sum of a String||66.3%|Easy||
|2244|Minimum Rounds to Complete All Tasks||62.9%|Medium||
|2245|Maximum Trailing Zeros in a Cornered Path||35.0%|Medium||
|2246|Longest Path With Different Adjacent Characters||55.8%|Hard||
|2247|Maximum Cost of Trip With K Highways||50.7%|Hard||
|2248|Intersection of Multiple Arrays||68.8%|Easy||
|2249|Count Lattice Points Inside a Circle||50.5%|Medium||
|2250|Count Number of Rectangles Containing Each Point||34.1%|Medium||
|2251|Number of Flowers in Full Bloom||51.3%|Hard||
|2252|Dynamic Pivoting of a Table||58.9%|Hard||
|2253|Dynamic Unpivoting of a Table||68.4%|Hard||
|2254|Design Video Sharing Platform||65.1%|Hard||
|2255|Count Prefixes of a Given String||72.9%|Easy||
|2256|Minimum Average Difference||43.1%|Medium||
|2257|Count Unguarded Cells in the Grid||52.2%|Medium||
|2258|Escape the Spreading Fire||35.0%|Hard||
|2259|Remove Digit From Number to Maximize Result||46.3%|Easy||
|2260|Minimum Consecutive Cards to Pick Up||51.2%|Medium||
|2261|K Divisible Elements Subarrays||48.0%|Medium||
|2262|Total Appeal of A String||56.8%|Hard||
|2263|Make Array Non-decreasing or Non-increasing||67.3%|Hard||
|2264|Largest 3-Same-Digit Number in String||59.5%|Easy||
|2265|Count Nodes Equal to Average of Subtree||85.2%|Medium||
|2266|Count Number of Texts||47.1%|Medium||
|2267|Check if There Is a Valid Parentheses String Path||38.0%|Hard||
|2268|Minimum Number of Keypresses||73.2%|Medium||
|2269|Find the K-Beauty of a Number||57.7%|Easy||
|2270|Number of Ways to Split Array||45.6%|Medium||
|2271|Maximum White Tiles Covered by a Carpet||33.0%|Medium||
|2272|Substring With Largest Variance||37.0%|Hard||
|2273|Find Resultant Array After Removing Anagrams||58.4%|Easy||
|2274|Maximum Consecutive Floors Without Special Floors||52.1%|Medium||
|2275|Largest Combination With Bitwise AND Greater Than Zero||72.4%|Medium||
|2276|Count Integers in Intervals||34.4%|Hard||
|2277|Closest Node to Path in Tree||63.8%|Hard||
|2278|Percentage of Letter in String||73.9%|Easy||
|2279|Maximum Bags With Full Capacity of Rocks||67.7%|Medium||
|2280|Minimum Lines to Represent a Line Chart||24.3%|Medium||
|2281|Sum of Total Strength of Wizards||27.4%|Hard||
|2282|Number of People That Can Be Seen in a Grid||49.6%|Medium||
|2283|Check if Number Has Equal Digit Count and Digit Value||73.3%|Easy||
|2284|Sender With Largest Word Count||56.1%|Medium||
|2285|Maximum Total Importance of Roads||60.7%|Medium||
|2286|Booking Concert Tickets in Groups||15.3%|Hard||
|2287|Rearrange Characters to Make Target String||58.0%|Easy||
|2288|Apply Discount to Prices||27.9%|Medium||
|2289|Steps to Make Array Non-decreasing||21.1%|Medium||
|2290|Minimum Obstacle Removal to Reach Corner||49.6%|Hard||
|2291|Maximum Profit From Trading Stocks||45.5%|Medium||
|2292|Products With Three or More Orders in Two Consecutive Years||39.3%|Medium||
|2293|Min Max Game||63.8%|Easy||
|2294|Partition Array Such That Maximum Difference Is K||72.7%|Medium||
|2295|Replace Elements in an Array||57.5%|Medium||
|2296|Design a Text Editor||40.7%|Hard||
|2297|Jump Game VIII||56.1%|Medium||
|2298|Tasks Count in the Weekend||84.7%|Medium||
|2299|Strong Password Checker II||56.3%|Easy||
|2300|Successful Pairs of Spells and Potions||32.4%|Medium||
|2301|Match Substring After Replacement||39.4%|Hard||
|2302|Count Subarrays With Score Less Than K||52.2%|Hard||
|2303|Calculate Amount Paid in Taxes||64.1%|Easy||
|2304|Minimum Path Cost in a Grid||65.7%|Medium||
|2305|Fair Distribution of Cookies||61.9%|Medium||
|2306|Naming a Company||47.1%|Hard||
|2307|Check for Contradictions in Equations||43.7%|Hard||
|2308|Arrange Table by Gender||72.7%|Medium||
|2309|Greatest English Letter in Upper and Lower Case||68.8%|Easy||
|2310|Sum of Numbers With Units Digit K||25.8%|Medium||
|2311|Longest Binary Subsequence Less Than or Equal to K||37.1%|Medium||
|2312|Selling Pieces of Wood||48.4%|Hard||
|2313|Minimum Flips in Binary Tree to Get Result||63.6%|Hard||
|2314|The First Day of the Maximum Recorded Degree in Each City||72.9%|Medium||
|2315|Count Asterisks||82.1%|Easy||
|2316|Count Unreachable Pairs of Nodes in an Undirected Graph||38.9%|Medium||
|2317|Maximum XOR After Operations||79.1%|Medium||
|2318|Number of Distinct Roll Sequences||56.3%|Hard||
|2319|Check if Matrix Is X-Matrix||66.7%|Easy||
|2320|Count Number of Ways to Place Houses||40.6%|Medium||
|2321|Maximum Score Of Spliced Array||55.8%|Hard||
|2322|Minimum Score After Removals on a Tree||50.8%|Hard||
|2323|Find Minimum Time to Finish All Jobs II||73.8%|Medium||
|2324|Product Sales Analysis IV||77.9%|Medium||
|2325|Decode the Message||84.7%|Easy||
|2326|Spiral Matrix IV||74.7%|Medium||
|2327|Number of People Aware of a Secret||44.8%|Medium||
|2328|Number of Increasing Paths in a Grid||47.6%|Hard||
|2329|Product Sales Analysis V||68.5%|Easy||
|2330|Valid Palindrome IV||75.7%|Medium||
|2331|Evaluate Boolean Binary Tree||78.5%|Easy||
|2332|The Latest Time to Catch a Bus||23.4%|Medium||
|2333|Minimum Sum of Squared Difference||25.5%|Medium||
|2334|Subarray With Elements Greater Than Varying Threshold||41.2%|Hard||
|2335|Minimum Amount of Time to Fill Cups||56.3%|Easy||
|2336|Smallest Number in Infinite Set||71.3%|Medium||
|2337|Move Pieces to Obtain a String||48.0%|Medium||
|2338|Count the Number of Ideal Arrays||26.3%|Hard||
|2339|All the Matches of the League||88.1%|Easy||
|2340|Minimum Adjacent Swaps to Make a Valid Array||75.2%|Medium||
|2341|Maximum Number of Pairs in Array||76.0%|Easy||
|2342|Max Sum of a Pair With Equal Sum of Digits||53.4%|Medium||
|2343|Query Kth Smallest Trimmed Number||41.3%|Medium||
|2344|Minimum Deletions to Make Array Divisible||56.8%|Hard||
|2345|Finding the Number of Visible Mountains||42.2%|Medium||
|2346|Compute the Rank as a Percentage||32.7%|Medium||
|2347|Best Poker Hand||60.7%|Easy||
|2348|Number of Zero-Filled Subarrays||57.8%|Medium||
|2349|Design a Number Container System||46.0%|Medium||
|2350|Shortest Impossible Sequence of Rolls||68.3%|Hard||
|2351|First Letter to Appear Twice||75.0%|Easy||
|2352|Equal Row and Column Pairs||70.7%|Medium||
|2353|Design a Food Rating System||34.7%|Medium||
|2354|Number of Excellent Pairs||46.0%|Hard||
|2355|Maximum Number of Books You Can Take||44.1%|Hard||
|2356|Number of Unique Subjects Taught by Each Teacher||90.3%|Easy||
|2357|Make Array Zero by Subtracting Equal Amounts||72.7%|Easy||
|2358|Maximum Number of Groups Entering a Competition||67.6%|Medium||
|2359|Find Closest Node to Given Two Nodes||46.2%|Medium||
|2360|Longest Cycle in a Graph||38.9%|Hard||
|2361|Minimum Costs Using the Train Line||76.1%|Hard||
|2362|Generate the Invoice||80.8%|Hard||
|2363|Merge Similar Items||75.5%|Easy||
|2364|Count Number of Bad Pairs||41.1%|Medium||
|2365|Task Scheduler II||46.8%|Medium||
|2366|Minimum Replacements to Sort the Array||40.9%|Hard||
|2367|Number of Arithmetic Triplets||83.6%|Easy||
|2368|Reachable Nodes With Restrictions||57.8%|Medium||
|2369|Check if There is a Valid Partition For The Array||40.1%|Medium||
|2370|Longest Ideal Subsequence||37.8%|Medium||
|2371|Minimize Maximum Value in a Grid||71.3%|Hard||
|2372|Calculate the Influence of Each Salesperson||85.6%|Medium||
|2373|Largest Local Values in a Matrix||83.3%|Easy||
|2374|Node With Highest Edge Score||46.5%|Medium||
|2375|Construct Smallest Number From DI String||74.5%|Medium||
|2376|Count Special Integers||36.6%|Hard||
|2377|Sort the Olympic Table||78.4%|Easy||
|2378|Choose Edges to Maximize Score in a Tree||60.3%|Medium||
|2379|Minimum Recolors to Get K Consecutive Black Blocks||57.7%|Easy||
|2380|Time Needed to Rearrange a Binary String||49.0%|Medium||
|2381|Shifting Letters II||34.9%|Medium||
|2382|Maximum Segment Sum After Removals||47.8%|Hard||
|2383|Minimum Hours of Training to Win a Competition||41.0%|Easy||
|2384|Largest Palindromic Number||31.1%|Medium||
|2385|Amount of Time for Binary Tree to Be Infected||56.6%|Medium||
|2386|Find the K-Sum of an Array||38.5%|Hard||
|2387|Median of a Row Wise Sorted Matrix||68.9%|Medium||
|2388|Change Null Values in a Table to the Previous Value||65.9%|Medium||
|2389|Longest Subsequence With Limited Sum||72.5%|Easy||
|2390|Removing Stars From a String||65.2%|Medium||
|2391|Minimum Amount of Time to Collect Garbage||84.4%|Medium||
|2392|Build a Matrix With Conditions||59.4%|Hard||
|2393|Count Strictly Increasing Subarrays||72.9%|Medium||
|2394|Employees With Deductions||41.7%|Medium||
|2395|Find Subarrays With Equal Sum||64.4%|Easy||
|2396|Strictly Palindromic Number||87.8%|Medium||
|2397|Maximum Rows Covered by Columns||52.8%|Medium||
|2398|Maximum Number of Robots Within Budget||32.9%|Hard||
|2399|Check Distances Between Same Letters||70.6%|Easy||
|2400|Number of Ways to Reach a Position After Exactly k Steps||33.0%|Medium||
|2401|Longest Nice Subarray||49.2%|Medium||
|2402|Meeting Rooms III||33.5%|Hard||
|2403|Minimum Time to Kill All Monsters||51.2%|Hard||
|2404|Most Frequent Even Element||50.9%|Easy||
|2405|Optimal Partition of String||75.2%|Medium||
|2406|Divide Intervals Into Minimum Number of Groups||46.0%|Medium||
|2407|Longest Increasing Subsequence II||21.7%|Hard||
|2408|Design SQL||84.0%|Medium||
|2409|Count Days Spent Together||43.5%|Easy||
|2410|Maximum Matching of Players With Trainers||60.9%|Medium||
|2411|Smallest Subarrays With Maximum Bitwise OR||41.6%|Medium||
|2412|Minimum Money Required Before Transactions||39.6%|Hard||
|2413|Smallest Even Multiple||87.7%|Easy||
|2414|Length of the Longest Alphabetical Continuous Substring||56.1%|Medium||
|2415|Reverse Odd Levels of Binary Tree||76.7%|Medium||
|2416|Sum of Prefix Scores of Strings||43.7%|Hard||
|2417|Closest Fair Integer||45.3%|Medium||
|2418|Sort the People||81.1%|Easy||
|2419|Longest Subarray With Maximum Bitwise AND||48.0%|Medium||
|2420|Find All Good Indices||37.5%|Medium||
|2421|Number of Good Paths||57.6%|Hard||
|2422|Merge Operations to Turn Array Into a Palindrome||69.7%|Medium||
|2423|Remove Letter To Equalize Frequency||18.0%|Easy||
|2424|Longest Uploaded Prefix||54.0%|Medium||
|2425|Bitwise XOR of All Pairings||58.6%|Medium||
|2426|Number of Pairs Satisfying Inequality||43.0%|Hard||
|2427|Number of Common Factors||79.8%|Easy||
|2428|Maximum Sum of an Hourglass||74.2%|Medium||
|2429|Minimize XOR||42.5%|Medium||
|2430|Maximum Deletions on a String||32.2%|Hard||
|2431|Maximize Total Tastiness of Purchased Fruits||67.6%|Medium||
|2432|The Employee That Worked on the Longest Task||49.4%|Easy||
|2433|Find The Original Array of Prefix Xor||85.6%|Medium||
|2434|Using a Robot to Print the Lexicographically Smallest String||38.5%|Medium||
|2435|Paths in Matrix Whose Sum Is Divisible by K||41.5%|Hard||
|2436|Minimum Split Into Subarrays With GCD Greater Than One||74.2%|Medium||
|2437|Number of Valid Clock Times||43.1%|Easy||
|2438|Range Product Queries of Powers||38.9%|Medium||
|2439|Minimize Maximum of Array||34.4%|Medium||
|2440|Create Components With Same Value||54.6%|Hard||
|2441|Largest Positive Integer That Exists With Its Negative||68.0%|Easy||
|2442|Count Number of Distinct Integers After Reverse Operations||78.7%|Medium||
|2443|Sum of Number and Its Reverse||45.7%|Medium||
|2444|Count Subarrays With Fixed Bounds||62.0%|Hard||
|2445|Number of Nodes With Value One||70.7%|Medium||
|2446|Determine if Two Events Have Conflict||49.9%|Easy||
|2447|Number of Subarrays With GCD Equal to K||48.6%|Medium||
|2448|Minimum Cost to Make Array Equal||34.8%|Hard||
|2449|Minimum Number of Operations to Make Arrays Similar||64.0%|Hard||
|2450|Number of Distinct Binary Strings After Applying Operations||67.6%|Medium||
|2451|Odd String Difference||60.1%|Easy||
|2452|Words Within Two Edits of Dictionary||60.4%|Medium||
|2453|Destroy Sequential Targets||37.7%|Medium||
|2454|Next Greater Element IV||39.8%|Hard||
|2455|Average Value of Even Numbers That Are Divisible by Three||59.1%|Easy||
|2456|Most Popular Video Creator||43.8%|Medium||
|2457|Minimum Addition to Make Integer Beautiful||37.0%|Medium||
|2458|Height of Binary Tree After Subtree Removal Queries||37.0%|Hard||
|2459|Sort Array by Moving Items to Empty Space||51.9%|Hard||
|2460|Apply Operations to an Array||67.2%|Easy||
|2461|Maximum Sum of Distinct Subarrays With Length K||34.3%|Medium||
|2462|Total Cost to Hire K Workers||37.6%|Medium||
|2463|Minimum Total Distance Traveled||40.6%|Hard||
|2464|Minimum Subarrays in a Valid Split||56.0%|Medium||
|2465|Number of Distinct Averages||58.3%|Easy||
|2466|Count Ways To Build Good Strings||42.1%|Medium||
|2467|Most Profitable Path in a Tree||48.3%|Medium||
|2468|Split Message Based on Limit||46.3%|Hard||
|2469|Convert the Temperature||89.1%|Easy||
|2470|Number of Subarrays With LCM Equal to K||38.4%|Medium||
|2471|Minimum Number of Operations to Sort a Binary Tree by Level||62.3%|Medium||
|2472|Maximum Number of Non-overlapping Palindrome Substrings||38.0%|Hard||
|2473|Minimum Cost to Buy Apples||63.0%|Medium||
|2474|Customers With Strictly Increasing Purchases||45.3%|Hard||
|2475|Number of Unequal Triplets in Array||70.8%|Easy||
|2476|Closest Nodes Queries in a Binary Search Tree||40.7%|Medium||
|2477|Minimum Fuel Cost to Report to the Capital||67.8%|Medium||
|2478|Number of Beautiful Partitions||30.7%|Hard||
|2479|Maximum XOR of Two Non-Overlapping Subtrees||46.9%|Hard||
|2480|Form a Chemical Bond||56.5%|Easy||
|2481|Minimum Cuts to Divide a Circle||52.5%|Easy||
|2482|Difference Between Ones and Zeros in Row and Column||79.8%|Medium||
|2483|Minimum Penalty for a Shop||55.7%|Medium||
|2484|Count Palindromic Subsequences||33.0%|Hard||
|2485|Find the Pivot Integer||80.0%|Easy||
|2486|Append Characters to String to Make Subsequence||64.0%|Medium||
|2487|Remove Nodes From Linked List||68.8%|Medium||
|2488|Count Subarrays With Median K||42.3%|Hard||
|2489|Number of Substrings With Fixed Ratio||58.7%|Medium||
|2490|Circular Sentence||64.6%|Easy||
|2491|Divide Players Into Teams of Equal Skill||58.8%|Medium||
|2492|Minimum Score of a Path Between Two Cities||46.5%|Medium||
|2493|Divide Nodes Into the Maximum Number of Groups||37.0%|Hard||
|2494|Merge Overlapping Events in the Same Hall||18.7%|Hard||
|2495|Number of Subarrays Having Even Product||65.5%|Medium||
|2496|Maximum Value of a String in an Array||72.0%|Easy||
|2497|Maximum Star Sum of a Graph||38.4%|Medium||
|2498|Frog Jump II||61.0%|Medium||
|2499|Minimum Total Cost to Make Arrays Unequal||44.1%|Hard||
|2500|Delete Greatest Value in Each Row||81.0%|Easy||
|2501|Longest Square Streak in an Array||39.5%|Medium||
|2502|Design Memory Allocator||52.2%|Medium||
|2503|Maximum Number of Points From Grid Queries||37.0%|Hard||
|2504|Concatenate the Name and the Profession||55.1%|Easy||
|2505|Bitwise OR of All Subsequence Sums||59.7%|Medium||
|2506|Count Pairs Of Similar Strings||70.5%|Easy||
|2507|Smallest Value After Replacing With Sum of Prime Factors||48.7%|Medium||
|2508|Add Edges to Make Degrees of All Nodes Even||32.3%|Hard||
|2509|Cycle Length Queries in a Tree||57.0%|Hard||
|2510|Check if There is a Path With Equal Number of 0's And 1's||55.5%|Medium||
|2511|Maximum Enemy Forts That Can Be Captured||37.4%|Easy||
|2512|Reward Top K Students||46.5%|Medium||
|2513|Minimize the Maximum of Two Arrays||26.6%|Medium||
|2514|Count Anagrams||34.2%|Hard||
|2515|Shortest Distance to Target String in a Circular Array||49.0%|Easy||
|2516|Take K of Each Character From Left and Right||33.9%|Medium||
|2517|Maximum Tastiness of Candy Basket||64.6%|Medium||
|2518|Number of Great Partitions||32.9%|Hard||
|2519|Count the Number of K-Big Indices||64.3%|Hard||
|2520|Count the Digits That Divide a Number||85.0%|Easy||
|2521|Distinct Prime Factors of Product of Array||50.1%|Medium||
|2522|Partition String Into Substrings With Values at Most K||46.7%|Medium||
|2523|Closest Prime Numbers in Range||37.0%|Medium||
|2524|Maximum Frequency Score of a Subarray||40.2%|Hard||
|2525|Categorize Box According to Criteria||33.9%|Easy||
|2526|Find Consecutive Integers from a Data Stream||45.0%|Medium||
|2527|Find Xor-Beauty of Array||70.1%|Medium||
|2528|Maximize the Minimum Powered City||32.1%|Hard||
|2529|Maximum Count of Positive Integer and Negative Integer||74.9%|Easy||
|2530|Maximal Score After Applying K Operations||43.3%|Medium||
|2531|Make Number of Distinct Characters Equal||25.5%|Medium||
|2532|Time to Cross a Bridge||53.6%|Hard||
|2533|Number of Good Binary Strings||64.7%|Medium||
|2534|Time Taken to Cross the Door||70.5%|Hard||
|2535|Difference Between Element Sum and Digit Sum of an Array||85.5%|Easy||
|2536|Increment Submatrices by One||48.3%|Medium||
|2537|Count the Number of Good Subarrays||47.9%|Medium||
|2538|Difference Between Maximum and Minimum Price Sum||34.8%|Hard||
|2539|Count the Number of Good Subsequences||77.4%|Medium||
|2540|Minimum Common Value||51.7%|Easy||
|2541|Minimum Operations to Make Array Equal II||30.9%|Medium||
|2542|Maximum Subsequence Score||38.0%|Medium||
|2543|Check if Point Is Reachable||42.2%|Hard||
|2544|Alternating Digit Sum||69.7%|Easy||
|2545|Sort the Students by Their Kth Score||85.8%|Medium||
|2546|Apply Bitwise Operations to Make Strings Equal||40.6%|Medium||
|2547|Minimum Cost to Split an Array||38.8%|Hard||
|2548|Maximum Price to Fill a Bag||69.0%|Medium||
|2549|Count Distinct Numbers on Board||59.8%|Easy||
|2550|Count Collisions of Monkeys on a Polygon||26.5%|Medium||
|2551|Put Marbles in Bags||52.8%|Hard||
|2552|Count Increasing Quadruplets||31.4%|Hard||
|2553|Separate the Digits in an Array||79.6%|Easy||
|2554|Maximum Number of Integers to Choose From a Range I||52.4%|Medium||
|2555|Maximize Win From Two Segments||30.1%|Medium||
|2556|Disconnect Path in a Binary Matrix by at Most One Flip||28.5%|Medium||
|2557|Maximum Number of Integers to Choose From a Range II||45.7%|Medium||
|2558|Take Gifts From the Richest Pile||66.4%|Easy||
|2559|Count Vowel Strings in Ranges||51.3%|Medium||
|2560|House Robber IV||39.6%|Medium||
|2561|Rearranging Fruits||34.2%|Hard||
|2562|Find the Array Concatenation Value||69.7%|Easy||
|2563|Count the Number of Fair Pairs||31.8%|Medium||
|2564|Substring XOR Queries||33.5%|Medium||
|2565|Subsequence With the Minimum Score||32.1%|Hard||
|2566|Maximum Difference by Remapping a Digit||60.1%|Easy||
|2567|Minimum Score by Changing Two Elements||47.5%|Medium||
|2568|Minimum Impossible OR||57.2%|Medium||
|2569|Handling Sum Queries After Update||25.6%|Hard||
|2570|Merge Two 2D Arrays by Summing Values||73.0%|Easy||
|2571|Minimum Operations to Reduce an Integer to 0||49.2%|Medium||
|2572|Count the Number of Square-Free Subsets||20.3%|Medium||
|2573|Find the String with LCP||35.2%|Hard||
|2574|Left and Right Sum Differences||88.7%|Easy||
|2575|Find the Divisibility Array of a String||30.9%|Medium||
|2576|Find the Maximum Number of Marked Indices||37.2%|Medium||
|2577|Minimum Time to Visit a Cell In a Grid||35.5%|Hard||
|2578|Split With Minimum Sum||69.3%|Easy||
|2579|Count Total Number of Colored Cells||57.0%|Medium||
|2580|Count Ways to Group Overlapping Ranges||34.6%|Medium||
|2581|Count Number of Possible Root Nodes||47.4%|Hard||
|2582|Pass the Pillow||46.2%|Easy||
|2583|Kth Largest Sum in a Binary Tree||46.6%|Medium||
|2584|Split the Array to Make Coprime Products||21.9%|Hard||
|2585|Number of Ways to Earn Points||60.0%|Hard||
|2586|Count the Number of Vowel Strings in Range||79.1%|Easy||
|2587|Rearrange Array to Maximize Prefix Score||39.2%|Medium||
|2588|Count the Number of Beautiful Subarrays||48.7%|Medium||
|2589|Minimum Time to Complete All Tasks||34.7%|Hard||
|2890|Design a Todo List||40.8%|Medium||
|------------|-------------------------------------------------------|-------| ----------------| ---------------|-------------|

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
