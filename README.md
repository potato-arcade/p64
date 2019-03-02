# p64
The P64 computer console 

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


## The Machine

The Potato console contains the following components :

- The CPU
- The ROM Cartridge Slot
- The Memory Banks
- Video Controller
- The IO controller
- The Audio Controller

## ROM Cartridges

## Roll your own ROM


## TODO

- Overlay graphical image of the console, like Andy's Beeb emulator

- Bootup static to look like real static

- Graphics primitives mapped to PotatoBASIC
    - LINE x,y,x2,y2,style
    - CIRCLE x,y,r
    - TEXT x,y,string
    - RECT x,y,x2,y2
    - FILLBOX x,y,x2,y2

- Sprites !

- Memory bank access to framebuffer

- Memory bank access to code

