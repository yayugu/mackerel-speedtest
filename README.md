# mackerel-speedtest

## preparation

```
$ wget https://install.speedtest.net/app/cli/ookla-speedtest-1.1.1-linux-x86_64.tgz
$ tar -xzvf ookla-speedtest-1.1.1-linux-x86_64.tgz
$ sudo mkdir -p /usr/local/src/ookla-speedtest/ookla-speedtest-1.1.1
$ sudo mv speedtest speedtest.5 speedtest.md /usr/local/src/ookla-speedtest/ookla-speedtest-1.1.1
$ sudo ln -s /usr/local/src/ookla-speedtest/ookla-speedtest-1.1.1/speedtest /usr/local/bin
$ exec $SHELL -l
$ speedtest --version
Speedtest by Ookla 1.1.1.28 (c732eb82cf) Linux/x86_64-linux-musl 5.10.102.1-microsoft-standard-WSL2 x86_64

The official command line client for testing the speed and performance
of your internet connection.
$ speedtest
==============================================================================

You may only use this Speedtest software and information generated
from it for personal, non-commercial use, through a command line
interface on a personal computer. Your use of this software is subject
to the End User License Agreement, Terms of Use and Privacy Policy at
these URLs:

    https://www.speedtest.net/about/eula
    https://www.speedtest.net/about/terms
    https://www.speedtest.net/about/privacy

==============================================================================

Do you accept the license? [type YES to accept]:
```
