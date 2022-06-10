# -*- coding:UTF-8 -*-

import os
import sys
import importlib
from pathlib import Path
import argparse

def tools_check():
    try:
        importlib.util.find_spec("virtualenv")
        print("系统已经安装virtualenv相关模块")
        return True
    except ImportError:
        os.system("python -m pip install virtualenv -i https://pypi.tuna.tsinghua.edu.cn/simple")
        print(sys.platform)
        os.system("python -m pip install virtualenvwrapper-win -i https://pypi.tuna.tsinghua.edu.cn/simple")
        return False
        
def venv_check():
    venvdir = Path(os.getcwd()+"/venv")
    if venvdir.exists():
        print("venv 已经存在，进入虚拟环境")
        if sys.platform == "win32":
            os.system("start "+os.getcwd()+"/venv/Scripts/"+"activate.bat")
        else:
            os.system(os.getcwd()+"/venv/Scripts/"+"activate.sh")

    else:
        print("新建python虚拟环境")
        os.system("python -m virtualenv venv")
        
def pargs_handle():
    parser = argparse.ArgumentParser()
    # parser.add_argument("-s", "--save", type=str, choices=["requirements.txt", "env.txt"], default="requirements.txt" help="Virtual env save to requirements.txt")
    # required=True/False
    # action="store_true,store_false,store_const,append,count"
    parser.add_argument("-s", "--save", type=str, help="Virtual env save to requirements.txt")
    parser.add_argument("-l", "--load", type=str, help="Virtual env reload from requirements.txt")
    args = parser.parse_args()
    print(args)
    if args.save:
        os.system("python -m pip freeze > "+args.save)
    if args.load:
        os.system("python -m pip install -r "+args.load+" -i https://pypi.tuna.tsinghua.edu.cn/simple")
        
if __name__ == "__main__":
    tools_check()
    pargs_handle()
    venv_check()