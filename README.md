# SzDiscordBot

SzDiscordBot is a Discord bot written in Go, designed to enhance your server experience with various features and commands.

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (Golang)


### Planned Features
 - [ ] **Game Management via Plugins**
    - [ ] Minecraft
    - [ ] Ark Survival Ascended
    - [ ] Rust
    - [ ] Conan Exiles
    - [ ] 7 Days to Die
- [ ] **Management of Server Status via:**
    - [ ] Pterodactyl
    - [ ] Docker





### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/SzDiscordBot.git
    ```
2. Navigate to the project directory:
    ```sh
    cd SzDiscordBot
    ```
3. Install the dependencies:
    ```sh
    go mod download
    ```

### Environment Variables

Create a `.env` file in the root directory of the project and add the following environment variables:

```env
DISCORD_BOT_TOKEN=your_discord_bot_token
GUILD_ID=your_guild_id
```

- `DISCORD_BOT_TOKEN`: Your Discord bot token.
- `GUILD_ID`: The GuildID of the server this bot will be on.

### Running the Bot

To start the bot, use the following command:

```sh
go run main.go
```


### Using Docker

To run the bot using Docker, follow these steps:

1. Ensure you have Docker and Docker Compose installed on your machine.
2. Create a `.env` file in the root directory of the project and add the necessary environment variables as mentioned above.
3. Use the following command to start the bot with Docker Compose:

    ```sh
    docker-compose up
    ```

This will build the Docker image and start the bot in a container.

## Contributing

Please read `CONTRIBUTING.md` for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the `LICENSE.md` file for details.

## Acknowledgments

- Hat tip to anyone whose code was used
- Inspiration
- etc.
