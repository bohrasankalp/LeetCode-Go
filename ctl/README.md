# LeetCode-Go ctl

## Configuration Method

1. In `.gitignore`, add a line `*.toml`
2. In the `LeetCode-Go` directory, add a text file `config.toml`.
3. Copy the following content into `config.toml`.
4. Change `test` in `config.toml` to your leetcode `username` and `password` respectively.
5. After logging in to leetcode, copy the website cookie (need to copy csrftoken and LEETCODE_SESSION) and
   replace `Cookie` in `config.toml`.

```toml
Username = "test"
Password = "test"
Cookie = "csrftoken=XXXXXXXXX; LEETCODE_SESSION=YYYYYYYY;"
CSRFtoken = "ZZZZZZZZ"
```

## PDF generation

Use the `leetcode-go pdf` command to generate the merged version pdf.md of the book content, and then use vscode or
other tools that support toc directory generation to generate toc. last use
Typora converts md files to pdf. You can release a new version.