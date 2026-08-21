# Task 04: The Pirate King's Challenge
## The 67th OEIS Problem
Most of my time went to figuring out the trick to this question. In layman's terms, multiplying primes consecutively will ensure that the gcds of the adjacent terms will be a prime. The constructive algorithm was such that the gcds taken in order will result in an array of sorted primes.
## Digit string
This is the last problem i did. ive been looking into this problem for a week now. I gave up and looked at the tutorial. My initial approach was to strip out the 1s and 3s on the left and strip out the 2s on the right and count and return the min of all strips. Dunno why it didnt work. the author basically did the reverse of my solution. it worked.
## Another Puzzle from Papyrus
Had to refer to the editorial. The trick here was the total no. of subtractions were independent of whichever permutation of the question. Rest was easy.
## Good times Good times
Couldn't figure out the algorithm for this either. Examples kinda decieved me. At the end got that the answer was 10^len(input)+1. Basically concatenating x two times. since x is already good, duplicating it will also give a good answer and 10^something+1 only have 1s and 0s.
## Duck Surplus
This one was pretty easy. I did what was exactly asked of me in the question. Did not have to "optimize" anything.