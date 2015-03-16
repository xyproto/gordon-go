# Flux #
The aim of this project is to provide a means to more efficiently write, read, and modify programs.  It remains to be seen whether it will succeed; it is far from complete and many of its details have yet to be worked out.  It is clear that plenty of experimentation and creativity will be needed.  Your thoughtful feedback is crucial.

I call it Flux because it was inspired by graphical dataflow programming languages like Pure Data, Max/MSP, LabView, and Reaktor, just to name a few.  However, it resembles these languages only on the surface.  Flux is essentially a graphical interface to the Go programming language.  Instead of using text to represent types, operations, and control flow, it uses graphics.  In Flux, text appears only in identifiers and comments.

I posit that text is not the best representation for writing, reading, and modifying programs by humans.  Text certainly has a lot going for it:  a standardized set of characters, the ubiquity of text editors, and relative familiarity.  Text is expressive, and it provides a great medium for quickly producing decent solutions to many problems, from writing emails to debugging programs to hunting the Wumpus.  In general, however, and in programming in particular, it is both too expressive and not enough.

Text is too expressive:  Only an infinitesimal fraction of character strings form valid programs.  This vastly overstates the problem; we are clever monkeys with fancy typewriters.  Even so, we waste a significant amount of time and effort trying to avoid syntax and type errors when programming.  Furthermore, when reading code, the effort of parsing is massively duplicated among all readers.

Text is not expressive enough:  It is only 1-dimensional (maybe 1.5 with line breaks) whereas a program's structure includes graphs that are better visualized in at least 2 dimensions.  Moreover, although text provides a broad selection of characters to choose from, they are not tailored to our task.

Flux provides an environment in which only valid (compilable) programs can be created and manipulated.  It displays them in a way that requires minimal interpretation by the viewer (at least this is the goal; "minimal interpretation" is hard to determine).  Because Flux embodies the semantics of Go, it should easily support refactorings and other structured manipulations.

It's not a new idea, of course, but I'm surprised that it has not caught on in general-purpose programming.  The rationale for it seems strong enough.  My guess is that past systems were not developed to their full potential, and that the general idea was dismissed when particular implementations failed to deliver.

Flux avoids some of the major weaknesses of other graphical languages.  Mouse editing is slow, so Flux is keyboard-centric (with optional mouse support).  It is tedious to manually arrange the nodes of a graph, so Flux does it automatically.  Even a fairly simple function can start to look like spaghetti, so Flux supports wireless (named) connections.

I mentioned that Flux is an interface to Go.  To elaborate:  Flux supports pretty much everything but the syntax of Go: its type system, control flow and concurrency mechanisms, builtins and operators.  It does so simply by generating Go code.  However, Flux is not an interface to arbitrary hand-written Go code.  Flux writes and reads only a subset of Go code that is structured in a verbose, easily machine readable (and poorly human readable) way.  That it generates Go code is an implementation detail; it might just as well compile to machine code.  The important thing to take away is that Flux is interoperable with Go:  Flux code can refer to objects defined in Go code and vice versa; they can even co-exist in the same package.

As I mentioned, Flux is incomplete; there are many features I would love to add and many more as yet unimagined.  It is also experimental; little is set in stone, so if something doesn't look right to you, please suggest an alternative.

## Install ##
Flux requires the [GLFW3](http://glfw.org) and [FTGL](http://sourceforge.net/projects/ftgl) libraries and header files.  If your system has a package manager, you can use it to install `libglfw3-dev` and `libftgl-dev`, if they are available.  Otherwise, you must build and install them from source:  FTGL requires [FreeType](http://freetype.org); install both libraries using the standard commands `./configure && make && sudo make install`.  On Linux, GLFW3 requires the `xorg-dev` and `libglu1-mesa-dev` packages.  To build GLFW3 you need [CMake](http://cmake.org).  Configure GLFW3 with the BUILD\_SHARED\_LIBS option: run `cmake -D BUILD_SHARED_LIBS=true . && make && sudo make install`.

Once you have GLFW3 and FTGL, you can `go get code.google.com/p/gordon-go/flux`.

## Usage ##
Please refer to the [package documentation](http://godoc.org/code.google.com/p/gordon-go/flux) for usage instructions.