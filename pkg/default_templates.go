package pkg

const defaultTemplate = `
*created: {{.FormattedCurrentTime}}*

`

const logTemplate = `
### Development Log
*created: {{.FormattedCurrentTime}}*


##### How did your development session go?



##### Did you learn anything new? If so, what did you learn?



##### What could have gone better?



##### What went well?


---

##### Notes

`

//TODO: get input on how many TD and use a loop to generate the TODOs
const tdTemplate = `
### TODO
*created: {{.FormattedCurrentTime}}*

- [ ]
- [ ]

`
