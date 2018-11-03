This requires changes to vscode-go to work, see https://github.com/Microsoft/vscode-go/pull/1410

Demo project to show how launch debug in docker using custom debug command could work. Before testing to debug be sure there is a `vscode-go-debug-test` image with `dlv`
installed.

```sh
docker build -t vscode-go-debug-test .
```
