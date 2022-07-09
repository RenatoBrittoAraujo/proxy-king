# proxy-king

Self-host a reliable and open source proxy server

## Why?

You need proxies. You don't want to pay for it, nor do you want to spend time looking for reliable proxies on google and testing them out.

Boot up an instance of proxy-king and let it handle that for you, for free!

## Design

### V1

![Untitled Diagram drawio](https://user-images.githubusercontent.com/45462822/178123292-f1f7bcc5-4e22-47d9-808a-d3101ebefc4e.png)

## Caveats 

There was not much though put into security for now, only ease of use. Use with caution.

## Usage

1. Install `proxychains`.

```bash
sudo apt install proxychains
```

2. Get IP and Port of your hosted instance.

3. Edit `/etc/proxychains.conf`, at the very end of the file.


```bash
vim /etc/proxychains.conf
```

```conf
...
[ProxyList]
# add proxy here ...
# meanwile
# defaults set to "tor"
http {your server ip} {port you picked} # <== line that was changed
```

4. Use proxychains to route any TCP over the server.

```
$ proxychains curl google.com
$ proxychains firefox
$ proxychains nmap -sV -p- {some ip}
...
```

5. Enjoy the realiability :)

## How to boot up my own instance?

This server can be hosted anywhere, but if you'd like a quick and easy AWS EC2 t2.micro setup fast, I have included a bash script that instances all for you very conveniently. Check it out on `scripts` folder.

