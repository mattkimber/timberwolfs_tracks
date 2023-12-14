#!/bin/bash

# Assemble the output using Roadie
../roadie/roadie.exe set.json

# NML
../jgr-nml/nmlc.exe -c timberwolfs_tracks.nml

echo "Building TAR"
mkdir -p timberwolfs_tracks
mv *.grf timberwolfs_tracks
cp grf_readme/* timberwolfs_tracks
tar -c timberwolfs_tracks > timberwolfs_tracks.tar
