# Markdown lexer syntax

## Markdown syntax

`# `- Heading 1
regex : `^# `

`## ` - Heading 2
regex : `^## `

`### ` - Heading 3
regex : `^### `

`#### ` - Heading 4
regex : `^#### `

`##### ` - Heading 5
regex : `^##### `

`###### ` - Heading 6
regex : `^###### `


`` - paragraph
If nothing is found then it is a normal char part of a paragraph

`\n` line  break
regex: `\n`

`-----------`  - horizontal line
regex: `^-{3,}`

`>` - block quote
`>>` - nested quote
regex: `^>+ `

`-` -  bullet point
regex: `^- `


`1. ` - list sequence
regex: `^[0-9]+. `

` ``` ` code blocks start and end
regex: ` ^``` `


` ![The San Juan Mountains are beautiful!](/assets/images/san-juan-mountains.jpg "San Juan Mountains") ` - image
regex: `^!\[.+\]\(.+\)`


### Inline 

`**text**` - bold
regex: `\*{2,2}`

`*text*` - italic
regex: `\*{1,1}`

`***text***` - bold and italic
regex: `\*{3,3}`

` `` ` - inline code
regex: ` ` `


` [Duck Duck Go](https://duckduckgo.com). ` hyper links
regex: ` \[.+\]\(.+\)`


`\` escape character
regex: `\\`

` ~~ ` strike through
regex: `~~{2,2}`