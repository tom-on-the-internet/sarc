#!/usr/bin/env bash

# This is the bash version of sarc. There is also a go version.

# Turns
# That's a good idea, Tom.
# Into
# tHaT's A gOoD iDeA, tOm

input="${*,,}"
character_index=0
should_uppercase="false"

while [ "$character_index" -lt "${#input}" ]
do
  char="${input:$character_index:1}"

  # only toggle (or not) the character if it's an alphabetical character
  if [[ $char =~ [[:alpha:]] ]]
  then
    if $should_uppercase
    then
      char="$(echo "$char" | tr '[:lower:]' '[:upper:]')"
      should_uppercase="false"
    else
      should_uppercase="true"
    fi
  fi

  output="${output}$char"
  character_index=$(( character_index + 1 ))
done

echo "$output" | xclip -sel clip &> /dev/null
echo "$output"
echo "(copied to system clipboard)"
