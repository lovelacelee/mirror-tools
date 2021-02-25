# -*- coding:UTF-8 -*-

import os
import sys
import importlib
from pathlib import Path
import argparse

pipresource = "https://pypi.tuna.tsinghua.edu.cn/simple"

def install(argslist):
    if not len(argslist):
        os.system("python -m pip install --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip install {} -i {}".format(args, pipresource)
    os.system(pipcmd)
    
def uninstall(argslist):
    if not len(argslist):
        os.system("python -m pip uninstall --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip uninstall {} ".format(args)
    os.system(pipcmd)

def download(argslist):
    if not len(argslist):
        os.system("python -m pip download --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip download {} -i {}".format(args, pipresource)
    os.system(pipcmd)
def list(argslist):
    if not len(argslist):
        os.system("python -m pip list --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip list {} -i {}".format(args, pipresource)
    os.system(pipcmd)
def check(argslist):
    if not len(argslist):
        os.system("python -m pip check --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip check {} ".format(args)
    os.system(pipcmd)
def show(argslist):
    if not len(argslist):
        os.system("python -m pip show --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip show {} ".format(args)
    os.system(pipcmd)
def search(argslist):
    if not len(argslist):
        os.system("python -m pip search --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip search {} -i {}".format(args, pipresource)
    os.system(pipcmd)
def cache(argslist):
    if not len(argslist):
        os.system("python -m pip cache --help")
        exit()
    args = " ".join(argslist)
    pipcmd = "python -m pip cache {} ".format(args)
    os.system(pipcmd)
    
def upgradepip(argslist):
    pipcmd = "python -m pip install --upgrade pip"
    os.system(pipcmd)

def sub_cmd_parser(subparsers, cmd, callback):
    parser = subparsers.add_parser(cmd)
    parser.set_defaults(func=callback)

def mt_support_list(cmd, help):
    print("  {:<20s} {}".format(cmd, help))
    
mt_support = [
    ('install', 'pip install', install),
    ('uninstall', 'pip uninstall', uninstall),
    ('download', 'pip download', download),
    ('list', 'pip list', list),
    ('show', 'pip show', show),
    ('check', 'pip check', check),
    ('search', 'pip search', search),
    ('cache', 'pip cache', cache),
    ('upgradepip', 'pip install --upgrade pip', upgradepip)
]
    
def pargs_handle():
    global pipresource
    parser = argparse.ArgumentParser()
    parser.add_argument("-i", type=str, default=pipresource, help="Base URL of the Python Package Index(default: https://pypi.tuna.tsinghua.edu.cn/simple)")

    subparsers = parser.add_subparsers(title="Commands", description="Redefined pip command tools.", metavar="Support command list:")
    
    for i in mt_support:
        sub_cmd_parser(subparsers, i[0], i[2])
    
    # 解析到不认识的参数时会产生错误
    #args = parser.parse_args()
    # 它的作用方式很类似 parse_args() 但区别在于当存在额外参数时它不会产生错误。 
    # 而是会返回一个由两个条目构成的元组，其中包含带成员的命名空间和剩余参数字符串的列表。
    args = parser.parse_known_args()

    try:
        # 手动指定的pip源
        pipresource = args[0].i
        args[0].func(args[1])
    except Exception as e:
        print(e)
        exit()
        
if __name__ == "__main__":

    if len(sys.argv) == 1:
        os.system("{} -h".format(sys.argv[0]))

        for i in mt_support:
            mt_support_list(i[0], i[1])
    else:
        pargs_handle()