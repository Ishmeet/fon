# fon
json flattener 

## Usage
```
ishmeets@mbp % cat test.json | ./fon
age: 30.000000
address.city: New York
address.state: NY
address.zip: 10001.000000
phone[0].type: home
phone[0].number: 555-1234
phone[1].type: work
phone[1].number: 555-5678
name: John
```

```
ishmeets@mbp % ./fon test.json 
age: 30.000000
address.city: New York
address.state: NY
address.zip: 10001.000000
phone[0].type: home
phone[0].number: 555-1234
phone[1].number: 555-5678
phone[1].type: work
name: John
```

## Installation
fon has no runtime dependencies. You can just [download a binary for Linux, Windows or Freebsd and run it](https://github.com/Ishmeet/fon/releases).
Put the binary in the `$PATH` (e.g. in `/usr/local/bin`) to make it easy to use:
```
▶ tar xzf fon-linux-amd64-v0.0.tgz
▶ sudo mv fon /usr/local/bin/
```

Or if you are a go user, you can install using
```
go install github.com/Ishmeet/fon@latest
```