<!-- <p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/6wj0hh6.jpg" alt="Project logo"></a>
</p> -->

<h1 align="center">pocketbase-vue-starter</h1>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)
<!-- [![GitHub Issues](https://img.shields.io/github/issues/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/kylelobo/The-Documentation-Compendium.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls) -->


</div>

---

<p align="center">
      A starter template for using Vue/Vuetify with Pocketbase as a backend.
    <br> 
</p>

## 📝 Table of Contents <!-- omit in toc -->
- [🧐 About ](#-about-)
- [🏁 Getting Started ](#-getting-started-)
  - [Prerequisites ](#prerequisites-)
  - [Installing ](#installing-)
    - [Without Docker ](#without-docker-)
    - [With Docker ](#with-docker-)
- [🎈 Usage ](#-usage-)
  - [Migrations ](#migrations-)
- [⛏️ Built Using ](#️-built-using-)
- [✍️ Authors ](#️-authors-)
- [📝 TODO ](#-todo-)


## 🧐 About <a name = "about"></a>

This project is a starter template for using Vue/Vuetify with Pocketbase as a backend.

This template includes:

- Vue 3
- Vuetify 3
- Vue Router
- Pinia
- PocketBase

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

<!-- See [deployment](#deployment) for notes on how to deploy the project on a live system. -->

### Prerequisites <a name="prerequisites"></a>

- [Node 18.14.0+](https://nodejs.org/en/download/)
- [Go 1.20.1+](https://go.dev/dl/)
- [Docker (optional)](https://docker.com/get-started)

### Installing <a name="installing"></a>

```shell
git pull https://github.com/mrsessions/pocketbase-vue-starter
```

#### Without Docker <a name="without-docker"></a>

Starting in the root directory

```shell
cd pocket-base
```

Serve PocketBase on port 8090

```shell
go run . serve
```

Back in the root directoy

```shell
cd vue-client
```

```shell
npm install
```

```shell
npm run dev
```

#### With Docker <a name="docker"></a>

The easiest way is run use Docker Compose
```shell
docker-compose up
```

To build image from scratch

```shell
docker-compose up --build --force-recreate
```

## 🎈 Usage <a name="usage"></a>

### Migrations <a name="migrations"></a>

There is an initial migration file in the migrations folder. This file was intentionally left blank. You can use this file to create any initial scripts, or you can leave it until you have your first auto migrations run.

To create migrations, you can view the PocketBase documentation here: [https://pocketbase.io/docs/migrations](https://pocketbase.io/docs/migrations)

In the pocket-base directory, you can run the following command to create a new migration:

```shell
go run . migrate create <migration_name>
```

After createing migrations or updating the schema within the PocketBase Admin, to generate typescript definitions from your pocketbase.io schema, you can use the [pocketbase-typegen](https://github.com/patmood/pocketbase-typegen):

```shell
npx pocketbase-typegen --db .pocket-base/pb_data/data.db
```

This will generate a typescript file in the root directory called `pocketbase-types.ts`. This file will be used to generate the typescript definitions for the PocketBase schema to use in your code.


<!-- ## 🚀 Deployment <a name = "deployment"></a>

Add additional notes about how to deploy this on a live system. -->

## ⛏️ Built Using <a name = "built_using"></a>

- [Vue 3](https://vuejs.org/) - Web Framework
- [Vuetify 3](https://next.vuetifyjs.com/) - Material Design Component Framework
- [PocketBase](https://pocketbase.io/) - Server/Database Framework

## ✍️ Authors <a name = "authors"></a>

- [@mrsessions](https://github.com/mrsessions) - Initial work

<!-- See also the list of [contributors](https://github.com/kylelobo/The-Documentation-Compendium/contributors) who participated in this project. -->

<!-- ## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- Hat tip to anyone whose code was used
- Inspiration
- References -->

<!-- add a todo section -->
## 📝 TODO <a name = "todo"></a>

- [x] [Add Pocketbase Typegen](https://github.com/patmood/pocketbase-typegen) (Generate typescript definitions from your pocketbase.io schema.) documentation
- [x] Add section on migrations.
- [ ] Create initialize PB page to create first PocketBase Admin
- [ ] Clean up README.md
- [ ] Add default layout
