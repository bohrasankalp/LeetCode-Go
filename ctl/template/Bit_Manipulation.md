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



{{.AvailableTagTable}}