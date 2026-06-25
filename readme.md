# Headache

A simple programming language that makes Brainfuck easier to read and write. It replaces Brainfuck's symbols with words like `increment`, `forward`, and `print`, while still compiling directly into standard Brainfuck code.

## Language Commands

| Command       | Brainfuck | Description                 |
| ---------------| -----------| -----------------------------|
| `forward`     | `>`       | Move pointer right 1 cell   |
| `forward N`   | `>` × N   | Move pointer right N cells  |
| `back`        | `<`       | Move pointer left 1 cell    |
| `back N`      | `<` × N   | Move pointer left N cells   |
| `increment`   | `+`       | Increment current cell by 1 |
| `increment N` | `+` × N   | Increment current cell by N |
| `decrement`   | `-`       | Decrement current cell by 1 |
| `decrement N` | `-` × N   | Decrement current cell by N |
| `input`       | `,`       | Read one byte               |
| `input N`     | `,` × N   | Read N bytes                |
| `print`       | `.`       | Output current cell         |
| `print N`     | `.` × N   | Output current cell N times |
| `loop`        | `[`       | Begin loop                  |
| `end`         | `]`       | End loop                    |

## Comments

```headache
// Full line comment

increment 65 // Inline comment
print
```

Everything after `//` is ignored.

## Case Insensitive

The translator converts everything to lowercase before parsing:

```headache
increment 65
Print

FoRwArD 2
increment 10
PRINT
```

is equivalent to:

```headache
increment 65
print

forward 2
increment 10
print
```

## Repeat Counts

Most commands accept an optional number:

```headache
increment 10
forward 5
decrement 3
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
increment 5
loop
    decrement
    forward
    increment
    back
end

forward
print
input
decrement 10
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
increment 3
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
increment 72
print

forward
increment 101
print

forward
increment 108
print

forward
increment 108
print

forward
increment 111
print

forward
increment 44
print

forward
increment 32
print

forward
increment 87
print

forward
increment 111
print

forward
increment 114
print

forward
increment 108
print

forward
increment 100
print

forward
increment 33
print

forward
increment 10
print
```