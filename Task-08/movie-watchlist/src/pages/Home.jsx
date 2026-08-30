import { useState } from "react";
import SearchBar from "../components/SearchBar";
import MovieGrid from "../components/MovieGrid";
import { useMovies } from "../hooks/useMovies";

export default function Home() {
  const [query, setQuery] = useState("");
  const { movies, loading, error } = useMovies(query);

  return (
    <div>
      <SearchBar onSearch={setQuery} />
      <h2>{query ? `Results for "${query}"` : "Trending This Week"}</h2>
      {loading && <p>Loading...</p>}
      {error && <p className="error">{error}</p>}
      {!loading && !error && <MovieGrid movies={movies} />}
    </div>
  );
}