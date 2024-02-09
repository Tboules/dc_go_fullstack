# Golang Fullstack Desert Collections

The plan is to build this application with 3 stacks, 1 in Go, 1 in Next(Server Actions), 1 in Sveltkit(maybe lol). I want to take the app from start to finish in each.
The issue has been that every time I start this project, I only get part way and so this three fold build is important to me.

> ## The Goal
>
> Overall it is important that I learn what backend / fullstack infrastructure I actually prefer.
> The only way to actually do that is to engage in that process from start to finish multiple times with the same project.

<br/>

---

### Technologies

1. Golang
2. [templ](https://templ.guide/)
3. [HTMX](https://htmx.org/)
4. [tailwindcss](https://tailwindcss.com/docs/installation)
5. [Echo](https://echo.labstack.com/docs)
6. [Goth](https://github.com/markbates/goth)
7. [PlanetScale](https://planetscale.com/docs/tutorials/planetscale-quick-start-guide)
8. [MySql](https://dev.mysql.com/doc/refman/8.0/en/programs.html)

### Steps

- [x] Create live server updates during coding
- [x] configure ~~templ~~ htmx css
- [x] configure goth
  - [x] sign in user
  - [x] create sign in page
  - [x] create jwt token and set http only cookie
  - [x] create middleware for error handling, protected route, isAuth
  - [x] create refresh
  - [ ] session logic in db
  - [ ] logout (clear cookies, clear db session)
- [ ] configure db
- [ ] Initialize db schema and migrations
- [ ] create auth page and persist user data
- [ ] create front end forms for CRUD quotes and authors
