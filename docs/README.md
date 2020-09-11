# pset-build

This is a small script and C++ build template suitable for small programming
challenges or proofs of concept.

## Usage

- Setup:

```console
$ source setup
Setup is complete
```

Now the alias `build` will direct to the correct script in the workspace.

- To create a new file from the template:

```console
$ build new xyz
New file "xyz.cpp" created from template
```

- To compile the file (or any file not called "new")

```console
$ build xyz.cpp
Compiling xyz.cpp...
Executable "xyz.ex" built without errors
```

I do notice that the `.ex` extension is associated with Elixir files. If this is
a problem for you, feel free to modify the `bin/build` script accordingly.

- Running the binary:

```console
$ ./xyz.ex
OK
```

- Clearing out the workspace

```console
$ build clean
All *.cpp and *.ex files are removed from the workspace
```

### TODO

- [ ] Testcase automation?

© 2020 Damien Stanton

See LICENSE for details.
