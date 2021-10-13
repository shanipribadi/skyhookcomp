Run 

```
go test
```

Output with Skyhook
```
$ go test .
2021/10/13 20:03:09 DEBUG: writing: "\x83\xa2ID\xabrecord:id:0\xa4Name\xb0name:record:id:0\xa5Value\xd9Jvalue:record:id:0 ts:2021-10-13 20:03:09.91467677 +0700 WIB m=+0.006259494" | [131 162 73 68 171 114 101 99 111 114 100 58 105 100 58 48 164 78 97 109 101 176 110 97 109 101 58 114 101 99 111 114 100 58 105 100 58 48 165 86 97 108 117 101 217 74 118 97 108 117 101 58 114 101 99 111 114 100 58 105 100 58 48 32 116 115 58 50 48 50 49 45 49 48 45 49 51 32 50 48 58 48 51 58 48 57 46 57 49 52 54 55 54 55 55 32 43 48 55 48 48 32 87 73 66 32 109 61 43 48 46 48 48 54 50 53 57 52 57 52]
2021/10/13 20:03:09 ERROR: reading: "��ID�record:id:0�Name�name:record:id:0�Value�Jvalue:record:id:0 ts:2021-10-13 20:03:09.91467677 +0700 WIB m=+0.006259494" | [239 191 189 239 191 189 73 68 239 191 189 114 101 99 111 114 100 58 105 100 58 48 239 191 189 78 97 109 101 239 191 189 110 97 109 101 58 114 101 99 111 114 100 58 105 100 58 48 239 191 189 86 97 108 117 101 239 191 189 74 118 97 108 117 101 58 114 101 99 111 114 100 58 105 100 58 48 32 116 115 58 50 48 50 49 45 49 48 45 49 51 32 50 48 58 48 51 58 48 57 46 57 49 52 54 55 54 55 55 32 43 48 55 48 48 32 87 73 66 32 109 61 43 48 46 48 48 54 50 53 57 52 57 52] | error: msgpack: unexpected code=ef decoding map length
2021/10/13 20:03:09 DEBUG: writing: "\x83\xa2ID\xabrecord:id:0\xa4Name\xb0name:record:id:0\xa5Value\xd9Jvalue:record:id:0 ts:2021-10-13 20:03:09.91850295 +0700 WIB m=+0.010085672" | [131 162 73 68 171 114 101 99 111 114 100 58 105 100 58 48 164 78 97 109 101 176 110 97 109 101 58 114 101 99 111 114 100 58 105 100 58 48 165 86 97 108 117 101 217 74 118 97 108 117 101 58 114 101 99 111 114 100 58 105 100 58 48 32 116 115 58 50 48 50 49 45 49 48 45 49 51 32 50 48 58 48 51 58 48 57 46 57 49 56 53 48 50 57 53 32 43 48 55 48 48 32 87 73 66 32 109 61 43 48 46 48 49 48 48 56 53 54 55 50]
--- FAIL: TestOnce (0.01s)
    main_test.go:24: v1: &skyhookcomp.Record{ID:"record:id:0", Name:"name:record:id:0", Value:"value:record:id:0 ts:2021-10-13 20:03:09.91467677 +0700 WIB m=+0.006259494"}, v2: &skyhookcomp.Record{ID:"record:id:0", Name:"name:record:id:0", Value:"value:record:id:0 ts:2021-10-13 20:03:09.91850295 +0700 WIB m=+0.010085672"}
FAIL
FAIL    skyhookcomp     0.013s
FAIL
```

```
Data being written
\x83  \xa2  I    D    \xab
[131] [162] [73] [68] [171]

Data read:
[239 181 189] [239 191 189] [73] [68] [239 191 189]
```