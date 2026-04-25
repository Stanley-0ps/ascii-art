# ASCII-Art
ASCII-Art is a command-line application written in **Go** that converts regular text into large decorative text using **ASCII characters**. Instead of printing normal strings, the program renders each character in an artistic banner style using predefined templates.

This project demonstrates how text can be transformed into graphical output directly in the terminal.

---

## Features

* Convert letters, numbers, symbols, and spaces into ASCII art
* Support for multiple lines using `\n`
* Uses banner template files such as:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Handles printable ASCII characters (32 to 126)
* Clean and modular Go code structure

---

## How It Works

Each character is represented by an **8-line high ASCII pattern** stored inside a banner file.
When the user enters text, the program reads the banner file, finds the matching pattern for each character, and prints the final result line by line.

Example input:

```bash id="gtwq3l"
go run . "Welcome"
```

Example output:

```text id="v4m2rx"
__        __        _                              
\ \      / /__  ___| | ___ ___  _ __ ___   ___    
 \ \ /\ / / _ \/ __| |/ __/ _ \| '_ ` _ \ / _ \   
  \ V  V /  __/ (__| | (_| (_) | | | | | |  __/   
   \_/\_/ \___|\___|_|\___\___/|_| |_| |_|\___|   
```

---

## Usage

Run the program with a string argument:

```bash id="hz3m0n"
go run . "Your Text"
```

Using multiple lines:

```bash id="yk0z7p"
go run . "Welcome\nWorld"
```

---

## Project Goals

This project helps practice:

* File handling in Go
* String manipulation
* Working with command-line arguments
* ASCII character encoding
* Loops and indexing
* Writing clean, maintainable code

---

## Allowed Packages

Only Go standard library packages are used.

---

## Example

```bash id="u1f9pa"
go run . "ASCII Art"
```

Transforms plain text into a stylized terminal output.

---

