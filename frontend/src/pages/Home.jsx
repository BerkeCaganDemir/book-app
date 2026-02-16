import { useState, useEffect } from "react";
import Header from "../components/Header";
import BookList from "../components/BookList";

function Home() {
  const [searchText, setSearchText] = useState("");
  const [viewType, setViewType] = useState("list");

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingBook, setEditingBook] = useState(null);

  const [formData, setFormData] = useState({
    title: "",
    author: "",
    note: "",
    url: "",
    image: null,
  });

  const [books, setBooks] = useState([]);

  
  // FETCH BOOKS
  
  const fetchBooks = async () => {
    const res = await fetch("http://localhost:8080/books");
    const data = await res.json();
    setBooks(data);
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const filteredBooks = books.filter(
    (book) =>
      book.title.toLowerCase().includes(searchText.toLowerCase()) ||
      book.author.toLowerCase().includes(searchText.toLowerCase())
  );

  
  // CREATE
  
  const handleSave = async () => {
    try {
      const response = await fetch("http://localhost:8080/books", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          title: formData.title,
          author: formData.author,
          notes: formData.note,
          buyUrl: formData.url,
        }),
      });

      if (!response.ok) {
        console.error("Create failed");
        return;
      }

      await fetchBooks();

      setIsModalOpen(false);
      setFormData({
        title: "",
        author: "",
        note: "",
        url: "",
        image: null,
      });
    } catch (err) {
      console.error(err);
    }
  };

  
  // DELETE
  
  const handleDelete = async (id) => {
    const response = await fetch(
      `http://localhost:8080/books/${id}`,
      { method: "DELETE" }
    );

    if (!response.ok) {
      console.error("Delete failed");
      return;
    }

    await fetchBooks();
  };

  
  // EDIT START
 
  const handleEdit = (book) => {
    setEditingBook(book);
    setFormData({
      title: book.title,
      author: book.author,
      note: book.notes,
      url: book.buyUrl,
      image: null,
    });
  };

  
  // UPDATE
  
  const handleUpdate = async () => {
    if (!editingBook) return;

    // 1 Text update
    await fetch(`http://localhost:8080/books/${editingBook.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: formData.title,
        author: formData.author,
        notes: formData.note,
        buyUrl: formData.url,
        imageUrl: editingBook.imageUrl,
      }),
    });

    //  Image upload 
    if (formData.image) {
      const imageData = new FormData();
      imageData.append("image", formData.image);

      await fetch(
        `http://localhost:8080/books/${editingBook.id}/image`,
        {
          method: "POST",
          body: imageData,
        }
      );
    }

    await fetchBooks();

    setEditingBook(null);
    setFormData((prev) => ({ ...prev, image: null }));
  };

  return (
    <div className="p-10 bg-gray-200 min-h-screen">
      <Header
        searchText={searchText}
        setSearchText={setSearchText}
        viewType={viewType}
        setViewType={setViewType}
        onAddClick={() => setIsModalOpen(true)}
      />

      <div className="mt-6">
        <BookList
          books={filteredBooks}
          viewType={viewType}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      </div>

      
      {/* CREATE MODAL */}
      
      {isModalOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center">
          <div className="bg-white p-8 rounded-xl w-96">
            <h2 className="text-xl font-semibold mb-4">Add Book</h2>

            <input
              type="text"
              placeholder="Title"
              value={formData.title}
              onChange={(e) =>
                setFormData({ ...formData, title: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <input
              type="text"
              placeholder="Author"
              value={formData.author}
              onChange={(e) =>
                setFormData({ ...formData, author: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <textarea
              placeholder="Note"
              value={formData.note}
              onChange={(e) =>
                setFormData({ ...formData, note: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <input
              type="text"
              placeholder="Buy URL"
              value={formData.url}
              onChange={(e) =>
                setFormData({ ...formData, url: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <div className="flex justify-end gap-2">
              <button
                onClick={() => setIsModalOpen(false)}
                className="px-4 py-2 border rounded"
              >
                Cancel
              </button>

              <button
                onClick={handleSave}
                className="px-4 py-2 bg-blue-600 text-white rounded"
              >
                Save
              </button>
            </div>
          </div>
        </div>
      )}

     
      {/* EDIT MODAL */}
      
      {editingBook && (
        <div className="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center">
          <div className="bg-white p-8 rounded-xl w-96">
            <h2 className="text-xl font-semibold mb-4">
              Edit Book
            </h2>

            <input
              type="text"
              value={formData.title}
              onChange={(e) =>
                setFormData({ ...formData, title: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <input
              type="text"
              value={formData.author}
              onChange={(e) =>
                setFormData({ ...formData, author: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <textarea
              value={formData.note}
              onChange={(e) =>
                setFormData({ ...formData, note: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <input
              type="text"
              value={formData.url}
              onChange={(e) =>
                setFormData({ ...formData, url: e.target.value })
              }
              className="w-full border p-2 mb-3 rounded"
            />

            <input
              type="file"
              accept="image/*"
              onChange={(e) =>
                setFormData({
                  ...formData,
                  image: e.target.files[0],
                })
              }
              className="w-full mb-4"
            />

            <div className="flex justify-end gap-2">
              <button
                onClick={() => setEditingBook(null)}
                className="px-4 py-2 border rounded"
              >
                Cancel
              </button>

              <button
                onClick={handleUpdate}
                className="px-4 py-2 bg-green-600 text-white rounded"
              >
                Save Changes
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default Home;
