#!/bin/bash

# team1_branch team2_branch amount

if [ ! $# -eq 3 ]; then
    exit
fi

./buildSpecificBranch.sh "$1" current > /dev/null 2>&1

if [ ! $? -eq 0 ]; then
    echo Failed to build first team
    exit
fi

if [ "$2" = "gliders" ]; then
    thunderleague_test $3 32 "$1" teams/current/start.sh gliders teams/gliders2d/start.sh > /dev/null 2>&1
else
    ./buildSpecificBranch.sh "$2" new > /dev/null 2>&1

    if [ ! $? -eq 0 ]; then
        echo Failed to build second team
        exit
    fi

    thunderleague_test $3 32 "$1" teams/current/start.sh "$2" teams/new/start.sh > /dev/null 2>&1
fi

if [ ! $? -eq 0 ]; then
    echo Failed to run
    exit
else
    nl match_outcome.csv | tail -1
fi

