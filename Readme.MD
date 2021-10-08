
<span style="display:block;text-align:center">

[![Made With][made-with-shield]][made-with-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

</span>

<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/kanitmann/Instagram_backend_API">
    <img src="" alt="Logo" width="280" height="280">
  </a>

  <h3 align="center">Instagram Backend API</h3>

  <p align="center">
    *To be Updated*
    <br />
    <a href="https://github.com/kanitmann/Instagram_backend_API"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/kanitmann/Instagram_backend_API">View Demo</a>
    ·
    <a href="https://github.com/kanitmann/Instagram_backend_API/issues">Report Bug</a>
    ·
    <a href="https://github.com/kanitmann/Instagram_backend_API/issues">Request Feature</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary><h2 style="display: inline-block">Table of Contents</h2></summary>
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgements">Meet the Team</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

<p>

The task is to develop a basic version of a Instagram. You are only required to ```develop the API``` for the system. Below are the details.

You are required to Design and Develop an HTTP JSON API capable of the following operations,
- Create an User
    - Should be a POST request
    - Use JSON request body
    - URL should be ``‘/users'``

- Get a user using id
    - Should be a GET request
    - Id should be in the url parameter
    - URL should be ``‘/users/<id here>’``

- Create a Post
    - Should be a POST request
    - Use JSON request body
    - URL should be ``‘/posts'``

- Get a post using id
    - Should be a GET request
    - Id should be in the url parameter
    - URL should be ``‘/posts/<id here>’``

- List all posts of a user
    - Should be a GET request
    - URL should be ``‘/posts/users/<Id here>'``

Additional Constraints/Requirements:

-   The API should be developed using Go.
-   MongoDB should be used for storage.
-   Only packages/libraries listed here and here can be used.

Scoring:
- Completion Percentage
    - No. of working endpoints among the ones listed above.
- Quality of Code
    - Reusability
    - Consistency in naming variables, methods, functions, types
    - Idiomatic i.e. in Go’s style
- Passwords should be securely stored such they can't be reverse engineered
- Make the server thread safe
- Add pagination to the list endpoint
- Add unit tests

</p>


### Built With

- [Go](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [HTML5 + CSS3]()

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

This list things you need, and how to install them.

- Go 1.17.1
- MongoDB 5.0.3

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/kanitmann/Instagram_backend_API.git
   ```

<!-- USAGE EXAMPLES 

## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

 ROADMAP -->
## Screenshots


<img src="./Asset/img2.jpeg" alt="SS2" width="280" height="420">

## Roadmap

See the [open issues](https://github.com/kanitmann/Instagram_backend_API/issues) for a list of proposed features (and known issues).

<!-- LICENSE -->

## License

Distributed under the GNU GPL3 License. See `LICENSE` for more information.

<!-- ACKNOWLEDGEMENTS -->

## Meet the Team

- [Kanit Mann](https://github.com/kanitmann) :wolf:

> Made with 💙 by Kanit Mann

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/kanitmann/Instagram_backend_API.svg?style=for-the-badge
[contributors-url]: https://github.com/kanitmann/Instagram_backend_API/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/kanitmann/Instagram_backend_API.svg?style=for-the-badge
[forks-url]: https://github.com/kanitmann/Instagram_backend_API/network/members
[stars-shield]: https://img.shields.io/github/stars/kanitmann/Instagram_backend_API.svg?style=for-the-badge
[stars-url]: https://github.com/kanitmann/Instagram_backend_API/stargazers
[issues-shield]: https://img.shields.io/github/issues/kanitmann/Instagram_backend_API.svg?style=for-the-badge
[issues-url]: https://github.com/kanitmann/Instagram_backend_API/issues
[license-shield]: https://img.shields.io/github/license/kanitmann/Instagram_backend_API?style=for-the-badge&logo=appveyor
[license-url]: https://github.com/kanitmann/Instagram_backend_API/blob/master/LICENSE
[made-with-shield]: https://img.shields.io/github/languages/top/kanitmann/Instagram_backend_API?style=for-the-badge
[made-with-url]: https://shields.io/github/languages/top/kanitmann/Instagram_backend_API.svg?style-for-the-badge
[product-screenshot]: (images/screenshot.png)
