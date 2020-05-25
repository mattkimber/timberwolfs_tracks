#!/bin/bash

mkdir -p intermediate

../cargopositor/cargopositor.exe -o intermediate -v voxels positor/ballast.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_underlay.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_bslope.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_crossing.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/diagonals.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/lc.json

../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/fours.json
../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/ones.json
../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/clip.json
../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/clip_rev.json


mkdir -p intermediate/1
mkdir -p intermediate/2
mkdir -p intermediate/4
mkdir -p intermediate/4f
mkdir -p intermediate/8

mv intermediate/*_1.vox intermediate/1
mv intermediate/*_2.vox intermediate/2
mv intermediate/*_4f.vox intermediate/4f
mv intermediate/*_4.vox intermediate/4

cp voxels/lc*open.vox intermediate/8
cp voxels/lc*closed.vox intermediate/8


# Render LC pieces
# Temporary for testing
#cp voxels/tile.vox intermediate/6/tile.vox
#mkdir -p intermediate/pieces
#cp voxels/tracks_2_level_crossing_pieces.vox intermediate/4/tracks_2_z_lc.vox



for i in `ls intermediate/1`; do 
    echo "$i"
	../gorender/renderobject.exe -i intermediate/1/$i -o $i -s 1,2 -u -m files/manifest_1x.json
done

for i in `ls intermediate/2`; do 
    echo "$i"
	../gorender/renderobject.exe -i intermediate/2/$i -o $i -s 1,2 -u -m files/manifest_2x.json
done

for i in `ls intermediate/4`; do 
    echo "$i"
	../gorender/renderobject.exe -i intermediate/4/$i -o $i -s 1,2 -u -m files/manifest_4x.json
done

for i in `ls intermediate/4f`; do 
    echo "$i"
	../gorender/renderobject.exe -i intermediate/4f/$i -o $i -s 1,2 -u -m files/manifest_4x_f.json
done

for i in `ls intermediate/8`; do 
    echo "$i"
	../gorender/renderobject.exe -i intermediate/8/$i -o $i -s 1,2 -u -m files/manifest_8x.json
done

# Remove mask files for now
#../splatter/splatter.exe -i 1x -o sheets_1x -d spritesheet.json -m 4 -k files/mask_1x.png
#../splatter/splatter.exe -i 2x -o sheets_2x -d spritesheet.json -m 8 -k files/mask_2x.png

mkdir -p sheets_1x
mkdir -p sheets_2x

../splatter/splatter.exe -i 1x -o sheets_1x -d spritesheet.json -m 4
../splatter/splatter.exe -i 2x -o sheets_2x -d spritesheet.json -m 8

# Assemble the output using Roadie
../roadie/roadie.exe set.json