# Go AI Gateway

A simple, robust, and scalable backend service written in Go that acts as a gateway to Large Language Models (LLMs) like Google's Gemini. It exposes a clean REST API for performing various NLP tasks, such as summarization, translation, and more, based on user-provided prompts.

## Features

* **Dynamic Prompting**: Users can supply their own prompts to perform a wide range of tasks, not just summarization.
* **Clean 3-Tier Architecture**: Enforces a strict separation of concerns (Controllers, Services, Repositories) for maintainability and scalability.
* **Database Logging**: All requests and AI-generated responses are logged to a MySQL database for auditing and analysis.
* **Ready for Production**: Built with a robust structure using Gin for routing, official Go SQL drivers, and environment-based configuration.

## Architecture

The project follows a classic 3-Tier Architecture to ensure the codebase is modular and easy to manage.

1.  **Controllers (Presentation Layer)**: Located in the `/controllers` directory. Responsible for handling incoming HTTP requests, validating the JSON payload, and formatting the final JSON response.
2.  **Services (Business Logic Layer)**: Located in the `/services` directory. Contains the core application logic, such as constructing the prompt for the AI, calling the external LLM API, and coordinating with the repository.
3.  **Repositories (Data Access Layer)**: Located in the `/repository` directory. This layer is responsible for all communication with the database. It abstracts the SQL `INSERT` queries from the rest of the application.

## Getting Started

Follow these instructions to get a local copy up and running.

### Prerequisites

* [Go](https://go.dev/doc/install) (version 1.18 or higher)
* [MySQL](https://dev.mysql.com/downloads/mysql/) (or another compatible SQL database)
* A Google Gemini API Key. You can get one from [Google AI Studio](https://aistudio.google.com/).

### Installation

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/AryaTabani/go-ai-gateway.git
    cd go-ai-gateway
    ```

2.  **Create a configuration file:**
    Create a file named `.env` in the root of the project and add your configuration details.
    ```env
    # Database Configuration
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_HOST=127.0.0.1
    DB_PORT=3306
    DB_NAME=ai_app_db

    # Google Gemini API Key
    GEMINI_API_KEY="your_gemini_api_key_here"
    ```
    *Ensure the database `ai_app_db` exists on your MySQL server.*

3.  **Install dependencies and run the server:**
    The application will automatically download the required modules.
    ```sh
    go run main.go
    ```
    The server will start and be available at `http://localhost:8080`.

## API Usage

The primary endpoint is `POST /api/v1/ai/summarize`. You can interact with it using any API client or `curl`.

### Example 1: Default Summarization

If you don't provide a `prompt`, the service will use a default summarization prompt.

```bash
curl -X POST \
  http://localhost:8080/api/v1/ai/summarize \
  -H 'Content-Type: application/json' \
  -d '{
    "text": "The Eiffel Tower is a wrought-iron lattice tower on the Champ de Mars in Paris, France."
}'
```


### Example 2: Custom Prompt (Changing Tone)

Provide a custom `prompt` to perform a different task, like making a sentence sound more enthusiastic for a team announcement.

```bash
curl -X POST \
  http://localhost:8080/api/v1/ai/summarize \
  -H 'Content-Type: application/json' \
  -d '{
    "text": "The quarterly report is finished and has been uploaded to the drive.",
    "prompt": "Rewrite the following text to sound more enthusiastic and celebratory for a team-wide message: "
}'
```
