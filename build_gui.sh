#!/bin/bash

# GUI
mkdir -p gui
mkdir -p gui/1x
mkdir -p gui/2x

# Depot GUI
../gorender/renderobject.exe $1 -i voxels/depot_1880.vox -o gui/1x/depot -s 1 -m files/manifest_gui.json
../gorender/renderobject.exe $1 -i voxels/depot_ng_1890.vox -o gui/1x/depot_ng -s 1 -m files/manifest_gui.json
../gorender/renderobject.exe $1 -i voxels/depot_1880.vox -o gui/2x/depot -s 2 -m files/manifest_gui.json
../gorender/renderobject.exe $1 -i voxels/depot_ng_1890.vox -o gui/2x/depot_ng -s 2 -m files/manifest_gui.json


# Rail types
for i in rail elrl tdrl f_rl ngrl ptwy; do
    echo $i
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/1x/${i} -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/1x/${i}_crossing -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/2x/${i} -s 2 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/2x/${i}_crossing -s 2 -m files/manifest_gui.json
    goo/goo.exe $i
done

# Rail types
for i in ngrl; do
    echo $i
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/1x/${i} -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/1x/${i}_crossing -s 1 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}.vox -o gui/2x/${i} -s 2 -m files/manifest_gui.json
    ../gorender/renderobject.exe $1 -i voxels/${i}_crossing.vox -o gui/2x/${i}_crossing -s 2 -m files/manifest_gui.json
    goo/goo.exe -depot depot_ng $i
done
