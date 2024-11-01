
# Devlog #5

## Idea

python 的使用經驗沒有很多，但是因為想要達成的功能算簡單，所以還是使用 python

這邊想要達成的功能是使用 python 定時去觸發後端提供的 API

## Survey & Impl


VS Code 環境安裝 [Python in Visual Studio Code](https://code.visualstudio.com/docs/languages/python)

會使用到的套件
- https://github.com/dbader/schedule
- https://github.com/psf/requests


```python
pip3 install -r requirements.txt
```

問題
```zsh
error: externally-managed-environment

× This environment is externally managed
╰─> To install Python packages system-wide, try brew install
    xyz, where xyz is the package you are trying to
    install.
```

[解決Linux pip install的"error: externally-managed-environment" 錯誤，改用虛擬環境安裝套件](https://ivonblog.com/posts/linux-solve-externally-managed-environment-error/)

```
python3 -m venv path/to/venv
```

```
source path/to/venv/bin/activate
```

```python
python3 main.py 
```

退出虛擬環境
```
deactivate
```


[Docker Hub: python](https://hub.docker.com/_/python)
範本
```dockerfile
FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

CMD [ "python", "./your-daemon-or-script.py" ]
```

## Summery



## References

