import { Link } from "react-router-dom";
import { imgUrl } from "../api/tmdb";
import { useWatchlist } from "../context/WatchlistContext";

export default function MovieCard({ movie }) {
  const { addToWatchlist, removeFromWatchlist, isInWatchlist } = useWatchlist();
  const saved = isInWatchlist(movie.id);

  return (
    <div className="movie-card">
      <Link to={`/movie/${movie.id}`}>
        <img src={imgUrl(movie.poster_path)} alt={movie.title} />
        <h3>{movie.title}</h3>
      </Link>
      <button onClick={() => (saved ? removeFromWatchlist(movie.id) : addToWatchlist(movie))}>
        {saved ? "− Remove" : "+ Watchlist"}
      </button>
    </div>
  );
}