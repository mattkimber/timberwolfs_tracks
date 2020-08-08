# Timberwolf's Tracks


A 2x zoom track set based on British railways, featuring multiple rail types.
Intended for use with Timberwolf's UK Trains, but can be used with most sets
that follow a standard railtype scheme.

## Features

* 2x zoom and 32bpp graphics compatible with original and OpenGFX style.
* Level crossings and fences update according to year.
* Depot graphics reflect build year.
* Multiple signal types for block/pre/post/combo/path signals
* Unobtrusive catenary wiring
* Overhead, third and fourth rail electric systems (with compatible
  train sets)
* Standard gauge, narrow gauge and plateway (with compatible train
  sets)

## Gameplay Notes

Track types feature balanced build and running costs.

- Unpowered standard gauge rail has average build and maintenance costs.
- Plateway and narrow gauge rail are cheap to build and maintain, but limited
  in speed and (depending on train set) vehicle selection.
- Fourth rail metro track is slightly more expensive than standard gauge rail,
  but far cheaper than other electrified types.
- Electrified track is expensive to build and maintain. 750V DC third rail is
  cheaper than 25kV overhead AC, but (depending on train set) typically has
  slower and less powerful trains available.
- Dual power track allows both 750V DC and 25kV AC locomotives to run, but
  is extremely expensive to build and maintain, so needs to be used
  sparingly.

## Developing/Building

Development notes, and how to build.

### Catenary

I expect you'd like to know about... catenary?

During development I couldn't find any good resources about how catenary worked,
so I ended up [documenting it here](catenary/README.md).

## Building from source

Building from source is unfortunately not user-friendly. You will need to build a lot of prerequisites and
have access to the GNU tools, either via a Linux or Mac environment or Windows via Git Bash or WSL.

(Note that the executables are expected to have Windows-style names, take note if you are building the
Go projects on a different platform)

### Prerequisites

You will need to obtain and build the following:

* [GoRender](https://github.com/mattkimber/gorender) - used for rendering voxel objects.
* [Cargopositor](https://github.com/mattkimber/cargopositor) - used for compositiing and transforming track tiles.
* [Roadie](https://github.com/mattkimber/roadie) - used for templating and creating NML files.

If you want to rebuild the GUI sprites, you will also need...

* Goo (included in the `goo` subdirectory) - used for generating GUI sprites

The build expects to find prerequisites in the following relative folder structure (note `.exe` extension):

* Roadie: `../roadie/roadie.exe`
* GoRender: `../gorender/renderobject.exe`
* Cargopositor: `../cargopositor/cargopositor.exe`
* Goo (if building GUI sprites): `goo/goo.exe`

### Building

To build the full set, run `./build.sh`. This will take a long time as it needs to render every track,
depot and signal sprite.
Future runs will not overwrite files, to re-render something you will need to delete its PNG file from
the `2x` directory.

To rebuild the GUI sprites, run `./build_gui.sh`. This will require Goo to be built and placed in the
correct folder.
