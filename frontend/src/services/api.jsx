const BASE_URL = "http://localhost:8080";

export const deleteBook = async (id) => {
  const res = await fetch(`${BASE_URL}/books/${id}`, {
    method: "DELETE",
  });

  if (!res.ok) {
    throw new Error("Failed to delete book");
  }
};
