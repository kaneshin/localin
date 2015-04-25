# Localin

## Installation

```
go get github.com/kaneshin/localin
```

## Support

- iOS (Plist)
- Android (XML)

## Usage

### Generate

Generate localized files.

```
localin gen foo.json bar.json --target ios --delimiter _ --base en_US --lang en_US
```

- `--target`: Choose supported OS.
    - ios
    - android
- `--delimiter`: Default `_`.
- `--base`:
- `--lang`:

### Output

Output localized files.

```
localin out foo.json bar.json --target ios --delimiter _ --base en_US --lang en_US
```

- `--target`: Choose supported OS.
    - ios
    - android
- `--delimiter`: Default `_`.
- `--base`:
- `--lang`:


### Lint

Verify json files as parsable.

```
localin lint foo.json bar.json
```

## Tutorial

Writing JSON file like below,

_foo.json_

```
{
  "com": {
    "example1": {
      "person": {
        "gender": {
          "en_US": "Gender",
          "ja_JP": "性別",
          "man": {
            "en_US": "Man",
            "ja_JP": "男"
          },
          "woman": {
            "en_US": "Woman",
            "ja_JP": "女"
          }
        }
      }
    }
  }
}
```

```
$ localin lint foo.json
$ localin gen foo.json
$ ls
android_en_US.xml	android_ja_JP.xml	foo.json	ios_en_US.strings	ios_ja_JP.strings
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
