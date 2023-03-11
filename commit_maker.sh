#!/bin/bash

for i in {0..32..1}
do
  i=$((i + 976947))

  touch algorithm_${i}.go
  git add .
  git commit -m "Start task No #$i"
done
