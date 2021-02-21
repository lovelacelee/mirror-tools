#!/bin/sh

python -m pip3 install virtualenv -i https://pypi.tuna.tsinghua.edu.cn/simple
	
python -m virtualenv venv

cd venv/Scripts

./activate