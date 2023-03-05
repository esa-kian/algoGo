#!/bin/bash

for i in {0..7..1}
do
  i=$((i + 6456595734))

  touch algorithm_${i}.go
  git add .
  git commit -m "Start task No #$i"
done
