//Topological Sort

/*
* Suppose we have some prerequisites [u, v] , where v depends on u
* We can consider "u" as parent and "v" child , its like child depends on parent.
* So parents should be processed first and then their children.
* Initially it is needed to identify the parents,
* Once those parents are visited add them to a visited list, then we need to add their
* children as new parents if their indegree is 0 i:e it has only edge from its parent.
*
* If the visited list length equals to the length of all prerequisite tasks, then
* Topological sort is possible. So Topological sort is not possible
* in case of cyclic dependency graph.
 */

package main
