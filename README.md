# negofetch
neofetch implementation written in Go

![A demo of the running negofetch](negofetch-macos.png "negofetch in action")

## About

This is a pure Go native implementation of neofetch, a known
tool to print information about the system with nice logo in
ascii art.

In the back there are some calls to shell in order to gather
data, but where it is possible, data is fetched via native
Go calls.

## Status

Right now a lot of stuff remains hard coded and sometimes only
for macOS, which is my main system right now.

I'm migrating slowly the features from neofetch to it, but I
don't believe 100% will be here.

| Object | status |
| --- | --- |
| Logo: macOS | done |
| Logo: AIX | done |
| Logo:  Hash | done |
| Logo:  alpine_small | done |
| Logo:  Amazon | done |
| Logo:  Arch | done |
| Option: ascii_art | done |
| Username @ hostname | done |
| OS | in progress |
| Host | done |
| Kernel | in progress |
| Uptime | in progress |
| Packages | in progress |
| Shell | in progress |
| Resolution | TBD |
| DE | TBD |
| WM | TBD |
| WM Theme | TBD |
| Terminal | TBD |
| Terminal Font | TBD |
| CPU | TBD |
| GPU | TBD |
| Memory | TBD |
| Color bar | TBD |
| Position information | Done|
| Color system | Done |
| --- | --- |

## ~~Warning~~

~~This project is built using AI coding assistant.  That means it might change
the license due copywrite issue. ~~

## AI conding assistant

I'd access to a coding assistant in my last job as experimental service.
Since I don't work there anymore, I'm not using it.  From the tag
"CodingAssistantFree" the coding assistant doesn't exist anymore.
