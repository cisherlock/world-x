<p align="center">
  <img src='https://nightofcode.com/github/craigsherlock/world-x/world-x.jpg' />
</p>

## World-X Monster Simulation
You are competing to be an evil overlord. World-X is your first target. Go forth and unleash as many monsters as you wish.

## About
This was a weekend project and my first time using Go. Built as part of an interview process.

## Instructions
##### 1. Compile
```
$ go build
```

##### 2. Run
```
$ world-x <number_of_monsters> <small|medium>
```

##### 3. Example
To run with 150 monsters using the small map
```
$ world-x 150 small
```
It takes time to create monsters. The more you add the longer you will have to wait before they can start destroying things.

## Simulation Stages
1.  All monsters are randomly spawned in cities across World-x
2.  Each turn a random number of monsters are chosen to move
3.  Each monster then chooses a random path to follow to the next city
4.  If a city is visited by 2 or more monsters at a time, the monsters fight and the city is immediately obliterated along with any mosters and paths in and out of the city


##### Stuck Monster
Monsters can get stuck. For example: a single monster may be trapped between 2 connected cities with no paths leading in or out.
The monster will run between the two cities until it reaches the maximum move count of 10,000 - at which point the monster will die from exhaustion.


##### End of Program
At the end of the program a map is given of the remaining cities.

## Author
Craig Sherlock [@nightofcode](https://twitter.com/nightofcode)