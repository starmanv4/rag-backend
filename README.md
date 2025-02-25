# RAG-Backend

This project is a **Retrieval-Augmented Generation (RAG) backend** implemented in **Go**. It provides an API to upload documents, retrieve relevant content based on user queries, and generate responses using OpenAI's GPT model.

## Features

-   **File Upload**: Upload text files to be stored in memory.
-   **Context Retrieval**: Extract relevant content from uploaded documents based on a user query.
-   **AI Response Generation**: Use OpenAI's GPT model to generate responses based on retrieved context.
-   **CORS Support**: Configured to allow frontend access.

## Project Structure

```
rag-backend/
│── rag/
│   ├── generator.go   # Handles OpenAI API interactions
│   ├── retriever.go   # Retrieves relevant context from uploaded documents
│── server/
│   ├── handlers.go    # Defines API handlers
│   ├── router.go      # Configures API routes
│   ├── server.go      # Initializes and starts the HTTP server
│── storage/
│   ├── memory.go      # Manages in-memory file storage
│── utils/
│   ├── parser.go      # Text processing utilities
│── .env               # Stores environment variables (not included in version control)
│── .gitignore         # Git ignore rules
│── go.mod             # Go module dependencies
│── go.sum             # Dependency checksums
│── main.go            # Entry point of the application
```

## Dependencies

This project uses the following Go packages:

-   [gorilla/mux](https://github.com/gorilla/mux) - HTTP router for request handling
-   [gorilla/handlers](https://github.com/gorilla/handlers) - Middleware for CORS support
-   [joho/godotenv](https://github.com/joho/godotenv) - Loads environment variables from `.env`
-   [sashabaranov/go-openai](https://github.com/sashabaranov/go-openai) - OpenAI API client

## Setup

### Prerequisites

-   Install **Go** (1.24 or higher)
-   Get an **OpenAI API key** from [OpenAI](https://platform.openai.com/)

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/starmanv4/rag-backend.git
    cd rag-backend
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

3. Create a `.env` file and add your **OpenAI API key**:
    ```
    OPENAI_API_KEY=your_openai_api_key_here
    ```

### Running the Server

Start the server with:

```sh
go run main.go
```

The server will start on **http://localhost:8080**.

## API Endpoints

### 1. **Upload File**

-   **Endpoint**: `/upload`
-   **Method**: `POST`
-   **Content-Type**: `multipart/form-data`
-   **Description**: Uploads a text file for processing.

#### Example Request (cURL)

```sh
curl -X POST -F "file=@sample.txt" http://localhost:8080/upload
```

#### Example Response

```json
{
	"message": "File uploaded successfully",
	"filename": "sample.txt"
}
```

### 2. **Query Document**

-   **Endpoint**: `/query`
-   **Method**: `POST`
-   **Content-Type**: `application/json`
-   **Description**: Submits a query and receives a response based on uploaded documents.

#### Example Request (cURL)

```sh
curl -X POST http://localhost:8080/query \
     -H "Content-Type: application/json" \
     -d '{"query": "What is the main topic of the document?"}'
```

#### Example Response

```json
{
	"response": "The document discusses artificial intelligence and its applications."
}
```

## Environment Variables

| Variable         | Description                                       |
| ---------------- | ------------------------------------------------- |
| `OPENAI_API_KEY` | OpenAI API key (required for response generation) |

## Future Improvements

-   **Persistent Storage**: Store uploaded files in a database instead of memory.
-   **Authentication**: Implement API key-based authentication.
-   **More Retrieval Methods**: Enhance document retrieval with embeddings.

## License

This project is licensed under the **MIT License**.

---

> **Author**: [starmanv4](https://github.com/starmanv4)  
> **GitHub**: [github.com/starmanv4/rag-backend](https://github.com/starmanv4/rag-backend)
