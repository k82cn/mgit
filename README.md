# mgit

The command lines to manage multiple repositories of a solution.

## Installation

```shell
$ go install -a github.com/k82cn/mgit@latest
```

## Configuration

The default configuration file is `$HOME/.mgit`, the following configuration demonstrate multiple solutions with multiple components.

```shell
current-solution: openbce
workspace: /home/k82cn/workspace
solutions:
  - name: ray
    git_server: "git@github.com:"
    user: k82cn
    components:
      - name: ray
        git_path: ray-project/ray
        module_path: ray.io/ray
  - name: openbce
    git_server: "git@github.com:"
    user: k82cn
    components:
      - name: device-manager
        git_path: openbce/device-manager
        module_path: openbce.io/device-manager
      - name: flame
        git_path: openbce/flame
        module_path: openbce.io/flame
      - name: kcache
        git_path: openbce/kcache
        module_path: openbce.io/kcache
```

## Examples

```shell
$ mgit download
Start to download device-manager: Done.
Start to download flame: Done.
Start to download kcache: Done.
```


```shell
$ mgit update
Start to update device-manager: Done.
Start to update flame: Done.
Start to update kcache: Done.
```

```shell
$ mgit list
  Name           User           GitServer                          Components#         
  ray            k82cn          git@github.com:                    1                   
* openbce        k82cn          git@github.com:                    3                                  
```
