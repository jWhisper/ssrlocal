# ssrlocal
a linux command tool for xx; Inspired by https://github.com/chenset/shadowsocksR-go

# how to use it
1. git clone 
2. cd ssrlocal/cmd/ssr/ && go build .
3. create a config: ./ssr config -c path_to_conifg and then modify the config
4. start proxy: ./ssr start -c path_to_config
5. curl --socks5://127.0.0.1:1080 http://www.google.com for test

# TODO:
too many...
找到一个不错:https://github.com/nadoo/glider
