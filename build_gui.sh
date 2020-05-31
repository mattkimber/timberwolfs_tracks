#!/bin/bash

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
