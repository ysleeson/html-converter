## Markdown to HTML converter

The HTML converter converts a Markdown `.md` file to `.html` file.

### Execution
```
cmd > html-converter filename.md
```

### Formatting Specifics
Markdown is a fairly rich specification; for this assignment, we’re only
looking for a small subset. This is the formatting we’d like you to implement:
| Markdown | HTML|
| -------------------------------------- | ------------------------------------------------- |
| `# Heading 1`   | `<h1>Heading 1</h1>`|
| `## Heading 2`  | `<h2>Heading 2</h2>`|
| `...` | `...`|
| `###### Heading 6` | `<h6>Heading 6</h6>`|
| `Unformatted text` | `<p>Unformatted text</p>`|
| `[Link text](https://www.example.com)` | `<a href="https://www.example.com">Linktext</a>` |
| `Blank line` | `Ignored`|

### Conversion Edge Cases
'#' for heading is placed in the beginning of a line, otherwise it is not considered as headings

If the number of # is more than 6 for headings, it will be considered as a paragraph

If parentheses for the link don't match, the perfectlly matched part will be considered as a link and the rest of them are considered as a paragraph.
for example `[[[Link text](https://www.example.com)))` -> `[[<a href="https://www.example.com">Linktext</a>))` 
