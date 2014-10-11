# Tiny Json Pretty Formatter

# Install

```
go get github.com/ikawaha/pretty
```

# Usage

* Before
```
$ curl -XPUT "http://localhost:8080/"  -d '{"input":"すもももももももものうち"}'
{"status":true,"tokens":[{"id":-1,"start":0,"end":0,"surface":"BOS","class":0,"features":null},{"id":36163,"start":0,"end":3,"surface":"すもも","class":1,"features":["名詞","一般","*","*","*","*","すもも","スモモ","スモモ"]},{"id":73244,"start":3,"end":4,"surface":"も","class":1,"features":["助詞","係助詞","*","*","*","*","も","モ","モ"]},{"id":74989,"start":4,"end":6,"surface":"もも","class":1,"features":["名詞","一般","*","*","*","*","もも","モモ","モモ"]},{"id":73244,"start":6,"end":7,"surface":"も","class":1,"features":["助詞","係助詞","*","*","*","*","も","モ","モ"]},{"id":74989,"start":7,"end":9,"surface":"もも","class":1,"features":["名詞","一般","*","*","*","*","もも","モモ","モモ"]},{"id":55829,"start":9,"end":10,"surface":"の","class":1,"features":["助詞","連体化","*","*","*","*","の","ノ","ノ"]},{"id":8024,"start":10,"end":12,"surface":"うち","class":1,"features":["名詞","非自立","副詞可能","*","*","*","うち","ウチ","ウチ"]},{"id":-1,"start":12,"end":12,"surface":"EOS","class":0,"features":null}]}
```

* After
```
$ curl -XPUT "http://localhost:8080/"  -d '{"input":"すもももももももものうち"}'|pretty
{
    "status": true,
    "tokens": [
        {
            "id": -1,
            "start": 0,
            "end": 0,
            "surface": "BOS",
            "class": 0,
            "features": null
        },
        {
            "id": 36163,
            "start": 0,
            "end": 3,
            "surface": "すもも",
            "class": 1,
            "features": [
                "名詞",
                "一般",
                "*",
                "*",
                "*",
                "*",
                "すもも",
                "スモモ",
                "スモモ"
            ]
        },
        {
            "id": 73244,
            "start": 3,
            "end": 4,
            "surface": "も",
            "class": 1,
            "features": [
                "助詞",
                "係助詞",
                "*",
                "*",
                "*",
                "*",
                "も",
                "モ",
                "モ"
            ]
        },
        {
            "id": 74989,
            "start": 4,
            "end": 6,
            "surface": "もも",
            "class": 1,
            "features": [
                "名詞",
                "一般",
                "*",
                "*",
                "*",
                "*",
                "もも",
                "モモ",
                "モモ"
            ]
        },
        {
            "id": 73244,
            "start": 6,
            "end": 7,
            "surface": "も",
            "class": 1,
            "features": [
                "助詞",
                "係助詞",
                "*",
                "*",
                "*",
                "*",
                "も",
                "モ",
                "モ"
            ]
        },
        {
            "id": 74989,
            "start": 7,
            "end": 9,
            "surface": "もも",
            "class": 1,
            "features": [
                "名詞",
                "一般",
                "*",
                "*",
                "*",
                "*",
                "もも",
                "モモ",
                "モモ"
            ]
        },
        {
            "id": 55829,
            "start": 9,
            "end": 10,
            "surface": "の",
            "class": 1,
            "features": [
                "助詞",
                "連体化",
                "*",
                "*",
                "*",
                "*",
                "の",
                "ノ",
                "ノ"
            ]
        },
        {
            "id": 8024,
            "start": 10,
            "end": 12,
            "surface": "うち",
            "class": 1,
            "features": [
                "名詞",
                "非自立",
                "副詞可能",
                "*",
                "*",
                "*",
                "うち",
                "ウチ",
                "ウチ"
            ]
        },
        {
            "id": -1,
            "start": 12,
            "end": 12,
            "surface": "EOS",
            "class": 0,
            "features": null
        }
    ]
}

```
