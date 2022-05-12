<h1 align="center">LEM-IN</h1>

<p align="center">A digital version of an ant farm.</p>

It will read from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

<h1 align="center">USAGE</h1>

First, run `./build.sh`

Then run `bin/lem-in maps/example00.txt` into the terminal replacing the file name as needed.

There is also the ability to run the following commands in the terminal:

PrintAllMoves will display the output from running all valid files maps, in the file 'allmoves':
`./printAllMoves.sh >allmoves`

RunAllMaps will display the output from running all files (including bad examples) in maps, in the file 'allmaps':
`./runAllMaps.sh >allmaps`

GetTimedMoves will display the time in which each of the valid files will take to run, in the terminal:
`./getTimedMoves.sh `