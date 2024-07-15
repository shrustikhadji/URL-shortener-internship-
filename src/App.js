import React, { useState } from "react";
import axios from "axios";
import "./App.css";

function App() {
  const [url, setURL] = useState("");
  const [shortURL, setShortURL] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setShortURL("");
    setError("");

    try {
      const response = await axios.post("http://localhost:8080/shorten", { url });
      setShortURL(response.data.short_url);
    } catch (error) {
      if (error.response) {
        switch (error.response.status) {
          case 400:
            setError("Invalid URL provided.");
            break;
          case 409:
            setError("Short URL already exists.");
            break;
          case 500:
            setError("Server error. Please try again later.");
            break;
          default:
            setError("An unexpected error occurred.");
        }
      } else if (error.request) {
        setError("No response from server. Please check your network connection.");
      } else {
        setError("Error setting up request. Please try again.");
      }
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>ShortLINKit!</h1>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={url}
            onChange={(e) => setURL(e.target.value)}
            placeholder="Enter URL"
          />
          <button type="submit">Shorten</button>
        </form>
        {error && <p className="error">{error}</p>}
        {shortURL && (
          <div className="short-url">
            <h2>Short URL:</h2>
            <a href={`http://localhost:8080/${shortURL}`} target="_blank" rel="noopener noreferrer">
              {`http://localhost:8080/${shortURL}`}
            </a>
          </div>
        )}
      </header>
    </div>
  );
}

export default App;
