#!/bin/bash

mkdir -p intermediate


../cargopositor/cargopositor.exe -o intermediate -v voxels positor/ballast.json

# Ballasting
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_underlay_stone.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_underlay_brick.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_ballast_stone.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_ballast_brick.json

# Crossings
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_crossing_brick.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_crossing_stone.json

# Slopes/diagonals
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/tracks_bslope.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/diagonals.json

# Level crossings
cp intermediate/ptwy_ballasted_2.vox intermediate/lc_plateway_track.vox
cp intermediate/rail_ballasted_2.vox intermediate/lc_modern_track.vox
cp intermediate/rail_ballasted_2.vox intermediate/lc_old_track.vox
cp intermediate/rail_ballasted_2.vox intermediate/lc_historic_track.vox
cp intermediate/ngrl_ballasted_2.vox intermediate/lc_ngrl_track.vox

../cargopositor/cargopositor.exe -o intermediate -v voxels positor/lc.json
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/lc_narrow_gauge.json

# Clip diagonals
../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/clip.json
../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/clip_rev.json

# Depots
../cargopositor/cargopositor.exe -o intermediate -v voxels positor/depots.json


#../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/fours.json
#../cargopositor/cargopositor.exe -o intermediate -v intermediate positor/ones.json


mkdir -p intermediate/1
mkdir -p intermediate/2
mkdir -p intermediate/4
mkdir -p intermediate/4f
mkdir -p intermediate/8
mkdir -p intermediate/fences
mkdir -p intermediate/depots

mv intermediate/*_1.vox intermediate/1
mv intermediate/*_2.vox intermediate/2
mv intermediate/*_4f.vox intermediate/4f
mv intermediate/*_4.vox intermediate/4

mv intermediate/*_depot_a.vox intermediate/depots
mv intermediate/*_depot_b.vox intermediate/depots
mv intermediate/*_depot_c.vox intermediate/depots
mv intermediate/*_depot_d.vox intermediate/depots
mv intermediate/*_depot_e.vox intermediate/depots

# Level crossing sprites
cp voxels/lc*open.vox intermediate/8
cp voxels/lc*closed.vox intermediate/8

# Fences
cp voxels/fence* intermediate/fences

cp voxels/lc_historic_open.vox intermediate/8/lc_plateway_open.vox
cp voxels/lc_historic_closed.vox intermediate/8/lc_plateway_closed.vox
cp voxels/lc_historic_open.vox intermediate/8/lc_ngrl_open.vox
cp voxels/lc_historic_closed.vox intermediate/8/lc_ngrl_closed.vox

# Render LC pieces
# Temporary for testing
#cp voxels/tile.vox intermediate/6/tile.vox
#mkdir -p intermediate/pieces
#cp voxels/tracks_2_level_crossing_pieces.vox intermediate/4/tracks_2_z_lc.vox

# Move catenary away from 4f sprites
mv intermediate/4f/catenary_narrow_slope_4f.vox intermediate/catenary_narrow_hill.vox
mv intermediate/4f/catenary_wide_1_slope_4f.vox intermediate/catenary_wide_1_hill.vox
mv intermediate/4f/catenary_wide_2_slope_4f.vox intermediate/catenary_wide_2_hill.vox

# Overrides
#cp overrides/* -r intermediate/


function render {

    for i in `ls $2`; do
        fn=`echo 2x/${i}_32bpp.png | sed -e s/.vox//`

        if [ -e $fn ]; then 
            echo "$i [cached]"
        else
            echo "$i"
	        ../gorender/renderobject.exe $1 -i $2/$i -o $i -s 1,2 -u -m files/manifest_$3.json
        fi
    done

}

# Signals
render "$1" voxels/signals signal

# General
render "$1" intermediate/1 1x
render "$1" intermediate/2 2x
render "$1" intermediate/4 4x
render "$1" intermediate/4f 4x_f
render "$1" intermediate/8 8x

# Fences (fixed, slopes are done by cargopositor and use 4x_f template)
render "$1" intermediate/fences fence

# Depots
render "$1" intermediate/depots depot

# GUI
mkdir -p gui

# Depot GUI
../gorender/renderobject.exe $1 -i voxels/depot_1880.vox -o gui/depot -s 1 -m files/manifest_gui.json
../gorender/renderobject.exe $1 -i voxels/depot_ng_1890.vox -o gui/depot_ng -s 1 -m files/manifest_gui.json

# Rail types
for i in rail elrl tdrl f_rl ngrl ptwy; do
    echo $i
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/${i} -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/${i}_crossing -s 1 -m files/manifest_gui.json
    goo/goo.exe $i
done

# Rail types
for i in ngrl; do
    echo $i
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/${i} -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/${i}_crossing -s 1 -m files/manifest_gui.json
    goo/goo.exe -depot depot_ng $i
done


# Catenary
../gorender/renderobject.exe $1 -i voxels/catenary_narrow.vox -o catenary_narrow -s 1,2 -u -m files/manifest_catenary_4.json
../gorender/renderobject.exe $1 -i intermediate/catenary_narrow_hill.vox -o catenary_narrow_hill -s 1,2 -u -m files/manifest_catenary_hill.json

../gorender/renderobject.exe $1 -i voxels/catenary_wide_1.vox -o catenary_wide_1 -s 1,2 -u -m files/manifest_catenary_4.json
../gorender/renderobject.exe $1 -i intermediate/catenary_wide_1_hill.vox -o catenary_wide_1_hill -s 1,2 -u -m files/manifest_catenary_hill.json
../gorender/renderobject.exe $1 -i voxels/catenary_wide_2.vox -o catenary_wide_2 -s 1,2 -u -m files/manifest_catenary_4.json
../gorender/renderobject.exe $1 -i intermediate/catenary_wide_2_hill.vox -o catenary_wide_2_hill -s 1,2 -u -m files/manifest_catenary_hill.json

../gorender/renderobject.exe $1 -i voxels/catenary_tunnel.vox -o catenary_x_tunnel -s 1,2 -u -m files/manifest_catenary_tunnel.json

# Catenary pylon
../gorender/renderobject.exe $1 -i voxels/pylon.vox -o pylon -s 1,2 -u -m files/manifest_pylon.json

# Sprite Overrides for 2x
#cp override_sprites/2x/* 2x/

mkdir -p sheets_1x
mkdir -p sheets_2x

../splatter/splatter.exe -i 1x -o sheets_1x -d spritesheet_masked.json -m 4 -k files/mask_1x.png
../splatter/splatter.exe -i 2x -o sheets_2x -d spritesheet_masked.json -m 8 -k files/mask_2x.png

../splatter/splatter.exe -i 1x -o sheets_1x -d spritesheet.json -m 4
../splatter/splatter.exe -i 2x -o sheets_2x -d spritesheet.json -m 8

# Assemble the output using Roadie
../roadie/roadie.exe set.json

# NML
../nml/nmlc.exe -c timberwolfs_tracks.nml

echo "Building TAR"
mkdir -p timberwolfs_tracks
mv *.grf timberwolfs_tracks
cp grf_readme/* timberwolfs_tracks
tar -c timberwolfs_tracks > timberwolfs_tracks.tar
