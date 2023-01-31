# braingo
Brainfuck interpreter written in Go.

### Brainfuck
If you are not familiar with brainfuck, consider checking out [brainfuck wiki page](https://en.wikipedia.org/wiki/Brainfuck),
especially [commands chapter](https://en.wikipedia.org/wiki/Brainfuck#Commands). 

### Usage
To run code from a .bf file run:
```./braingo -r <file_path>```

Braingo also comes with an interactive mode, which comes with a few interactive commands. To enter interactive mode, just run `./braingo`
Following message should appear:
```
-----
Braingo - brainfuck interpreter written in go
Available interactive commands:
\f <file>  - run code from <file>;
\r <size>  - reset pointer & memory and set its size to <size>;
\m <d/c>   - change current IO format: 
                d - print values as digits, 
                c - print ascii symbol represented by value;
\d <f> <t> - print memory values at cells with indices from <f> to <t>;
\h         - print this message;
-----
```

### Examples
There are few [examples](https://github.com/dimaglushkov/braingo/tree/main/example) in the corresponding directory.


```
> \f example/sum.bf                 # run code from file example/sum.bf
17 25                               # it takes two numbers 
42                                  # and prints their sum

> \d 0 10                           # print memory values from 0 to 10
0 42 0 0 0 0 0 0 0 0 

> \r 10                             # reset pointer and memory & resize mem to 10
> \d 0 10                           # validate 
0 0 0 0 0 0 0 0 0 0
 
> \f example/hello_world.bf         # run code from file example/sum.bf
7210110810811132871111141081003310  # it prints "Hello World!", 
                                    # but in our case it's codes instead of characters

> \r 10                             # reset memory and pointer 
> \m c                              # change io format from digital to character
> \f example/hello_world.bf         # run code from file example/sum.bf one more time
Hello World!
```