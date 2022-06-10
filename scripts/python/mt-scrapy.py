# -*- coding: UTF-8 -*-
import sys
import os

def help():
    print("Usage: {} [project name] [spider name] [url]".format(sys.argv[0]))

def param_check():
    if len(sys.argv) != 4:
        help()
        return False
    else:
        return True

if __name__ == "__main__":
    if not param_check():
        exit()
    else:
        print("Create scrapy project:  [{}] [{}] [{}]".format(sys.argv[1], sys.argv[2], sys.argv[3]))
        os.system("python -m scrapy startproject {}".format(sys.argv[1]))
        os.chdir(sys.argv[1])
        os.system("python -m scrapy genspider {} {}".format(sys.argv[2], sys.argv[3]))


