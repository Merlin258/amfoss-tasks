import { useState, useEffect } from "react";
import { fetchTrending, searchMovies } from "../api/tmdb";

export function useMovies(query) {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let cancelled = false;
    setLoading(true);
    const load = query ? searchMovies(query) : fetchTrending();

    load
      .then((data) => !cancelled && setMovies(data))
      .catch((err) => !cancelled && setError(err.message))
      .finally(() => !cancelled && setLoading(false));

    return () => { cancelled = true; }; // cleanup: avoid state update on unmounted component
  }, [query]);

  return { movies, loading, error };
}