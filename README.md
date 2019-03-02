# Model P-64 Programmable Graphics Console

Introducing the Potato-64 emulator, written in 100% Go

(with a bit of OpenGL C code under the hood)


![potato](potato.jpg)

A fyne app  
http://fyne.io/fyne

## What is it, and why should I care ?

![pong](pong.gif)

The P64 is a virtual piece of hardware, intended as a minimalist platform to write and execute nifty little graphics programs on.

The capablities of the machine are deliberately limited 

- like the resolution (64x64)
- the refresh rate (24Hz fixed)
- the onboard memory store (64 'banks' of memory to store variables in)
- and of course ... the colour palette of the display  (Monochrome).

![bw](bw/bw2.jpg)

But like shooting black and white film in the digital age, these extremely limited constraints really force you to open up your creative side. To use the machine to tell a story.


You can write complete, and compelling games in an lunch break ....  but only if you think creatively, and code around the limitations of the hardware.  Getting your code to work well needs to you to think about the code from the machine's point of view.


Its literally BASIC, so its easy to get started with and easy to learn. But deceptively hard to master. 

![bw](bw/bw6.jpg)

Like a photographer armed with an old film camera, you have the same challenge of creating something compelling that jumps off the page. Its about the image, not the equipment.


Can YOU think up an addictive game that fits within the limitations of the hardware ?


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


Run the Basic Interpreter to test some code
```
p64basic filename
```


![bw](bw/bw7.jpg)

## The Machine

The Potato console contains the following components :

- The CPU
- The ROM Cartridge Slot
- The Memory Banks
- Video Controller
- The IO controller
- The Audio Controller

Because we are not trying to emulate any known machine, the definition of the P64 itself is entirely software driven - so we can give it additional concepts, features, bugs and limitations that do not even exist in any real hardware.

![bw](bw/bw8.jpg)

## ROM Cartridges

## Roll your own ROM

![bw](bw/bw3.jpg)

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

- Extra framebuffers, with BLIT commands to fast copy between them

- Memory bank access to framebuffer

- Memory bank access to the ROM code, so you can write self-modifying code

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