# Sample markdown

This document contains sample markdown to test the markdown compiler implementation. 

There will be 2 corresponding files containing expected output of each compilation step:
1. `sample_markdown.tokens.json` for lexer output
2. `sample_markdown.elements.json` for parser output

The format of each output is given below.

## sample_markdown.tokens.json

Each token will be seperated by **<< % >>**.
First continous word in each seperated token will be the token name and rest remaining exlcuding the immediate space is the value

## sample_markdown.elements.json

```
[
    {
        DefName string
        V   string
        C   []*self
    }
]
```

-----

## Elements covered

[.] Headings
[ ] Checkboxes

> Note: This document is in progress

