import MovieCard from "./MovieCard";

export default function MovieGrid({ movies }) {
  if (!movies.length) return <p>No movies found.</p>;
  return (
    <div className="movie-grid">
      {movies.map((m) => <MovieCard key={m.id} movie={m} />)}
    </div>
  );
}