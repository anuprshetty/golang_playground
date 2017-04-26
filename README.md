# Golang Playground

Build simple, fast, reliable, and efficient software at scale.  
It's concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines.

## Why Golang?

- Low latency
- Garbage collection
- GPU power and it's full utilization
- Concurrency support --> it's not coming from any library. It's actually built into the langauge.
- Why new language? Golang is 21st century C programming language.

## Installation

### Additional tools which get installed when you install Go package

- MOD --> Go Package Manager
- ENV --> Lists environment variables
  - GoROOT --> Where all the binaries to run the program are stored in the Go root directory.
  - GoWorkSpace --> Folder where you put your code.
    - Go tells you to specifically place and organize your code in a very strict directory structure that you should follow.
    - This directory structure is expected to be followed if you are pushing up your project in production, or you are moving that on to a public repository where anybody can take advantage of your code.
    - A special directory structure:
      - project
        - bin --> binaries
        - pkg --> packages
        - src --> code
          - <repo_hosting_website>
            - <user_name>
              - <code_file>
- RUN -->
- GET --> Get some things from repository
- FMT --> Format package. This package is all about formatting input and output.

## Golang Commands

- go version
- go help
- go env
- go run main.go

## References

- [Golang official website](https://go.dev/)
