## Running command

```sh
python3 evaluateShared.py --cmd "go run main.go" --problemDir problems
```

### Notes
It's not the best solution, not the most effecient one 
To implement the most effecient, I would probably need to use some kind of dp algorithm which would iterate over all possible combinations

I am familiar with problem of minimal spanning trees, and had in mind kruskal's algorithm and Prim's algorithm (this solution is kind of Prim's algorithm)

https://en.wikipedia.org/wiki/Prim%27s_algorithm
https://en.wikipedia.org/wiki/Kruskal%27s_algorithm

I've spent more than I planned initially, but I would be glad to tinker around and come up with a really performant + advanced alogirhtm (implementation of which would take considerably more time)
I was thinking about dynamic-programming + some kind of clustering of closes nodes...or chains of nodes (dropoff + pick-up points)

It's really interesting challenge and I've enjoyed it, though spent more time without really "impressive" outcome (just greedy + Prim's algorithm) 