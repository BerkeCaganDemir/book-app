function BookList({ books, viewType, onEdit, onDelete }) {
  if (!books.length) {
    return (
      <div className="text-gray-500 mt-10">
        No books found.
      </div>
    );
  }

  if (viewType === "grid") {
    return (
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {books.map((book) => (
          <div
            key={book.id}
            className="bg-white rounded-xl shadow-sm hover:shadow-md transition overflow-hidden"
          >
            {book.imageUrl ? (
              <img
                src={`http://localhost:8080${book.imageUrl}`}
                alt={book.title}
                className="w-full h-48 object-cover"
              />
            ) : (
              <div className="w-full h-48 bg-gray-100 flex items-center justify-center text-gray-400 text-sm">
                No Image
              </div>
            )}

            <div className="p-6">
              <h3 className="text-lg font-semibold">
                {book.title}
              </h3>

              <p className="text-sm text-gray-500 mb-2">
                {book.author}
              </p>

              {book.notes && (
                <p className="text-sm text-gray-400 mb-2">
                  {book.notes}
                </p>
              )}

              {book.buyUrl && (
                <a
                  href={book.buyUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="text-blue-600 text-sm hover:underline block mb-2"
                >
                  Buy Link
                </a>
              )}

              <div className="flex justify-between mt-4">
                <button
                  onClick={() => onEdit(book)}
                  className="text-sm text-blue-600 hover:underline"
                >
                  Edit
                </button>

                <button
                  onClick={() => onDelete(book.id)}
                  className="text-sm text-red-600 hover:underline"
                >
                  Delete
                </button>
              </div>
            </div>
          </div>
        ))}
      </div>
    );
  }

  // LIST VIEW
  return (
    <div className="space-y-4">
      {books.map((book) => (
        <div
          key={book.id}
          className="bg-white p-4 rounded-lg shadow-sm flex justify-between items-center"
        >
          <div>
            <h3 className="font-medium">{book.title}</h3>
            <p className="text-sm text-gray-500">
              {book.author}
            </p>
          </div>

          <div className="flex gap-4">
            <button
              onClick={() => onEdit(book)}
              className="text-blue-600 text-sm hover:underline"
            >
              Edit
            </button>

            <button
              onClick={() => onDelete(book.id)}
              className="text-red-600 text-sm hover:underline"
            >
              Delete
            </button>
          </div>
        </div>
      ))}
    </div>
  );
}

export default BookList;
