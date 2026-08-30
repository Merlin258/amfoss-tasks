const BASE_URL = "https://api.themoviedb.org/3";
const token = import.meta.env.VITE_TMDB_TOKEN;

const headers = {
  Authorization: `Bearer ${token}`,
  "Content-Type": "application/json",
};

export async function fetchTrending() {
  const res = await fetch(`${BASE_URL}/trending/movie/week`, { headers });
  if (!res.ok) throw new Error("Failed to fetch trending movies");
  return (await res.json()).results;
}

export async function searchMovies(query) {
  const res = await fetch(
    `${BASE_URL}/search/movie?query=${encodeURIComponent(query)}`,
    { headers }
  );
  if (!res.ok) throw new Error("Search failed");
  return (await res.json()).results;
}

export async function fetchMovieDetails(id) {
  const res = await fetch(`${BASE_URL}/movie/${id}`, { headers });
  if (!res.ok) throw new Error("Failed to fetch movie details");
  return await res.json();
}

export const imgUrl = (path, size = "w500") =>
  path ? `https://image.tmdb.org/t/p/${size}${path}` : "/no-poster.png";