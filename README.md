# GTK3-Learning-Easy
My Easy Code Samples for learning GTK3 (golang and c++)

This repo provides samples to get started with GTK3 developement.

# Installation of GTK3 Library

## Windows

Compiling GTK3 in windows requires mingw and gtk libs. It's suggested to use [MSYS2](https://www.msys2.org/) to install all of these as manually installing it is a lot of work.

### Using Chocolatey

Refering to gotk3 wiki,

> If you are using [Chocolatey](https://chocolatey.org), the following is all you need to do to get started developing:
> PS C:\> choco install golang
> PS C:\> choco install git
> PS C:\> choco install msys2
> PS C:\> mingw64
> $ pacman -S mingw-w64-x86_64-gtk3 mingw-w64-x86_64-toolchain base-devel glib2-devel
> $ echo 'export PATH=/c/Go/bin:$PATH' >> ~/.bashrc
> $ echo 'export PATH=/c/Program\ Files/Git/bin:$PATH' >> ~/.bashrc
> $ source ~/.bashrc
> $ sed -i -e 's/-Wl,-luuid/-luuid/g' /mingw64/lib/pkgconfig/gdk-3.0.pc # This fixes a bug in pkgconfig
> $ go get github.com/gotk3/gotk3/gtk

The wiki suggests setting go path to .bashrc inside MSYS2 but its not required. We wont be using MSYS2 shell for compiling.
installing packages with pacman and running the pkgconfig fix should be sufficient.


### Manually with MSYS2 shell
* To install MSYS2 follow instructions from [MSYS2 installation page](https://www.msys2.org/wiki/MSYS2-installation/). It should install required mingw64 packages.

* After installation is complete run:
`pacman -S mingw-w64-x86_64-gtk3 mingw-w64-x86_64-toolchain base-devel glib2-devel`
type `y` whenever it asks to install packages.

I already had MSYS2 installed so i decided to manually install the GTK3 libs right away.

### Installing Glade (GUI Design Tool)

Glade is an excellent tool to desgin GUI with drag and drop. It's recommended for everyone :)

To install Glade in MSYS2:
* Open MSYS2 shell
* Type `pacman -S mingw-w64-x86_64-glade`
* Go to `mingw64` folder inside msys2, then inside `bin` folder you'll find glade.exe.
* Run Glade.exe

## Linux

To install GTK3 libraries in linux run the following command is terminal accordingly

### Debian / Ubuntu

`sudo apt install libgtk-3-dev libcairo2-dev libglib2.0-dev`

### Fedora

`sudo dnf install gtk3-devel cairo-devel glib-devel`

### Alpine

`sudo apk add gtk+3.0-dev cairo-dev glib-dev`

### Arch

Install GTK3 and you get the dev part free :P

## MacOS

To Install GTK3 library in MacOS you'll need homebrew. If homebrew is installed run the follow commands

`% brew install pkg-config gtk+3 adwaita-icon-theme`

`% go get github.com/gotk3/gotk3/gtk`


# Installing Libraries for Languages

## Golang

Golang has pretty nice GTK3 wrapper [gotk3](https://github.com/gotk3/gotk3)

To install gotk3 run:

`go get github.com/gotk3/gotk3`

## C++

C++ requires official wrapper gtkmm along with GTK3 libraries.

To install gtkmm follow instuctions on the [download page](https://www.gtkmm.org/en/download.html). For windows you need to run `pacman -S mingw-w64-x86_64-gtkmm` in MSYS2 shell.

After you have the required libraries you can now compile GTK3 applications.
