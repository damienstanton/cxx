# pset-build

This is a small script and C++ build template suitable for small programming
challenges or proofs of concept.

## Usage

- To create a new file from the template:

```console
$ make new f=xyz.cpp
New file "xyz.cpp" created from template
```

- To compile the file

```console
$ ./build xyz.cpp
Compiling xyz.cpp...
Executable "xyz" built without errors
```

- Running the binary:

```console
$ ./xyz
OK
```

- Clearing out the workspace

```console
$ make clean
All *.cpp files are removed from the workspace
```

© 2020-2021 Damien Stanton

See LICENSE for details.
