
<a name="readme-top"></a>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/omar4f6d6172/Weather-Data-Display-System">

  </a>

  <h3 align="center">Weather-Data-Display-System</h3>

  <p align="center">
     Fetching and Displaying Weather Data via Raspberry Pi and STM32
    <br />
    <a href="https://github.com/omar4f6d6172/Weather-Data-Display-System"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/omar4f6d6172/Weather-Data-Display-System">View Demo</a>
    ·
    <a href="https://github.com/omar4f6d6172/Weather-Data-Display-System/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/omar4f6d6172/Weather-Data-Display-System/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
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
        <li><a href="#installation">Installation</a></li>
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



## Overview

The Weather Data Display System is an innovative project designed to fetch, process, and display weather data on a hardware setup involving a Raspberry Pi and STM32 microcontroller. It leverages OpenMeteo API to retrieve weather data, which is then sent through SPI communication from a Raspberry Pi to an STM32, ultimately displaying the information on an EA DOGM162L-A LCD display.

![](diagram.svg)

## Technologies Used

- Raspberry Pi: Manages API requests, processes data, and handles SPI data transmission.
- STM32 (ST Micro 'Nucleo' L432KC): Receives and processes SPI data to display.
- OpenMeteo API: Provides real-time weather data.
- Cloudflare Tunnel: Ensures secure access to the Raspberry Pi server.
- SPI Communication: Facilitates data transfer between Raspberry Pi and STM32.
- HTML/CSS: For the web interface hosted on the Raspberry Pi.


## Getting Started

To get started with the Weather Data Display System, follow these steps to set up the project on your Raspberry Pi and STM32 microcontroller.

### Prerequisites

Ensure you have the following prerequisites installed:

- Git
- Go (for compiling the Raspberry Pi application)
- TinyGo (for compiling and flashing the STM32 application)
- SSH access to your Raspberry Pi


### Installation

Follow these steps to install and run the application:

1. Clone the repository:
    ```sh
    git clone https://github.com/omar4f6d6172/Weather-Data-Display-System.git
    ```

2. Navigate to the project directory:
    ```sh
    cd Weather-Data-Display-System
    ```

3. Run the deployment script for the Raspberry Pi:
    ```sh
    ./deploy.sh
    ```

4. For the STM32 microcontroller, use TinyGo to flash the device:
    ```sh
    tinygo flash -target=nucleo-l432kc stm32/main.go
    ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
## Usage
Raspberry Pi

After deploying with the script, the Raspberry Pi application will start automatically and host a web interface accessible via the configured domain or IP address. This interface allows you to view and manage weather data collected through the OpenMeteo API.
STM32

The STM32 microcontroller displays the weather data sent from the Raspberry Pi on the EA DOGM162L-A LCD. The microcontroller should be connected as specified in the hardware setup documentation, and it will start displaying data once flashed and powered.

For more details on interacting with the system and examples of the interface, please refer to the Documentation.
<p align="right">(<a href="#readme-top">back to top</a>)</p>
## Roadmap

- [x] Implement weather data fetching via OpenMeteo API.
- [x] Develop deployment scripts for semi automated setup.
- [ ] Enable SPI communication between Raspberry Pi and STM32.
- [ ] Enhance the user interface for better interaction.



<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>






<!-- CONTACT -->
## Contact

[Omar Altanbakji](mailto:omar.4f6d72@gmail.com) 

Project Link: [https://github.com/omar4f6d6172/Weather-Data-Display-System](https://github.com/omar4f6d6172/Weather-Data-Display-System)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


