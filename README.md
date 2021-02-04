# GoCache

A Simple Cache Write By Golang.

# Future

- 添加更多缓存策略 TinyLFU，LFU，支持缓存过期
- 添加热点互备，加快热点数据查询
  - 思路1：每次从远程节点获取，随机一定概率将数据放到 hotcache。（groupcache 是这样做的，10%）
  - 思路2：添加一个 hotcache 结构，储存 qps 大于一定阈值的 key。

# Todos

[] LFU

[] TinyLFU

[] Expirable Cache

[] hotcache

# Usage
https://capotej.com/blog/2013/07/28/playing-with-groupcache/#how-to-use-it

# References
<https://capotej.com/blog/2013/07/28/playing-with-groupcache/#how-to-use-it>
<https://github.com/golang/groupcache>
<https://github.com/imlgw/gacache/>
