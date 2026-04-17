# GREEN-API Web UI

A web-based interface for interacting with GREEN-API WhatsApp API methods. This application provides a user-friendly HTML interface to test and use GREEN-API endpoints.

## Features

- **Connection Parameters**: Input fields for `idInstance` and `ApiTokenInstance`
- **API Methods**:
  - `getSettings` - Retrieve instance configuration
  - `getStateInstance` - Get instance state
  - `sendMessage` - Send text messages via WhatsApp
  - `sendFileByUrl` - Send files via URL
- **Response Display**: Read-only field showing API responses in JSON format
- **Modern UI**: Clean, responsive design with gradient styling

## Requirements

- Go 1.26 or higher
- Fiber v3 web framework

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/green-api-web-ui.git
cd green-api-web-ui
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

### Local Development

1. Start the server:
```bash
go run cmd/app/main.go
```

2. Open your browser and navigate to:
```
http://localhost:3000
```

### Building for Production

1. Build the executable:
```bash
go build -o green-api-web-ui cmd/app/main.go
```

2. Run the executable:
```bash
./green-api-web-ui
```

## Usage

1. **Create GREEN-API Instance**:
   - Go to [GREEN-API](https://green-api.com) and create a free developer account
   - Create a new instance in your dashboard
   - Scan the QR code to connect your WhatsApp number

2. **Get Credentials**:
   - Copy `idInstance` from your GREEN-API dashboard
   - Copy `ApiTokenInstance` from your GREEN-API dashboard

3. **Use the Web UI**:
   - Enter your `idInstance` and `ApiTokenInstance` in the connection parameters section
   - Click on the method buttons to test different API calls
   - View the response in the read-only response field

### Method Parameters

#### Get Settings
- No additional parameters required

#### Get State Instance
- No additional parameters required

#### Send Message
- **Chat ID**: WhatsApp chat ID (e.g., `79001234567@c.us`)
- **Message**: Text message to send

#### Send File by URL
- **Chat ID**: WhatsApp chat ID (e.g., `79001234567@c.us`)
- **File URL**: Public URL of the file to send
- **File Name**: Name of the file (e.g., `document.pdf`)
- **Caption** (optional): File description

## Deployment

### Deploy to Vercel/Netlify (Static HTML)

Since this is a Go application, you can deploy it to various cloud platforms:

### Option 1: Render.com
1. Create a `render.yaml` file:
```yaml
services:
  - type: web
    name: green-api-web-ui
    env: go
    buildCommand: go build -o bin/main cmd/app/main.go
    startCommand: ./bin/main
    envVars:
      - key: PORT
        value: 10000
```

2. Push to GitHub and connect to Render

### Option 2: Railway
1. Install Railway CLI
2. Login: `railway login`
3. Deploy: `railway up`

### Option 3: Heroku
1. Create a `Procfile`:
```
web: green-api-web-ui
```

2. Deploy:
```bash
heroku create
git push heroku main
```

### Option 4: VPS (DigitalOcean, AWS, etc.)
1. SSH into your server
2. Clone the repository
3. Build and run:
```bash
go build -o green-api-web-ui cmd/app/main.go
./green-api-web-ui
```

4. Use a process manager like systemd or PM2 to keep it running

## Project Structure

```
green-api-web-ui/
├── cmd/
│   └── app/
│       └── main.go          # Main application file
├── public/
│   └── index.html          # HTML interface
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
├── .gitignore
└── README.md
```

## API Endpoints

The application exposes the following API endpoints:

- `POST /api/getSettings` - Get instance settings
- `POST /api/getStateInstance` - Get instance state
- `POST /api/sendMessage` - Send a text message
- `POST /api/sendFileByUrl` - Send a file by URL

## Troubleshooting

### Connection Errors
- Ensure your GREEN-API instance is active and authorized
- Check that your `idInstance` and `ApiTokenInstance` are correct
- Verify your internet connection

### Message Not Sending
- Ensure the chat ID format is correct (e.g., `79001234567@c.us`)
- Check that your WhatsApp number is connected to the instance
- Verify the instance state is `authorized`

## License

This project is created as a test assignment for GREEN-API.

## Contact

For questions about this project, please contact the developer.

## GREEN-API Documentation

For more information about GREEN-API, visit:
- [Official Documentation](https://green-api.com/en/docs/)
- [API Reference](https://green-api.com/en/docs/api/)