//这是一个滑动窗口（sliding window)问题
//The Sliding Problem contains a sliding window which is a sub – list that runs over a Large Array which is an underlying collection of elements.
//“滑动问题”包含一个滑动窗口，该窗口是一个子列表，它运行在一个大数组上，该大数组是一个基础元素集合。
//A brute force solution to this problem will give a runtime complexity of O(nw) which is not efficient. 
//Implementation of Sliding Window Protocol Problem using Heap Data Structure will give a runtime complexity of O(n Log w).
//If a Queue is used to solve obtain solution to Sliding Window Problem, O(n) runtime complexity can be achieved.
//这个问题的蛮力解决方案会给运行时带来 o (nw)的复杂性，这是不高效的。 
//使用堆数据结构实现滑动窗口协议会给运行时带来 o (nlogw)的复杂性。 如果使用队列来求解滑动窗口问题，则可以获得 o (n)运行时复杂度。
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:

// 输入: "abcabcbb"
// 输出: 3 
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
package main



func lengthOfLongestSubstring(s string) int {
    // 定义游标尺寸大小,游标的左边位置
    window,start := 0,0

    // 循环字符串
    for key := 0; key < len(s); key++ {
        // 查看当前字符串是否在游标内
        isExist := strings.Index(string(s[start:key]), string(s[key]));
        
        // 如果不存在游标内部,游标长度重新计算并赋值
        if (isExist == -1) {
            if (key - start + 1 > window) {
                window = key - start + 1
            }
        } else { //存在，游标开始位置更换为重复字符串位置的下一个位置
            start = start + 1 + isExist
        }
    }
    return window
}
