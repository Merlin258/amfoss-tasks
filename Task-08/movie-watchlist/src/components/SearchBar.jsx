import { useState } from "react";

export default function SearchBar({ onSearch }) {
  const [value, setValue] = useState("");

  function handleSubmit(e) {
    e.preventDefault();
    onSearch(value.trim());
  }

  return (
    <form onSubmit={handleSubmit} className="search-bar">
      <input
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Search movies..."
      />
      <button type="submit">Search</button>
    </form>
  );
}