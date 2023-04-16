# time complexity and space complexity

## 1. Time complexity data scale

The data size that can solve the problem within 1s: 10^6 ~ 10^7

-   O(n^2) algorithm can handle data scale of 10^4 level (conservative estimate, it is definitely no problem to deal with 1000 level problems)
-   O(n) algorithm can handle data scale of 10^8 level (conservative estimate, it is definitely no problem to deal with 10^7 level problems)
-   O(nlog n) algorithm can handle data scale of 10^7 level (conservative estimate, it is definitely no problem to deal with 10^6 level problems)

|     | Data size | Time complexity | Algorithm example                                                                        |
|:---:|:----------|:----------------|:-----------------------------------------------------------------------------------------|
|  1  | 10        | O(n!)           | permutations                                                                             |
|  2  | 20~30     | O(2^n)          | combination                                                                              |
|  3  | 50        | O(n^4)          | DFS search,<br/> DP dynamic programming                                                  |
|  4  | 100       | O(n^3)          | The shortest path between any two points,<br/>DP dynamic programming                     |
|  5  | 1000      | O(n^2)          | Dense graph, DP dynamic programming                                                      |
|  6  | 10^6      | O(nlog n)       | sorting,<br/>heap,<br/>recursion and divide and conquer                                  |
|  7  | 10^7      | O(n)            | DP dynamic programming,<br/>graph traversal,<br/>topological sorting,<br/>tree traversal |
|  8  | 10^9      | O(sqrt(n))      | Sieve prime numbers,<br/>find square root                                                |
|  9  | 10^10     | O(log n)        | Binary Search                                                                            |
| 10  | +∞        | O(1)            | Mathematics-related algorithms                                                           |
|     |           |                 |                                                                                          |

Some deceptive examples:

```c
void hello (int n){

    for( int sz = 1 ; sz < n ; sz += sz)
        for( int i = 1 ; i < n ; i ++)
            cout << "Hello" << endl;
}
```

The time complexity of the above code is O(nlog n) instead of O(n^2)

```c
bool isPrime(int n){
for( int x = 2 ; x *x <= n ; x ++ )
        if( n % x == 0)
            return false;
    return true;
}
```

The time complexity of the above code is O(sqrt(n)) instead of O(n).

To give another example, there is an array of strings, sort each string in the array alphabetically, and then sort the entire string array lexicographically. What is the overall time complexity of the two-step operation?

If the answer is O(n_nlog n + nlog n) = O(n^2log n), the answer is wrong. The length of the string and the length of the array are irrelevant, so these two variables should be calculated separately. Suppose the longest string has length s and there are n strings in the array. The time complexity of sorting each string is O(slog s), and the time complexity of sorting each string in alphabetical order in the array is O(n \_slog s).
The time complexity of sorting the entire string array lexicographically is O(s \_nlog n). O(nlog n) in the sorting algorithm is the number of comparisons. Since the comparison is an integer number, each comparison is O(1). But strings are compared lexicographically, and the time complexity is O(s). So the time complexity of sorting an array of strings lexicographically is O(s \_nlog n). So the overall complexity is O(n \_slog s) + O(s \_nlog n) = O(n\*slog s + s\*nlogn) = O(n\*s\*(log s + log n) )

## Two. Space complexity

The recursive call has a space cost, and the recursive algorithm needs to save the recursive stack information, so the space complexity spent will be higher than that of the non-recursive algorithm.

```c
int sum( int n ){
    assert( n >= 0 )
    int ret = 0;
    for ( int i = 0 ; i <= n ; i++)
        ret += i;
    return ret;
}
```

The time complexity of the above algorithm is O(n), and the space complexity is O(1).

```c
int sum( int n ){
    assert( n >= 0 )
    if ( n == 0 )
        return 0;
    return n + sum( n -1);
}
```

The time complexity of the above algorithm is O(n), and the space complexity is O(n).

## 3. Time complexity of recursion

### Only one recursive call

If only one recursive call is made in the recursive function, and the recursive depth is depth, and in each recursive function, the time complexity is T, then the overall time complexity is O(T \*depth)

for example:

```c
int binarySearch(int arr[], int l, int r, int target){
if( l > r)
return -1;
    int mid = l + (r-l)/2;//anti-overflow
    if(arr[mid] == target)
        return mid;
    else if (arr[mid]>target)
        return binarySearch(arr,l,mid-1,target);
    else
        return binarySearch(arr, mid+1, r, target);
}

```

In the recursive implementation of binary search, only itself is called recursively. The recursion depth is log n, and the complexity of each recursion is O(1), so the time complexity of the recursive implementation of binary search is O(log n).

### Only multiple recursive calls

For the case of multiple recursive calls, you need to look at the number of times it is calculated and called. Usually you can draw a recursive tree to see. Example:

```c
int f(int n){
    assert( n >= 0 );
    if( n ==0 )
        return 1;
    return f( n -1 ) + f ( n -1 );

```

The number of recursive calls above is 2^0^ + 2^1^ + 2^2^ + ... + 2^n^ = 2^n+1^ -1 = O(2^n)

> For the complexity analysis of more complex recursion, please refer to, the main theorem. The main theorem gives correct conclusions for various complex situations.
