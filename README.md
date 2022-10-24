<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/romodeus/open-funding">
    <img src="" alt="Logo">
  </a>

  <h3 align="center">Open Funding</h3>

  <p align="center">
    Open Funding for community
    <br />
    <br />
    <a href="https://github.com/romodeus/open-funding/issues">Report Bug</a>
    ·
    <a href="https://github.com/romodeus/open-funding/issues">Request Feature</a>
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
      <ul>
       <li><a href="#feature">Feature</a></li>
      </ul>
      <ul>
       <li><a href="#sequence-diagram">Sequence Diagram</a></li>
      </ul>
      <ul>
       <li><a href="#high-level-architecture">High Level Architecture</a></li>
      </ul>
      <ul>
        <li><a href="#api-documentation">API Documentation</a></li>
      </ul>
    </li>
    <li>
        <a href="#developement">Developement</a>
        <ul>
            <li><a href="#required-stuff">Required Stuff</a></li>
        </ul>
        <ul>
            <li><a href="#required-step">Required Step</a></li>
        </ul>
        <ul>
            <li><a href="#usage">Usage</a></li>
        </ul>
    </li>
    <li><a href="#contributors">Contributors</a></li>
  
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Open funding is a service that aims to fund projects ranging from individual and community scale and large, this service is like a broker but with free funding and the target is the community.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This the tech stack that we used for developing Open Funding service

* [![AWS][aws]][aws-url]
* [![Golang][golang]][golang-url]
* [![Postgre][postgre]][postgre-url]
* [![Docker][docker]][docker-url]
* [![Fiber][fiber]][fiber-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- FEATURE -->
### Feature
- Creating Campaign
    - a community can add campaigns for the project it wants to fund, one community can create multiple campaigns.
- Payment
    - anyone can donate to a campaign, it's easygoing cause we build using a payment gateway
- Disbursement Automatic
    - The community admin automatically got the fund to their bank account or E-wallet that connected to Open Funding Service.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Sequence Diagram -->
### Sequence Diagram

<br />
<div align="center">
  <a href="https://drive.google.com/file/d/1VeXE4rAx5lLIBSRvYfI97deVdpDE8PzD/view?usp=sharing">
    <img src="https://i.ibb.co/XCYNc6B/Patungan-sequence-diagram-drawio-1.png" alt="hla">
  </a>
</div>
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- HLA -->
### High Level Architecture

<br />
<div align="center">
  <a href="https://drive.google.com/file/d/1fj3VTV65RnXKy032kT4l5ZWvpOijo46F/view?usp=sharing">
    <img src="https://i.ibb.co/gvqXRwR/High-Level-Diagram-drawio.png" alt="hla">
  </a>
</div>
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Api Documentation -->
### API Documentation

<br />
<div align="center">
  <a href="https://documenter.getpostman.com/view/11975231/2s84LGVuo8">
  Open Funding RESTful API Docs
  </a>
</div>
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Developement -->
## Developement

<br />

### Required Stuff:

- Docker
- Docker Compose

### Required Step:

Copy `.env-example` and rename to `.env`

### Usage:

Command to run dev environment :
```bash
go mod tidy
```

Command to run dev environment :
```bash
./run-dev.sh
```

Command to start testing :
```bash
./test.sh
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Contributors -->
## Contributors

Havis Iqbal Syahrunizar - [@hirasakavizu](https://twitter.com/hirasakavizu) - havisikkubaru@gmail.com

Github: [https://github.com/arch-havis](https://github.com/arch-havis)

[![linkedinhavis][linkedinhavis-shield]][linkedinhavis-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/romodeus/manifesto-backend.svg?style=for-the-badge
[contributors-url]: https://github.com/romodeus/manifesto-backend/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/romodeus/manifesto-backend.svg?style=for-the-badge
[forks-url]: https://github.com/romodeus/manifesto-backend/network/members

[stars-shield]: https://img.shields.io/github/stars/romodeus/manifesto-backend.svg?style=for-the-badge
[stars-url]: https://github.com/romodeus/manifesto-backend/stargazers

[issues-shield]: https://img.shields.io/github/issues/romodeus/manifesto-backend.svg?style=for-the-badge
[issues-url]: https://github.com/romodeus/manifesto-backend/issues

[linkedinhavis-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedinhavis-url]: https://www.linkedin.com/in/havis-iqbal/


[fiber]: https://img.shields.io/badge/Fiber-gray?style=for-the-badge&logo=images/fiber-logo&logoColor=00ADD8
[fiber-url]: https://gofiber.io/

[s3]: https://img.shields.io/badge/S3-gray?style=for-the-badge&logo=amazons3&logoColor=569A31
[s3-url]: https://aws.amazon.com/s3/

[aws]: https://img.shields.io/badge/AWS-gray?style=for-the-badge&logo=amazonaws&logoColor=FF9900
[aws-url]: https://aws.amazon.com/

[golang]: https://img.shields.io/badge/golang-gray?style=for-the-badge&logo=go&logoColor=00ADD8
[golang-url]: https://go.dev/

[docker]: https://img.shields.io/badge/Docker-gray?style=for-the-badge&logo=docker&logoColor=00ADD8
[docker-url]: https://www.docker.com/

[postgre]: https://img.shields.io/badge/Postgre-gray?style=for-the-badge&logo=postgresql&logoColor=00ADD8
[postgre-url]: https://www.postgresql.org/
