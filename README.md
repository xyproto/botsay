# botsay [![Build Status](https://travis-ci.org/xyproto/botsay.svg?branch=master)](https://travis-ci.org/xyproto/botsay) [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/botsay/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/botsay)](https://goreportcard.com/report/github.com/xyproto/botsay)

Like cowsay, but with ASCII robots generated with [go-asciibot](https://github.com/mattes/go-asciibot).

### Examples

Output of `botsay -- $(fortune)`:

```
       .-.           .------------------------------------------.
    ._(u u)_.        | Nothing increases your golf score like   |
      (_n_)       --<| witnesses.                               |
    .=[::+]=.        '------------------------------------------'
  ]=' [___] '=[
      /| |\
     (0) (0)
```

Output of `botsay -- $(botsay;botsay;botsay)`:

```
     _._._          .--------------------------------------------.
    -)O O(-         | )_( .--. |- -| '--' |_^_| o==|ooo|==o      |
     \_0_/       --<| |___| // \\ _\\ //_ .===./` .--. /.- -.    |
  .==|>o<|==:=L     | \ '--' "\_n_/" ,"|+ |". _\|+__|/_ ( )      |
  '=c|___|          | __) (__ _._._ .--. -)ooo(- '--' \_=_/      |
     ]| |[          | (m9\:::/\ /___\6 . /___\ . . ..:::::::. .  |
    [_| |_]         '--------------------------------------------'
```

Output of `date --iso-8601 | botsay -`:

```
      .---.          .-------------.
     } - - {         | 2019-06-03  |
      \_O_/          '-------------'
  .-._/___\_.-.
  ;   \___/   ;
      |_|_|
      /_|_\
```

Output of `botsay 'Build complete'`:

```
     .===.          .----------------.
    //o o\\         | Build complete |
    \\_v_//         '----------------'
  .==|>o<|==:=L
  '=c|___|
    .'._.'.
    |_| |_|
```

### Color support

Use `-c` if you wish to colorize the output text with a rainbow-like effect:

![](img/botsay.png)

### Installation

    go get -u github.com/xyproto/botsay

### Flags

* `-c` - color the output with rainbow-like colors
* `--version` - output the current version
* `--help` - output brief help

## General info

* Version: 1.2.1
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
