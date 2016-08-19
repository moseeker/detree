# dtree

WIP.

Build the dependency tree for javascript and sass/less files.

## How it works

dtree reads a list of entry files, process each entry file, to calculate
the `deps` and `refs` of each sub modules. The `refs` stands for the 
entry files that use it or indirect use it.

## TODO

* Prevent cyclic dependencies.

---
&copy; 2016 Towry Wang
