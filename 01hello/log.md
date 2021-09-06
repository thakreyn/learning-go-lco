# Folder 1: 01Hello

* 3rd September 2021
    
    1. Installed Golang and Set up all the Go tools available in the VScod extensions.
    2. Created a new Go file.
    3. Ran `go mod init hello`. The purpose of this command is still unknown. It creates a module of some sort.
    4. Go uses `fmt` for input and output.
    5. Libraries are imported automatically.
    6. Function decleration

        ```go
        func <name>() {
            . . . 
        }
        ```
    7. `go run <filename>` runs the file, but does not build the executables.
    8. The main essence of go is that the code should not have many interpretations. It should be easily readable. No 
    ambiguity and only one obvious way of doing things (like python).
    
    > Simplicity is Complex. It needs to have its own identity and not converge.

* 4th September 2021

    1. Details about the Go lexer.
    2. How Go lexer automatically imports the requried packages, ignores semicolons etc.
    3. Everything in Go is a type.
    4. For access modifiers, first letter is important. lowercase -> private and uppercase -> public