# echoerr

    Print a message to STDOUT and exit with the specified code.
    
    USAGE:
            echoerr [exitcode] <echo text...>

### Examples

    $ echoerr 1 Pretty standard stuff
    Pretty standard stuff
    $ echo $?
    1

    $ echoerr 2 Something more interesting
    Something more interesting
    $ echo $?
    2

    $ echoerr 128 "Standard shell stuff" even vars: $PATH
    Standard shell stuff even vars: /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin
    $ echo $?
    128

    $ echoerr No exit code 
    No exit code
    $ echo $?
    1

### Installation

For gophers:

    go get github.com/buildertools/echoerr

For Darwin (OSX):

    curl -L -o ~/bin/echoerr https://github.com/buildertools/echoerr/releases/download/v1.0/echoerr-darwin64
    chmod a+x ~/bin/echoerr

### LICENSE

Released under [AL2](LICENSE) and copyright held by "Jeff Nickoloff."
