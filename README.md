# cloudflare-ddns

A simple DDNS service interacted with [Cloudflare](https://cloudflare.com) API.

## Features

1. Written in Golang and quite easy/small
2. Use embedded timer (let you get rid of crontab)
3. Designed for systemd
4. Support IPv4/IPv6

## How to use

Download latest release assets and set executable permission.

```
chmod +x cloudflare-ddns
```

Edit systemd configuration file (copy from ddns-systemd-template.service)

```
cp ddns.service /etc/systemd/system
systemctl daemon-reload
systemctl enable ddns
systemctl start ddns
```

# License

See the [LICENSE](https://github.com/i3h/cloudflare-ddns/blob/master/LICENSE.md) file for license rights and limitations (MIT).
