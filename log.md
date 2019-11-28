Thought process while follwing the tutorial

## Part 1: Building a simple REPL

- Pretty much straight forward from the tutorial.

## Part 2: World's simplest SQL compiler and VM

- Using `const` with `iota` in Go to emulate the enum structures in the tutorial
- structs are defined in the same manner as in the tutorial; No other changes

## Part 3: An In-Memory, Append only, Single-Table Database

- First challenge is this code block:
```c
+#define size_of_attribute(Struct, Attribute) sizeof(((Struct*)0)->Attribute)
+
+const uint32_t ID_SIZE = size_of_attribute(Row, id);
+const uint32_t USERNAME_SIZE = size_of_attribute(Row, username);
+const uint32_t EMAIL_SIZE = size_of_attribute(Row, email);
+const uint32_t ID_OFFSET = 0;
+const uint32_t USERNAME_OFFSET = ID_OFFSET + ID_SIZE;
+const uint32_t EMAIL_OFFSET = USERNAME_OFFSET + USERNAME_SIZE;
+const uint32_t ROW_SIZE = ID_SIZE + USERNAME_SIZE + EMAIL_SIZE;
```
It seems that Go does not have `macros` in the language so this block can't be implemented in the exact same way.
- What we're trying to do is create the idea of a `page` and fitting our rows and the data contained, into these pages. 
- What `size_of_attributes` does is determine the size of a single member in any arbitary struct
- The whole point of doing this is to be able to serialize/deserialize rows into 