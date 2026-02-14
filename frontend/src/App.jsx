import { useEffect, useState } from "react";
import "./index.css";

function App() {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/books")
      .then(res => res.json())
      .then(data => setBooks(data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div className="container">
      <h1 className="title">Book App</h1>

      <div className="grid">
        {books.map((book) => (
          <div key={book.id} className="card">
            <h3>{book.title}</h3>
          </div>
        ))}
      </div>
    </div>
  );
}

export default App;
