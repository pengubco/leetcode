I practice them occasionally and this repo contains solutions to Leetcode problems.

## Go
There is a Go module at `solutions/go`. Each problem corresponds to a folder and thus one top-level package of the module. In the folder, there is a source file and test file. For example. 
- solutions/go/add-two-numbers/add_two_numbers.go
- solutions/go/add-two-numbers/add_two_numbers_test.go

There is no `main.go`. In order to verify the solution, run go test. For example, 
```sh
cd solutions/go
go mod tidy
cd add-two-numbers
go test
```

## Java
There is a Maven project at `solutions/java/leetcode`. Each problem corresponds to a Java package. There is a `Solution.java` and `SolutionTest.java`. For example. 
- solutions/java/leetcode/src/main/java/fyi/peng/addTwoNumbers/Solution.java
- solutions/java/leetcode/src/test/java/fyi/peng/addTwoNumbers/SolutionTest.java

In order to verify the solution, run Junit5. For example, 
```sh
cd solutions/java/leetcode
mvn compile
mvn test
```

## Some background
I studied Software Engineering at East China Normal University from 2006 to 2010. During that time, I developed interests in Algorithms and Data Structures. Together with two amazing teammates, we competed in the 
[International Collegiate Programming Contest (ICPC)](https://icpc.global/) and won a Gold Medal. It was fun!
Back then, I didn't know that ICPC problems would later become popular in software engineering interviews. 
I don't think Engineer A is better than Engineer B just because A is better at solving ICPC problems. However, solving ICPC problems shows that an engineer 
can code and talk about algorithms and data structures. So, in my opinion, using ICPC problems in interviews is like 
using credit scores for loan applications; you don't compare individuals, but rather groups.

To the memorable summer training camp of 2010! 
![AMC-ICPC 2010](./acm-icpc-2010.jpg)
