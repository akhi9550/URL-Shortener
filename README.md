# URL SHORTENER
This code provides a simple implementation of a URL shortening service with basic URL validation and short URL generation.

1. Imports and Variables:
- **math/rand** for generating random strings.
- **net/http** and **net/url** for handling HTTP requests and URL validation.
- **sync** for managing concurrent access to the URL store.
- **gin** from the Gin Gonic library for web server functionalities.

2. Global Variables:
- **urlStore**: A map to store the mapping between shortened URLs and original URLs.
- **mu**: A read-write mutex to handle concurrent access to urlStore.
- **baseURL**: The base URL for generating full shortened URLs.

3. Main Function:
- **POST /shorten**: Handles URL shortening requests.
- **GET /:shortUrl**: Retrieves the original URL from a shortened URL.

4. Request Handling:
- **shortenURLHandler**: Takes a URL from the request body, validates it, generates a short URL, stores the mapping, and returns the shortened URL.
- **retrieveURLHandler**: Retrieves the original URL associated with a shortened URL and redirects to it.

5. Helper Functions:

- **isValidURL**: Checks if a given URL is valid.
- **generateShortURL**: Creates a random short URL using a defined character set.

- ## main_test.go
- This test file, url_test.go, is designed to test the functionality of the URL shortener service implemented in the main package of the Go application. The tests are organized in a separate package (main_test) to maintain a clean separation between the test logic and the main application logic.

## Postman Api documentation

[API Documentation](https://documenter.getpostman.com/view/29514478/2sAXjQ2AEi)