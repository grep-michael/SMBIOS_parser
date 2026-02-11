#!/bin/bash

echo $1

case $1 in
2.7)
./parser_ai.py ./data/DSP0134_2.7.1.md > ./data/2_7.jsonc && ./code_gen.py ./data/2_7.jsonc
;;
3.3)
./parser_ai.py ./data/DSP0134_3.3.0.md > ./data/3_3.jsonc && ./code_gen.py ./data/3_3.jsonc
;;
esac