<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Contributors][contributors-shield]][contributors-url] -->
<div align="center">

[![Status][status-shield]][project-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

[![Build][build-shield]]()
[![Docker-Release][docker-release-shield]][release-url]
[![release-date][release-date-shield]][release-url]


</div>
<!-- [![LinkedIn][linkedin-shield]][linkedin-url] -->



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/MRSessions/pocketbase-vue-starter">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

  <!-- <h1 align="center"><s>PocketBase Vue Starter</s></h1> -->
  <h1 align="center">PocketBase Vue Template</h1>

  <p align="center">
    A starter template for using Vue/Vuetify with Pocketbase as a backend.
    <!-- <br /> -->
    <!-- <a href="https://github.com/MRSessions/pocketbase-vue-starter"><strong>Explore the docs »</strong></a> -->
    <br />
    <br />
    <!-- <a href="https://github.com/MRSessions/pocketbase-vue-starter">View Demo</a>
    · -->
    <a href="https://github.com/MRSessions/pocketbase-vue-starter/issues">Report Bug</a>
    ·
    <a href="https://github.com/MRSessions/pocketbase-vue-starter/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li>
          <a href="#installation-and-setup">Installation and Setup</a>
            <ul><a href="#docker-dev-environments">Docker Dev Environments</a></ul>
            <ul><a href="#local-non-docker">Local (Non-Docker)</a></ul>
            <ul><a href="#docker">Docker</a></ul>
        </li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

I started working with PocketBase and Vue/Vuetify on a couple of side projects and found that I was referencing one when creating a new project from scratch. After that, I decided that I wanted to try and create a starter template for anyone to be able to use.

Here's why:
* Your time should be focused on the core of your application, not having to create everything from scratch.
* Comes out of the box with a Vue Admin Setup page (but is still customizable through the `pocketbase.go` file)
* Comes with a default layout to get you up and started quickly

Of course, no one template will serve all projects since your needs may be different. I'll be adding more customizability in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue.

<!-- Thanks to all the people who have contributed to expanding this template! -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Docker][docker-shield]][docker-url]
* [![PocketBase][PocketBase.io]][Pocketbase-url]
* [![Vue][Vue.js]][Vue-url]
* [![Vuetify][Vuetify.js]][Vuetify-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Node 20.11.0+](https://nodejs.org/en/download/)
- [Go 1.21.6+](https://go.dev/dl/)
- [Docker (Recommended)](https://docker.com/get-started)

### Installation and Setup

1. Clone the repo
   ```sh
   git clone https://github.com/MRSessions/pocketbase-vue-starter.git
    ```

#### Docker Dev Environments

> Coming soon!

#### Local (Non-Docker)

> *Note: You can run the project separately with the default ports 8090 (PocketBase) and 3000 (Vue). The defult .env file is using `VITE_POCKETBASE_URL` to set the PocketBase URL. You can change this to point to a different PocketBase instance or if you change the port.*

1. In the pocketbase directory, run the following command to start the PocketBase server
    ```sh
    go run . serve #Runs PocketBase on default port 8090
    ```
2. In the vue-client directory, install NPM packages
    ```sh
    npm install
    ```
3. In the vue-client directory, run the following command to start the Vue server
    ```sh
    npm run dev #Runs Vue on default port 3000
    ```

#### Docker

- The easiest way to run is use Docker Compose
    ```sh
    docker-compose up --build
    ```
    or to recreate the container
    ```sh
    docker-compose up --build --force-recreate
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

### PocketBase

#### Defaults

I have setup PocketBase to remove(rewrite) the PocketBase default routes. By default, it is allowed. If you want to disable PocketBase routes, you can set the environment variable `POCKETBASE_DISABLE_UI` to `true`. This will keep users from accessing the PocketBase UI. Find details in below code sections.

<details>
  <summary>pocketbase.go</summary>

  ```go
  func main() {
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
      if getenvBool("POCKETBASE_DISABLE_UI") {
        e.Router.Pre(middleware.Rewrite(map[string]string{
          "/_":  "/",
          "/_*": "/",
        }))
        log.Default().Println("PocketBase UI is disabled")
      }
      e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS(publicDir), indexFallback))
      return nil
    })
  }

  func getenvBool(key string) bool {
    val := os.Getenv(key)
    ret, err := strconv.ParseBool(val)
    if err != nil {
      return false
    }
    return ret
  }
  ```

</details>

<details>
  <summary>docker-compose.yml</summary>

  ```yml
  version: "3"

  pocketbase-vue-starter:
    image: ghcr.io/mrsessions/pocketbase-vue-starter:latest
    container_name: pocketbase-vue-starter
    restart: unless-stopped
    environment:
      - POCKETBASE_DISABLE_UI=true # Set to true to disable the PocketBase UI
      - POCKETBASE_ADMIN_EMAIL=admin@example.com #This is the default if you don't set it or this value is removed
      - POCKETBASE_ADMIN_PASSWORD=1234567890 #This is the default if you don't set it or this value is removed
    volumes:
      - ./pocketbase-db:/app/pb_data
    ports:
      - 8090:80

  volumes:
    pocketbase-db:
  ```
</details>

<details>
  <summary>Dockerfile (Build Final Image Section)</summary>

  ```dockerfile
  # build final image
  FROM golang:1.21.6-alpine3.19 AS final

  WORKDIR /app

  COPY --from=builder /app/pocketbase ./

  COPY --from=node-builder /app/dist ./dist

  # Set to true to disable the PocketBase UI if not using Docker Compose
  ENV POCKETBASE_DISABLE_UI=false

  EXPOSE 8090

  CMD ["/app/pocketbase", "serve", "--http=0.0.0.0:80"]
  ```

</details>

#### Migrations

- PocketBase has a built-in migration system. You can create a migration file by running the following command:
    ```sh
    go run . migration create <migration-name>
    ```

- PocketBase will automatically migrate the database when the server starts. You can also run the migrations manually by running the following command:
    ```sh
    go run . migration up
    ```

- You can also rollback the migrations by running the following command:
    ```sh
    go run . migration down
    ```

- After creating migrations or updating the schema with PocketBase, if you want to generate typescript definitions, you can use the [pocketbase-typegen](https://github.com/patmood/pocketbase-typegen) commands:
    ```sh
    npx pocketbase-typegen --db .pocket-base/pb_data/data.db
    ```
> *Note: This will generate a typescript file in the rot directory called `pocketbase-types.ts`. This file will be used to generate the typescript definitions for the PocketBase schema to use in your code.*

### Vue

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap


- [x] Add section on migrations.
- [x] Add default layout
- [x] Create initialize PB page in Vue to create first PocketBase Admin
- [x] Clean up README.md
  - [x] [Add Pocketbase Typegen](https://github.com/patmood/pocketbase-typegen) (Generate typescript definitions from your pocketbase.io schema.) documentation
- [ ] Create a Docker Dev Environment
- [ ] Add a section for a quick how to use the PocketBase API in Vue (Refer to the [PocketBase API Docs](https://pocketbase.io/docs/api))

See the [open issues](https://github.com/MRSessions/pocketbase-vue-starter/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature (`git checkout -b feature/AmazingFeature`) or Bug Fix (`git checkout -b bug/AmazingBugFix`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature or some AmazingBugFix'`)
4. Push to the Branch (`git push origin feature/AmazingFeature` or `git push origin bug/AmazingBugFix`)
5. Open a Pull Request
   1. If your change includes quite a bit of change, please document the changes in detail in the pull request.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Matt Sessions - [@MRSessions](https://github.com/MRSessions) - pbvuestarter@sessionstech.com

Project Link: [https://github.com/MRSessions/pocketbase-vue-starter](https://github.com/MRSessions/pocketbase-vue-starter)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
<!-- ## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
* [Malven's Flexbox Cheatsheet](https://flexbox.malven.co/)
* [Malven's Grid Cheatsheet](https://grid.malven.co/)
* [Img Shields](https://shields.io)
* [GitHub Pages](https://pages.github.com)
* [Font Awesome](https://fontawesome.com)
* [React Icons](https://react-icons.github.io/react-icons/search)

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[project-url]: https://github.com/MRSessions/pocketbase-vue-starter
[status-shield]: https://img.shields.io/badge/status-active-success.svg?style=for-the-badge
[forks-shield]: https://img.shields.io/github/forks/MRSessions/pocketbase-vue-starter.svg?style=for-the-badge
[forks-url]: https://github.com/MRSessions/pocketbase-vue-starter/network/members
[stars-shield]: https://img.shields.io/github/stars/MRSessions/pocketbase-vue-starter.svg?style=for-the-badge
[stars-url]: https://github.com/MRSessions/pocketbase-vue-starter/stargazers
[issues-shield]: https://img.shields.io/github/issues/MRSessions/pocketbase-vue-starter.svg?style=for-the-badge
[issues-url]: https://github.com/MRSessions/pocketbase-vue-starter/issues
[license-shield]: https://img.shields.io/github/license/MRSessions/pocketbase-vue-starter.svg?style=for-the-badge
[license-url]: https://github.com/MRSessions/pocketbase-vue-starter/blob/master/LICENSE
[build-shield]: https://img.shields.io/github/actions/workflow/status/MRsessions/pocketbase-vue-starter/build-single-docker-image.yml?style=for-the-badge
[build-url]: https://github.com/MRSessions/pocketbase-vue-starter/actions
[prerelease-shield]: https://img.shields.io/github/v/release/MRSessions/pocketbase-vue-starter?color=s&include_prereleases&label=Pre-release&logo=s&logoColor=s&style=for-the-badge
[release-date-shield]: https://img.shields.io/github/release-date-pre/mrsessions/pocketbase-vue-starter?label=Released&style=for-the-badge
[docker-release-shield]: https://img.shields.io/github/v/tag/mrsessions/pocketbase-vue-starter?include_prereleases&label=docker&style=for-the-badge
[release-url]: https://github.com/MRSessions/pocketbase-vue-starter/pkgs/container/pocketbase-vue-starter


[Vue.js]: https://img.shields.io/badge/Vue.js-3.2.38+-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Vuetify.js]: https://img.shields.io/badge/Vuetify-3.0.0+-green?style=for-the-badge&logo=vuetify&logoColor=blue
[Vuetify-url]: https://next.vuetifyjs.com/en/
[PocketBase.io]: https://img.shields.io/badge/PocketBase-0.13.2+-b8dbe4?style=for-the-badge&logo=pocketbase&logoColor=b8dbe4
[Pocketbase-url]: https://pocketbase.io
[docker-shield]: https://img.shields.io/badge/Docker-latest-blue?style=for-the-badge&logo=docker&logoColor=blue
[docker-url]: https://docker.com


<!-- Created with the help of https://github.com/othneildrew/Best-README-Template/pull/73 -->
