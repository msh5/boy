<!-- markdownlint-disable MD033 -->
<!-- markdownlint-disable MD041 -->

<h1 align="center">
  Boy
</h1>

<div align="center">
  <strong>
    On-Cloud Snippet Runner.
  </strong>
  <br />
  <br />
  <img src="https://github.com/msh5/boy/workflows/Build/badge.svg"> <!-- CI badge for Build workflow -->
  <img src="https://github.com/msh5/boy/workflows/Release/badge.svg"> <!-- CI badge for Release workflow -->
</div>

## Usage

```console
$ boy exec gist.github.com/msh5/965b29269fee385ee2a082101f247a26
Hello World !
```

## Installation

We distribute pre-built binaries for major CPUs and OSs,
so it is good to download latest-version binary matched your system:

* [Releases · msh5/boy](https://github.com/msh5/boy/releases)

If you cannot find the binary matched with your system, please `go get` this repository:

```console
env GO111MODULE=off go get -u github.com/msh5/boy
```

<!--

```shell
# Install via Homebrew.
brew tap msh5/boy
brew install boy

# Test with version flag.
$ boy --version
v0.0.1

# Finally, register your GitHub credential.
boy config --add gist.gh_access_token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

-->
