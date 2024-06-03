#!/bin/bash

file='.git/hooks/pre-push'

if [ ! -f $file ]
then
    echo 'Setting up pre-push hooks...'
    touch .git/hooks/pre-push
    echo "make lint" >> .git/hooks/pre-push
    echo "make unit-tests" >> .git/hooks/pre-push
else
    echo 'A pre-push hook already exists. Please ensure pre-push hooks have everything in place'
fi