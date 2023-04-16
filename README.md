# site_status

> check if a site is offline or online

**WHY? why not**

Installation and Usage
---

```sh
git clone git@gthub.com:musaubrian/site_status

cd site_status

go build . -o status_go

./status_go -f <path_to_file_containing_sites>

---
help

./status_go -h

Usage of ./status_go:
  -f string
        The file location (full or relative path)
        Example: ~/Desktop/some_file.txt (default "./site.txt")
```

It does support both absolute paths or relative paths.

![demo gif](./examples/demo.gif)
**License** -> [MIT](./License)
