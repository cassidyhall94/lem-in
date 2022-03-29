You're reading the data about the *structure* of an Ant Farm from a text file

This structure contains:

- Rooms: ID X_pos Y_pos
    - Ant Farm's are normally 2D, X_pos and Y_pos are positions in space and bear no relation from one room to the next. Technically if 2 rooms have the same X & Y then they intersect and this should probably throw an error but that's a problem for the visualiser...probably
- Links: Room_ID-Room_ID
- Comments: #some text, this is a comment
- Start Comment, A special comment that says "the next line that is a room is the room that ants start in": ##start
- End Comment, A special comment that says "the next line that is a room is the room that ants end in": ##end

The overall task is:

```
There is an ant farm, it has ants in it that will move from a start room to an end room. Each room can hold only one ant, each room can connect to n other rooms, connections between rooms cannot hold ants.

Find the shortest path for the ants to take and print out the journey for each any through the farm.
```

Your jobs are:

- Read the file and structure the data in it to be useful for the task
- Find the shortest path between the start room and the end room in terms of numbers of rooms moved through
- Print that movement for each ant