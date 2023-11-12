# Personal Information Agent

This project implements a Personal Information Agent, a web app that serves personalized content to professionals based on their interests. It utilizes server-side rendering with Go and the Colly scraping library to fetch and display the latest articles from specified domains.

## Features

- Personalized content based on user input
- Server-side rendering for dynamic content updates
- Web scraping using Colly for up-to-date information retrieval

## Prerequisites

Before running this project, you need to have the following installed:
- Go (version 1.21 or later)
- Git (for cloning the repository)

## Installation

Follow these steps to get your development environment set up:

1. Clone the repository to your local machine using Git:

```bash
git clone https://github.com/yourusername/personal-information-agent.git
cd personal-information-agent
```

2. Download and install the dependencies:

```bash
go mod tidy
```

## Configuration

Create a `.env` file in the root directory and specify the following variables:

```env
PORT=8080
TARGET_DOMAIN=example.com
TARGET_URL=https://www.example.com/latest
```

Replace the `TARGET_DOMAIN` and `TARGET_URL` with the domain and URL you wish to scrape content from.

## Running the Application

To run the application, execute the following command in the terminal:

```bash
go run main.go
```

The server will start and listen on the port specified in your configuration (default is `8080`). You can access the web application by navigating to `http://localhost:8080` in your web browser.

## Usage

Once the application is running:
- Navigate to the homepage.
- Enter your interests and desired update frequency in the form provided.
- Click the "Customize" button to see personalized content based on your input.

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch: `git checkout -b my-new-feature`.
3. Make your changes and commit them: `git commit -am 'Add some feature'`.
4. Push to the branch: `git push origin my-new-feature`.
5. Submit a pull request.

## Support

If you encounter any problems, please file an issue along with a detailed description.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Acknowledgments

- This project uses the [Colly](http://go-colly.org/) scraping framework.
