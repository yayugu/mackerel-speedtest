# mackerel-speedtest

## preparation
Install `speedtest` (the official one. not `speedtest-cli`)

```

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

Put toml file (default path: $HOME/.mackerel-speedtest.conf)

```
speedtest_path = "speedtest"
speedtest_server_id = 48463 # you can find your nearby server with `speedtest -L` command
apikey = "" # Please copy from /etc/mackerel-agent/*.conf or Mackerel web conxole
service_name = "speedtest"
```
