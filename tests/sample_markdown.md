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
> Hello too
> > Nested hehe


### Inline elements

This sentence contains a **Bold**, and *Italic* and a ***Bold and italic*** ~~strikethrough also~~.
Trying a **Bold** *Italic* and ~~**strikethrough with bold**~~ and **~~vice versa~~**.
Deep nested ***~~italic bold and strikthrough~~***
Some *`italic emphasis`* and `*vice versa*`. 
Multiple **nested ~~elements~~ `*mdelements*` in one *parent* ** element.

#### *Italic Unordered lists*

Image: 

![The San Juan Mountains are beautiful](https://user-images.githubusercontent.com/9877795/143689169-e3386847-46ad-4747-9934-2293f3d39abd.png)
Another ![Hyperlink](www.google.com) in line.

- item 1 with **bold**
- item 2

##### Heading 5

> ** kbjkbafe **

###### Heading 

A __underlined__ element and a nested __*underline*__

