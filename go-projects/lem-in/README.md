This is lem-in

The program distributes a number of ants through an ant farm consisting of linked rooms.
It uses Depth First Search to find the optimal path combinations for getting all the ants to the end in the least amount of moves.

Errors are handled carefully.

A valid ant farm must have number of ants, rooms and links between them from start to end.

Here is an example of this in practice :
``
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
``

Usage: ./lem-in <antfarm.txt>

Made by Joman, 62oz, Bluesky, and Ville