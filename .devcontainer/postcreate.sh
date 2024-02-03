#!/bin/bash

#source any custom dotfiles
[ -d $HOME/.mydevcontainer/dotfiles ] && cp -r $HOME/.mydevcontainer/dotfiles/. $HOME || echo No custom dotfiles

#exec any init scripts for the devcontainer
[ -f $HOME/.mydevcontainer/init.sh ] && $HOME/.mydevcontainer/init.sh || echo 'No init'
