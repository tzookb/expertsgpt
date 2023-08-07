# Experts GPT

This repo is just in its infancy. It's a playground for me to learn about GPT-3 and OpenAI's API.
The plan is to create "experts" that work together as a team to solve problems.
For example a software development team that can plan a project:

- a product manager that can write user stories and spec the project
- a software architect that can design the system
- a program manager that can plan the project

Other options are possible:

- online marketing plan
- SEO plan
- research project
- etc

It's all built around "roles". We define the exact role for that persona and what he needs to do and output.
We can extract the output and forward it to other experts.

The actual outputs should be in different formats: csv, markdown, pdf, etc.

```bash
go run cmd/maind.go \
 --openaiToken=TOKEN
```