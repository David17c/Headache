# Headache

Headache is a simple programming language that makes Brainfuck easier to read and write. It replaces Brainfuck's symbols with words like `add`, `forward`, and `print`, while still compiling directly into standard Brainfuck code.

## Language Commands

| Command     | Brainfuck | Description                 |
| -------------| -----------| -----------------------------|
| `forward`   | `>`       | Move pointer right 1 cell   |
| `forward N` | `>` × N   | Move pointer right N cells  |
| `back`      | `<`       | Move pointer left 1 cell    |
| `back N`    | `<` × N   | Move pointer left N cells   |
| `add`       | `+`       | Increment current cell by 1 |
| `add N`     | `+` × N   | Increment current cell by N |
| `remove`    | `-`       | Decrement current cell by 1 |
| `remove N`  | `-` × N   | Decrement current cell by N |
| `input`     | `,`       | Read one byte               |
| `input N`   | `,` × N   | Read N bytes                |
| `print`     | `.`       | Output current cell         |
| `print N`   | `.` × N   | Output current cell N times |
| `loop`      | `[`       | Begin loop                  |
| `end`       | `]`       | End loop                    |

## Comments

```headache
// Full line comment

add 65 // Inline comment
print
```

Everything after `//` is ignored.

## Case Insensitive

The translator converts everything to lowercase before parsing:

```headache
ADD 65
Print

FoRwArD 2
aDd 10
PRINT
```

is equivalent to:

```headache
add 65
print

forward 2
add 10
print
```

## Repeat Counts

Most commands accept an optional number:

```headache
add 10
forward 5
remove 3
print 2
input 4
```

Compiles to:

```brainfuck
++++++++++>>>>>---..,,,,
```

---

## Complete Syntax Example

```headache
add 5
loop
    remove
    forward
    add
    back
end

forward
print
input
remove 10
print 2
back 1
forward 3
```


## CLI Usage

### Run a Headache program

```bash
headache run program.ha
```

Translates the file to Brainfuck and executes it.

### Run a Brainfuck program

```bash
headache run program.bf
```

Executes a Brainfuck file directly.

### Compile to Brainfuck and print to terminal

```bash
headache program.ha
```

Example:

```headache
add 3
print
```

Output:

```brainfuck
+++.
```

### Compile to a Brainfuck file

```bash
headache program.ha output.bf
```

Example:

```bash
headache hello.ha hello.bf
```

Generates:

```text
hello.bf
```

containing the translated Brainfuck code.

## "Hello, World!" in Headache

```headache
add 72
print

forward
add 101
print

forward
add 108
print

forward
add 108
print

forward
add 111
print

forward
add 44
print

forward
add 32
print

forward
add 87
print

forward
add 111
print

forward
add 114
print

forward
add 108
print

forward
add 100
print

forward
add 33
print

forward
add 10
print
```