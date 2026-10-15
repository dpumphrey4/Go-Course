# Installation Guide

Go here to download Go: https://go.dev/doc/install. Just follow the instructions given for your target OS.

Go here to download VSCode https://code.visualstudio.com/download. On Windows and Mac, download and run the installer. For Debian Linux, run ```sudo dpkg -i ./path_to.deb```.

## Linux Specifics
Add Golang permanently to your user PATH variable by editing a file with ```nano ~/.profile```. Go to the bottom and put ```export PATH=$PATH:/usr/local/go/bin``` at the bottom of the file. Then save and close the file. (You don't have to use nano, use whatever you want to edit your profile). Restarting your Linux machine should now include the path to the Golang binaries on all terminal sessions.


## Final Steps
Open VSCode and try to run ```go version``` in the terminal, if it doesn't work then verify that Go is installed and included in the path.

Once VSCode is downloaded, go download the Go extension "Go".

Now you have Go.