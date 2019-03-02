# Model P-64 Programmable Graphics Console

Introducing the Potato-64 emulator, written in 100% Go

(with a bit of OpenGL C code under the hood)


![potato](potato.jpg)

A fyne app  
http://fyne.io/fyne


## Install it

```
go get -u github.com/potato-arcade/p64
```

## Usage

Run the Potato Console

```
potato64 filename
```

eg - `potato64 ~/go/src/github.com/potato-arcade/p64/ROM/BOUNCY.BAS` to get up and running

![pong](pong.gif)

Run the Basic Interpreter to test some code
```
p64basic filename
```


## The Machine

The Potato console contains the following components :

- The CPU
- The ROM Cartridge Slot
- The Memory Banks
- Video Controller
- The IO controller
- The Audio Controller

Because we are not trying to emulate any known machine, the definition of the P64 itself is entirely software driven - so we can expand the capabilities of the P64 to release new models all the time by adding more virtual hardware.

## ROM Cartridges

## Roll your own ROM


## TODO

- Overlay graphical image of the console, like Andy's Beeb emulator

- Bootup static to look like real static

- Extra Video Modes

- Text Mode

- Colors !!  Colored border

- Graphics primitives mapped to PotatoBASIC
    - LINE x,y,x2,y2,style
    - CIRCLE x,y,r
    - TEXT x,y,string
    - RECT x,y,x2,y2
    - FILLBOX x,y,x2,y2

- Sprites !

- Memory bank access to framebuffer

- Memory bank access to code

## If you like this, you should like these repos too

https://github.com/skx/gobasic

https://github.com/fyne-io/fyne

https://github.com/go-gl/glfw

https://gitlab.com/rastersoft/fbzx

https://github.com/remogatto/gospeccy

https://github.com/ichikaway/goNES/


## Further reading - some vids you might like, to get you in the mood for P64 coding

PONG - First documented Video Ping-Pong game - 1969
https://www.youtube.com/watch?v=XNRx5hc4gYc


c64 peeks and pokes
https://www.youtube.com/watch?v=k4BCyfpP38Q
https://www.youtube.com/watch?v=zAndQn1p5L8


Incredible story of how they recovered the Apollo mission control software from the dumpster
https://www.youtube.com/watch?v=WquhaobDqLU