package pkg

const defaultTemplate = `{{ if .FormattedCurrentTime }}created: {{.FormattedCurrentTime}}{{ end }}
{{ if .Tags }}tags: {{.Tags}}{{ end }}

### Note

`

const logTemplate = `{{ if .FormattedCurrentTime }}created: {{.FormattedCurrentTime}}{{ end }}
{{ if .Tags }}tags: {{.Tags}}{{ end }}

### Development Log


##### How did your development session go?



##### Did you learn anything new? If so, what did you learn?



##### What could have gone better?



##### What went well?


---

##### Notes

`

//TODO: get input on how many TDs and use a loop to generate the TODOs
const tdTemplate = `{{ if .FormattedCurrentTime }}created: {{.FormattedCurrentTime}}{{ end }}
{{ if .Tags }}tags: {{.Tags}}{{ end }}

### TODO


- [ ]
- [ ]

`
