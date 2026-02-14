function Header({ searchText, setSearchText, viewType, setViewType, onAddClick }) {
  return (
    <div className="flex items-center gap-4">
      <input
        type="text"
        placeholder="Search books..."
        value={searchText}
        onChange={(e) => setSearchText(e.target.value)}
        className="flex-1 p-3 rounded-lg border"
      />

      <button
        onClick={onAddClick}
        className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
      >
        +
      </button>

      <button
        onClick={() => setViewType("list")}
        className={viewType === "list" ? "bg-black text-white px-3 py-2 rounded" : ""}
      >
        List
      </button>

      <button
        onClick={() => setViewType("grid")}
      >
        Grid
      </button>
    </div>
  );
}

export default Header;
