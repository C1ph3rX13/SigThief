# SigThief
SigThief by Go

## Description

[variant](https://github.com/C1ph3rX13/variant) 项目衍生的工具

## Usage

### Update v2

#### Help

```powershell
SigThief.exe -h

        ██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
        ▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
        ░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
        ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
        ▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
        ▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
        ░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
        ░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
                ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

        @Auth: C1ph3rX13
        @Blog: https://c1ph3rx13.github.io
        @Note：SigThief - Go
        @Warn: The code is for learning purposes only, please do not use it for other purposes

Usage SigThief.exe:
  -act string
        操作类型 (save, exe, sign)
  -cert string
        导出的证书文件
  -dir string
        指定目录
  -dst string
        已签名的目标文件
  -out string
        窃取签名后的输出文件
  -src string
        未签名的源文件
```

#### Save (*.exe) Cert

```powershell
SigThief.exe -act save -dst code.exe -cert code.der

        ██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
        ▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
        ░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
        ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
        ▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
        ▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
        ░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
        ░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
                ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

        @Auth: C1ph3rX13
        @Blog: https://c1ph3rx13.github.io
        @Note：SigThief - Go
        @Warn: The code is for learning purposes only, please do not use it for other purposes

2024/10/18 13:57:30 INFO saved cert to code.der
```

#### Cert Signed

```powershell
SigThief.exe -act sign -src o.exe -cert code.der -out signed.exe

        ██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
        ▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
        ░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
        ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
        ▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
        ▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
        ░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
        ░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
                ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

        @Auth: C1ph3rX13
        @Blog: https://c1ph3rx13.github.io
        @Note：SigThief - Go
        @Warn: The code is for learning purposes only, please do not use it for other purposes

2024/10/18 13:58:53 INFO saved signed to signed.exe
```

#### Exe Signed

```powershell
SigThief.exe -act exe -src o.exe -dst code.exe -out signed.exe

        ██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
        ▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
        ░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
        ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
        ▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
        ▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
        ░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
        ░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
                ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

        @Auth: C1ph3rX13
        @Blog: https://c1ph3rx13.github.io
        @Note：SigThief - Go
        @Warn: The code is for learning purposes only, please do not use it for other purposes

2024/10/18 14:00:30 INFO saved signed to signed.exe
```

### Update v1

#### Help

```powershell
SigThief.exe -h

██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
  ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
      ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

@Auth: C1ph3rX13
@Blog: https://c1ph3rx13.github.io
@Note：SigThief - Go
@Warn: 代码仅供学习使用，请勿用于其他用途

Usage of SigThief.exe:
  -c string // 证书文件: -c 证书文件
        Cert File
  -f string // 带有签名的文件: -f 带有签名的文件
        Signed Exe
  -o string // 窃取签名并成果签名后输出的新文件(必填): -o 自定义命名
        Output Signed File / Cert
  -s string // 保存目标文件的证书: -s 目标文件
        Save Cert
  -t string // 未签名的文件: -t 未签名的文件
        Target (Unsign) Exe
```

#### Save (*.exe) Cert

```powershell
SigThief.exe -s vs.exe -o vs.der

██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
  ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
      ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

@Auth: C1ph3rX13
@Blog: https://c1ph3rx13.github.io
@Note：SigThief - Go
@Warn: 代码仅供学习使用，请勿用于其他用途

[*] Saved: vs.der
```

#### Cert Signed

```powershell
SigThief.exe -t a.exe -c vs.der -o signed.exe

██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
  ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
      ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

@Auth: C1ph3rX13
@Blog: https://c1ph3rx13.github.io
@Note：SigThief - Go
@Warn: 代码仅供学习使用，请勿用于其他用途

[*] Cert Signed: signed.exe
```

#### Exe Signed

```powershell
SigThief.exe -t a.exe -f vs.exe -o signed.exe

██████  ██▓  ▄████  ███▄    █ ▄▄▄█████▓ ██░ ██  ██▓▓█████   █████▒
▒██    ▒ ▓██▒ ██▒ ▀█▒ ██ ▀█   █ ▓  ██▒ ▓▒▓██░ ██▒▓██▒▓█   ▀ ▓██   ▒
░ ▓██▄   ▒██▒▒██░▄▄▄░▓██  ▀█ ██▒▒ ▓██░ ▒░▒██▀▀██░▒██▒▒███   ▒████ ░
  ▒   ██▒░██░░▓█  ██▓▓██▒  ▐▌██▒░ ▓██▓ ░ ░▓█ ░██ ░██░▒▓█  ▄ ░▓█▒  ░
▒██████▒▒░██░░▒▓███▀▒▒██░   ▓██░  ▒██▒ ░ ░▓█▒░██▓░██░░▒████▒░▒█░
▒ ▒▓▒ ▒ ░░▓   ░▒   ▒ ░ ▒░   ▒ ▒   ▒ ░░    ▒ ░░▒░▒░▓  ░░ ▒░ ░ ▒ ░
░ ░▒  ░ ░ ▒ ░  ░   ░ ░ ░░   ░ ▒░    ░     ▒ ░▒░ ░ ▒ ░ ░ ░  ░ ░
░  ░  ░   ▒ ░░ ░   ░    ░   ░ ░   ░       ░  ░░ ░ ▒ ░   ░    ░ ░
      ░   ░        ░          ░           ░  ░  ░ ░     ░  ░

@Auth: C1ph3rX13
@Blog: https://c1ph3rx13.github.io
@Note：SigThief - Go
@Warn: 代码仅供学习使用，请勿用于其他用途

[*] SigThief: signed.exe
```

## Thanks

https://github.com/TideSec/GoBypassAV

https://github.com/Tylous/Mangle
