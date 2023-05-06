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