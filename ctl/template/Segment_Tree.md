---
title: 2.18 ✅ Segment Tree
type: docs
weight: 18
---

# Segment Tree

![](https://img.halfrost.com/Leetcode/Segment_Tree.png)
- The classic array implementation of the line segment tree. The pushUp logic of merging two nodes is abstracted, and any operation can be realized (common operations include: addition, max, min, etc.). Question 218, Question 303, Question 307, Question 699.
- The classic way of counting line segment trees. Question 315, Question 327, Question 493.
- The implementation of the tree of the line segment tree. Question 715, question 732.
- Range lazy update. Question 218, question 699.
- Discretization. Discretization needs to pay attention to a special case: if the three intervals are [1,10] [1,4] [6,10], after discretization x[1]=1,x[2]=4,x[3] =6,x[4]=10. The first interval is [1,4], the second interval is [1,2], and the third interval is [3,4]. In this way, interval one = interval two + interval three, which is the same as before discrete The model does not match. Before the discretization, it is obvious that interval 1 > interval 2 + interval 3. The correct way is: add a number between numbers with a difference greater than 1, for example, add 5 in the middle of 1 4 6 10 above, then x[1]=1, x[2]=4, x[3]=5, x[4]=6,x[5]=10. After this processing, interval one is 1-5, interval two is 1-2, and interval three is 4-5.
- Flexible construction of segment trees. Line segment tree nodes can store multiple pieces of information, and the pushUp operation for merging two nodes can also be varied. Question 850, question 1157.


Line segment tree [question type] (https://blog.csdn.net/xuechelingxiao/article/details/38313105) from simple to difficult:

1. Single point update:
[HDU 1166 Enemy Soldier Arrangement](http://acm.hdu.edu.cn/showproblem.php?pid=1166) update: single-point increase and decrease query: interval summation
[HDU 1754 I Hate It](http://acm.hdu.edu.cn/showproblem.php?pid=1754) update: single point replacement query: maximum value of interval
[HDU 1394 Minimum Inversion Number](http://acm.hdu.edu.cn/showproblem.php?pid=1394) update: single-point increase and decrease query: interval summation
[HDU 2795 Billboard](http://acm.hdu.edu.cn/showproblem.php?pid=2795) query: find the position of the maximum value in the interval (directly do the update operation in the query)
2. Interval update:
[HDU 1698 Just a Hook](http://acm.hdu.edu.cn/showproblem.php?pid=1698) update: segment replacement (since the total interval is only queried once, the information of 1 node can be output directly )
[POJ 3468 A Simple Problem with Integers](http://poj.org/problem?id=3468) update: Segment increase and decrease query: Interval summation
[POJ 2528 Mayor’s posters](http://poj.org/problem?id=2528) discretization + update: segment replacement query: simple hash
[POJ 3225 Help with Intervals](http://poj.org/problem?id=3225) update: segment replacement, interval XOR query: simple hash
3. Interval merging (this type of question will ask the longest continuous interval that meets the conditions in the interval, so when PushUp needs to merge the intervals of the left and right sons):
[POJ 3667 Hotel](http://poj.org/problem?id=3667) update: interval replacement query: query the leftmost endpoint that meets the condition
4. Scanning line (this kind of question needs to sort some operations, and then use a scanning line to scan from left to right. The most typical one is rectangular area union, perimeter length union and other questions):
[HDU 1542 Atlantis](http://acm.hdu.edu.cn/showproblem.php?pid=1542) update: Interval increase and decrease query: Take the value of the root node directly
[HDU 1828 Picture](http://acm.hdu.edu.cn/showproblem.php?pid=1828) update: interval increase and decrease query: take the value of the root node directly
{{.AvailableTagTable}}