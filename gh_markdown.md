Learning GitHub Markdown
========================

So this file is being use to figure out the markdown syntax.

- [x] @mentions, #refs, [links](), **formatting**, and <del>tags</del> are supported 
- [x] list syntax is required (any unordered or ordered list supported) 
- [x] this is a complete item 
- [ ] this is an incomplete item

Reference

https://help.github.com/articles/writing-on-github


Hey this emoji cheat sheet is really cool: http://www.emoji-cheat-sheet.com/

~~Mistaken text. Strike-through~~

Code block

```ruby
require 'redcarpet'
markdown = Redcarpet.new("Hello World!")
puts markdown.to_html
```

Tables

First Header  | Second Header
------------- | -------------
Content Cell  | Content Cell
Content Cell  | Content Cell



| First Header  | Second Header |
| ------------- | ------------- |
| Content Cell  | Content Cell  |
| Content Cell  | Content Cell  |



| Name | Description          |
| ------------- | ----------- |
| Help      | Display the help window.|
| Close     | Closes a window     |


| Left-Aligned  | Center Aligned  | Right Aligned |
| :------------ |:---------------:| -----:|
| col 3 is      | some wordy text | $1600 |
| col 2 is      | centered        |   $12 |
| zebra stripes | are neat        |    $1 |
