# mirror-tools

python mirror tools, 以清华源为pip源https://pypi.tuna.tsinghua.edu.cn/simple



## Windows下使用python虚拟环境

执行pythonvenv.bat后

将mkvirtualenv.bat所在的目录加入Windows的系统PATH目录中

之后workon命令即可使用



## 环境恢复

1. pip导出所有依赖名及版本号到txt文件中

```
pip freeze > requirements.txt
```

2. 将对应依赖导出到指定文件夹中

```
pip download -d ./py_packages -r requirements.txt
```

3. 在无网络环境下导入依赖

```
pip install --no-index --find-links=py_packages/ -r requiments.txt
```