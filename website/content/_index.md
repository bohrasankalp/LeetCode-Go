---
title: 序
type: docs
---

# sequence

{{< columns >}}
## About LeetCode

Speaking of LeetCode, as a programmer, you should be familiar with it, and it will be mentioned in interviews in recent years. Programmers at home and abroad use it to brush questions mainly for interviews. According to historical records, this website was established in 2011, and it is about to reach its 10th birthday. Weekly competitions, biweekly competitions, and monthly competitions are held every week, and coding within a limited time is really a test of people's algorithmic ability. In addition to prizes for the top few winners of some big companies sponsoring title competitions, they can also directly get the opportunity to be promoted internally.

<--->
## What is Cookbook

The literal translation is a cookbook, a book that teaches you to make various recipes and delicacies. Students who often read O'Reilly technical books will be familiar with this term. Generally hands-on, practical books will have this name.

{{< /columns >}}

<img src="https://books.halfrost.com/leetcode/logo.png" alt="logo" height="600" align="right" style="padding-left: 30px;"/>
## Why did you write this open source book

The author has been brushing the questions for a year, and I want to share with you some experience in doing the questions and methods of solving the questions. I want to make friends with people who have the same hobbies, and exchange and study together. For myself, writing problem solutions is also an improvement. Telling a esoteric topic to someone who has no idea at all, and allowing him to fully understand it, can exercise people's ability to express. During the explanation, you may also encounter some questions from the listeners. These questions may be your own knowledge loopholes, and you are forced to make up for them. The author has done relevant sharing in the company, and I feel deeply, and both parties benefited.

> In addition, when I was in college, I hated writing problem solutions the most when I was doing questions. I felt that it was a waste of time and spent more time doing more questions. Now I don't know if it counts as "you have to pay back if you come out to mess around".


## About the book cover
Students who often read O'Reilly's animal books know that this cover is a tribute to them. It is indeed for this purpose. O'Reilly's cover animals are all rare animals, and the style of painting is black and white sketch style. These animals are all copyrighted, so I can only find black-and-white sketch-style pictures without copyright on the Internet. It is common to find 40 images in this style. However, there are too many people using it, and the author struggled to find several other pictures of this kind, and this peacock with its tail open is one of them. The meaning of Peacock Open Screen is to hope that after finishing LeetCode, everyone can improve their algorithmic ability and open their own "screen" on the stage of life. The color scheme of the whole book is also green, because this is the color of AC.


## About the author
The author is a gopher newcomer who has just been in the industry for a year and a half. Please give me some advice. The university participated in ACM-ICPC for 3 years, but did not get a gold medal due to low qualifications. So in terms of algorithms, I'm a novice in my own evaluation. The biggest gain from participating in ACM-ICPC is to train thinking ability, which will also be used in life. Secondly, I got to know many very smart players in China and saw the gap between myself and them. Finally, there are more than 200 pages, some of which I did not fully understand, and the densely printed [algorithm template](https://github.com/halfrost/leetcode-go/releases/tag/Special). If you learn knowledge, you will own it for life. If you don't learn it, that knowledge is something outside your body.

The author started writing questions from March 25, 2019 to March 25, 2020, a whole year. The original plan was one question per day. In fact, there are sometimes more than one question every day, and finally completed 600+:

![](https://img.halfrost.com/Blog/ArticleImage/2019_leetcode.png)
> A warm reminder: The author thought that doing one question every day would make the submissions picture all green, but I found out that I was wrong. If you also want to insist and make this picture all green, you must pay attention to the following issues: LeetCode server is in +0 time zone, and this picture is also calculated according to this time zone. In other words, before 8 o'clock in the morning in China, the previous day is counted! It is also because of the time zone that I left these 22 grids blank. For example, there is a hard question that is difficult, and there is a lot of work that day, and it will not be until the next morning when I come home from get off work at night and figure it out. So do another question as the next day's amount. It turns out that these 2 questions are counted as the previous day. Sometimes the author gets up at 6 o'clock in the morning to brush up the questions, and the submission is also the day before.
>
> (Of course, these are all in the past, they are not important anymore, they are just some small episodes on the road of struggle)

In 2020, the author will definitely continue to write questions, because some of my goals have not been achieved yet. You may work towards 1000 questions, or you may go back and start doing the second or third time when you reach 800 questions. (Don't give up if you don't reach your goal~)

## About the code in the book
The codes are all placed in [github repo](https://github.com/halfrost/leetcode-go/tree/master/leetcode), and the questions can be searched by question number.
The code for the title of this book has been beats 100%. If there is no beats 100% solution, it will not be included in this book. The author will continue to optimize those topics to 100% before putting them in.

Readers may ask, why pursue beats 100%. The author thinks that optimizing to 100% beats can be regarded as a feeling for this question. There are several Hard questions, and the author used brute force to solve AC, and then only beats 5%. This question is like not doing it. And if such an answer is given during the interview, the interviewer will not be satisfied, "Is there a better solution?". If you can give a better solution through your own thinking, the interviewer will be more satisfied.

LeetCode statistics code runtime will fluctuate, the same code submitted 10 times may beats 100%. The author didn't find this problem at first, and I submitted many questions with the correct code many times in a row, submitting 3400+ times a year, which caused my correct rate to become extremely high. 😢
Of course, if there are other more elegant solutions that can beat 100%, welcome to submit a PR, and the author will learn with everyone.

## target audience

Programming enthusiasts who want to improve their algorithmic ability through LeetCode.


## Programming language

The algorithms in this book are all implemented in Go language.

## Instructions for use

-There is a search bar in the upper left corner of this e-book, which can quickly help you find the chapter and question number you want to read.
-Each page of this e-book is connected to Gitalk, and there is a comment box at the bottom of each page for comments. If it is not displayed, please check your own network.
-Regarding the problem solution, the author recommends using it in this way: first read the problem by yourself, and think about how to solve the problem. If you don't have an idea in 15 minutes, then look at the author's problem-solving ideas first, but don't look at the code. After you have an idea, you can implement it with code yourself. If you can't write at all, then look at the code provided by the author, find out where you can't write, find out the problem and write it down, this is the knowledge gap that you want to make up for. If you realize it by yourself and there is an error after submission, you should debug it yourself first. Before AC reaches 100% in the future, I will first think about how to optimize it. If each question can be optimized to 100%, then the improvement will be great after a period of time. So in general, if you really have no idea, look at the problem-solving ideas; if you really can’t optimize 100%, look at the code.

## Interaction and Errata
If there are any missing articles in the book, please click the edit button at the bottom of the page to comment and interact. Thank you for your support and help.

## at last

Let's start brushing questions together~

![](https://img.halfrost.com/Blog/ArticleImage/hello_leetcode.png)

This work adopts [Intellectual Attribution-Non-Commercial Use-No Derivatives (BY-NC-ND) 4.0 International License Agreement](https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode.zh-Hans ) for permission.

The copyright of all the questions in the solution is owned by [LeetCode](https://leetcode.com/) and [Leetcode China](https://leetcode-cn.com/)
the