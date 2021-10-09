#!/bin/bash

CSV_FILENAME="teste.csv"

./team/start.sh -t team >/dev/null 2>&1 &
./gliders/start.sh -t gliders >/dev/null 2>&1 &
rcssserver server::synch_mode=true CSVSaver::save=true CSVSaver::filename=$CSV_FILENAME server::auto_mode=true server::penalty_shoot_outs=false server::nr_extra_halfs=0 >/dev/null 2>&1
sleep 1

tail -n1 $CSV_FILENAME
rm $CSV_FILENAME
