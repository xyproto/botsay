# botsay

Like cowsay, but with a small ASCII bot from [go-asciibot](https://github.com/mattes/go-asciibot).

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

### Installation

    go get -u github.com/xyproto/botsay

### Flags

* `-c` - color the output with rainbow-like colors
* `--version` - output the current version
* `--help` - output brief help

## General info

* License: MIT
* Version: 1.2.1
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
